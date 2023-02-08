package kafka

import (
	"encoding/json"
	"log"
	"os"
	"time"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/leandrolorente/imersao-FullCycle/application/route"
	"github.com/leandrolorente/imersao-FullCycle/infra/kafka"
)

// {"clienteId":"1", "routeId":"1"}
func Produce(msg *ckafka.Message) {
	producer := kafka.NewKafkaProducer()
	route := route.NewRoute()
	json.Unmarshal(msg.Value, &route)
	route.LoadPositions()
	posititions, err := route.ExportJsonPositions()
	if err != nil {
		log.Println(err.Error())
	}

	for _, p := range posititions {
		kafka.Publish(p, os.Getenv("KafkaProduceTopic"), producer)
		time.Sleep(time.Millisecond * 500)
	}

}
