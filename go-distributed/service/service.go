package service

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func Start(ctx context.Context, serviceName, host, port string, registerHandlesFunc func()) (context.Context, error) {
	registerHandlesFunc()
	ctx = startSevvice(ctx, serviceName, host, port)
	return ctx, nil
}

func startSevvice(ctx context.Context, serviceName, host, port string) context.Context {
	ctx, cancel := context.WithCancel(ctx)
	var srv http.Server
	srv.Addr = ":" + port
	go func() {
		log.Panicln(srv.ListenAndServe())
		cancel()
	}()

	go func() {
		fmt.Printf("%v start press any key to stop\n", serviceName)
		var s string
		fmt.Scanln(&s)
		srv.Shutdown(ctx)
		cancel()
	}()

	return ctx
}
