package main

import (
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	startup2 "github.com/wallet-tracky/Golang-backend/startup"
	util2 "github.com/wallet-tracky/Golang-backend/util"
)

func main() {

	configSetup := startup2.InitializeConfig()
	config := configSetup.GetConfig()

	startup2.InitializeLogger(config)

	util2.InitializeDatabase(config)

	server := gin.Default()
	startup2.NewRouter().InitializeRouting(server)

	port := config.GetString("app.port")
	err := server.Run(":" + port)
	if err != nil {
		logger.Fatalf("Failed to run server. %s",err.Error())
	}
}
