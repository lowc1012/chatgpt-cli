package commands

import (
	"bytes"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Version is the version info for chatgpt-cli.
type Version struct {
	Major, Minor, Patch int
}

var (
	// CurrentVersion is the current version of chatgpt-cli.
	CurrentVersion = Version{
		Major: 0,
		Minor: 0,
		Patch: 1,
	}

	// rootCmd is the entry point for chatgpt-cli
	rootCmd = &cobra.Command{
		Use:   "chatgpt-cli",
		Short: "chatgpt-cli - a simple CLI to use ChatGPT",
		Long:  "chatgpt-cli is a command-line app using Cobra and OpenAI Chat API to use ChatGPT in terminal",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
)

func (v Version) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch))
	return buffer.String()
}

// Execute executes the current command using rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
