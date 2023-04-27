package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "ChatGpt CLI",
		Short: "chatgpt-cli - a simple CLI to use ChatGPT",
		Long:  "chatgpt-cli is a command-line app using Cobra and OpenAI Chat API to use ChatGPT in terminal",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
