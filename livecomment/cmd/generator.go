package cmd

import (
	"errors"
	"github.com/common-nighthawk/go-figure"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"livecomments/commentgenerator"
	"log"
	"os"
	"strings"
)

type GeneratorCommand struct{}

func (c *GeneratorCommand) Execute() error {
	figure.NewColorFigure("Generator server", "doom", "green", false).Print()

	video := uuid.New().String()

	dispatcherURL := os.Getenv("DISPATCHER_URL")
	if dispatcherURL == "" {
		return errors.New("no dispatcher URL found in env variable DISPATCHER_URL")
	}

	gatewayURL := os.Getenv("GATEWAY_URL")
	if gatewayURL == "" {
		return errors.New("no gateway URL found in env variable GATEWAY_URL")
	}

	log.Printf("generate requests to URL: %s, video: %s\n\n", dispatcherURL, strings.TrimSpace(video))
	log.Printf("please, open this URL in your browser: %s%s", gatewayURL, video)

	commentgenerator.CommentGenerator(dispatcherURL, video)

	return nil
}

func newGeneratorCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generator",
		Short: "Starts the generator program",
		RunE: func(cmd *cobra.Command, args []string) error {
			command := &GeneratorCommand{}

			return command.Execute()
		},
	}

	return cmd
}
