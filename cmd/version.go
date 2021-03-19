package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Paich",
	Long:  `This is your Paich's version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Paich's Go code Generator v0.1")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
