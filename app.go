package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"

	"github.com/spf13/pflag"
)

const version = "0.0.1"

var checkVar bool
var printLine bool
var selfDepDepth int
var depOnly bool
var fnOnly bool

func init() {
	pflag.BoolVarP(&checkVar, "version", "v", false, "Checks Version")
	pflag.BoolVarP(&printLine, "line", "l", false, "[For Self-Dep] Prints Line")
	pflag.IntVarP(&selfDepDepth, "depth", "d", -1, "[For Self-Dep] Depth on Dependency Check")
	pflag.BoolVarP(&depOnly, "dep", "r", false, "[For Self-Dep] Prints Dependency Only")
	pflag.BoolVarP(&fnOnly, "fn", "f", false, "[For Self-Dep] Prints Function Only")
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("usage: go-analyze <mode> <file> [flags...]")
		fmt.Println("use -h | --help for help")
		return
	}

	pflag.Parse()
	if checkVar {
		fmt.Printf("Version %v \n", version)
		return
	}

	if pflag.NArg() == 0 {
		fmt.Println("No arguments given, please use the following format: go-analyze <mode> <file> [flags...]")
		return
	}

	mode := pflag.Arg(0)
	file := pflag.Arg(1)

	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, file, nil, parser.ParseComments)
	if err != nil {
		fmt.Println("Failed to open file:", file)
		return
	}

	switch mode {
	case "self-dep":
		modeStructSelfDependency(f, fset)
	}

}

func modeStructSelfDependency(f *ast.File, fset *token.FileSet) {
	for _, v := range f.Decls {
		if fn, ok := v.(*ast.FuncDecl); ok {
			fmt.Println("Function:", fn.Name.Name+"()")
			dep := selfDepChecker(fn, fset)
			for _, r := range dep {
				fmt.Println(r)
			}
		}
	}
}
