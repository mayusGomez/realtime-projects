package cmd

import (
	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
	"livecomments/gateway"
)

var Port string

type GatewayCommand struct{}

func (c *GatewayCommand) Execute() error {
	figure.NewColorFigure("Gateway server", "doom", "blue", false).Print()

	serviceContainer := gateway.NewService()
	err := serviceContainer.Run(Port)
	if err != nil {
		return err
	}

	return nil
}

func newGatewayCobraCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gateway",
		Short: "Starts the gateway program",
		RunE: func(cmd *cobra.Command, args []string) error {
			command := &GatewayCommand{}

			return command.Execute()
		},
	}

	initGatewayFlags(cmd)

	return cmd
}

func initGatewayFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&Port, "port", "p", "8080", "Port to listen on")
}
