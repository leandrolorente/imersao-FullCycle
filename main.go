package main

import (
	"fmt"
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
	"github.com/leandrolorente/imersao-FullCycle/application/kafka"
	kafka2 "github.com/leandrolorente/imersao-FullCycle/infra/kafka"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading env file")
	}

}

func main() {
	msgChan := make(chan *ckafka.Message)
	consumer := kafka2.NewKafkaConsumer(msgChan)
	go consumer.Consume()

	/*producer := kafka2.NewKafkaProducer()
	kafka2.Publish("Ola", "readtest", producer)*/
	for msg := range msgChan {

		fmt.Println(string(msg.Value))
		go kafka.Produce(msg)

	}
	/*route := route.Route{
		ID:       "1",
		ClientID: "1",
	}

	route.LoadPositions()

	stringjson, _ := route.ExportJsonPositions()
	fmt.Println(stringjson[0])*/
}
