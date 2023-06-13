package ast

import (
	bp "course/compilers/parser" // baseparser
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

type Visitor struct {
	bp.BasetinycVisitor

	Module      *ir.Module
	curBlock    *ir.Block
	curDeclType types.Type
}

func NewVisitor() *Visitor {
	return &Visitor{}
}

func (v *Visitor) VisitCompilationUnit(ctx *bp.CompilationUnitContext) interface{} {
	v.Module = ir.NewModule()

	for _, decl := range ctx.AllExternalDeclaration() {
		v.VisitExternalDeclaration(decl.(*bp.ExternalDeclarationContext))
	}
	return nil
}

func (v *Visitor) VisitExternalDeclaration(ctx *bp.ExternalDeclarationContext) interface{} {
	if d := ctx.FunctionDefinition(); d != nil {
		v.VisitFunctionDefinition(d.(*bp.FunctionDefinitionContext))
	} else {
		panic(UnimplementedError(ctx.GetText()))
	}
	return nil
}

func (v *Visitor) VisitFunctionDefinition(ctx *bp.FunctionDefinitionContext) interface{} {
	retType := types.I32           // TODO: VisitDeclarationSpecifiers
	name := "main"                 // TODO: VisitDeclarator
	params := make([]*ir.Param, 0) // TODO: VisitDeclarationList
	function := v.Module.NewFunc(name, retType, params...)

	compoundStatement := ctx.CompoundStatement()
	block := v.VisitCompoundStatement(compoundStatement.(*bp.CompoundStatementContext))
	function.Blocks = append(function.Blocks, block)

	return nil
}

func (v *Visitor) VisitCompoundStatement(ctx *bp.CompoundStatementContext) *ir.Block {
	block := ir.NewBlock("")

	v.curBlock = block
	for _, nodeCtx := range ctx.AllBlockItem() {
		v.VisitBlockItem(nodeCtx.(*bp.BlockItemContext))
	}
	v.curBlock = nil

	return block
}

func (v *Visitor) VisitBlockItem(ctx *bp.BlockItemContext) interface{} {
	if nodeCtx := ctx.Statement(); nodeCtx != nil {
		v.VisitStatement(nodeCtx.(*bp.StatementContext))
	} else if nodeCtx := ctx.Declaration(); nodeCtx != nil {
		v.VisitDeclaration(nodeCtx.(*bp.DeclarationContext))
	}
	return nil
}

// statement
//
//	:   labeledStatement
//	|   compoundStatement
//	|   expressionStatement
//	|   selectionStatement
//	|   iterationStatement
//	|   jumpStatement
//	;
func (v *Visitor) VisitStatement(ctx *bp.StatementContext) interface{} {
	if nodeCtx := ctx.JumpStatement(); nodeCtx != nil {
		v.VisitJumpStatement(nodeCtx.(*bp.JumpStatementContext))
	} else if nodeCtx := ctx.ExpressionStatement(); nodeCtx != nil {
		v.VisitExpressionStatement(nodeCtx.(*bp.ExpressionStatementContext))
	} else {
		panic(UnimplementedError(ctx.GetText()))
	}
	return nil
}

func (v *Visitor) VisitExpressionStatement(ctx *bp.ExpressionStatementContext) interface{} {
	if nodeCtx := ctx.Expression(); nodeCtx != nil {
		v.VisitExpression(nodeCtx.(*bp.ExpressionContext))
	}
	return nil
}

// jumpStatement: 'return' expression ';' ;
func (v *Visitor) VisitJumpStatement(ctx *bp.JumpStatementContext) interface{} {
	expCtx := ctx.Expression()
	retValue := v.VisitExpression(expCtx.(*bp.ExpressionContext))
	v.curBlock.NewRet(v.dereference(retValue))
	return nil
}

func (v *Visitor) VisitExpression(ctx *bp.ExpressionContext) value.Value {
	var expr value.Value
	for index, nodeCtx := range ctx.AllAssignmentExpression() {
		if index > 0 {
			panic(UnimplementedError(ctx.GetText()))
		}
		expr = v.VisitAssignmentExpression(nodeCtx.(*bp.AssignmentExpressionContext))
	}
	return expr
}

// assignmentExpression
//
//	:   logicalOrExpression
//	|   postfixExpression assignmentOperator assignmentExpression
//	;
func (v *Visitor) VisitAssignmentExpression(ctx *bp.AssignmentExpressionContext) value.Value {
	if nodeCtx := ctx.LogicalOrExpression(); nodeCtx != nil {
		return v.VisitLogicalOrExpression(nodeCtx.(*bp.LogicalOrExpressionContext))
	} else if assigmentCtx := ctx.AssignmentExpression(); assigmentCtx != nil {
		postfixCtx := ctx.PostfixExpression()
		postfix := v.VisitPostfixExpression(postfixCtx.(*bp.PostfixExpressionContext))
		assigment := v.VisitAssignmentExpression(assigmentCtx.(*bp.AssignmentExpressionContext))
		v.curBlock.NewStore(v.dereference(assigment), postfix)
		return assigment
	} else {
		panic(UnimplementedError(ctx.GetText()))
	}
}

func (v *Visitor) VisitLogicalOrExpression(ctx *bp.LogicalOrExpressionContext) value.Value {
	var expr value.Value
	for index, nodeCtx := range ctx.AllLogicalAndExpression() {
		nextExpr := v.VisitLogicalAndExpression(nodeCtx.(*bp.LogicalAndExpressionContext))
		if index == 0 {
			expr = nextExpr
		} else {
			expr = v.curBlock.NewOr(expr, nextExpr)
		}
	}
	return expr
}

func (v *Visitor) VisitLogicalAndExpression(ctx *bp.LogicalAndExpressionContext) value.Value {
	var expr value.Value
	for index, nodeCtx := range ctx.AllEqualityExpression() {
		nextExpr := v.VisitEqualityExpression(nodeCtx.(*bp.EqualityExpressionContext))
		if index == 0 {
			expr = nextExpr
		} else {
			expr = v.curBlock.NewAnd(expr, nextExpr)
		}
	}
	return expr
}

func (v *Visitor) VisitEqualityExpression(ctx *bp.EqualityExpressionContext) value.Value {
	// TODO: support float
	var expr value.Value
	var pred enum.IPred
	for index, nodeCtx := range ctx.GetChildren() {
		if index%2 == 1 {
			node := nodeCtx.(antlr.TerminalNode)
			switch node.GetText() {
			case "==":
				pred = enum.IPredEQ
			case "!=":
				pred = enum.IPredNE
			}
		} else if index == 0 {
			expr = v.VisitRelationalExpression(nodeCtx.(*bp.RelationalExpressionContext))
		} else {
			nextExpr := v.VisitRelationalExpression(nodeCtx.(*bp.RelationalExpressionContext))
			expr = v.curBlock.NewICmp(pred, expr, nextExpr)
		}
	}
	return expr
}

func (v *Visitor) VisitRelationalExpression(ctx *bp.RelationalExpressionContext) value.Value {
	// TODO: support float
	var expr value.Value
	var pred enum.IPred
	for index, nodeCtx := range ctx.GetChildren() {
		if index%2 == 1 {
			node := nodeCtx.(antlr.TerminalNode)
			switch node.GetText() {
			case "<":
				pred = enum.IPredSLT
			case ">":
				pred = enum.IPredSGT
			case "<=":
				pred = enum.IPredSLE
			case ">=":
				pred = enum.IPredSGE
			}
		} else if index == 0 {
			expr = v.VisitAdditiveExpression(nodeCtx.(*bp.AdditiveExpressionContext))
		} else {
			nextExpr := v.VisitAdditiveExpression(nodeCtx.(*bp.AdditiveExpressionContext))
			expr = v.curBlock.NewICmp(pred, expr, nextExpr)
		}
	}
	return expr
}

func (v *Visitor) VisitAdditiveExpression(ctx *bp.AdditiveExpressionContext) value.Value {
	// TODO: support float
	var expr value.Value
	var operation string
	for index, nodeCtx := range ctx.GetChildren() {
		if index%2 == 1 {
			node := nodeCtx.(antlr.TerminalNode)
			operation = node.GetText()
		} else if index == 0 {
			expr = v.VisitMultiplicativeExpression(nodeCtx.(*bp.MultiplicativeExpressionContext))
		} else {
			nextExpr := v.VisitMultiplicativeExpression(nodeCtx.(*bp.MultiplicativeExpressionContext))
			switch operation {
			case "+":
				expr = v.curBlock.NewAdd(v.sameT(expr, nextExpr))
			case "-":
				expr = v.curBlock.NewSub(v.sameT(expr, nextExpr))
			}
		}
	}
	return expr
}

func (v *Visitor) VisitMultiplicativeExpression(ctx *bp.MultiplicativeExpressionContext) value.Value {
	// TODO: support float
	var expr value.Value
	var operation string
	for index, nodeCtx := range ctx.GetChildren() {
		if index%2 == 1 {
			node := nodeCtx.(antlr.TerminalNode)
			operation = node.GetText()
		} else if index == 0 {
			expr = v.VisitCastExpression(nodeCtx.(*bp.CastExpressionContext))
		} else {
			nextExpr := v.VisitCastExpression(nodeCtx.(*bp.CastExpressionContext))
			switch operation {
			case "*":
				expr = v.curBlock.NewMul(v.sameT(expr, nextExpr))
			case "/":
				expr = v.curBlock.NewSDiv(v.sameT(expr, nextExpr))
			case "%":
				expr = v.curBlock.NewSRem(v.sameT(expr, nextExpr))
			}
		}
	}
	return expr
}

// castExpression
//
//	:   '(' typeName ')' castExpression
//	|   postfixExpression
//	;
func (v *Visitor) VisitCastExpression(ctx *bp.CastExpressionContext) value.Value {
	if nodeCtx := ctx.PostfixExpression(); nodeCtx != nil {
		return v.VisitPostfixExpression(nodeCtx.(*bp.PostfixExpressionContext))
	} else {
		panic(UnimplementedError(ctx.GetText()))
	}
}

// postfixExpression
//
//	:
//	(   primaryExpression
//	|   '(' typeName ')' '{' initializerList ','? '}'
//	)
//	('[' expression ']'
//	| '(' argumentExpressionList? ')'
//	)*
//	;
func (v *Visitor) VisitPostfixExpression(ctx *bp.PostfixExpressionContext) value.Value {
	var expr value.Value
	if nodeCtx := ctx.PrimaryExpression(); nodeCtx != nil {
		expr = v.VisitPrimaryExpression(nodeCtx.(*bp.PrimaryExpressionContext))
	} else {
		panic(UnimplementedError(ctx.GetText()))
	}

	for range ctx.AllArgumentExpressionList() {
		panic(UnimplementedError(ctx.GetText()))
	}

	for _, nodeCtx := range ctx.AllExpression() {
		idx := v.VisitExpression(nodeCtx.(*bp.ExpressionContext))
		arrPtr := toPtr(expr)
		expr = v.curBlock.NewGetElementPtr(arrPtr.ElemType, arrPtr, constant.NewInt(types.I64, 0), v.dereference(idx))
	}
	return expr
}

// primaryExpression
//
//	:   Identifier
//	|   Constant
//	|   StringLiteral+
//	|   '(' expression ')'
//	;
func (v *Visitor) VisitPrimaryExpression(ctx *bp.PrimaryExpressionContext) value.Value {
	if nodeCtx := ctx.Constant(); nodeCtx != nil {
		val, err := strconv.Atoi(nodeCtx.GetText())
		if err != nil {
			panic(err)
		}
		return constant.NewInt(types.I32, int64(val))
	} else if nodeCtx := ctx.Expression(); nodeCtx != nil {
		return v.VisitExpression(nodeCtx.(*bp.ExpressionContext))
	} else if nodeCtx := ctx.Identifier(); nodeCtx != nil {
		name := ctx.GetText()
		id := getVariableByName(v.curBlock, name)
		if id == nil {
			panic(UndefindedError(name))
		}
		return id
	} else {
		panic(UnimplementedError(ctx.GetText()))
	}
}
