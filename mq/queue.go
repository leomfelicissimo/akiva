package mq

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

// QueueOptions ...
type QueueOptions struct {
	Durable    bool
	AutoDelete bool
	Exclusive  bool
	NoWait     bool
	Args       amqp.Table
}

func getQueueOptions(opts *QueueOptions) *QueueOptions {
	if opts == nil {
		return &QueueOptions{
			AutoDelete: false,
			Durable:    false,
			Exclusive:  false,
			NoWait:     false,
			Args:       nil,
		}
	}

	return opts
}

// DeclareQueue ...
func DeclareQueue(queueName string, opts *QueueOptions, handler func(ch *amqp.Channel) error) error {
	log.Println("Connecting to RabbitMQ")
	conn, err := amqp.Dial("amqp://" + getConnectionURI())
	if err != nil {
		return err
	}

	defer conn.Close()
	log.Println("Connected to Rabbit")

	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	qOpts := getQueueOptions(opts)

	q, err := ch.QueueDeclare(
		queueName,
		qOpts.Durable,
		qOpts.AutoDelete,
		qOpts.Exclusive,
		qOpts.NoWait,
		qOpts.Args,
	)
	if err != nil {
		return err
	}

	log.Printf("Queued created: %s", q.Name)

	handler(ch)

	return nil
}

// PublishOptions ...
type PublishOptions struct {
	Exchange   string
	Queue      string
	RoutingKey string
	Mandatory  bool
	Immediate  bool
	Body       string
}

func getPublishOptions(queue string, opts *PublishOptions) *PublishOptions {
	if opts == nil {
		return &PublishOptions{
			Exchange:   "",
			Immediate:  false,
			Mandatory:  false,
			RoutingKey: queue,
		}
	}

	return opts
}

// Publish ...
func Publish(queue string, body string, opts *PublishOptions) error {
	po := getPublishOptions(queue, opts)
	return DeclareQueue(queue, nil, func(ch *amqp.Channel) error {
		err := ch.Publish(
			po.Exchange,
			po.RoutingKey,
			po.Mandatory,
			po.Immediate,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			},
		)
		if err != nil {
			return err
		}

		return nil
	})
}

// ConsumeOptions ...
type ConsumeOptions struct {
	Consumer  string
	AutoAck   bool
	Exclusive bool
	NoLocal   bool
	NoWait    bool
	Args      amqp.Table
}

func consumeOptionsValue(opts *ConsumeOptions) *ConsumeOptions {
	if opts == nil {
		return &ConsumeOptions{
			AutoAck:   false,
			Exclusive: false,
			NoLocal:   false,
			NoWait:    false,
			Args:      nil,
			Consumer:  "",
		}
	}

	return opts
}

// Consume ...
func Consume(qname string, opts *ConsumeOptions) error {
	f := consumeOptionsValue(opts)
	return DeclareQueue(qname, nil, func(ch *amqp.Channel) error {
		msgs, err := ch.Consume(
			qname,
			f.Consumer,
			f.AutoAck,
			f.Exclusive,
			f.NoLocal,
			f.NoWait,
			f.Args,
		)
		if err != nil {
			return err
		}

		forever := make(chan bool)

		go func() {
			for m := range msgs {
				log.Printf("Received a message: %s", m.Body)
			}
		}()

		log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
		<-forever

		return nil
	})
}
