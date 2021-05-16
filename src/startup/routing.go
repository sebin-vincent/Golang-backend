package startup

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wallet-tracky/Golang-backend/src/router"
)

type Router struct {
	expenseRouter *router.ExpenseRouter
}

func (routing *Router) InitializeRouting(server *gin.Engine) {
	fmt.Println("Initializing routing..")

	routing.expenseRouter.InitializeExpenseRouting(server)

	fmt.Println("Routing initialized")
}

func NewRouter() *Router {

	return &Router{
		expenseRouter: router.NewExpenseRouter(),
	}
}
