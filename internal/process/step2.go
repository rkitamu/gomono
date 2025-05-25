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
