package main

import (
	"github.com/antlr4-go/antlr/v4"
	"course/compilers/parser"
	"os"
	"fmt"
)

type TreeShapeListener struct {
	*parser.BasetinycListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (listener *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewtinycLexer(input)
	stream := antlr.NewCommonTokenStream(lexer,0)
	p := parser.NewtinycParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Program()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}
