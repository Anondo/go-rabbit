package cmd

import (
	"rabbiting/helper"
	"rabbiting/tasks"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	PublishCmd = &cobra.Command{
		Use:   "publish",
		Short: "publishes a message to the task queue",
		Run:   tasks.Produce,
	}
)

func init() {
	PublishCmd.Flags().StringP("message", "m", "Hello world", "The message to send to the queue")

	helper.FailOnError(viper.BindPFlags(PublishCmd.Flags()), "Failed to bind publish flags with viper")
}
