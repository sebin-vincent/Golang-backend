package middlewares

import (
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	"github.com/wallet-tracky/Golang-backend/util"
)

func Authenticate(privilege string) func(ctx *gin.Context){
	return func (context *gin.Context) {
		token := context.GetHeader("token")


		claims, err := util.VerifyToken(token)
		if err!=nil{
			logger.Error("Error while verifying verifying token. ",err.Error())
			context.AbortWithStatusJSON(401,"UnAuthorized")
			return
		}

		//userPrevileges=db.FetchUSerprevileges
		//if userPrivileges.notcontain(previlege) return error

		logger.Info("Checking user privilege for ",privilege,"...")

		context.Set("userId",claims.UserId)
		context.Next()
	}

}
