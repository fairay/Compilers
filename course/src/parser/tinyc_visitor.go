// Code generated from tinyc.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // tinyc
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by tinycParser.
type tinycVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by tinycParser#primaryExpression.
	VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{}

	// Visit a parse tree produced by tinycParser#genericAssocList.
	VisitGenericAssocList(ctx *GenericAssocListContext) interface{}

	// Visit a parse tree produced by tinycParser#genericAssociation.
	VisitGenericAssociation(ctx *GenericAssociationContext) interface{}

	// Visit a parse tree produced by tinycParser#postfixExpression.
	VisitPostfixExpression(ctx *PostfixExpressionContext) interface{}

	// Visit a parse tree produced by tinycParser#argumentExpressionList.
	VisitArgumentExpressionList(ctx *ArgumentExpressionListContext) interface{}

	// Visit a parse tree produced by tinycParser#unaryExpression.
	VisitUnaryExpression(ctx *UnaryExpressionContext) interface{}

	// Visit a parse tree produced by tinycParser#unaryOperator.
	VisitUnaryOperator(ctx *UnaryOperatorContext) interface{}

	// Visit a parse tree produced by tinycParser#castExpression.
	VisitCastExpression(ctx *CastExpressionContext) interface{}

	// Visit a parse tree produced by tinycParser#multiplicativeExpression.
	VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) interface{}

	// Visit a parse tree produced by tinycParser#additiveExpression.
	VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{}

	// Visit a parse tree produced by tinycParser#relationalExpression.
	VisitRelationalExpression(ctx *RelationalExpressionContext) interface{}

	// Visit a parse tree produced by tinycParser#equalityExpression.
	VisitEqualityExpression(ctx *EqualityExpressionContext) interface{}

	// Visit a parse tree produced by tinycParser#logicalAndExpression.
	VisitLogicalAndExpression(ctx *LogicalAndExpressionContext) interface{}

	// Visit a parse tree produced by tinycParser#logicalOrExpression.
	VisitLogicalOrExpression(ctx *LogicalOrExpressionContext) interface{}

	// Visit a parse tree produced by tinycParser#assignmentExpression.
	VisitAssignmentExpression(ctx *AssignmentExpressionContext) interface{}

	// Visit a parse tree produced by tinycParser#assignmentOperator.
	VisitAssignmentOperator(ctx *AssignmentOperatorContext) interface{}

	// Visit a parse tree produced by tinycParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by tinycParser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by tinycParser#declarationSpecifiers.
	VisitDeclarationSpecifiers(ctx *DeclarationSpecifiersContext) interface{}

	// Visit a parse tree produced by tinycParser#initDeclaratorList.
	VisitInitDeclaratorList(ctx *InitDeclaratorListContext) interface{}

	// Visit a parse tree produced by tinycParser#initDeclarator.
	VisitInitDeclarator(ctx *InitDeclaratorContext) interface{}

	// Visit a parse tree produced by tinycParser#typeSpecifier.
	VisitTypeSpecifier(ctx *TypeSpecifierContext) interface{}

	// Visit a parse tree produced by tinycParser#specifierQualifierList.
	VisitSpecifierQualifierList(ctx *SpecifierQualifierListContext) interface{}

	// Visit a parse tree produced by tinycParser#structDeclaratorList.
	VisitStructDeclaratorList(ctx *StructDeclaratorListContext) interface{}

	// Visit a parse tree produced by tinycParser#structDeclarator.
	VisitStructDeclarator(ctx *StructDeclaratorContext) interface{}

	// Visit a parse tree produced by tinycParser#declarator.
	VisitDeclarator(ctx *DeclaratorContext) interface{}

	// Visit a parse tree produced by tinycParser#gccAttribute.
	VisitGccAttribute(ctx *GccAttributeContext) interface{}

	// Visit a parse tree produced by tinycParser#nestedParenthesesBlock.
	VisitNestedParenthesesBlock(ctx *NestedParenthesesBlockContext) interface{}

	// Visit a parse tree produced by tinycParser#parameterTypeList.
	VisitParameterTypeList(ctx *ParameterTypeListContext) interface{}

	// Visit a parse tree produced by tinycParser#parameterList.
	VisitParameterList(ctx *ParameterListContext) interface{}

	// Visit a parse tree produced by tinycParser#parameterDeclaration.
	VisitParameterDeclaration(ctx *ParameterDeclarationContext) interface{}

	// Visit a parse tree produced by tinycParser#identifierList.
	VisitIdentifierList(ctx *IdentifierListContext) interface{}

	// Visit a parse tree produced by tinycParser#typeName.
	VisitTypeName(ctx *TypeNameContext) interface{}

	// Visit a parse tree produced by tinycParser#abstractDeclarator.
	VisitAbstractDeclarator(ctx *AbstractDeclaratorContext) interface{}

	// Visit a parse tree produced by tinycParser#initializer.
	VisitInitializer(ctx *InitializerContext) interface{}

	// Visit a parse tree produced by tinycParser#initializerList.
	VisitInitializerList(ctx *InitializerListContext) interface{}

	// Visit a parse tree produced by tinycParser#designation.
	VisitDesignation(ctx *DesignationContext) interface{}

	// Visit a parse tree produced by tinycParser#designatorList.
	VisitDesignatorList(ctx *DesignatorListContext) interface{}

	// Visit a parse tree produced by tinycParser#designator.
	VisitDesignator(ctx *DesignatorContext) interface{}

	// Visit a parse tree produced by tinycParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by tinycParser#labeledStatement.
	VisitLabeledStatement(ctx *LabeledStatementContext) interface{}

	// Visit a parse tree produced by tinycParser#compoundStatement.
	VisitCompoundStatement(ctx *CompoundStatementContext) interface{}

	// Visit a parse tree produced by tinycParser#blockItem.
	VisitBlockItem(ctx *BlockItemContext) interface{}

	// Visit a parse tree produced by tinycParser#expressionStatement.
	VisitExpressionStatement(ctx *ExpressionStatementContext) interface{}

	// Visit a parse tree produced by tinycParser#selectionStatement.
	VisitSelectionStatement(ctx *SelectionStatementContext) interface{}

	// Visit a parse tree produced by tinycParser#iterationStatement.
	VisitIterationStatement(ctx *IterationStatementContext) interface{}

	// Visit a parse tree produced by tinycParser#jumpStatement.
	VisitJumpStatement(ctx *JumpStatementContext) interface{}

	// Visit a parse tree produced by tinycParser#compilationUnit.
	VisitCompilationUnit(ctx *CompilationUnitContext) interface{}

	// Visit a parse tree produced by tinycParser#externalDeclaration.
	VisitExternalDeclaration(ctx *ExternalDeclarationContext) interface{}

	// Visit a parse tree produced by tinycParser#functionDefinition.
	VisitFunctionDefinition(ctx *FunctionDefinitionContext) interface{}

	// Visit a parse tree produced by tinycParser#declarationList.
	VisitDeclarationList(ctx *DeclarationListContext) interface{}
}
