package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	"github.com/wallet-tracky/Golang-backend/src/startup"
	"github.com/wallet-tracky/Golang-backend/src/util"
)

func main() {

	startup.InitializeLogger()

	logger.Info("Logger Initialized")

	server := gin.Default()

	util.InitializeDatabase()


	router := startup.NewRouter()
	router.InitializeRouting(server)

	err := server.Run(":8080")
	if err != nil {

		fmt.Println(err)
		return
	}
}
