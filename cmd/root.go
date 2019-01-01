package cmd

import "github.com/spf13/cobra"

var (
	RootCmd = &cobra.Command{
		Use:   "rabbit",
		Short: "rabbit is cli app to demonstrate rabbitmq",
	}
)

func ExecuteRootCommand() error {
	err := RootCmd.Execute()
	return err
}

func init() {
	RootCmd.AddCommand(PublishCmd)
	RootCmd.AddCommand(WorkerCmd)
}
