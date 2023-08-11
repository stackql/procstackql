// cmd/root.go

package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "procstack",
	Short: "Procedural language parser and executor",
}

// Execute initializes and runs the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(execCmd)
	rootCmd.AddCommand(shellCmd)
}
