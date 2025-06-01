package deps

import (
	"errors"
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
		return "", errors.New("go.mod not found")
	}
	return findGoModRecursive(parent)
}
