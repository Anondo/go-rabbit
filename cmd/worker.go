package cmd

import (
	"gorabbit/helper"
	"gorabbit/workers"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	WorkerCmd = &cobra.Command{
		Use:   "worker",
		Short: "starts the worker",
		Run:   workers.Work,
	}
)

func init() {

	WorkerCmd.Flags().IntP("concurrency", "c", 1, "The number of concurrent workers")

	helper.FailOnError(viper.BindPFlags(WorkerCmd.Flags()), "Failed to bind worker flags with viper")

}
