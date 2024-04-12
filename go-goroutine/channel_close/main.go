package main

import "fmt"

func Test(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
}
func main() {
	ch := make(chan int)
	go Test(ch)
	for {
		// ok如果为true表示channel没有关闭，如果为false表示channel已经关闭
		if data, ok := <-ch; ok {
			fmt.Println(data)
		} else {
			break
			fmt.Println("main Finished.. ")
		}
	}
}
