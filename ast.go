package main

import (
	"fmt"
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

type Package struct {
	*doc.Package
}

type StructField struct {
	ast.Field
}

func (f *StructField) Name() string {
	return f.Names[0].String()
}

func (f *StructField) TypeName() string {
	return ExprString(f.Type)
}

type StructInfo struct {
	ast.StructType

	Package Package

	Name   string
	Fields []StructField
}

func getPackageStructs(path string) (structs []*StructInfo, err error) {
	pkgs, err := parser.ParseDir(token.NewFileSet(), path, func(fileInfo os.FileInfo) bool {
		return strings.HasSuffix(fileInfo.Name(), ".go") ||
			!strings.HasSuffix(fileInfo.Name(), "_gen.go") ||
			!strings.HasSuffix(fileInfo.Name(), "_test.go")
	}, parser.ParseComments)
	if err != nil {
		err = fmt.Errorf("%w: parser.ParseDir", err)
		return
	}

	for _, astPkg := range pkgs {
		docPkg := doc.New(astPkg, path, doc.AllDecls)
		pkg := Package{docPkg}
		for _, ptyp := range docPkg.Types {
			for _, spec := range ptyp.Decl.Specs {
				typeSpec := spec.(*ast.TypeSpec)

				structTypePtr, ok := typeSpec.Type.(*ast.StructType)
				if !ok {
					continue
				}

				newStruct := &StructInfo{
					Package: pkg,
					Name:    typeSpec.Name.String(),
				}
				newStruct.StructType = *structTypePtr
				for _, filed := range newStruct.StructType.Fields.List {
					newField := *filed
					newStruct.Fields = append(newStruct.Fields, StructField{newField})
				}
				structs = append(structs, newStruct)
			}
		}
	}

	return
}
