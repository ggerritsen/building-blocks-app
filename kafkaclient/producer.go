package kafkaclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

type producer struct {
	w *kafka.Writer
}

// NewProducer creates a new producer for a kafka topic
func NewProducer(brokers []string, topic string) *producer {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: brokers,
		Topic:   topic,
	})

	return &producer{w}
}

// Close closes the underlying kafka connection
func (p *producer) Close() {
	p.w.Close()
}

// Produce sends an object to the kafka topic
func (p *producer) Produce(r *Record) error {
	b := &bytes.Buffer{}
	if err := json.NewEncoder(b).Encode(r); err != nil {
		return err
	}

	return p.w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(fmt.Sprintf("%d", time.Now().Unix())),
			Value: b.Bytes(),
		},
	)
}
