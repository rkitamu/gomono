package astutil

import (
	"go/parser"
	"go/token"
	"os"
)

func ExtractImports(filePath string) ([]string, error) {
	fset := token.NewFileSet()

	src, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	node, err := parser.ParseFile(fset, filePath, src, parser.ImportsOnly)
	if err != nil {
		return nil, err
	}

	var imports []string
	for _, imp := range node.Imports {
		path := imp.Path.Value
		imports = append(imports, trimQuotes(path))
	}

	return imports, nil
}

func trimQuotes(s string) string {
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		return s[1 : len(s)-1]
	}
	return s
}
