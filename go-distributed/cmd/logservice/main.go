package main

import (
	"context"
	"fmt"
	"go-stacks/log"
	"go-stacks/service"
	stLog "log"
)

func main() {
	log.Run("/distributed.log")
	host, port := "localhost", "4000"
	ctx, err := service.Start(
		context.Background(),
		"Log Service",
		host,
		port,
		log.RegisterHandlers,
	)
	if err != nil {
		stLog.Fatalln()
	}
	<-ctx.Done()
	fmt.Println("Shutting dowm Log Service\n")
}
