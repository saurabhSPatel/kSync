package main

import (
	"time"

	"github.com/kSync/consumer"
	"github.com/kSync/producer"
)

func main() {
	go producer.InitProducer()
	go consumer.InitConsumer()

	time.Sleep(time.Second * 100)
}
