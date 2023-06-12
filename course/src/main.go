package main

import (
	"course/compilers/compiler"
	baseparser "course/compilers/parser"
	"course/compilers/utils"
	"github.com/antlr4-go/antlr/v4"
	"log"
	"os"
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
	walkTree(utils.Config().SourcePath)
	compiler.Compile(utils.CompileConfig())
}

func walkTree(fileName string) {
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
}

func init() {
	utils.Config(os.Args[1:]...)
}
