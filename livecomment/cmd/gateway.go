package cmd

import (
	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
	"livecomments/cmd/adapters"
	"livecomments/gateway"
	"log"
)

type GatewayCommand struct{}

func NewWGatewayCommand() (*GatewayCommand, error) {
	return &GatewayCommand{}, nil
}

func (c *GatewayCommand) Execute() error {
	figure.NewColorFigure("Gateway server", "doom", "blue", false).Print()

	serviceContainer := gateway.NewService()

	appAdapter := adapters.NewAppAdapters()
	appAdapter.AddAdapters(
		adapters.NewWebAdapter(Port, serviceContainer),
	)

	appAdapter.Run()

	return nil
}

func newGatewayCobraCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gateway",
		Short: "Starts the gateway program",
		RunE: func(cmd *cobra.Command, args []string) error {
			command, err := NewWGatewayCommand()
			if err != nil {
				log.Println("error trying to start worker command ", err)

				return err
			}

			return command.Execute()
		},
	}

	initGatewayFlags(cmd)

	return cmd
}

func initGatewayFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&Port, "port", "p", "8080", "Port to listen on")
}
