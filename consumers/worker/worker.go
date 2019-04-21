package main

import (
	"bytes"
	"log"
	"os"
	"time"

	"github.com/leomfelicissimo/akiva/mq"
	"github.com/streadway/amqp"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("receive [queue_name]")
	}

	qname := os.Args[1]

	err := mq.DeclareQueue(qname, func(ch *amqp.Channel) error {
		msgs, err := ch.Consume(
			qname,
			"",
			true,
			false,
			false,
			false,
			nil,
		)

		if err != nil {
			log.Fatalf("%s", err)
		}

		forever := make(chan bool)

		go func() {
			for m := range msgs {
				log.Printf("Received a message: %s", m.Body)
				dotCount := bytes.Count(m.Body, []byte("."))
				t := time.Duration(dotCount)
				time.Sleep(t * time.Second)
				log.Printf("Done")
			}
		}()

		log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
		<-forever
		return nil
	})

	if err != nil {
		log.Fatalf("%s", err)
	}
}
