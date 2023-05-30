package main

import (
	"fmt"
	"log"
	"time"
	"github.com/Shopify/sarama"
)

func main() {
    // 创建一个Kafka配置实例
    config := sarama.NewConfig()
    // 设置消费者组
    config.Consumer.Group.Session.Timeout = 10 * time.Second
    config.Consumer.Group.Heartbeat.Interval = 3 * time.Second
    // 创建一个Kafka消费者实例
    consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
    if err != nil {
        log.Fatalf("Failed to create consumer: %s", err)
    }
    defer func() {
        if err := consumer.Close(); err != nil {
            log.Fatalf("Failed to close consumer: %s", err)
        }
    }()
    // 创建一个Kafka生产者实例
    producer, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, config)
    if err != nil {
        log.Fatalf("Failed to create producer: %s", err)
    }
    defer func() {
        if err := producer.Close(); err != nil {
            log.Fatalf("Failed to close producer: %s", err)
        }
    }()
    // 发送一条消息到Kafka
    producer.Input() <- &sarama.ProducerMessage{
        Topic: "my-topic",
        Value: sarama.StringEncoder("Hello, Kafka!"),
    }
    // 从Kafka消费消息
    partitionConsumer, err := consumer.ConsumePartition("my-topic", 0, sarama.OffsetOldest)
    if err != nil {
        log.Fatalf("Failed to create partition consumer: %s", err)
    }
    defer func() {
        if err := partitionConsumer.Close(); err != nil {
            log.Fatalf("Failed to close partition consumer: %s", err)
        }
    }()
    for msg := range partitionConsumer.Messages() {
        fmt.Printf("Received message: %s", string(msg.Value))
    }
}