package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "live-comments-service",
	Short: "Live Comments Service",
}

func init() {
	gatewayCmd := newGatewayCobraCommand()
	dispatcherCmd := newDispatcherCobraCommand()
	generator := newGeneratorCommand()

	rootCmd.AddCommand(gatewayCmd)
	rootCmd.AddCommand(dispatcherCmd)
	rootCmd.AddCommand(generator)
}

// Execute executes the root command of the application
func Execute() error {
	return rootCmd.Execute()
}
