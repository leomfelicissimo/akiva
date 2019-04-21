package main

import (
	"log"
	"os"

	"github.com/leomfelicissimo/akiva/mq"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("receive [queue_name]")
	}

	qname := os.Args[1]

	mq.Consume(qname, nil)
}
