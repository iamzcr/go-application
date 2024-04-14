package main

import (
	"context"
	"fmt"
	"go-stacks/log"
	"go-stacks/register_store"
	"go-stacks/service"
	stLog "log"
)

func main() {
	log.Run("/distributed.log")
	host, port := "127.0.0.1", "4000"
	serviceAddress := fmt.Sprintf("http://%s:%s", host, port)
	//新建日志服务
	r := register_store.Registration{
		ServiceName: "Log Service",
		ServiceURL:  serviceAddress,
	}
	//启动日志服务，并注册到注册服务
	ctx, err := service.Start(
		context.Background(),
		host,
		port,
		r,
		//log注册函数
		log.RegisterHandlers,
	)
	if err != nil {
		stLog.Fatalln()
	}
	//监听go的cancel
	<-ctx.Done()
	fmt.Println("Shutting dowm Log Service\n")
}
