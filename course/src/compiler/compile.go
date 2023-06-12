package compiler

import (
	"course/compilers/ast"
	baseparser "course/compilers/parser"
	"course/compilers/utils"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/antlr4-go/antlr/v4"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
)

func Compile(config *utils.CompileConfigT) {
	module, err := buildModule(config.SourcePath)
	if err != nil {
		panic(err)
	}

	dumpModule(module, config.IRPath)
	buildExe(config.IRPath, config.ObjectPath, config.ExecutablePath)
}

func dumpModule(module *ir.Module, filePath string) {
	// Create missing dirs
	dirPath := filepath.Dir(filePath)
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		log.Panicf("Failed to create dirs for build path: %v", err)
	}

	err = os.WriteFile(
		filePath,
		[]byte(module.String()),
		0644,
	)
	if err != nil {
		log.Panicf("Failed to write LLVM IR to file: %s", err.Error())
	}
}

func buildExe(irFile string, objFile string, exeFile string) {
	cmd := exec.Command("clang", "-c", "-o", objFile, irFile)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	cmd = exec.Command(
		"lld-link",
		"/subsystem:console",
		fmt.Sprintf("/out:%s", exeFile),
		objFile,
	)
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
}

// addMainStartup adds a startup function "mainCRTStartup" to the given module.
// It looks for a function named "main" in the module and generates a call to it
// from the startup function. The startup function is responsible for invoking
// the main function when the program starts.
func addMainStartup(module *ir.Module) {
	startup := module.NewFunc("mainCRTStartup", types.I32)
	entry := startup.NewBlock("")

	var mainFunc *ir.Func
	for _, f := range module.Funcs {
		if f.Name() == "main" {
			mainFunc = f
			break
		}
	}
	if mainFunc == nil {
		panic(ast.NoMainError())
	}

	mainRes := entry.NewCall(mainFunc)
	entry.NewRet(mainRes)
}

func buildModule(fileName string) (*ir.Module, error) {
	input, _ := antlr.NewFileStream(fileName)
	lexer := baseparser.NewtinycLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
	parser := baseparser.NewtinycParser(stream)
	tree := parser.CompilationUnit()

	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	parser.BuildParseTrees = true

	visitor := ast.NewVisitor()
	visitor.VisitCompilationUnit(tree.(*baseparser.CompilationUnitContext))

	addMainStartup(visitor.Module)
	return visitor.Module, nil
}
