package kafka

import (
	"encoding/json"
	"log"
	"os"
	"time"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	route "github.com/jrcamelo/fullcycle-simulator/application/route"
	kafkaManager "github.com/jrcamelo/fullcycle-simulator/infra/kafka-manager"
)

// {"clientId":"1","routeId":"1"}
// {"clientId":"2","routeId":"3"}
// {"clientId":"3","routeId":"3"}
func Produce(msg *ckafka.Message) {
	producer := kafkaManager.NewKafkaProducer()
	route := route.NewRoute()
	json.Unmarshal(msg.Value, &route)
	route.LoadPositions()
	positions, err := route.ExportJsonPositions()
	if err != nil {
		log.Println(err.Error())
	}
	for _, p := range positions {
		kafkaManager.Publish(p, os.Getenv("KafkaProduceTopic"), producer)
		time.Sleep(time.Millisecond * 500)
	}
}
