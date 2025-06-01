package deps

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/rkitamu/gomono/internal/util"
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

var parsedPackages map[string]struct{}
var dependPackages []*DependPackage
var module string
var baseDir string

func AnalyzeLocalDependencies(goFilePath, goModPath, moduleName string) ([]*DependPackage, error) {
	parsedPackages = make(map[string]struct{})
	dependPackages = make([]*DependPackage, 0)
	module = moduleName
	baseDir = filepath.Dir(goModPath)

	// get starting point package name
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, goFilePath, nil, parser.PackageClauseOnly)
	if err != nil {
		return nil, fmt.Errorf("parse failed %v", err)
	}
	packageName := f.Name.Name

	// get starting point module path
	rel, err := filepath.Rel(baseDir, filepath.Dir(goFilePath))
	if err != nil {
		return nil, err
	}
	packagePath := filepath.ToSlash(rel)

	// recurcive analyze
	err = analyzeLocalDependenciesRecursive(packagePath, packageName)
	if err != nil {
		return nil, err
	}

	return dependPackages, nil
}

func analyzeLocalDependenciesRecursive(packagePath, packageName string) error {
	packagePathName := packagePath + "/" + packageName
	if _, ok := parsedPackages[packagePathName]; ok {
		return nil
	} else {
		parsedPackages[packagePathName] = struct{}{}
	}

	dir := baseDir + "/" + packagePath
	files, err := util.ListFileFromDir(dir)
	if err != nil {
		return err
	}
	goFiles := make([]string, 0)
	for _, file := range files {
		if strings.HasSuffix(file, ".go") {
			goFiles = append(goFiles, file)
		}
	}

	parsedList := make([]*ParsedFile, 0)
	for _, v := range goFiles {
		parsed, err := parseGoFile(v)
		if err != nil {
			return err
		}
		parsedList = append(parsedList, parsed)
	}

	dependPackages = append(dependPackages, &DependPackage{
		Name:  packageName,
		Path:  packagePath,
		Files: parsedList,
	})

	return nil
}

func parseGoFile(goFilePath string) (*ParsedFile, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, goFilePath, nil, parser.ParseComments|parser.AllErrors)
	if err != nil {
		return nil, fmt.Errorf("parse failed %v", err)
	}

	for _, imp := range f.Imports {
		s, err := strconv.Unquote(imp.Path.Value)
		if err != nil {
			return nil, err
		}
		if strings.HasPrefix(s, module) {
			nPackagePath := strings.TrimPrefix(s, module+"/")
			dir := baseDir + "/" + nPackagePath
			files, err := util.ListFileFromDir(dir)
			if err != nil {
				return nil, err
			}
			goFiles := make([]string, 0)
			for _, file := range files {
				if strings.HasSuffix(file, ".go") {
					goFiles = append(goFiles, file)
				}
			}
			nPackageName, err := GetPackageName(goFiles[0])
			if err != nil {
				return nil, err
			}
			if err := analyzeLocalDependenciesRecursive(nPackagePath, nPackageName); err != nil {
				return nil, err
			}
		}
	}

	return &ParsedFile{
		Path: goFilePath,
		FSet: fset,
		AST:  f,
	}, nil
}

func GetPackageName(goFilePath string) (string, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, goFilePath, nil, parser.PackageClauseOnly)
	if err != nil {
		return "", fmt.Errorf("parse failed %v", err)
	}
	return f.Name.Name, nil
}

func IsLocalImport(importPath, moduleName string) bool {
	return strings.HasPrefix(importPath, moduleName)
}
