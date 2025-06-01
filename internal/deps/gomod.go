package deps

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
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

// GetModuleName extracts the module name from a go.mod file.
func GetModuleName(goModFilePath string) (string, error) {
	file, err := os.Open(goModFilePath)
	if err != nil {
		return "", fmt.Errorf("failed to open go.mod: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module ")), nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading go.mod: %w", err)
	}

	return "", fmt.Errorf("module name not found in %s", goModFilePath)
}
