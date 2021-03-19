package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "paich",
	Short: "A simple Go code generator based on text/template (EXPERIMENTAL)",
	Long:  `A simple Go code generator based on text/template`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Stop it, get some help by run paich help. Thank you.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
