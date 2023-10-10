package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
    // Set up a Kafka producer configuration
    producer, err := kafka.NewProducer(&kafka.ConfigMap{
        "bootstrap.servers": "localhost:9092",
    })

    if err != nil {
        fmt.Printf("Error creating Kafka producer: %v\n", err)
        return
    }

    defer producer.Close()

    // Produce a message
    topic := "your_topic"
    message := "Hello, Kafka!"

    err = producer.Produce(&kafka.Message{
        TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
        Value:          []byte(message),
    }, nil)

    if err != nil {
        fmt.Printf("Error producing message: %v\n", err)
    } else {
        fmt.Println("Message sent successfully.")
    }

    //producer.Flush(15 * 1000) // Wait for any outstanding messages to be delivered and delivery reports to be received
}