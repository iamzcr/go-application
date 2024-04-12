package main

import "fmt"

// 自定义类型
type myInt int

// 自定义方法类型
type myFunc func(int, int) int

// 1.9加入类型别名
type myFloat = float64

// 定义结构体，首字母大写表示公有，小写表示私有
type Person struct {
	name string
	age  int
}

type Test struct {
	string
	Temp
	int
}
type Temp struct {
	temp int
}

func main() {
	var test = &Test{}
	fmt.Printf("test值是%v,全部信息是%#v,类型是%T,地址是%p\n", test, test, test, &test)

	//键值对实例化
	//var person = Person{
	//	name: "test3",
	//	age:  88,
	//}
	//fmt.Printf("person值是%v,全部信息是%#v,类型是%T,地址是%p\n", person, person, person, &person)

	//new实例化，但是返回的时指针类型，实际上该方式出来的实例，理论上需要通过结构体指针访问成员变量的，(*person1).name，
	// 而实际上person.name这样访问，底层也是(*person).name,而不需要显式声明
	//var person = new(Person)
	//person.name = "test1"
	//(*person).age = 2
	//fmt.Printf("person值是%v,全部信息是%#v,类型是%T,地址是%p\n", person, person, person, &person)

	//引用方式实例化
	//var person = &Person{
	//	name: "test3",
	//	age:  88,
	//}
	//fmt.Printf("person值是%v,全部信息是%#v,类型是%T,地址是%p\n", person, person, person, &person)

	//不指定实例化，但是顺序要和定义的类型里面的属性顺序一样
	var person = &Person{
		"test3",
		1,
	}
	fmt.Printf("person值是%v,全部信息是%#v,类型是%T,地址是%p\n", person, person, person, &person)
}
