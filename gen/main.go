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
	dst := targetPath + sep + "gen_" + sourceFile

	_, err := os.Stat(dst)
	if err == nil {
		os.Remove(dst)
	}

	os.MkdirAll(targetPath, 0777)
	os.Link(sourcePath + sep + sourceFile, dst)

	rewrite(dst)
}

func rewrite(path string) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, 0)

	if err != nil {
		panic(err)
	}

	// Change package name
	f.Name.Name = f.Name.Name + pkgSuffix

	writeAst(f, fset, path)
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