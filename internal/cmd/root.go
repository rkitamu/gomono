package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/rkitamu/gomono/internal/astutil"
)

var (
	rootPath string
	mainPath string
)

// rootCmd represents the base command
var rootCmd = &cobra.Command{
	Use:   "gomono",
	Short: "Flatten modular Go code into a single-file script",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Project Root:", rootPath)
		fmt.Println("Main.go Path:", mainPath)

		imports, err := astutil.ExtractImports(mainPath)
		if err != nil {
			return err
		}
		fmt.Println("Imports:", imports)

		
		return nil
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().StringVar(&rootPath, "root", "", "Path to project root directory (required)")
	rootCmd.Flags().StringVar(&mainPath, "main", "", "Path to main.go file (required)")

	_ = rootCmd.MarkFlagRequired("root")
	_ = rootCmd.MarkFlagRequired("main")
}
