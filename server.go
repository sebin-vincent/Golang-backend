package main

import (
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	"github.com/wallet-tracky/Golang-backend/startup"
	"github.com/wallet-tracky/Golang-backend/util"
)

func main() {

	configSetup := startup.InitializeConfig()
	config := configSetup.GetConfig()

	startup.InitializeLogger(config)

	util.InitializeDatabase(config)

	server := gin.Default()
	startup.NewRouter().InitializeRouting(server)

	port := config.GetString("app.port")
	err := server.Run(":" + port)
	if err != nil {
		logger.Fatalf("Failed to run server. %s", err.Error())
	}
}
