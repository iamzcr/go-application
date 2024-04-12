package main

import "fmt"

// 定义结构体，首字母大写表示公有，小写表示私有
type Person struct {
	name string
	age  int
}

func (p Person) PrintInfo() {
	fmt.Printf("Person值是%v,全部信息是%#v,类型是%T,地址是%p\n", p, p, p, &p)
}

func (p *Person) PrintPointInfo() {
	fmt.Printf("p值是%v,全部信息是%#v,类型是%T,地址是%p\n", p, p, p, &p)
}

func (p *Person) SetInfo(name string, age int) {
	p.name = name
	p.age = age
}

func main() {
	////实例化一个对象person
	//var person = Person{
	//	name: "小李",
	//	age:  1,
	//}
	////使用person调用对象的方法
	//person.PrintPointInfo()
	//person.PrintInfo()
	//person.SetInfo("小明", 555)
	//person.PrintPointInfo()
	//person.PrintInfo()
	//
	//实例化一个指针变量对象person
	var person = &Person{
		name: "小张",
		age:  1,
	}
	person.PrintPointInfo()
	person.PrintInfo()
	person.SetInfo("小明", 555)
	person.PrintPointInfo()
	person.PrintInfo()
}
