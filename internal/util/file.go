package util

import (
	"fmt"
	"os"
	"path/filepath"
)

func ListFileFromDir(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return nil, err
	}

	filePaths := make([]string, 0)
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		filePaths = append(filePaths, filepath.Join(dir, entry.Name()))
	}

	return filePaths, nil
}

// NormalizePath converts the given path to an absolute and cleaned path.
// This ensures consistent comparison across relative/absolute paths.
func NormalizePath(path string) (string, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	return filepath.Clean(abs), nil
}
