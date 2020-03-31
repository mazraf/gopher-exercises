package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd is the task command, it the parent cmd.
// It has no run method, and all the other cmd are added into it.
var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is a simple to do list CLI tool.",
}
