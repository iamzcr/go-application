package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// 创建可达对象
	person := &Person{Name: "Alice", Age: 25}
	// 使用person变量
	fmt.Println(person.Name, person.Age)
	// 创建不可达对象
	_ = &Person{Name: "Bob", Age: 30}
}
