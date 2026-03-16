package main

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
	"github.com/Petersheg/go-kafka-microservices/pkg/kafka"
)

func main() {

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Consumer.Return.Errors = true

	brokers := "kafka:9092"

	consumer := kafka.NewConsumer(brokers)
	producer := kafka.NewProducer(brokers)

	defer consumer.Close()
	defer producer.Close()

	partitionConsumer, err := consumer.ConsumePartition(
		"order_created",
		0,
		sarama.OffsetNewest,
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Payment Service listening for orders...")

	for msg := range partitionConsumer.Messages() {

		orderID := string(msg.Value)

		fmt.Println("Received order:", orderID)
		fmt.Println("Processing payment for:", orderID)

		paymentMsg := &sarama.ProducerMessage{
			Topic: "payment_completed",
			Value: sarama.StringEncoder(orderID),
		}

		producer.SendMessage(paymentMsg)
		
		fmt.Println("Payment completed event published for:", orderID)
	}
}