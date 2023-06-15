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
	mainBlock   *ir.Block
	blocks      []*ir.Block
	varScopes   []VarScope
	curDeclType types.Type
}

func NewVisitor() *Visitor {
	return &Visitor{}
}

func testModule() {
	// Create a new LLVM IR module.
	module := ir.NewModule()

	// Create the main function.
	mainFunc := module.NewFunc("main", types.I32)
	block := mainFunc.NewBlock("entry")

	// Allocate memory for variable 'a'.
	a := block.NewAlloca(types.I32)

	// Initialize 'a' with value 2.
	block.NewStore(constant.NewInt(types.I32, 2), a)

	// Create the loop condition block.
	loopCondBlock := mainFunc.NewBlock("loop.cond")
	loopBodyBlock := mainFunc.NewBlock("loop.body")
	loopExitBlock := mainFunc.NewBlock("loop.exit")

	// Branch from entry block to loop condition block.
	block.NewBr(loopCondBlock)

	// Loop condition block.
	loopCondBlock.NewBr(loopCondBlock) // Infinite loop until break condition.

	// Load the current value of 'a'.
	current := loopCondBlock.NewLoad(types.I32, a)

	// Compare 'current' with 10.
	cmp := loopCondBlock.NewICmp(enum.IPredSLT, current, constant.NewInt(types.I32, 10))

	// Branch to either loop body or loop exit based on the comparison result.
	loopCondBlock.NewCondBr(cmp, loopBodyBlock, loopExitBlock)

	// Loop body block.
	loopBodyBlock.NewBr(loopCondBlock) // Jump back to loop condition block.

	// Multiply 'current' by 2.
	mul := loopBodyBlock.NewMul(current, constant.NewInt(types.I32, 2))

	// Store the updated value back to 'a'.
	loopBodyBlock.NewStore(mul, a)

	// Branch back to loop condition block.
	loopBodyBlock.NewBr(loopCondBlock)

	// Loop exit block.
	exitValue := loopExitBlock.NewLoad(types.I32, a)

	// Return the final value of 'a'.
	loopExitBlock.NewRet(exitValue)

	// Print the LLVM IR assembly code.
	fmt.Println(module)
}

func (v *Visitor) VisitCompilationUnit(ctx *bp.CompilationUnitContext) interface{} {
	// testModule()
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

	block := function.NewBlock("main")
	v.pushBlock(block)
	compoundStatement := ctx.CompoundStatement()
	v.VisitCompoundStatement(compoundStatement.(*bp.CompoundStatementContext))
	v.popBlock(block)

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
//     :   '(' typeName ')' castExpression
//     |   postfixExpression
//     |   unaryExpression
//     ;
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
		expr = v.block().NewGetElementPtr(arrPtr.ElemType, arrPtr, constant.NewInt(types.I64, 0), v.dereference(idx))
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
		id := v.variable(name)
		if id == nil {
			panic(UndefindedError(name))
		}
		return id
	} else {
		panic(UnimplementedError(ctx.GetText()))
	}
}
