package deps

import (
	"fmt"
	"os"
	"path/filepath"
)

// FindGoModPath recursively searches for go.mod starting from startDir upward.
func FindGoModPath(startFilePath string) (string, error) {
	dir := filepath.Dir(startFilePath)
	return findGoModRecursive(dir)
}

func findGoModRecursive(dir string) (string, error) {
	goModPath := filepath.Join(dir, "go.mod")
	if _, err := os.Stat(goModPath); err == nil {
		return goModPath, nil
	}
	parent := filepath.Dir(dir)
	if parent == dir {
		return "", fmt.Errorf("go.mod not found; reached filesystem root: %s", dir)
	}
	return findGoModRecursive(parent)
}
