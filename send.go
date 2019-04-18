package main

import (
	"log"
	"os"

	"github.com/leomfelicissimo/akiva/amqpc"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	body := os.Args[1]
	qname := "comments"
	ch, err := amqpc.DeclareQueue(qname)
	err = ch.Publish(
		"",
		qname,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)

	failOnError(err, "Failed to publish a message")
	log.Println("Message, published!")
}
