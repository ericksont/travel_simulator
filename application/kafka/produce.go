package kafka

import (
	"encoding/json"
	"log"
	"os"
	"time"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	routeApp "github.com/ericksont/travel_simulator/application/route"
	"github.com/ericksont/travel_simulator/infra/kafka"
)

//{"clientId":"1","routeId":"1"} (1, 2 and 3 - id files) 
func Produce(msg *ckafka.Message) {
	producer := kafka.NewKafkaProducer()
	route := routeApp.NewRoute()
	json.Unmarshal(msg.Value, &route)
	route.LoadPositions()
	positions, err := route.ExportJsonPositions()
	if err != nil {
		log.Println(err.Error())
	}
	for _, p := range positions {
		kafka.Publish(p, os.Getenv("KafkaProduceTopic"), producer)
		time.Sleep(time.Millisecond * 500)
	}
}
