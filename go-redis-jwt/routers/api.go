package routers

import (
	"github.com/gin-gonic/gin"
	"jwt/service/api"
)

func ApiRoutersInit(r *gin.Engine) {
	ApiRouter := r.Group("/api") //可以加载这个后面
	{
		ApiRouter.GET("/login", api.Login)
		ApiRouter.GET("/getCode", api.GetCode)
	}
}
