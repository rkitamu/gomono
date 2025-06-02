package codegen

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/token"
	"os"
)

func GenerateToFile(fset *token.FileSet, node ast.Node, filePath string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	return format.Node(f, fset, node)
}

func GenerateToStdout(fset *token.FileSet, node ast.Node) error {
	var buf bytes.Buffer
	if err := format.Node(&buf, fset, node); err != nil {
		return err
	}
	_, err := buf.WriteTo(os.Stdout)
	if err != nil {
		return err
	}
	return nil
}
