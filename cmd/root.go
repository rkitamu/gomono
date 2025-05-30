package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type Context struct {
	Args Args
}

type Args struct {
	OutputFilePath string
	Verbose        bool
	Version        bool
}

var ctx Context

var rootCmd = &cobra.Command{
	Use:   "gomono",
	Short: "gomono - Bundle local packages into a single file",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test")
	},
}

func init() {
	ctx = Context{
		Args: Args{},
	}
	// --output, -o
	rootCmd.Flags().StringVarP(&ctx.Args.OutputFilePath, "output", "o", "", "Output file path (default: stdout)")

	// --verbose, -v
	rootCmd.Flags().BoolVarP(&ctx.Args.Verbose, "verbose", "v", false, "Enable verbose output")

	// --version
	rootCmd.Flags().BoolVarP(&ctx.Args.Version, "version", "", false, "Print version information")
}

func Execute() error {
	return rootCmd.Execute()
}
