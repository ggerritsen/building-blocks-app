package kafkaclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/ggerritsen/building-blocks-app/model"
	"github.com/segmentio/kafka-go"
)

type consumer struct {
	r *kafka.Reader
	s docStore
}

type docStore interface {
	Store(name string) (*model.Document, error)
}

// NewConsumer creates a consumer that consumes from a kafka topic
func NewConsumer(brokers []string, topic string, s docStore) *consumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   topic,
	})
	r.SetOffset(0)

	return &consumer{r, s}
}

// Close closes the underlying kafka connection
func (c *consumer) Close() {
	c.r.Close()
}

// Consume consumes from the topic and sends to docStore
// It returns (i.e. stops consuming) when an error is encountered
func (c *consumer) Consume() error {
	for {
		m, err := c.r.ReadMessage(context.Background())
		if err != nil {
			return err
		}

		d := &model.Document{}
		if err := json.NewDecoder(bytes.NewReader(m.Value)).Decode(d); err != nil {
			return fmt.Errorf("could not deserialize %q: %s", string(m.Value), err)
		}

		_, err = c.s.Store(d.Name)
		if err != nil {
			return err
		}
	}
}
