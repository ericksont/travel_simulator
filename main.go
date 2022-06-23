package main

import (
	"fmt"
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	//kafkaApp "github.com/ericksont/travel_simulator/application/kafka"
	//routeApp "github.com/ericksont/travel_simulator/application/route"
	"github.com/ericksont/travel_simulator/infra/kafka"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {

	producer := kafka.NewKafkaProducer()
	log.Println(producer)
	kafka.Publish("ola","readtest", producer)

	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()

	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		//go kafkaApp.Produce(msg)
	}
}
