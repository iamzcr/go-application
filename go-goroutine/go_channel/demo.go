package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		defer fmt.Println("sub go end")

		fmt.Println("sub go run")

	}()
	time.Sleep(time.Second)
	fmt.Println("main go end")
}

/*
func main() {
	ch := make(chan int)
	go func() {
		defer fmt.Println("sub go end")

		fmt.Println("sub go run")

		ch <- 999
	}()

	num := <-ch
	fmt.Println("num=", num)
	fmt.Println("main go end")
} */
