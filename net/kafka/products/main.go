package main

import (
	"log"
	"os"
	"github.com/Shopify/sarama"
	"strings"
)

var (
	logger = log.New(os.Stderr, "[srama]", log.LstdFlags)
)

func main() {
	sarama.Logger = logger

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	msg := &sarama.ProducerMessage{}
	msg.Topic = "hello"
	msg.Partition = int32(-1)
	msg.Key = sarama.StringEncoder("key")
	msg.Value = sarama.ByteEncoder("你好, 世界!")

	producer, err := sarama.NewSyncProducer(strings.Split("localhost:9092", ","), config)
	if err != nil {
		logger.Println("Failed to produce message1:", err)
		os.Exit(500)
	}
	defer producer.Close()

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		logger.Println("Failed to produce message2: ", err)
	}
	logger.Printf("partition=%d, offset=%d\n", partition, offset)
}