package main

import (
	kafkaConfig "go-kafka/config"
	"go-kafka/producer"
)

func main() {
	topic := kafkaConfig.CONST_TOPIC
	producer.Produce(topic, 1000)
}
