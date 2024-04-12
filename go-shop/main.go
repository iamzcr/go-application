package main

import (
	"fmt"
	"go-shop/conf"
	"go-shop/db"
	"go-shop/lib"
	"sync"
)

func main() {

	//初始化配置
	err := conf.InitConfig()
	if err != nil {
		fmt.Println("config init:", err)
	}
	//初始化redis
	err = db.InitRedis()
	if err != nil {
		fmt.Println("redis init:", err)
	}

	// 定义并发测试的并发数
	concurrency := 10
	// 使用 WaitGroup 等待所有并发任务完成
	var wg sync.WaitGroup
	wg.Add(concurrency)

	for i := 0; i < concurrency; i++ {
		go func() {
			defer wg.Done()

			// 调用 RedisUniqueId 函数
			uniqueId, err := lib.RedisUniqueId("order_id")
			if err != nil {
				fmt.Println(err)
				return
			}
			// 输出每个并发任务生成的 uniqueId
			fmt.Println("uniqueId:", uniqueId)
		}()
	}
	// 等待所有并发任务完成
	wg.Wait()

}
