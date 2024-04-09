package consum

import (
	"github.com/IBM/sarama"
	"log"
)

func InitConsumer() {
	// 创建 Kafka 消费者
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		log.Fatal("Failed to create consum: ", err)
	}
	defer consumer.Close()
}
