package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/rkitamu/gomono/internal/codegen"
	"github.com/rkitamu/gomono/internal/deps"
	"github.com/rkitamu/gomono/internal/logutil"
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

	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		logutil.SetupLogger(arguments.Verbose)
	}
}

func Execute() error {
	return rootCmd.Execute()
}

func runGomono(cmd *cobra.Command, args []string) {
	slog.Debug("checking if input file exists", "path", arguments.InputFilePath)
	if _, err := os.Stat(arguments.InputFilePath); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	slog.Debug("searching for go.mod", "from", arguments.InputFilePath)
	goModPath, err := deps.FindGoModPath(arguments.InputFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	slog.Debug("extracting for module name", "path", goModPath)
	moduleName, err := deps.GetModuleName(goModPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	slog.Debug("analyze dependencies candidate", "path", arguments.InputFilePath)
	deps, err := deps.AnalyzeLocalDependencies(arguments.InputFilePath, goModPath, moduleName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	if arguments.OutputFilePath == "" {
		slog.Debug("generate merged code stdout")
		err := codegen.GenerateToStdout(
			deps[0].Files[0].FSet,
			deps[0].Files[0].AST)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
	} else {
		slog.Debug("generate merged code", "path", arguments.OutputFilePath)
		err := codegen.GenerateToFile(
			deps[0].Files[0].FSet,
			deps[0].Files[0].AST,
			arguments.OutputFilePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
	}
}
