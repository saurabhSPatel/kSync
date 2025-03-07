package producer

import (
	"fmt"
	"hash/fnv"
	"log"

	"github.com/IBM/sarama"
)

func hashKey(key string, numPartitions uint32) uint32 {
	h := fnv.New32a()
	h.Write([]byte(key))
	return h.Sum32() % numPartitions
}

func InitProducer() {
	brokers := []string{"localhost:9092"}
	topic := "example-topic"

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatalf("Failed to start producer: %v", err)
	}
	defer producer.Close()

	keys := []string{"user1", "user2", "user3", "user4"}
	for _, key := range keys {

		msg := &sarama.ProducerMessage{
			Topic:     topic,
			Key:       sarama.StringEncoder(key), // Key ensures ordering
			Value:     sarama.StringEncoder("Message for " + key),
			Partition: int32(hashKey(key, 10)), // Optional, but useful for direct partitioning //using 10 partiotion
		}

		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			log.Printf("Failed to send message: %v", err)
		} else {
			fmt.Printf("Sent message to partition %d at offset %d\n", partition, offset)
		}
	}
}
