package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {

	topic := "example-123"

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		GroupID: "group-2",
		Topic:   topic,
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("Failed to read message : ", err)
			break
		}

		fmt.Printf("group-2 consumer : Message at offset %d : %s = %s\n", m.Offset, string(m.Key), string(m.Value))

	}

	if err := r.Close(); err != nil {
		log.Fatal("Failed to close reader : ", err)
	}

}
