package config

import (
	"errors"
	"github.com/spf13/viper"
)

var (
	ErrConfigFileNotFound  = errors.New("config file not found")
	ErrConfigUnmarshalling = errors.New("config unmarshalling failed")
)

type AppConfig struct {
	Db DbConfig `mapstructure:"db"`
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
