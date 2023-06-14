package ast

import (
	bp "course/compilers/parser" // baseparser

	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// declaration
//
//	:   declarationSpecifiers initDeclaratorList? ';'
//	;
func (v *Visitor) VisitDeclaration(ctx *bp.DeclarationContext) interface{} {
	declCtx := ctx.DeclarationSpecifiers()
	v.curDeclType = v.VisitDeclarationSpecifiers(declCtx.(*bp.DeclarationSpecifiersContext))
	initCtx := ctx.InitDeclaratorList()
	v.VisitInitDeclaratorList(initCtx.(*bp.InitDeclaratorListContext))
	v.curDeclType = nil
	return nil
}

// declarationSpecifiers
//
//	:   typeSpecifier+
//	;
func (v *Visitor) VisitDeclarationSpecifiers(ctx *bp.DeclarationSpecifiersContext) types.Type {
	if len(ctx.AllTypeSpecifier()) != 1 {
		panic(UnimplementedError(ctx.GetText()))
	}
	typeCtx := ctx.AllTypeSpecifier()[0]
	return v.VisitTypeSpecifier(typeCtx.(*bp.TypeSpecifierContext))
}

// typeSpecifier
//
//	:   'char'
//	|   'int'
//	|   'float'
//	;
func (v *Visitor) VisitTypeSpecifier(ctx *bp.TypeSpecifierContext) types.Type {
	switch ctx.GetText() {
	case "int":
		return types.I32
	default:
		panic(UnimplementedError(ctx.GetText()))
	}
}

// initDeclaratorList
//
//	:   initDeclarator (',' initDeclarator)*
//	;
func (v *Visitor) VisitInitDeclaratorList(ctx *bp.InitDeclaratorListContext) interface{} {
	for _, nodeCtx := range ctx.AllInitDeclarator() {
		v.VisitInitDeclarator(nodeCtx.(*bp.InitDeclaratorContext))
	}
	return nil
}

// initDeclarator
//
//	:   declarator ('=' initializer)?
//	;
func (v *Visitor) VisitInitDeclarator(ctx *bp.InitDeclaratorContext) interface{} {
	declCtx := ctx.Declarator()
	t, name := v.VisitDeclarator(declCtx.(*bp.DeclaratorContext))
	ptr := v.newVariable(t, name)
	if nodeCtx := ctx.Initializer(); nodeCtx != nil {
		init := v.VisitInitializer(nodeCtx.(*bp.InitializerContext))
		v.block().NewStore(init, ptr)
	}
	return nil
}

// declarator
//
//	:   Identifier
//	|   '(' declarator ')'
//	|   declarator '[' assignmentExpression? ']'
//	|   declarator '(' parameterTypeList ')'
//	|   declarator '(' identifierList? ')'
//	;
func (v *Visitor) VisitDeclarator(ctx *bp.DeclaratorContext) (types.Type, string) {
	t := v.curDeclType
	if nodeCtx := ctx.Identifier(); nodeCtx != nil {
		name := ctx.GetText()
		return t, name
	} else if exprCtx := ctx.AssignmentExpression(); exprCtx != nil {
		declCtx := ctx.Declarator()
		t, name := v.VisitDeclarator(declCtx.(*bp.DeclaratorContext))
		expr := v.VisitAssignmentExpression(exprCtx.(*bp.AssignmentExpressionContext))

		size, ok := expr.(*constant.Int)
		if !ok {
			panic(VariableSizeError(exprCtx.GetText()))
		}
		t = types.NewArray(size.X.Uint64(), t)
		return t, name
	} else {
		panic(UnimplementedError(ctx.GetText()))
	}
}

// initializer
//
//	:   assignmentExpression
//	|   '{' initializerList ','? '}'
//	;
func (v *Visitor) VisitInitializer(ctx *bp.InitializerContext) value.Value {
	if nodeCtx := ctx.AssignmentExpression(); nodeCtx != nil {
		return v.VisitAssignmentExpression(nodeCtx.(*bp.AssignmentExpressionContext))
	} else {
		panic(UnimplementedError(ctx.GetText()))
	}
}
