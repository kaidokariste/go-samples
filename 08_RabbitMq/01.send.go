package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
	"encoding/json"
)

// JSON datastructure to send out
type League struct {
	Teams []Footballteam
}

type Footballteam struct {
	Name      string `json:"name"`
	ShortName string `json:"shortName"`
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	//Define slice of structs conatining two documents
	list := []Footballteam{
		{"Levadia", "LEV"},
		{"Paide", "PAID"},
	}

	//Marshal them to  Parent struct into Teams variable
	jsonString, err := json.Marshal(League{Teams: list})
	if err != nil {
		panic(err)
	}

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
			ContentType: "application/json",
			Body:        jsonString,
		})
	log.Printf(" [x] Sent %s", jsonString)
	failOnError(err, "Failed to publish a message")
}
