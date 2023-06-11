package main

import (
	"course/compilers/ast"
	baseparser "course/compilers/parser"
	"course/compilers/utils"
	"log"
	"os"
	"path/filepath"

	"github.com/antlr4-go/antlr/v4"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
)

// little trick to define parser type
var parser = baseparser.NewtinycParser(nil)

type TreeShapeListener struct {
	*baseparser.BasetinycListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (listener *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	name := parser.RuleNames[ctx.GetRuleIndex()]
	log.Printf("[%30s] %s\n", name, ctx.GetText())
}

func main() {
	module, err := walkTree(utils.Config().SourcePath)
	if err != nil {
		panic(err)
	}

	addMainStartup(module)
	dumpModule(module)
}

func dumpModule(module *ir.Module) {
	filePath := utils.Config().OutputPath

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

func walkTree(fileName string) (*ir.Module, error) {
	input, _ := antlr.NewFileStream(fileName)
	lexer := baseparser.NewtinycLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
	parser = baseparser.NewtinycParser(stream)

	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	parser.BuildParseTrees = true
	tree := parser.CompilationUnit()
	log.Println("Walking all tree nodes")
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
	log.Println("All tree nodes traversed")

	visitor := ast.NewVisitor()
	visitor.VisitCompilationUnit(tree.(*baseparser.CompilationUnitContext))

	return visitor.Module, nil
}

func init() {
	utils.Config(os.Args[1:]...)
}
