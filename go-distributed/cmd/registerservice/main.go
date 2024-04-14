package main

import (
	"context"
	"fmt"
	"go-stacks/register_store"
	"log"
	"net/http"
)

// #注册服务的服务启动
func main() {
	http.Handle("/services", &register_store.RegistryService{})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var srv http.Server
	srv.Addr = register_store.ServerPort
	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()
	go func() {
		fmt.Println("Registry service started. Press any key to stop.")
		var s string
		fmt.Scanln(&s)
		srv.Shutdown(ctx)
		cancel()
	}()
	<-ctx.Done()
	fmt.Println(" Shutting down registry service ")
}
