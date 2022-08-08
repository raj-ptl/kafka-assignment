package main

import (
	"context"
	"fmt"
	"log"

	kafka "github.com/segmentio/kafka-go"
)

func main() {

	ctx := context.Background()

	w := &kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092"),
		Topic:                  "example-123",
		RequiredAcks:           kafka.RequireAll,
		AllowAutoTopicCreation: true,
		Async:                  true,
		Completion: func(messages []kafka.Message, err error) {

			if err != nil {
				fmt.Println("Compl. error : ", err)
				return
			}

			for _, val := range messages {
				fmt.Printf("Messages sent, offset %d, key %s, val %s\n", val.Offset, val.Key, val.Value)
			}
		},
	}

	err := w.WriteMessages(ctx,
		kafka.Message{
			Key:   []byte("dummyKey1"),
			Value: []byte("dummyValue1"),
		},
		kafka.Message{
			Key:   []byte("dummyKey2"),
			Value: []byte("dummyValue2"),
		},
		kafka.Message{
			Key:   []byte("dummyKey3"),
			Value: []byte("dummyValue3"),
		},
	)

	if err != nil {
		log.Fatal("Failed to write messages : ", err)
	}

	if err := w.Close(); err != nil {
		log.Fatal("Failed to close writer: ", err)
	}
}
