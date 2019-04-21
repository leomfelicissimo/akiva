package mq

import (
	"log"

	"github.com/streadway/amqp"
)

// ExchangeOptions ...
type ExchangeOptions struct {
	Durable    bool
	AutoDelete bool
	Internal   bool
	NoWait     bool
	Args       amqp.Table
}

func getExchangeOptions(opts *ExchangeOptions) *ExchangeOptions {
	if opts == nil {
		return &ExchangeOptions{
			Durable:    true,
			AutoDelete: false,
			Internal:   false,
			NoWait:     false,
			Args:       nil,
		}
	}

	return opts
}

// DeclareExchange ...
func DeclareExchange(name string, kind string, opts *ExchangeOptions, handler func(ch *amqp.Channel) error) error {
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

	eOpts := getExchangeOptions(opts)
	err = ch.ExchangeDeclare(
		name,
		kind,
		eOpts.Durable,
		eOpts.AutoDelete,
		eOpts.Internal,
		eOpts.NoWait,
		eOpts.Args,
	)

	handler(ch)

	return nil
}

// PublishToExchange ...
func PublishToExchange(name string, kind string, body string, opts *PublishOptions) error {
	pOpts := getPublishOptions("", opts)
	return DeclareExchange(name, kind, nil, func(ch *amqp.Channel) error {
		err := ch.Publish(
			name,
			pOpts.RoutingKey,
			pOpts.Mandatory,
			pOpts.Immediate,
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
