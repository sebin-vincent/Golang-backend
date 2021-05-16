package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wallet-tracky/Golang-backend/src/startup"
	"github.com/wallet-tracky/Golang-backend/src/util"
)

func main() {

	configSetup := startup.InitializeConfig()
	config := configSetup.GetConfig()

	util.InitializeDatabase(config)

	server := gin.Default()
	startup.NewRouter().InitializeRouting(server)

	port := config.GetString("app.port")
	err := server.Run(":" + port)
	if err != nil {
		fmt.Println(err)
		return
	}
}
