package cmd

import (
	"rabbiting/helper"
	"rabbiting/workers"

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

	WorkerCmd.Flags().IntP("workers", "w", 2, "The number of workers")

	helper.FailOnError(viper.BindPFlags(WorkerCmd.Flags()), "Failed to bind worker flags with viper")

}
