package main

import (
	config "go-kafka/config"
	"go-kafka/consumer"
)

func main() {
	topic := config.CONST_TOPIC
	consumer.Consume(topic)
}
