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
	mainBlock	*ir.Block
	blocks      []*ir.Block
	varScopes   []VarScope
	curDeclType types.Type
}

func NewVisitor() *Visitor {
	return &Visitor{}
}

func testModule() {
	// Создаем модуль
	module := ir.NewModule()

	// Добавляем функцию main без параметров и возвращаемого значения
	mainFunc := module.NewFunc("main", types.I32, ir.NewParam("", types.Void))
	block := mainFunc.NewBlock("")

	// Создаем переменную a и присваиваем ей значение 0
	a := block.NewAlloca(types.I32)
	block.NewStore(constant.NewInt(types.I32, 0), a)

	// Создаем условие 1 == 2
	cmp := block.NewICmp(enum.IPredEQ, constant.NewInt(types.I32, 1), constant.NewInt(types.I32, 2))

	// Создаем блоки ifTrue и ifFalse
	ifTrue := mainFunc.NewBlock("if_true")
	ifFalse := mainFunc.NewBlock("if_false")
	nextBlock := mainFunc.NewBlock("next")

	// Создаем условный переход
	block.NewCondBr(cmp, ifTrue, ifFalse)

	// В блоке ifTrue увеличиваем a на 1
	ifTrue.NewStore(ifTrue.NewAdd(ifTrue.NewLoad(types.I32, a), constant.NewInt(types.I32, 1)), a)
	ifTrue.NewBr(nextBlock)

	// В блоке ifFalse просто переходим к следующей инструкции
	ifFalse.NewBr(nextBlock)

	// Продолжаем в следующем блоке
	nextBlock.NewStore(nextBlock.NewSDiv(nextBlock.NewLoad(types.I32, a), constant.NewInt(types.I32, 2)), a)
	nextBlock.NewStore(nextBlock.NewSub(nextBlock.NewLoad(types.I32, a), constant.NewInt(types.I32, 1)), a)
	nextBlock.NewRet(nextBlock.NewLoad(types.I32, a))

	// Выводим LLVM IR на экран
	fmt.Println(module)

	// Компилируем и выполняем модуль
	// engine, err := ir.NewJITCompiler(module)
	// if err != nil {
	// 	fmt.Println("Failed to create JIT compiler:", err)
	// 	return
	// }
	// result, err := engine.RunFunc(mainFunc, nil)
	// if err != nil {
	// 	fmt.Println("Failed to run main function:", err)
	// 	return
	// }
	// fmt.Println("Result:", result.Int())
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
	
	v.block().NewCondBr(cond, ifBlock, elseBlock)
	v.replaceBlock(nextBlock)
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
