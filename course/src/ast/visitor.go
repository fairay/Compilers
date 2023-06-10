package ast

import (
	"course/compilers/parser"
	"fmt"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
)

type Visitor struct {
	parser.BasetinycVisitor

	Module *ir.Module
	curBlock *ir.Block
}

func NewVisitor() *Visitor {
	return &Visitor{}
}

func (v *Visitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
	fmt.Printf("VisitProgram\n")
	v.Module = ir.NewModule()
	main := v.Module.NewFunc("main", types.I32)
	v.curBlock = main.NewBlock("")

	for _, spec := range ctx.AllStatement() {
		v.VisitStatement(spec.(*parser.StatementContext))
	}

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitStatement(ctx *parser.StatementContext) interface {} {
	fmt.Printf("VisitStatement: %s\n", ctx.GetText())
	return nil
}

func (v *Visitor) VisitInteger(ctx *parser.IntegerContext) interface{} {
	fmt.Printf("Visiting integer: %d\n", 0)
	return v.VisitChildren(ctx)
}
