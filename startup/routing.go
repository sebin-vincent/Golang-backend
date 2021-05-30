package startup

import (
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	"github.com/wallet-tracky/Golang-backend/router"
)

type Router struct {
	expenseRouter *router.ExpenseRouter
	userRouter *router.UserRouter
}

func (routing *Router) InitializeRouting(server *gin.Engine) {
	logger.Info("Initializing routing..")

	routing.expenseRouter.InitializeExpenseRouting(server)
	routing.userRouter.InitializeUserRouter(server)

	logger.Info("Routing initialized")
}

func NewRouter() *Router {

	return &Router{
		expenseRouter: router.NewExpenseRouter(),
		userRouter: router.NewUserRouter(),
	}
}
