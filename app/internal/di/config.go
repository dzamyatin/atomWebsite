package di

import (
	"github.com/dzamyatin/atomWebsite/internal/service/config"
)

var conf *config.AppConfig

func CreateConfig(path string) (err error) {
	if conf != nil {
		panic("config is already initialized")
	}

	conf, err = config.NewAppConfig(path)

	return
}

func getConfig() *config.AppConfig {
	if conf == nil {
		panic("config not initialized")
	}

	return conf
}
