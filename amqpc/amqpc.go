package amqpc

import (
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func getConnectionURI() string {
	return "guest:guest@localhost:5672"
}

func DeclareQueue(queueName string) (*amqp.Channel, error) {
	log.Println("Connecting to RabbitMQ")
	conn, err := amqp.Dial("amqp://" + getConnectionURI())
	if err != nil {
		return nil, err
	}

	defer conn.Close()
	log.Println("Connected to Rabbit")

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	log.Printf("Queued created: %s", q.Name)

	return ch, nil
}
