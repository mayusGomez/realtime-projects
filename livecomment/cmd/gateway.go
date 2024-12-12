package cmd

import (
	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
	"livecomments/gateway"
	"os"
)

type GatewayCommand struct{}

func (c *GatewayCommand) Execute() error {
	figure.NewColorFigure("Gateway server", "doom", "blue", false).Print()

	serviceContainer := gateway.NewService(os.Getenv("DISPATCHER_URL"), os.Getenv("QUEUE"))
	err := serviceContainer.Run(os.Getenv("PORT"), os.Getenv("RABBIT_MQ"), os.Getenv("QUEUE"))
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

	return cmd
}
