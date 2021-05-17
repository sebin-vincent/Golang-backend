package startup

import (
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	"github.com/wallet-tracky/Golang-backend/src/router"
)

type Router struct {
	expenseRouter *router.ExpenseRouter
}

func (routing *Router) InitializeRouting(server *gin.Engine) {
	logger.Info("Initializing routing..")

	routing.expenseRouter.InitializeExpenseRouting(server)

	logger.Info("Routing initialized")
}

func NewRouter() *Router {

	return &Router{
		expenseRouter: router.NewExpenseRouter(),
	}
}
