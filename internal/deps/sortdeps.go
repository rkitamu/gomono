package deps

import (
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

// TopoSort performs a topological sort of local Go source files
// based on their internal import dependencies.
func TopoSort(files map[string]struct{}, rootPath, moduleName string) ([]string, error) {
	graph := map[string][]string{} // adjacency list: file -> list of imported files
	inDegree := map[string]int{}   // count of incoming edges for each file

	for file := range files {
		imports, err := extractLocalImports(file, moduleName)
		if err != nil {
			return nil, err
		}

		for _, imp := range imports {
			importDir := filepath.Join(rootPath, imp)
			err := filepath.Walk(importDir, func(p string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if info.IsDir() || !strings.HasSuffix(p, ".go") {
					return nil
				}
				graph[file] = append(graph[file], p)
				inDegree[p]++
				return nil
			})
			if err != nil {
				return nil, err
			}
		}
		if _, ok := inDegree[file]; !ok {
			inDegree[file] = 0
		}
	}

	var queue []string
	for f, deg := range inDegree {
		if deg == 0 {
			queue = append(queue, f)
		}
	}

	var result []string
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		result = append(result, curr)
		for _, dep := range graph[curr] {
			inDegree[dep]--
			if inDegree[dep] == 0 {
				queue = append(queue, dep)
			}
		}
	}

	if len(result) != len(inDegree) {
		return nil, ErrCyclicDependency
	}

	return result, nil
}

// ErrCyclicDependency indicates a circular import between local packages.
var ErrCyclicDependency = os.ErrInvalid

// extractLocalImports returns relative import paths within the same module.
func extractLocalImports(path string, moduleName string) ([]string, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, path, nil, parser.ImportsOnly)
	if err != nil {
		return nil, err
	}

	var result []string
	for _, imp := range node.Imports {
		importPath := strings.Trim(imp.Path.Value, `"`)
		if strings.HasPrefix(importPath, moduleName) {
			rel := strings.TrimPrefix(importPath, moduleName)
			rel = strings.TrimPrefix(rel, "/")
			result = append(result, rel)
		}
	}
	return result, nil
}
