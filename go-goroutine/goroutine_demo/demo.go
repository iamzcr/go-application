package go_channel

import (
	"fmt"
	"time"
)

func TestFunc() {
	for i := 0; i < 3; i++ {
		fmt.Println("Test() hello", i)
	}
}
func main() {
	go TestFunc() //开启一个协程
	for i := 0; i < 3; i++ {
		fmt.Println("main() hello", i)
		time.Sleep(time.Millisecond * 20)
	}
}
