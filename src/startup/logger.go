package startup

import (logger "github.com/sirupsen/logrus")

func InitializeLogger(){
	logger.SetFormatter(&logger.JSONFormatter{})
}
