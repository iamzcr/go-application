package main

import (
	"fmt"
	"go-kafka/product"
	"go-kafka/tail_logs"
)

func main() {
	err := product.InitProdcuer([]string{"localhost:9092"})
	if err != nil {
		fmt.Println("product error:", err)
		return
	}
	err = tail_logs.InitTail("./test.log")
	if err != nil {
		fmt.Println("tail error:", err)
		return
	}
	tail_logs.GetLogs()
}
