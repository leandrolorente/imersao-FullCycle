package main

import (
    "log"
	"github.com/joho/godotenv"
	"github.com/leandrolorente/imersao-FullCycle/infra/kafka"
)

func init(){
	err := godotenv.Load()
	if err != nil{
		log.Fatalf("error loading env file")
	}
	
}

func main() {
	producer := kafka.NewKafkaProducer()
	kafka.Publish("Ola", "readtest", producer)

	/*route := route.Route{
		ID:       "1",
		ClientID: "1",
	}

	route.LoadPositions()

	stringjson, _ := route.ExportJsonPositions()
	fmt.Println(stringjson[0])*/
}
