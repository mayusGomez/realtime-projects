package cmd

import (
	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
	"livecomments/dispatcher"
	"os"
	"strings"
)

type DispatcherCommand struct{}

func (cmd *DispatcherCommand) Execute() error {
	figure.NewColorFigure("Dispatcher server", "doom", "red", false).Print()

	queues := strings.Split(os.Getenv("QUEUES"), ",")

	serviceContainer, err := dispatcher.NewServiceContainer(os.Getenv("RABBIT_MQ"), queues)
	if err != nil {
		return err
	}

	err = serviceContainer.Run(os.Getenv("PORT"))
	if err != nil {
		return err
	}

	return nil
}

func newDispatcherCobraCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dispatcher",
		Short: "Starts the dispatcher program",
		RunE: func(cmd *cobra.Command, args []string) error {
			command := &DispatcherCommand{}

			return command.Execute()
		},
	}

	return cmd
}
