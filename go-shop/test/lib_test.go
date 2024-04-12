package test

import (
	"go-shop/lib"
	"sync"
	"testing"
)

func TestRedisUniqueId(t *testing.T) {
	// 定义并发测试的并发数
	concurrency := 100

	// 使用 WaitGroup 等待所有并发任务完成
	var wg sync.WaitGroup
	wg.Add(concurrency)

	for i := 0; i < concurrency; i++ {
		go func() {
			defer wg.Done()

			// 调用 RedisUniqueId 函数
			uniqueId, err := lib.RedisUniqueId("test_key")
			if err != nil {
				t.Errorf("RedisUniqueId error: %v", err)
				return
			}

			// 在测试中进行自定义的验证逻辑
			// ...

			// 输出每个并发任务生成的 uniqueId
			t.Logf("Generated uniqueId: %d", uniqueId)
		}()
	}

	// 等待所有并发任务完成
	wg.Wait()
}
