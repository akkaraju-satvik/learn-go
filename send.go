package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/joho/godotenv"
	mq "github.com/rabbitmq/amqp091-go"
)

func main() {
	godotenv.Load(".env")
	conn, err := mq.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
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
		panic(err)
	}
	var x = map[string]interface{}{
		"message": "Hello, World!",
		"success": map[string]interface{}{
			"code":    int32(200),
			"message": "OK",
		},
	}
	v, err := json.Marshal(x)
	if err != nil {
		panic(err)
	}
	err = ch.PublishWithContext(
		context.Background(),
		"",
		q.Name,
		false,
		false,
		mq.Publishing{
			ContentType: "application/json",
			Body:        v,
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("Message sent")
}
