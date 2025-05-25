package merge

import (
	"go/parser"
	"go/token"
	"os"
	"sort"
	"strings"
)

// ExtractExternalImports collects all non-local import paths from the given files.
func ExtractExternalImports(filePaths []string, moduleName string) ([]string, error) {
	importSet := map[string]struct{}{}

	for _, path := range filePaths {
		fset := token.NewFileSet()
		src, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}

		node, err := parser.ParseFile(fset, path, src, parser.ImportsOnly)
		if err != nil {
			return nil, err
		}

		for _, imp := range node.Imports {
			path := strings.Trim(imp.Path.Value, `"`)
			// Exclude local imports
			if !strings.HasPrefix(path, moduleName) {
				importSet[path] = struct{}{}
			}
		}
	}

	var imports []string
	for path := range importSet {
		imports = append(imports, path)
	}
	sort.Strings(imports)
	return imports, nil
}
