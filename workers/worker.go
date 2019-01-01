package workers

import (
	"log"
	"rabbiting/helper"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

func Work(cmd *cobra.Command, args []string) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	helper.FailOnError(err, "Failed to connect to the rabbitmq server")

	ch, err := conn.Channel()
	helper.FailOnError(err, "Failed to create channel")

	q, err := ch.QueueDeclare(
		"rabbit",
		false,
		false,
		false,
		false,
		nil,
	)

	helper.FailOnError(err, "Failed to create task queue")

	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	inf := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received message: %s\n", d.Body)
		}
	}()

	workers := viper.GetInt("workers")

	log.Printf("%d workers are waiting for tasks, Press CTRL+C to kill\n", workers)

	<-inf

}
