package merge

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
)

// ExtractCodeWithoutPackageAndImports parses a .go file and returns only the declarations,
// excluding the package clause and import specs.
func ExtractCodeWithoutPackageAndImports(filePath string) (string, error) {
	fset := token.NewFileSet()

	src, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	// Parse full file including declarations
	node, err := parser.ParseFile(fset, filePath, src, parser.AllErrors)
	if err != nil {
		return "", err
	}

	// Remove import declarations
	node.Imports = nil
	node.Decls = filterNonImportDecls(node.Decls)

	// Buffer output
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, fset, node.Decls); err != nil {
		return "", err
	}

	return buf.String(), nil
}

// Remove import declarations from declarations
func filterNonImportDecls(decls []ast.Decl) []ast.Decl {
	var result []ast.Decl
	for _, decl := range decls {
		// Skip import declarations
		if gen, ok := decl.(*ast.GenDecl); ok && gen.Tok == token.IMPORT {
			continue
		}
		result = append(result, decl)
	}
	return result
}
