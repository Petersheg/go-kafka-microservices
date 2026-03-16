package kafka

import (
	"log"
	"time"

	"github.com/IBM/sarama"
)

func NewProducer(broker string) sarama.SyncProducer {

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	var producer sarama.SyncProducer
	var err error

	for i := 0; i < 10; i++ {

		producer, err = sarama.NewSyncProducer([]string{broker}, config)

		if err == nil {
			log.Println("Connected to Kafka")
			return producer
		}

		log.Println("Kafka not ready, retrying in 3 seconds...")
		time.Sleep(3 * time.Second)
	}

	log.Fatal("Could not connect to Kafka:", err)

	return nil
}

func NewConsumer(broker string) sarama.Consumer {

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	var consumer sarama.Consumer
	var err error

	for i := 0; i < 10; i++ {

		consumer, err = sarama.NewConsumer([]string{broker}, config)

		if err == nil {
			log.Println("Connected to Kafka")
			return consumer
		}

		log.Println("Kafka not ready, retrying in 3 seconds...")
		time.Sleep(3 * time.Second)
	}

	log.Fatal("Failed to connect to Kafka:", err)

	return nil
}