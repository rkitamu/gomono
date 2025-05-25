package process

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/rkitamu/gomono/internal/deps"
	"github.com/rkitamu/gomono/internal/merge"
)

func Step2AnalyzeDependencies(rootPath, mainPath string) error {
	goModPath := filepath.Join(rootPath, "go.mod")
	moduleName, err := deps.ParseGoModModuleName(goModPath)
	if err != nil {
		return err
	}

	files, err := deps.CollectLocalGoFiles(mainPath, rootPath, moduleName)
	if err != nil {
		return err
	}

	fmt.Println("Local Go Files:")
	for f := range files {
		fmt.Println(" -", f)
	}

	// step3: sort dependencies
	sortedFiles, err := deps.TopoSort(files, rootPath, moduleName)
	if err != nil {
		return fmt.Errorf("failed to sort dependencies: %w", err)
	}
	fmt.Println("Sorted Go Files:")
	for _, f := range sortedFiles {
		fmt.Println(" -", f)
	}

	// step4: merge files
	mergedCode, err := MergeFiles(sortedFiles)
	if err != nil {
		return fmt.Errorf("failed to merge files: %w", err)
	}
	fmt.Println("Merged Go Code:")
	fmt.Println(mergedCode)

	// step5: write merged code to output file
	outputPath := filepath.Join(rootPath, "merged_output.go")
	outputSource, err := EmitFinalOutput(sortedFiles, moduleName)
	if err != nil {
		return fmt.Errorf("failed to write merged code to file: %w", err)
	}
	fmt.Println("Merged code written to:", outputPath)
	fmt.Println("output source:\n", outputSource)
	return nil
}

// MergeFiles joins all ordered files into a single Go source code (excluding package/import).
func MergeFiles(ordered []string) (string, error) {
	var builder strings.Builder
	for _, path := range ordered {
		code, err := merge.ExtractCodeWithoutPackageAndImports(path)
		if err != nil {
			return "", err
		}
		builder.WriteString(code)
		builder.WriteString("\n\n")
	}
	return builder.String(), nil
}

// EmitFinalOutput generates the full Go code: package + import + merged declarations
func EmitFinalOutput(ordered []string, moduleName string) (string, error) {
	imports, err := merge.ExtractExternalImports(ordered, moduleName)
	if err != nil {
		return "", err
	}

	code, err := MergeFiles(ordered)
	if err != nil {
		return "", err
	}

	var builder strings.Builder
	builder.WriteString("package main\n\n")

	if len(imports) > 0 {
		builder.WriteString("import (\n")
		for _, imp := range imports {
			builder.WriteString(fmt.Sprintf("\t%q\n", imp))
		}
		builder.WriteString(")\n\n")
	}

	builder.WriteString(code)
	return builder.String(), nil
}
