package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type Arguments struct {
	InputFilePath  string
	OutputFilePath string
	Verbose        bool
}

var arguments Arguments

var rootCmd = &cobra.Command{
	Use:   "gomono",
	Short: "gomono - Bundle local packages into a single file",
	Run:   runGomono,
}

func init() {
	rootCmd.Flags().StringVarP(&arguments.InputFilePath, "input", "i", "./main.go", "Target file path (default: ./main.go)")
	rootCmd.Flags().StringVarP(&arguments.OutputFilePath, "output", "o", "", "Output file path (default: stdout)")
	rootCmd.Flags().BoolVarP(&arguments.Verbose, "verbose", "v", false, "Enable verbose output")
}

func Execute() error {
	return rootCmd.Execute()
}

func runGomono(cmd *cobra.Command, args []string) {
	fmt.Println(arguments)
}
