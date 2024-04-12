// 饿汉模式
package main

import (
	"fmt"
)

/*
保证这个类非公有化,外界不能通过这个类直接创建一个对象
那么这个类就应该变得非公有访问类名称首字母要小写
*/
type singleton struct {
	data string
}

/*
但是还要有一个指针可以指向这个唯一对象，但是这个指针永远不能改变方向
Golang中没有常指针概念，所以只能通过将这个指针私有化不让外部模块访
*/
var instance *singleton = new(singleton)

/*
如果全部都是私有化,那么外部模块将永远无法访问到这个对象.所需要对外提供一个方法来获取到这个对象
GetInstanc不能定义为singelton一个成员方法,因为如果为成员方法就必须要先访问对象、再访问函数，但是类和对象县前都已经私有化..外界无法访问.所以这个方法一定是一个全局普通函数
*/
func GetInstance() *singleton {
	return instance
}

func (s *singleton) Test() {
	fmt.Println("test func")
}
func main() {
	singleton1 := GetInstance()
	singleton2 := GetInstance()
	singleton1.Test()
	singleton2.Test()
	fmt.Println(singleton1 == singleton2) // Output: true
}
