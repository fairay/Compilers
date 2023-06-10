package ast

import (
	bp "course/compilers/parser"		// baseparser
	"fmt"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
)

type Visitor struct {
	bp.BasetinycVisitor

	Module *ir.Module
	curBlock *ir.Block
}

func NewVisitor() *Visitor {
	return &Visitor{}
}

func (v *Visitor) VisitCompilationUnit(ctx *bp.CompilationUnitContext) interface{} {
	fmt.Printf("VisitCompilationUnit\n")
	v.Module = ir.NewModule()
	main := v.Module.NewFunc("main", types.I32)
	v.curBlock = main.NewBlock("")

	// for _, spec := range ctx.GetChildren() {
	// 	v.VisitStatement(spec.(*bp.Ext))
	// }
	return nil
}
