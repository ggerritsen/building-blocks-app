package kafkaclient

import (
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Printf("Start of demonstration.\n")

	brokers := []string{"localhost:9092"}
	topic := "test"

	p := NewProducer(brokers, topic)
	defer p.Close()

	c := NewConsumer(brokers, topic)
	defer c.Close()

	go c.Consume()

	for i := 0; i < 10; i++ {
		if err := p.Produce(&Record{Name: fmt.Sprintf("record-%d", i)}); err != nil {
			log.Fatal(err)
		}
	}

	time.Sleep(15 * time.Second)

	fmt.Printf("End of demonstration.\n")
}
