package main

import (
	"github.com/gin-gonic/gin"
	"jwt/routers"
)

func main() {
	r := gin.Default()
	//注册路由
	routers.ApiRoutersInit(r)
	//启动web服务
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
