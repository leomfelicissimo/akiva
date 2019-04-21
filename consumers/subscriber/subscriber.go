package main

import (
	"log"
	"os"

	"github.com/leomfelicissimo/akiva/mq"
	"github.com/streadway/amqp"
)

func failOnError(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %s", message, err)
	}
}

func main() {
	if len(os.Args) < 3 {
		log.Fatalln("subscriber [name] [kind] [body]")
	}

	eName := os.Args[1]
	eKind := os.Args[2]

	err := mq.DeclareExchange(
		eName,
		eKind,
		&mq.ExchangeOptions{
			Durable:    true,
			AutoDelete: false,
			Internal:   false,
			NoWait:     false,
			Args:       nil,
		},
		func(ch *amqp.Channel) error {
			q, err := ch.QueueDeclare(
				"",
				false,
				false,
				true,
				false,
				nil,
			)
			failOnError(err, "Failed to declare a queue")

			err = ch.QueueBind(
				q.Name,
				"",
				eName,
				false,
				nil,
			)
			failOnError(err, "Failed to bind a queue")

			msgs, err := ch.Consume(
				q.Name,
				"",
				true,
				false,
				false,
				false,
				nil,
			)
			failOnError(err, "Failed to register a consumer")

			forever := make(chan bool)

			go func() {
				for m := range msgs {
					log.Printf("[x] %s", m.Body)
				}
			}()

			log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
			<-forever

			return nil
		},
	)

	failOnError(err, "Failed to declare exchange")
}
