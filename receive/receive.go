package main

import (
	"log"

	"github.com/leomfelicissimo/akiva/amqpc"
)

func main() {
	qname := "comments"
	ch, err := amqpc.DeclareQueue(qname)

	msgs, err := ch.Consume(
		qname,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Reveive a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	<-forever
}
