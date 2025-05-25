package deps

import (
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

func ParseGoModModuleName(goModPath string) (string, error) {
	data, err := os.ReadFile(goModPath)
	if err != nil {
		return "", err
	}
	for _, line := range strings.Split(string(data), "\n") {
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module ")), nil
		}
	}
	return "", nil
}

func CollectLocalGoFiles(mainPath string, rootPath string, moduleName string) (map[string]struct{}, error) {
	visited := map[string]struct{}{}
	var walk func(string) error
	walk = func(path string) error {
		if _, ok := visited[path]; ok {
			return nil
		}
		visited[path] = struct{}{}

		fset := token.NewFileSet()
		node, err := parser.ParseFile(fset, path, nil, parser.ImportsOnly)
		if err != nil {
			return err
		}

		for _, imp := range node.Imports {
			importPath := strings.Trim(imp.Path.Value, `"`)
			if strings.HasPrefix(importPath, moduleName) {
				rel := strings.TrimPrefix(importPath, moduleName)
				rel = strings.TrimPrefix(rel, "/")
				pkgDir := filepath.Join(rootPath, rel)

				err := filepath.Walk(pkgDir, func(f string, info os.FileInfo, err error) error {
					if err != nil {
						return err
					}
					if info.IsDir() || !strings.HasSuffix(info.Name(), ".go") {
						return nil
					}
					return walk(f)
				})
				if err != nil {
					return err
				}
			}
		}
		return nil
	}

	if err := walk(mainPath); err != nil {
		return nil, err
	}
	return visited, nil
}
