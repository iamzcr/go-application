package product

import (
	"github.com/IBM/sarama"
	"log"
)

var producerObj sarama.SyncProducer

func InitProdcuer(addrs []string) (err error) {
	config := sarama.NewConfig()

	// 设置消息发送确认机制
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 3
	config.Producer.Return.Successes = true

	// 创建生产者对象
	producerObj, err = sarama.NewSyncProducer(addrs, config)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func SendMsgToKafka(topic string, msg string) (err error) {
	// 构造消息
	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(msg),
	}
	// 发送消息到 Kafka
	_, _, err = producerObj.SendMessage(message)
	if err != nil {
		log.Fatal(err)
	}
	return
}
