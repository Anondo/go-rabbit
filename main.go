package main

import (
	"gorabbit/cmd"
	"gorabbit/helper"
)

func main() {
	if err := cmd.ExecuteRootCommand(); err != nil {
		helper.FailOnError(err, "Failed to execute root command")
	}
}
