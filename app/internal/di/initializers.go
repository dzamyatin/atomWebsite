package di

import (
	"context"
	"database/sql"
	"github.com/dzamyatin/atomWebsite/internal/entity"
	atomWebsite "github.com/dzamyatin/atomWebsite/internal/grpc/generated"
	grpcservice2 "github.com/dzamyatin/atomWebsite/internal/grpc/grpc"
	"github.com/dzamyatin/atomWebsite/internal/repository"
	serviceauth "github.com/dzamyatin/atomWebsite/internal/service/auth"
	"github.com/dzamyatin/atomWebsite/internal/service/config"
	servicemail "github.com/dzamyatin/atomWebsite/internal/service/mail"
	servicemessenger "github.com/dzamyatin/atomWebsite/internal/service/messenger"
	messengertelegram "github.com/dzamyatin/atomWebsite/internal/service/messenger/telegram"
	"github.com/dzamyatin/atomWebsite/internal/service/metric"
	"github.com/dzamyatin/atomWebsite/internal/service/process"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	_ "github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/exp/slices"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"time"
)

func newTelegramBotServer(
	logger *zap.Logger,
) *messengertelegram.TelegramDriver {
	return messengertelegram.NewTelegramDriver(
		getConfig().TelegramBotConfig.Token,
		logger,
	)
}

func newMessenger(
	logger *zap.Logger,
) servicemessenger.ISenderService {
	return servicemessenger.NewAggrigatorSender(
		logger,
		[]servicemessenger.ISenderService{
			//servicemessenger.NewTelegramSender(logger),
		},
	)
}

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
		process.Process{
			Name:   "grpc server",
			Object: serv,
		},
		process.Process{
			Name:   "http server",
			Object: server,
		},
		process.Process{
			Name:   "signal listener",
			Object: listener,
		},
		process.Process{
			Name: "database shutdowner",
			Object: process.NewShutdowner(
				logger,
				func() error {
					return db.Close()
				},
			),
		},
		process.Process{
			Name:   "http listener",
			Object: newHttpMetricProcessor(metric),
		},
	)
}

func newDb(ctx context.Context) (*sql.DB, error) {
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
	grpcServer *grpc.Server,
) *grpcservice2.GRPCServer {
	lis, err := net.Listen("tcp", getConfig().AddrGrpc)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	return grpcservice2.NewGRPCServer(
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
	server grpcservice2.AuthServer,
) *grpcservice2.HTTPServer {
	return grpcservice2.NewHTTPServer(
		server,
		getConfig().AddHttp,
	)
}
