package deps

import (
	"go/ast"
	"go/token"
	"strings"
)

type DependPackage struct {
	Name  string
	Path  string // Path is the import path relative to the module root, e.g. "internal/foo"
	Files []*ParsedFile
}

type ParsedFile struct {
	Path string
	FSet *token.FileSet
	AST  *ast.File
}

func AnalyzeLocalDependencies(filePath, moduleName string) *DependPackage {
	// TODO: Do it
	return &DependPackage{}
}

func IsLocalImport(importPath, moduleName string) bool {
	return strings.HasPrefix(importPath, moduleName)
}
