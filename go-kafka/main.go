package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/IBM/sarama"
)

func main() {
	// 创建 Kafka 生产者
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, nil)
	if err != nil {
		log.Fatal("Failed to create producer: ", err)
	}
	defer producer.Close()

	// 创建 Kafka 消费者
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		log.Fatal("Failed to create consumer: ", err)
	}
	defer consumer.Close()

	// 创建一个新的 topic
	topic := "my_topic"

	// 向 Kafka 发送消息
	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder("Hello, Kafka!"),
	}
	partition, offset, err := producer.SendMessage(message)
	if err != nil {
		log.Fatal("Failed to send message: ", err)
	}
	fmt.Printf("Message sent to partition %d at offset %d\n", partition, offset)

	// 从 Kafka 接收消息
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatal("Failed to create partition consumer: ", err)
	}
	defer partitionConsumer.Close()

	// 使用信号通知来优雅地退出
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			select {
			case msg := <-partitionConsumer.Messages():
				fmt.Printf("Received message: %s\n", string(msg.Value))
			case <-signals:
				return
			}
		}
	}()

	wg.Wait()
}
