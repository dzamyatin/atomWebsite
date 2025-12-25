package di

import (
	"context"
	"database/sql"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/dzamyatin/atomWebsite/internal/entity"
	atomWebsite "github.com/dzamyatin/atomWebsite/internal/grpc/generated"
	grpcservice2 "github.com/dzamyatin/atomWebsite/internal/grpc/grpc"
	"github.com/dzamyatin/atomWebsite/internal/repository"
	serviceauth "github.com/dzamyatin/atomWebsite/internal/service/auth"
	"github.com/dzamyatin/atomWebsite/internal/service/config"
	servicemail "github.com/dzamyatin/atomWebsite/internal/service/mail"
	servicemessengerdriver "github.com/dzamyatin/atomWebsite/internal/service/messenger/driver"
	messengerserver "github.com/dzamyatin/atomWebsite/internal/service/messenger/server"
	messengertelegram "github.com/dzamyatin/atomWebsite/internal/service/messenger/telegram"
	"github.com/dzamyatin/atomWebsite/internal/service/metric"
	"github.com/dzamyatin/atomWebsite/internal/service/process"
	servicetrace "github.com/dzamyatin/atomWebsite/internal/service/trace"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	_ "github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/exp/slices"
	"google.golang.org/grpc"
)

func newTelegramBotServer(
	logger *zap.Logger,
) *messengertelegram.TelegramDriver {
	return messengertelegram.NewTelegramDriver(
		getConfig().TelegramBotConfig.Token,
		logger,
	)
}

//
//func newMessenger(
//	logger *zap.Logger,
//) servicemessengersender.ISenderService {
//	return servicemessengersender.NewAggrigatorSender(
//		logger,
//		[]servicemessengersender.ISenderService{
//			//servicemessenger.NewTelegramSender(logger),
//		},
//	)
//}

func newMailer(
	logger *zap.Logger,
) servicemail.IMailer {
	smtpMap := getConfig().Mailer.Smtp

	smtps := make([]config.SmtpConfig, len(smtpMap))

	i := 0
	for _, smtp := range smtpMap {
		smtps[i] = smtp
		i++
	}

	slices.SortFunc(
		smtps,
		func(a, b config.SmtpConfig) int {
			if a.Weight > b.Weight {
				return 1
			}

			if a.Weight < b.Weight {
				return -1
			}

			return 0
		},
	)

	mailers := make([]servicemail.IMailer, len(smtps))
	for i, smtp := range smtps {
		mailers[i] = servicemail.NewMailerGomailSmtp(
			smtp.Host,
			smtp.Port,
			smtp.Username,
			smtp.Password,
			logger,
			smtp.Sender,
			smtp.SSL,
			smtp.Timeout,
		)
	}

	return servicemail.NewMailerAtLeastOneSuccess(
		logger,
		mailers,
	)
}

func newRegistry(
	logger *zap.Logger,
) *metric.Registry {
	registry := metric.NewRegistry(
		logger,
	)
	err := registry.RunGCMetrics()
	if err != nil {
		panic("failed to run gc metrics")
	}

	return registry
}

func newGRPCProcessManager(
	logger *zap.Logger,
	serv *grpcservice2.GRPCServer,
	listener *process.SignalListener,
	db *sql.DB,
	metric *metric.Registry,
	server *grpcservice2.HTTPServer,
) *process.ProcessManager {
	return process.NewProcessManager(
		logger,
		process.NewProcess("grpc server", serv),
		process.NewProcess("http server", server),
		process.NewProcess("signal listener", listener),
		process.NewProcess("database shutdowner", process.NewShutdowner(
			logger,
			func() error {
				return db.Close()
			},
		),
		),
		process.NewProcess("http listener", newHttpMetricProcessor(metric)),
	)
}

func newTracer(server *servicetrace.TraceServer) *servicetrace.Trace {
	return servicetrace.NewTrace(server.GetTrace())
}

func newDb(
	ctx context.Context,
	registry *process.ShutdownerRegistry,
) (*sql.DB, error) {
	config := getConfig().Db

	query := "dbname=" + config.Database
	if !config.SSL {
		query += "&sslmode=disable"
	}

	u := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(config.Username, config.Password),
		Host:     config.Host + ":" + config.Port,
		RawQuery: query,
	}

	db, err := sql.Open("postgres", u.String())

	if err != nil {
		return nil, errors.Wrap(err, "could not connect to postgres")
	}

	registry.Add("db", func() error {
		return db.Close()
	})

	ctx, _ = context.WithTimeout(ctx, 5*time.Second)

	err = db.PingContext(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "could not connect to postgres")
	}

	return db, nil
}

func newDbx(db *sql.DB) (*sqlx.DB, error) {
	dbx := sqlx.NewDb(db, "pgx")

	if err := dbx.Ping(); err != nil {
		return nil, errors.Wrap(err, "sqlx could not connect to postgres")
	}

	return dbx, nil
}

func newLogger() *zap.Logger {
	return zap.New(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
			zapcore.AddSync(os.Stdout),
			zapcore.InfoLevel,
		),
	)
}

func newServer(
	logger *zap.Logger,
	grpcServer *grpc.Server,
) *grpcservice2.GRPCServer {
	lis, err := net.Listen("tcp", getConfig().AddrGrpc)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	return grpcservice2.NewGRPCServer(
		logger,
		lis,
		grpcServer,
	)
}

func newGrpcServer(
	server grpcservice2.AuthServer,
	m *metric.Metric,
) *grpc.Server {
	grpcServer := grpc.NewServer()

	grpc.WithIdleTimeout(time.Minute * 1)
	grpc.WithTimeout(time.Minute * 1)

	grpc.WithChainUnaryInterceptor(
		func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
			var err error

			m.IncomingRequestHistogram(
				func() {
					err = invoker(ctx, method, req, reply, cc, opts...)
				},
			)

			return err
		},
	)

	atomWebsite.RegisterAuthServer(grpcServer, server)

	return grpcServer
}

func newHttpMetricProcessor(r *metric.Registry) *process.Processor {
	mux := http.NewServeMux()
	mux.Handle("/metrics", r.HTTPHandler())

	ctx := context.Background()

	server := &http.Server{
		Addr:    getConfig().AddrMetric,
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			return ctx
		},
	}

	return process.NewProcessor(
		func(_ context.Context) error {
			return server.ListenAndServe()
		},
		func() error {
			return server.Shutdown(ctx)
		},
	)
}

func newSequentialProvider(
	userRepository repository.IUserRepository,
	logger *zap.Logger,
	passwordComparator entity.PasswordComparator,
) *serviceauth.SequentialProvider {
	return serviceauth.NewSequentialProvider(
		logger,
		serviceauth.NewPasswordProvider(
			userRepository,
			logger,
			passwordComparator,
		),
	)
}

func newJWT(logger *zap.Logger) *serviceauth.JWT {
	return serviceauth.NewJWT(
		"hella1245912dasdas",
		128*time.Hour,
		logger,
	)
}

func newHTTPServer(
	logger *zap.Logger,
	router *grpcservice2.HttpRouter,
	metric *metric.Metric,
	trace *servicetrace.Trace,
) *grpcservice2.HTTPServer {
	server := grpcservice2.NewHTTPServer(
		logger,
		getConfig().AddHttp,
		router,
		metric,
		trace,
		grpcservice2.WithTimeout(
			getConfig().HttpServerTimeout,
			getConfig().HttpServerTimeout,
			getConfig().HttpServerTimeout,
		),
	)

	return server
}

func newMessengerServerRegistry(
	telegramDriver *messengertelegram.TelegramDriver,
) *messengerserver.MessengerServerRegistry {
	return messengerserver.NewMessengerServerRegistry(
		[]servicemessengerdriver.IMessengerDriver{
			telegramDriver,
		},
	)
}
