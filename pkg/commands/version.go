package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the current version",
	Long:  "The `chatgpt-cli version` command displays the version of the chatgpt-cli  software.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("chatgpt-cli version: ", CurrentVersion.String())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
