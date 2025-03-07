package consumer

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
)

type ConsumerGroupHandler struct{}

func (ConsumerGroupHandler) Setup(sarama.ConsumerGroupSession) error   { return nil }
func (ConsumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error { return nil }
func (ConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("Consumed message: key=%s, value=%s, partition=%d, offset=%d\n",
			string(msg.Key), string(msg.Value), msg.Partition, msg.Offset)
		sess.MarkMessage(msg, "")
	}
	return nil
}

func InitConsumer() {
	brokers := []string{"0.0.0.0:9092"}
	groupID := "example-group"
	topics := []string{"example-topic"}

	config := sarama.NewConfig()
	config.Version = sarama.MaxVersion
	config.Consumer.Group.Rebalance.Strategy = sarama.NewBalanceStrategyRoundRobin()

	consumerGroup, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		log.Fatalf("Error creating consumer group: %v", err)
	}
	defer consumerGroup.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		for {
			if err := consumerGroup.Consume(ctx, topics, ConsumerGroupHandler{}); err != nil {
				log.Fatalf("Error consuming messages: %v", err)
			}
		}
	}()

	// Handle shutdown
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt, syscall.SIGTERM)
	<-sigchan
	fmt.Println("Shutting down consumer...")
}
