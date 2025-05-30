package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gomono",
	Short: "gomono - Bundle local packages into a single file",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test")
	},
}

func Execute() error {
	return rootCmd.Execute()
}
