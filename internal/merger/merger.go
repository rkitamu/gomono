package merger

import (
	"go/ast"
	"go/token"

	"github.com/rkitamu/gomono/internal/deps"
	"github.com/rkitamu/gomono/internal/util"
)

func MergeLocalDependencies(startPointFilePath string, dependencies []*deps.DependPackage) (*token.FileSet, *ast.File, error) {
	absStartPath, err := util.NormalizePath(startPointFilePath)
	if err != nil {
		return nil, nil, err
	}

	var sp *deps.ParsedFile
	for _, d := range dependencies {
		for _, f := range d.Files {
			fp, err := util.NormalizePath(f.Path)
			if err != nil {
				return nil, nil, err
			}
			if fp == absStartPath {
				sp = f
				break
			}
		}
		if sp != nil {
			break
		}
	}

	// TODO: merge logic here

	return sp.FSet, sp.AST, nil
}
