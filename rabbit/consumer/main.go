package main

import (
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	checkErrors(err)

	defer conn.Close()

	ch, err := conn.Channel()
	checkErrors(err)

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	checkErrors(err)

	del, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	checkErrors(err)

	for d := range del {
		logrus.Printf("Received a message: %s", d.Body)
		//d.Ack(true)
	}
}

func checkErrors(err error) {
	if err != nil {
		logrus.Fatal(err)
	}
}
