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
