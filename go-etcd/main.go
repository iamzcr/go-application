package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

func main() {
	// 创建 etcd 客户端
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://127.0.0.1:2379"}, // etcd 集群的地址
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// 设置键值对
	_, err = client.Put(context.Background(), "key", "zcr")
	if err != nil {
		log.Fatal(err)
	}

	// 获取值
	resp, err := client.Get(context.Background(), "key")
	if err != nil {
		log.Fatal(err)
	}

	// 处理响应
	for _, kv := range resp.Kvs {
		fmt.Printf("Key: %s, Value: %s\n", kv.Key, kv.Value)
	}
}
