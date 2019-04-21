package main

import (
	"log"
	"os"
	"strings"

	"github.com/leomfelicissimo/akiva/mq"
	"github.com/streadway/amqp"
)

func bodyFrom(arr []string) string {
	if len(arr) == 3 {
		return arr[2]
	} else if len(arr) > 2 {
		return strings.Join(arr[2:], " ")
	} else {
		return ""
	}
}

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("send [queue_name] [content]")
	}

	qname := os.Args[1]
	body := bodyFrom(os.Args)

	err := mq.DeclareQueue(qname, func(ch *amqp.Channel) error {
		err := ch.Publish(
			"",
			qname,
			false,
			false,
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "text/plain",
				Body:         []byte(body),
			},
		)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Fatalln("Error publishing data", err)
	}

	log.Println("Message, published!")
}
