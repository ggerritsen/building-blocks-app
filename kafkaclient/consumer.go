package kafkaclient

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

// Record is the object stored in the kafka topic
type Record struct {
	Name string `json:"name"`
}

type consumer struct {
	r *kafka.Reader
}

// NewConsumer creates a consumer that consumes from a kafka topic
func NewConsumer(brokers []string, topic string) *consumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   topic,
	})
	r.SetOffset(0)

	return &consumer{r}
}

// Close closes the underlying kafka connection
func (c *consumer) Close() {
	c.r.Close()
}

// Consume consumes from the topic until an error is encountered
func (c *consumer) Consume() error {
	for {
		m, err := c.r.ReadMessage(context.Background())
		if err != nil {
			return err
		}
		fmt.Printf("Received message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}
}
