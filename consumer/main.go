package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
    // Set up a Kafka consumer configuration
    consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
        "bootstrap.servers": "localhost:9092",
        "group.id":          "your_consumer_group_id",
        "auto.offset.reset": "earliest",
    })

    if err != nil {
        fmt.Printf("Error creating Kafka consumer: %v\n", err)
        return
    }

    defer consumer.Close()

    // Subscribe to a topic
    topic := "your_topic"
    consumer.SubscribeTopics([]string{topic}, nil)

    for {
        msg, err := consumer.ReadMessage(-1)
        if err == nil {
            fmt.Printf("Received message: %s\n", string(msg.Value))
        } else {
            fmt.Printf("Error while consuming: %v\n", err)
        }
    }
}