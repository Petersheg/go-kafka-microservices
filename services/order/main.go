package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/IBM/sarama"
	"github.com/Petersheg/go-kafka-microservices/pkg/kafka"
)

func main() {

	producer := kafka.NewProducer("kafka:9092")
	defer producer.Close()

	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		orderID := fmt.Sprintf("order-%d", time.Now().Unix())

		message := &sarama.ProducerMessage{
			Topic: "order_created",
			Value: sarama.StringEncoder(orderID),
		}

		partition, offset, err := producer.SendMessage(message)

		if(err != nil) {
			http.Error(w, err.Error(), 500)
		}

		fmt.Println("Order created:", orderID)
		fmt.Fprintf(w, "Order %s created (Partition %d, Offset %d)\n", orderID, partition, offset)
	})

	fmt.Println("Order service running on port 8080")

	http.ListenAndServe(":8080", nil)
}