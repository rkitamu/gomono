package process

import (
	"fmt"
	"path/filepath"

	"github.com/rkitamu/gomono/internal/deps"
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

	return nil
}
