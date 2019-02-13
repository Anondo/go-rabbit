package workers

import (
	"gorabbit/helper"
	"log"
	"sync"

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

	workerNumber := viper.GetInt("concurrency")
	wg := sync.WaitGroup{}
	wg.Add(workerNumber)

	for i := 0; i < workerNumber; i++ {
		go func() {
			for d := range msgs {
				log.Printf("Received message: %s\n", d.Body)
			}
			wg.Done()
		}()
	}

	log.Printf("%d workers are waiting for tasks, Press CTRL+C to kill\n", workerNumber)

	wg.Wait()

}
