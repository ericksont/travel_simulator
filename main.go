package main

import (
	"fmt"
	//"log"

	//ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	//kafkaApp "github.com/ericksont/travel_simulator/application/kafka"
	routeApp "github.com/ericksont/travel_simulator/application/route"
	//"github.com/ericksont/travel_simulator/infra/kafka"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {

	route := routeApp.Route{
		ID: "1",
		ClientID: "1",
	}
	route.LoadPositions()
	stringjson, _ := route.ExportJsonPositions()

	fmt.Println(stringjson[0])

	/*msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()

	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		go kafkaApp.Produce(msg)
	}*/
}
