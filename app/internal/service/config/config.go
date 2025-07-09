package config

import (
	"errors"
	"github.com/spf13/viper"
	"time"
)

var (
	ErrConfigFileNotFound  = errors.New("config file not found")
	ErrConfigUnmarshalling = errors.New("config unmarshalling failed")
)

type SmtpConfig struct {
	Weight   uint64        `mapstructure:"weight"`
	Host     string        `mapstructure:"host"`
	Port     uint32        `mapstructure:"port"`
	Username string        `mapstructure:"username"`
	Password string        `mapstructure:"password"`
	Sender   string        `mapstructure:"sender"`
	SSL      bool          `mapstructure:"ssl"`
	Timeout  time.Duration `mapstructure:"timeout"`
}

type MailerConfig struct {
	Smtp map[string]SmtpConfig `mapstructure:"smtp"`
}

type AppConfig struct {
	Db                DbConfig          `mapstructure:"db"`
	AddrMetric        string            `mapstructure:"addr_metric"`
	AddrGrpc          string            `mapstructure:"addr_grpc"`
	AddHttp           string            `mapstructure:"addr_http"`
	Mailer            MailerConfig      `mapstructure:"mailer"`
	TelegramBotConfig TelegramBotConfig `mapstructure:"telegram_bot"`
}

type TelegramBotConfig struct {
	Token string `mapstructure:"token"`
}

type DbConfig struct {
	SSL      bool   `mapstructure:"ssl"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

func NewAppConfig(path string) (*AppConfig, error) {
	v := viper.New()

	v.SetConfigFile(path)

	err := v.ReadInConfig()
	if err != nil {
		return nil, errors.Join(
			ErrConfigFileNotFound,
			err,
		)
	}

	cnfg := &AppConfig{}

	err = v.Unmarshal(cnfg)
	if err != nil {
		return nil, errors.Join(
			ErrConfigUnmarshalling,
			err,
		)
	}

	return cnfg, nil
}
