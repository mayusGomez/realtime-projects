package cmd

import (
	"bufio"
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
	"livecomments/commentgenerator"
	"log"
	"os"
)

type GeneratorCommand struct{}

func (c *GeneratorCommand) Execute() error {
	figure.NewColorFigure("Generator server", "doom", "green", false).Print()

	video := os.Getenv("VIDEO")
	if video == "" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter the videoId: ")
		video, _ = reader.ReadString('\n')
	}
	dispatcherURL := os.Getenv("DISPATCHER_URL")
	log.Printf("generate requests to URL: %s, video: %s", dispatcherURL, video)

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
