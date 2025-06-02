package merger

import (
	"go/ast"
	"go/token"

	"github.com/rkitamu/gomono/internal/deps"
)

var imports []ast.ImportSpec

func MergeLocalDependencies(dependencies []*deps.DependPackage) (*token.FileSet, *ast.File, error) {
	asts := make([]*ast.File, 0)
	for i := 0; i < len(dependencies); i++ {
		for j := 0; j < len(dependencies[i].Files); j++ {
			removeLocalImport(dependencies[i].Files[j].AST)
			popImport(dependencies[i].Files[j].AST)
			renameIdent(dependencies[i].Files[j].AST)
			addMergedComment(dependencies[i], j)
			asts = append(asts, dependencies[i].Files[j].AST)
		}
	}

	// TODO: merge asts
	merged := *ast.File{}
	for _, a := range asts {
		
	}

	return sp.FSet, sp.AST, nil
}

func removeLocalImport(a *ast.File) {
	// TODO
}

func popImport(a *ast.File) {
	// TODO
	imports = append(imports, ast.ImportSpec{})
}

func renameIdent(a *ast.File) {
	// TODO
	// rename ident
	// {cur_package_name}_Ident
	// {reference_package_name}_Ident
}

func addMergedComment(a *deps.DependPackage, j int) {
	// TODO: add comment to ast

	// -----------------------------------------
	// package: a.Name, path: a.Files[j].Path
	// -----------------------------------------
}