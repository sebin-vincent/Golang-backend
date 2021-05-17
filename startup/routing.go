package startup

import (
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	router2 "github.com/wallet-tracky/Golang-backend/router"
)

type Router struct {
	expenseRouter *router2.ExpenseRouter
}

func (routing *Router) InitializeRouting(server *gin.Engine) {
	logger.Info("Initializing routing..")

	routing.expenseRouter.InitializeExpenseRouting(server)

	logger.Info("Routing initialized")
}

func NewRouter() *Router {

	return &Router{
		expenseRouter: router2.NewExpenseRouter(),
	}
}
