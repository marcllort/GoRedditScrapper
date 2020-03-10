package rabbitmq

import (
	"github.com/streadway/amqp"
)

func connect() {
	//https://www.rabbitmq.com/tutorials/tutorial-one-go.html
	q, ch := connectChannel()

	body := "Hello World!"
	err := ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
}
