package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wallet-tracky/Golang-backend/src/startup"
	"github.com/wallet-tracky/Golang-backend/src/util"
)

func main() {

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
