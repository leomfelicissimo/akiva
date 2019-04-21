package main

import (
	"log"
	"os"

	"github.com/leomfelicissimo/akiva/mq"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("send [queue_name] [content]")
	}

	qname := os.Args[1]
	body := os.Args[2]

	err := mq.Publish(qname, body, nil)
	if err != nil {
		log.Fatalln("Error publishing data", err)
	}

	log.Println("Message, published!")
}
