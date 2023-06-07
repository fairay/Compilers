// Code generated from tinyc.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // tinyc
import "github.com/antlr4-go/antlr/v4"

type BasetinycVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasetinycVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetinycVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetinycVisitor) VisitParen_expr(ctx *Paren_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetinycVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetinycVisitor) VisitTest(ctx *TestContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetinycVisitor) VisitSum_(ctx *Sum_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetinycVisitor) VisitTerm(ctx *TermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetinycVisitor) VisitId_(ctx *Id_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetinycVisitor) VisitInteger(ctx *IntegerContext) interface{} {
	return v.VisitChildren(ctx)
}
