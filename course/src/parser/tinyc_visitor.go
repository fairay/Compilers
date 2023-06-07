// Code generated from tinyc.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // tinyc
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by tinycParser.
type tinycVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by tinycParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by tinycParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by tinycParser#paren_expr.
	VisitParen_expr(ctx *Paren_exprContext) interface{}

	// Visit a parse tree produced by tinycParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by tinycParser#test.
	VisitTest(ctx *TestContext) interface{}

	// Visit a parse tree produced by tinycParser#sum_.
	VisitSum_(ctx *Sum_Context) interface{}

	// Visit a parse tree produced by tinycParser#term.
	VisitTerm(ctx *TermContext) interface{}

	// Visit a parse tree produced by tinycParser#id_.
	VisitId_(ctx *Id_Context) interface{}

	// Visit a parse tree produced by tinycParser#integer.
	VisitInteger(ctx *IntegerContext) interface{}
}
