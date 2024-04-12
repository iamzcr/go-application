package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "success",
	})
}
func GetCode(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"msg":     "detail",
	})
}
