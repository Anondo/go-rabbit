package cmd

import (
	"gorabbit/helper"

	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
		Use:   "rabbit",
		Short: "rabbit is cli app to demonstrate rabbitmq",
	}
)

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		helper.FailOnError(err, "Could not execute root command")
	}
}

func init() {
	RootCmd.AddCommand(PublishCmd)
	RootCmd.AddCommand(WorkerCmd)
}
