package ast

import (
	bp "course/compilers/parser" // baseparser
	"log"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

type Visitor struct {
	bp.BasetinycVisitor

	Module   *ir.Module
	curBlock *ir.Block
}

func NewVisitor() *Visitor {
	return &Visitor{}
}

func (v *Visitor) VisitCompilationUnit(ctx *bp.CompilationUnitContext) interface{} {
	log.Println("VisitCompilationUnit")
	v.Module = ir.NewModule()

	for _, decl := range ctx.AllExternalDeclaration() {
		v.VisitExternalDeclaration(decl.(*bp.ExternalDeclarationContext))
	}
	return nil
}

func (v *Visitor) VisitExternalDeclaration(ctx *bp.ExternalDeclarationContext) interface{} {
	log.Println("VisitExternalDeclaration")
	if d := ctx.FunctionDefinition(); d != nil {
		v.VisitFunctionDefinition(d.(*bp.FunctionDefinitionContext))
	} else {
		panic(UnimplementedError(ctx.GetText()))
	}
	return nil
}

func (v *Visitor) VisitFunctionDefinition(ctx *bp.FunctionDefinitionContext) interface{} {
	log.Println("VisitFunctionDefinition")
	retType := types.I32		// TODO: VisitDeclarationSpecifiers
	name := "main"				// TODO: VisitDeclarator
	params := make([]*ir.Param, 0)	// TODO: VisitDeclarationList
	function := v.Module.NewFunc(name, retType, params...)

	v.curBlock = function.NewBlock("")
	{
		// TODO: VisitCompoundStatement
		v.curBlock.NewRet(constant.NewInt(types.I32, 42))
	}
	v.curBlock = nil

	return nil
}
