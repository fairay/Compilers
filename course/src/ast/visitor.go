package ast

import (
	bp "course/compilers/parser" // baseparser
	"fmt"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

type VarScope map[string]*ir.InstAlloca

type Visitor struct {
	bp.BasetinycVisitor

	Module      *ir.Module
	Function	*ir.Func
	mainBlock   *ir.Block
	blocks      []*ir.Block
	varScopes   []VarScope
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

// functionDefinition
//     :   declarationSpecifiers? Identifier '(' parameterList ')' declarationList? compoundStatement
//     ;
func (v *Visitor) VisitFunctionDefinition(ctx *bp.FunctionDefinitionContext) interface{} {

	retType := v.VisitDeclarationSpecifiers(
		ctx.DeclarationSpecifiers().(*bp.DeclarationSpecifiersContext),
	)
	name := ctx.Identifier().GetText()

	params := v.VisitParameterList(ctx.ParameterList().(*bp.ParameterListContext))
	// params := make([]*ir.Param, 0) // TODO: VisitDeclarationList

	v.Function = v.Module.NewFunc(name, retType, params...)
	block := v.Function.NewBlock("main")
	v.pushBlock(block)

	compoundStatement := ctx.CompoundStatement()
	v.VisitCompoundStatement(compoundStatement.(*bp.CompoundStatementContext))

	v.popBlock(block)
	v.Function = nil
	return nil
}

func (v *Visitor) VisitCompoundStatement(ctx *bp.CompoundStatementContext) interface{} {
	v.pushScope()
	for _, nodeCtx := range ctx.AllBlockItem() {
		v.VisitBlockItem(nodeCtx.(*bp.BlockItemContext))
	}
	v.popScope()
	return nil
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
	} else if nodeCtx := ctx.SelectionStatement(); nodeCtx != nil {
		v.VisitSelectionStatement(nodeCtx.(*bp.SelectionStatementContext))
	} else if nodeCtx := ctx.CompoundStatement(); nodeCtx != nil {
		v.VisitCompoundStatement(nodeCtx.(*bp.CompoundStatementContext))
	} else if nodeCtx := ctx.IterationStatement(); nodeCtx != nil {
		v.VisitIterationStatement(nodeCtx.(*bp.IterationStatementContext))
	} else {
		panic(UnimplementedError(ctx.GetText()))
	}
	return nil
}

// selectionStatement
//
//	:   'if' '(' expression ')' statement ('else' statement)?
//	;
func (v *Visitor) VisitSelectionStatement(ctx *bp.SelectionStatementContext) interface{} {
	exprCtx := ctx.Expression()
	cond := v.VisitExpression(exprCtx.(*bp.ExpressionContext))

	blockNum := len(v.block().Parent.Blocks)
	ifBlock := v.block().Parent.NewBlock(fmt.Sprintf("if-%d", blockNum))
	elseBlock := v.block().Parent.NewBlock(fmt.Sprintf("else-%d", blockNum))
	nextBlock := v.block().Parent.NewBlock(fmt.Sprintf("main-%d", blockNum))

	preBlock := v.replaceBlock(nextBlock)
	preBlock.NewCondBr(v.castCond(cond), ifBlock, elseBlock)
	ifBlock.NewBr(nextBlock)
	elseBlock.NewBr(nextBlock)

	{
		// if
		v.pushBlock(ifBlock)
		stCtx := ctx.AllStatement()[0]
		v.VisitStatement(stCtx.(*bp.StatementContext))
		v.popBlock(ifBlock)
	}

	if len(ctx.AllStatement()) > 1 {
		// else (optional)
		v.pushBlock(elseBlock)
		stCtx := ctx.AllStatement()[1]
		v.VisitStatement(stCtx.(*bp.StatementContext))
		v.popBlock(elseBlock)
	}

	return nil
}

// iterationStatement
//
//	:   While '(' expression ')' statement
//	;
func (v *Visitor) VisitIterationStatement(ctx *bp.IterationStatementContext) interface{} {

	blockNum := len(v.block().Parent.Blocks)
	condBlock := v.block().Parent.NewBlock(fmt.Sprintf("while.cond-%d", blockNum))
	bodyBlock := v.block().Parent.NewBlock(fmt.Sprintf("while.body-%d", blockNum))
	nextBlock := v.block().Parent.NewBlock(fmt.Sprintf("main-%d", blockNum))

	preBlock := v.replaceBlock(nextBlock)
	preBlock.NewBr(condBlock)
	bodyBlock.NewBr(condBlock)

	{
		// condition
		v.pushBlock(condBlock)
		exprCtx := ctx.Expression()
		cond := v.VisitExpression(exprCtx.(*bp.ExpressionContext))
		condBlock.NewCondBr(v.castCond(cond), bodyBlock, nextBlock)
		v.popBlock(condBlock)
	}

	{
		// statement
		v.pushBlock(bodyBlock)
		stCtx := ctx.Statement()
		v.VisitStatement(stCtx.(*bp.StatementContext))
		v.popBlock(bodyBlock)
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
	v.block().NewRet(v.dereference(retValue))
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
		v.block().NewStore(v.dereference(assigment), postfix)
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
			expr = v.block().NewOr(expr, nextExpr)
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
			expr = v.block().NewAnd(expr, nextExpr)
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
			expr = v.block().NewICmp(pred, v.dereference(expr), v.dereference(nextExpr))
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
			expr = v.block().NewICmp(pred, v.dereference(expr), v.dereference(nextExpr))
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
				expr = v.block().NewAdd(v.sameT(expr, nextExpr))
			case "-":
				expr = v.block().NewSub(v.sameT(expr, nextExpr))
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
				expr = v.block().NewMul(v.sameT(expr, nextExpr))
			case "/":
				expr = v.block().NewSDiv(v.sameT(expr, nextExpr))
			case "%":
				expr = v.block().NewSRem(v.sameT(expr, nextExpr))
			}
		}
	}
	return expr
}

// castExpression
//
//	:   '(' typeName ')' castExpression
//	|   postfixExpression
//	|   unaryExpression
//	;
func (v *Visitor) VisitCastExpression(ctx *bp.CastExpressionContext) value.Value {
	if nodeCtx := ctx.PostfixExpression(); nodeCtx != nil {
		return v.VisitPostfixExpression(nodeCtx.(*bp.PostfixExpressionContext))
	} else if nodeCtx := ctx.UnaryExpression(); nodeCtx != nil {
		return v.VisitUnaryExpression(nodeCtx.(*bp.UnaryExpressionContext))
	} else {
		panic(UnimplementedError(ctx.GetText()))
	}
}

// unaryExpression: unaryOperator castExpression;
func (v *Visitor) VisitUnaryExpression(ctx *bp.UnaryExpressionContext) value.Value {
	exprCtx := ctx.CastExpression()
	value := v.VisitCastExpression(exprCtx.(*bp.CastExpressionContext))

	switch ctx.UnaryOperator().GetText() {
	case "-":
		value = v.block().NewSub(constant.NewInt(types.I32, 0), value)
		// TODO: support float
	case "!":
		value = v.block().NewICmp(
			enum.IPredEQ,
			constant.False,
			v.castCond(value),
		)
	case "+":
		// pass
	}
	return value
}

// postfixExpression
//     :
//     (   primaryExpression
//     |   '(' typeName ')' '{' initializerList ','? '}'
//     )
//     ('[' expression ']' | funcCall)*
//     ;
func (v *Visitor) VisitPostfixExpression(ctx *bp.PostfixExpressionContext) value.Value {
	var expr value.Value
	if nodeCtx := ctx.PrimaryExpression(); nodeCtx != nil {
		expr = v.VisitPrimaryExpression(nodeCtx.(*bp.PrimaryExpressionContext))
	} else {
		panic(UnimplementedError(ctx.GetText()))
	}

	if len(ctx.AllFuncCall()) > 0 {
		argsCtx := ctx.FuncCall(0)
		args := v.VisitFuncCall(argsCtx.(*bp.FuncCallContext))
		expr = v.block().NewCall(expr, args...)
	} else if len(ctx.AllExpression()) > 0 {
		variable, ok := expr.(*ir.InstAlloca)
		if !ok {
			panic(ExpectedPtrError(expr))
		}

		indexes := []value.Value{constant.NewInt(types.I32, 0)}
		for _, nodeCtx := range ctx.AllExpression() {
			idx := v.VisitExpression(nodeCtx.(*bp.ExpressionContext))
			indexes = append(indexes, v.dereference(idx))
		}
		expr = v.block().NewGetElementPtr(variable.ElemType, variable, indexes...)
	}
	return expr
}

// funcCall: '(' (assignmentExpression (',' assignmentExpression)*)? ')';
func (v *Visitor) VisitFuncCall(ctx *bp.FuncCallContext) []value.Value {
	args := make([]value.Value, len(ctx.AllAssignmentExpression()))
	for i, argCtx := range ctx.AllAssignmentExpression() {
		args[i] = v.VisitAssignmentExpression(argCtx.(*bp.AssignmentExpressionContext))
	}
	return args
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
		id := v.identifier(name)
		if id == nil {
			panic(UndefindedError(name))
		}
		return id
	} else {
		panic(UnimplementedError(ctx.GetText()))
	}
}
