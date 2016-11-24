package main

import (
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"go/parser"
	"go/printer"
	"go/token"
	"bufio"
	"go/ast"
)

var (
	sourceDir string
	targetDir string
	pkgSuffix string
	sep = string(filepath.Separator)
)

func init() {
	flag.StringVar(&targetDir, "o", "", "Output dir")
	flag.StringVar(&pkgSuffix, "s", "", "Generated package suffix")
}

func main() {
	flag.Parse()
	sourceDir = flag.Args()[0]

	generateAll(sourceDir)
}

func generateAll(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		childPath := path + sep + f.Name()
		if isDir(childPath) {
			generateAll(childPath)
		} else  {
			generate(path, f.Name())
		}
	}
}

func generate(sourcePath, sourceFile string) {
	targetPath := strings.Replace(sourcePath, sourceDir, targetDir, 1) + pkgSuffix
	src := sourcePath + sep + sourceFile
	dst := targetPath + sep + "gen_" + sourceFile

	_, err := os.Stat(dst)
	if err == nil {
		os.Remove(dst)
	}

	os.MkdirAll(targetPath, 0777)

	rewrite(src, dst)
}

func rewrite(src, dst string) {
	find := "TYPE"
	replace := "int32"

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, src, nil, 0)
	if err != nil {
		panic(err)
	}

	ast.Print(fset, f)

	// Change package name
	f.Name.Name = f.Name.Name + pkgSuffix

	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.TypeSpec:
			if x.Name.Name == find {
				x.Type.(*ast.Ident).Name = replace
				return false
			}
		case *ast.Ident:
			if x.Name == find {
				x.Name = replace
			}
		}
		return true
	})

	writeAst(f, fset, dst)
}

func writeAst(f interface{}, fset *token.FileSet, dst string) {
	fOut, err := os.Create(dst)
	if err != nil {
		panic(err)
	}
	defer fOut.Close()
	output := bufio.NewWriter(fOut)
	defer output.Flush()
	printer.Fprint(output, fset, f)
}

func isDir(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		panic(err)
	}
	return fi.IsDir()
}