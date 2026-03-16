package main

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
	"github.com/Petersheg/go-kafka-microservices/pkg/kafka"

)

func main() {

	consumer := kafka.NewConsumer("kafka:9092")
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition(
		"payment_completed",
		0,
		sarama.OffsetNewest,
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Notification Service listening for completed payments...")

	for msg := range partitionConsumer.Messages() {

		orderID := string(msg.Value)

		fmt.Println("Sending notification for order:", orderID)

	}
}