package di

import (
	"github.com/dzamyatin/atomWebsite/internal/service/config"
)

var globalAppConf *config.AppConfig

func CreateConfig(path string) (*config.AppConfig, error) {
	if globalAppConf != nil {
		panic("config is already initialized")
	}

	var err error
	globalAppConf, err = config.NewAppConfig(path)

	return globalAppConf, err
}

func getConfig() *config.AppConfig {
	if globalAppConf == nil {
		panic("config not initialized")
	}

	return globalAppConf
}
