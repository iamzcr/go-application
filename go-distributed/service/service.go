package service

import (
	"context"
	"fmt"
	"go-stacks/register_store"
	"log"
	"net/http"
)

// 启动服务
func Start(ctx context.Context, host, port string, reg register_store.Registration,
	registerHandlesFunc func()) (context.Context, error) {
	registerHandlesFunc()
	//启动某个服务
	ctx = startService(ctx, reg.ServiceName, host, port)
	//启动某个服务后注册到服务里面
	err := register_store.RegisterService(reg)
	if err != nil {
		return ctx, nil
	}
	return ctx, nil
}

func startService(ctx context.Context, serviceName register_store.ServiceName, host,
	port string) context.Context {
	ctx, cancel := context.WithCancel(ctx)
	var srv http.Server
	srv.Addr = ":" + port
	//启动某个服务
	go func() {
		log.Println(srv.ListenAndServe())
		err := register_store.ShutdownService(fmt.Sprintf("http://%s:%s", host, port))
		if err != nil {
			log.Println(err)
		}
		cancel()
	}()
	//手动停止某个服务
	go func() {
		fmt.Printf("%v start press any key to stop\n", serviceName)
		var s string
		fmt.Scanln(&s)
		err := register_store.ShutdownService(fmt.Sprintf("http://%s:%s", host, port))
		if err != nil {
			log.Println(err)
		}
		srv.Shutdown(ctx)
		cancel()
	}()

	return ctx
}
