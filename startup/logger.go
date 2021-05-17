package startup

import (logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitializeLogger(config *viper.Viper){
	env := config.Get("env")

	switch env {
	case "dev":
		logger.SetFormatter(&logger.TextFormatter{})
	case "prod":
		logger.SetFormatter(&logger.JSONFormatter{})
	default:
		logger.SetFormatter(&logger.TextFormatter{})


	}

	logger.Info("Logger initialised")
}
