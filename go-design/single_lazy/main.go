// 懒汉模式
package main

import (
	"fmt"
	"sync"
)

type singleton struct {
	data string
}

var instance *singleton

// var lock sync.Mutex
var once sync.Once

//var atomicCount uint32

func GetInstance() *singleton {
	/*
		//加锁处理并发线性安全问题
		lock.Lock()
		defer lock.Unlock()
	*/
	/*
		//原子地读取值
		if atomic.LoadUint32(&atomicCount) == 1 {
			return instance
		}
		if instance == nil {
			instance = new(singleton)
			//原子地设置值
			atomic.StoreUint32(&atomicCount, 1)
		}
	*/
	//利用golang提供类库sync.Once的Do方法控制，其实实现底层也是原子操作。
	once.Do(func() {
		instance = new(singleton)
	})
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
