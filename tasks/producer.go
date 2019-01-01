package tasks

import (
	"log"
	"rabbiting/helper"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

func Produce(cmd *cobra.Command, args []string) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	helper.FailOnError(err, "Failed to connect with rabbitmq server")
	defer conn.Close()

	ch, err := conn.Channel()
	helper.FailOnError(err, "Failed to create channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"rabbit", //name
		false,    //durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)

	helper.FailOnError(err, "Failed to create task queue")

	message := viper.GetString("message")

	err = ch.Publish(
		"",     //exchange
		q.Name, // key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)

	helper.FailOnError(err, "Failed to publish message")

	log.Printf("Successfully published message:%s to the queue:%s\n", message, q.Name)

}
