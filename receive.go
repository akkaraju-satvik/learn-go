package main

import (
	"encoding/json"
	"log"

	mq "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := mq.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Error dialing AMQP server: %v", err)
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Error opening channel: %v", err)
	}
	defer ch.Close()
	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Error declaring queue: %v", err)
	}
	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Error consuming queue: %v", err)
	}
	for msg := range msgs {
		var m map[string]interface{}
		err := json.Unmarshal(msg.Body, &m)
		if err != nil {
			log.Fatalf("Error unmarshalling JSON: %v", err)
		}
		log.Printf("Received message: %s", m["success"].(map[string]interface{})["message"])
	}
}
