package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "live-comments-service",
		Short: "Live Comments Service",
	}

	Port string
)

func init() {
	startCmd := newGatewayCobraCommand()
	rootCmd.AddCommand(startCmd)
}

// Execute executes the root command of the application
func Execute() error {
	return rootCmd.Execute()
}
