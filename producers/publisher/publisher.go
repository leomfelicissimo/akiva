package main

import (
	"log"
	"os"
	"strings"

	"github.com/leomfelicissimo/akiva/mq"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalln("publisher [name] [kind] [body]")
	}

	eName := os.Args[1]
	eKind := os.Args[2]
	var body string
	if len(os.Args) > 4 {
		body = strings.Join(os.Args[3:], " ")
	} else {
		body = os.Args[3]
	}

	err := mq.PublishToExchange(eName, eKind, body, nil)
	if err != nil {
		log.Fatalf("%s", err)
	}
}
