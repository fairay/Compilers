// Code generated from tinyc.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // tinyc
import "github.com/antlr4-go/antlr/v4"

// BasetinycListener is a complete listener for a parse tree produced by tinycParser.
type BasetinycListener struct{}

var _ tinycListener = &BasetinycListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasetinycListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasetinycListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasetinycListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasetinycListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPrimaryExpression is called when production primaryExpression is entered.
func (s *BasetinycListener) EnterPrimaryExpression(ctx *PrimaryExpressionContext) {}

// ExitPrimaryExpression is called when production primaryExpression is exited.
func (s *BasetinycListener) ExitPrimaryExpression(ctx *PrimaryExpressionContext) {}

// EnterGenericAssocList is called when production genericAssocList is entered.
func (s *BasetinycListener) EnterGenericAssocList(ctx *GenericAssocListContext) {}

// ExitGenericAssocList is called when production genericAssocList is exited.
func (s *BasetinycListener) ExitGenericAssocList(ctx *GenericAssocListContext) {}

// EnterGenericAssociation is called when production genericAssociation is entered.
func (s *BasetinycListener) EnterGenericAssociation(ctx *GenericAssociationContext) {}

// ExitGenericAssociation is called when production genericAssociation is exited.
func (s *BasetinycListener) ExitGenericAssociation(ctx *GenericAssociationContext) {}

// EnterPostfixExpression is called when production postfixExpression is entered.
func (s *BasetinycListener) EnterPostfixExpression(ctx *PostfixExpressionContext) {}

// ExitPostfixExpression is called when production postfixExpression is exited.
func (s *BasetinycListener) ExitPostfixExpression(ctx *PostfixExpressionContext) {}

// EnterArgumentExpressionList is called when production argumentExpressionList is entered.
func (s *BasetinycListener) EnterArgumentExpressionList(ctx *ArgumentExpressionListContext) {}

// ExitArgumentExpressionList is called when production argumentExpressionList is exited.
func (s *BasetinycListener) ExitArgumentExpressionList(ctx *ArgumentExpressionListContext) {}

// EnterUnaryExpression is called when production unaryExpression is entered.
func (s *BasetinycListener) EnterUnaryExpression(ctx *UnaryExpressionContext) {}

// ExitUnaryExpression is called when production unaryExpression is exited.
func (s *BasetinycListener) ExitUnaryExpression(ctx *UnaryExpressionContext) {}

// EnterUnaryOperator is called when production unaryOperator is entered.
func (s *BasetinycListener) EnterUnaryOperator(ctx *UnaryOperatorContext) {}

// ExitUnaryOperator is called when production unaryOperator is exited.
func (s *BasetinycListener) ExitUnaryOperator(ctx *UnaryOperatorContext) {}

// EnterCastExpression is called when production castExpression is entered.
func (s *BasetinycListener) EnterCastExpression(ctx *CastExpressionContext) {}

// ExitCastExpression is called when production castExpression is exited.
func (s *BasetinycListener) ExitCastExpression(ctx *CastExpressionContext) {}

// EnterMultiplicativeExpression is called when production multiplicativeExpression is entered.
func (s *BasetinycListener) EnterMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {}

// ExitMultiplicativeExpression is called when production multiplicativeExpression is exited.
func (s *BasetinycListener) ExitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {}

// EnterAdditiveExpression is called when production additiveExpression is entered.
func (s *BasetinycListener) EnterAdditiveExpression(ctx *AdditiveExpressionContext) {}

// ExitAdditiveExpression is called when production additiveExpression is exited.
func (s *BasetinycListener) ExitAdditiveExpression(ctx *AdditiveExpressionContext) {}

// EnterRelationalExpression is called when production relationalExpression is entered.
func (s *BasetinycListener) EnterRelationalExpression(ctx *RelationalExpressionContext) {}

// ExitRelationalExpression is called when production relationalExpression is exited.
func (s *BasetinycListener) ExitRelationalExpression(ctx *RelationalExpressionContext) {}

// EnterEqualityExpression is called when production equalityExpression is entered.
func (s *BasetinycListener) EnterEqualityExpression(ctx *EqualityExpressionContext) {}

// ExitEqualityExpression is called when production equalityExpression is exited.
func (s *BasetinycListener) ExitEqualityExpression(ctx *EqualityExpressionContext) {}

// EnterInclusiveOrExpression is called when production inclusiveOrExpression is entered.
func (s *BasetinycListener) EnterInclusiveOrExpression(ctx *InclusiveOrExpressionContext) {}

// ExitInclusiveOrExpression is called when production inclusiveOrExpression is exited.
func (s *BasetinycListener) ExitInclusiveOrExpression(ctx *InclusiveOrExpressionContext) {}

// EnterLogicalAndExpression is called when production logicalAndExpression is entered.
func (s *BasetinycListener) EnterLogicalAndExpression(ctx *LogicalAndExpressionContext) {}

// ExitLogicalAndExpression is called when production logicalAndExpression is exited.
func (s *BasetinycListener) ExitLogicalAndExpression(ctx *LogicalAndExpressionContext) {}

// EnterLogicalOrExpression is called when production logicalOrExpression is entered.
func (s *BasetinycListener) EnterLogicalOrExpression(ctx *LogicalOrExpressionContext) {}

// ExitLogicalOrExpression is called when production logicalOrExpression is exited.
func (s *BasetinycListener) ExitLogicalOrExpression(ctx *LogicalOrExpressionContext) {}

// EnterAssignmentExpression is called when production assignmentExpression is entered.
func (s *BasetinycListener) EnterAssignmentExpression(ctx *AssignmentExpressionContext) {}

// ExitAssignmentExpression is called when production assignmentExpression is exited.
func (s *BasetinycListener) ExitAssignmentExpression(ctx *AssignmentExpressionContext) {}

// EnterAssignmentOperator is called when production assignmentOperator is entered.
func (s *BasetinycListener) EnterAssignmentOperator(ctx *AssignmentOperatorContext) {}

// ExitAssignmentOperator is called when production assignmentOperator is exited.
func (s *BasetinycListener) ExitAssignmentOperator(ctx *AssignmentOperatorContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasetinycListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasetinycListener) ExitExpression(ctx *ExpressionContext) {}

// EnterConstantExpression is called when production constantExpression is entered.
func (s *BasetinycListener) EnterConstantExpression(ctx *ConstantExpressionContext) {}

// ExitConstantExpression is called when production constantExpression is exited.
func (s *BasetinycListener) ExitConstantExpression(ctx *ConstantExpressionContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BasetinycListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BasetinycListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterDeclarationSpecifiers is called when production declarationSpecifiers is entered.
func (s *BasetinycListener) EnterDeclarationSpecifiers(ctx *DeclarationSpecifiersContext) {}

// ExitDeclarationSpecifiers is called when production declarationSpecifiers is exited.
func (s *BasetinycListener) ExitDeclarationSpecifiers(ctx *DeclarationSpecifiersContext) {}

// EnterDeclarationSpecifiers2 is called when production declarationSpecifiers2 is entered.
func (s *BasetinycListener) EnterDeclarationSpecifiers2(ctx *DeclarationSpecifiers2Context) {}

// ExitDeclarationSpecifiers2 is called when production declarationSpecifiers2 is exited.
func (s *BasetinycListener) ExitDeclarationSpecifiers2(ctx *DeclarationSpecifiers2Context) {}

// EnterDeclarationSpecifier is called when production declarationSpecifier is entered.
func (s *BasetinycListener) EnterDeclarationSpecifier(ctx *DeclarationSpecifierContext) {}

// ExitDeclarationSpecifier is called when production declarationSpecifier is exited.
func (s *BasetinycListener) ExitDeclarationSpecifier(ctx *DeclarationSpecifierContext) {}

// EnterInitDeclaratorList is called when production initDeclaratorList is entered.
func (s *BasetinycListener) EnterInitDeclaratorList(ctx *InitDeclaratorListContext) {}

// ExitInitDeclaratorList is called when production initDeclaratorList is exited.
func (s *BasetinycListener) ExitInitDeclaratorList(ctx *InitDeclaratorListContext) {}

// EnterInitDeclarator is called when production initDeclarator is entered.
func (s *BasetinycListener) EnterInitDeclarator(ctx *InitDeclaratorContext) {}

// ExitInitDeclarator is called when production initDeclarator is exited.
func (s *BasetinycListener) ExitInitDeclarator(ctx *InitDeclaratorContext) {}

// EnterTypeSpecifier is called when production typeSpecifier is entered.
func (s *BasetinycListener) EnterTypeSpecifier(ctx *TypeSpecifierContext) {}

// ExitTypeSpecifier is called when production typeSpecifier is exited.
func (s *BasetinycListener) ExitTypeSpecifier(ctx *TypeSpecifierContext) {}

// EnterStructOrUnionSpecifier is called when production structOrUnionSpecifier is entered.
func (s *BasetinycListener) EnterStructOrUnionSpecifier(ctx *StructOrUnionSpecifierContext) {}

// ExitStructOrUnionSpecifier is called when production structOrUnionSpecifier is exited.
func (s *BasetinycListener) ExitStructOrUnionSpecifier(ctx *StructOrUnionSpecifierContext) {}

// EnterStructOrUnion is called when production structOrUnion is entered.
func (s *BasetinycListener) EnterStructOrUnion(ctx *StructOrUnionContext) {}

// ExitStructOrUnion is called when production structOrUnion is exited.
func (s *BasetinycListener) ExitStructOrUnion(ctx *StructOrUnionContext) {}

// EnterStructDeclarationList is called when production structDeclarationList is entered.
func (s *BasetinycListener) EnterStructDeclarationList(ctx *StructDeclarationListContext) {}

// ExitStructDeclarationList is called when production structDeclarationList is exited.
func (s *BasetinycListener) ExitStructDeclarationList(ctx *StructDeclarationListContext) {}

// EnterStructDeclaration is called when production structDeclaration is entered.
func (s *BasetinycListener) EnterStructDeclaration(ctx *StructDeclarationContext) {}

// ExitStructDeclaration is called when production structDeclaration is exited.
func (s *BasetinycListener) ExitStructDeclaration(ctx *StructDeclarationContext) {}

// EnterSpecifierQualifierList is called when production specifierQualifierList is entered.
func (s *BasetinycListener) EnterSpecifierQualifierList(ctx *SpecifierQualifierListContext) {}

// ExitSpecifierQualifierList is called when production specifierQualifierList is exited.
func (s *BasetinycListener) ExitSpecifierQualifierList(ctx *SpecifierQualifierListContext) {}

// EnterStructDeclaratorList is called when production structDeclaratorList is entered.
func (s *BasetinycListener) EnterStructDeclaratorList(ctx *StructDeclaratorListContext) {}

// ExitStructDeclaratorList is called when production structDeclaratorList is exited.
func (s *BasetinycListener) ExitStructDeclaratorList(ctx *StructDeclaratorListContext) {}

// EnterStructDeclarator is called when production structDeclarator is entered.
func (s *BasetinycListener) EnterStructDeclarator(ctx *StructDeclaratorContext) {}

// ExitStructDeclarator is called when production structDeclarator is exited.
func (s *BasetinycListener) ExitStructDeclarator(ctx *StructDeclaratorContext) {}

// EnterDeclarator is called when production declarator is entered.
func (s *BasetinycListener) EnterDeclarator(ctx *DeclaratorContext) {}

// ExitDeclarator is called when production declarator is exited.
func (s *BasetinycListener) ExitDeclarator(ctx *DeclaratorContext) {}

// EnterDirectDeclarator is called when production directDeclarator is entered.
func (s *BasetinycListener) EnterDirectDeclarator(ctx *DirectDeclaratorContext) {}

// ExitDirectDeclarator is called when production directDeclarator is exited.
func (s *BasetinycListener) ExitDirectDeclarator(ctx *DirectDeclaratorContext) {}

// EnterGccAttribute is called when production gccAttribute is entered.
func (s *BasetinycListener) EnterGccAttribute(ctx *GccAttributeContext) {}

// ExitGccAttribute is called when production gccAttribute is exited.
func (s *BasetinycListener) ExitGccAttribute(ctx *GccAttributeContext) {}

// EnterNestedParenthesesBlock is called when production nestedParenthesesBlock is entered.
func (s *BasetinycListener) EnterNestedParenthesesBlock(ctx *NestedParenthesesBlockContext) {}

// ExitNestedParenthesesBlock is called when production nestedParenthesesBlock is exited.
func (s *BasetinycListener) ExitNestedParenthesesBlock(ctx *NestedParenthesesBlockContext) {}

// EnterPointer is called when production pointer is entered.
func (s *BasetinycListener) EnterPointer(ctx *PointerContext) {}

// ExitPointer is called when production pointer is exited.
func (s *BasetinycListener) ExitPointer(ctx *PointerContext) {}

// EnterParameterTypeList is called when production parameterTypeList is entered.
func (s *BasetinycListener) EnterParameterTypeList(ctx *ParameterTypeListContext) {}

// ExitParameterTypeList is called when production parameterTypeList is exited.
func (s *BasetinycListener) ExitParameterTypeList(ctx *ParameterTypeListContext) {}

// EnterParameterList is called when production parameterList is entered.
func (s *BasetinycListener) EnterParameterList(ctx *ParameterListContext) {}

// ExitParameterList is called when production parameterList is exited.
func (s *BasetinycListener) ExitParameterList(ctx *ParameterListContext) {}

// EnterParameterDeclaration is called when production parameterDeclaration is entered.
func (s *BasetinycListener) EnterParameterDeclaration(ctx *ParameterDeclarationContext) {}

// ExitParameterDeclaration is called when production parameterDeclaration is exited.
func (s *BasetinycListener) ExitParameterDeclaration(ctx *ParameterDeclarationContext) {}

// EnterIdentifierList is called when production identifierList is entered.
func (s *BasetinycListener) EnterIdentifierList(ctx *IdentifierListContext) {}

// ExitIdentifierList is called when production identifierList is exited.
func (s *BasetinycListener) ExitIdentifierList(ctx *IdentifierListContext) {}

// EnterTypeName is called when production typeName is entered.
func (s *BasetinycListener) EnterTypeName(ctx *TypeNameContext) {}

// ExitTypeName is called when production typeName is exited.
func (s *BasetinycListener) ExitTypeName(ctx *TypeNameContext) {}

// EnterAbstractDeclarator is called when production abstractDeclarator is entered.
func (s *BasetinycListener) EnterAbstractDeclarator(ctx *AbstractDeclaratorContext) {}

// ExitAbstractDeclarator is called when production abstractDeclarator is exited.
func (s *BasetinycListener) ExitAbstractDeclarator(ctx *AbstractDeclaratorContext) {}

// EnterDirectAbstractDeclarator is called when production directAbstractDeclarator is entered.
func (s *BasetinycListener) EnterDirectAbstractDeclarator(ctx *DirectAbstractDeclaratorContext) {}

// ExitDirectAbstractDeclarator is called when production directAbstractDeclarator is exited.
func (s *BasetinycListener) ExitDirectAbstractDeclarator(ctx *DirectAbstractDeclaratorContext) {}

// EnterTypedefName is called when production typedefName is entered.
func (s *BasetinycListener) EnterTypedefName(ctx *TypedefNameContext) {}

// ExitTypedefName is called when production typedefName is exited.
func (s *BasetinycListener) ExitTypedefName(ctx *TypedefNameContext) {}

// EnterInitializer is called when production initializer is entered.
func (s *BasetinycListener) EnterInitializer(ctx *InitializerContext) {}

// ExitInitializer is called when production initializer is exited.
func (s *BasetinycListener) ExitInitializer(ctx *InitializerContext) {}

// EnterInitializerList is called when production initializerList is entered.
func (s *BasetinycListener) EnterInitializerList(ctx *InitializerListContext) {}

// ExitInitializerList is called when production initializerList is exited.
func (s *BasetinycListener) ExitInitializerList(ctx *InitializerListContext) {}

// EnterDesignation is called when production designation is entered.
func (s *BasetinycListener) EnterDesignation(ctx *DesignationContext) {}

// ExitDesignation is called when production designation is exited.
func (s *BasetinycListener) ExitDesignation(ctx *DesignationContext) {}

// EnterDesignatorList is called when production designatorList is entered.
func (s *BasetinycListener) EnterDesignatorList(ctx *DesignatorListContext) {}

// ExitDesignatorList is called when production designatorList is exited.
func (s *BasetinycListener) ExitDesignatorList(ctx *DesignatorListContext) {}

// EnterDesignator is called when production designator is entered.
func (s *BasetinycListener) EnterDesignator(ctx *DesignatorContext) {}

// ExitDesignator is called when production designator is exited.
func (s *BasetinycListener) ExitDesignator(ctx *DesignatorContext) {}

// EnterStatement is called when production statement is entered.
func (s *BasetinycListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BasetinycListener) ExitStatement(ctx *StatementContext) {}

// EnterLabeledStatement is called when production labeledStatement is entered.
func (s *BasetinycListener) EnterLabeledStatement(ctx *LabeledStatementContext) {}

// ExitLabeledStatement is called when production labeledStatement is exited.
func (s *BasetinycListener) ExitLabeledStatement(ctx *LabeledStatementContext) {}

// EnterCompoundStatement is called when production compoundStatement is entered.
func (s *BasetinycListener) EnterCompoundStatement(ctx *CompoundStatementContext) {}

// ExitCompoundStatement is called when production compoundStatement is exited.
func (s *BasetinycListener) ExitCompoundStatement(ctx *CompoundStatementContext) {}

// EnterBlockItemList is called when production blockItemList is entered.
func (s *BasetinycListener) EnterBlockItemList(ctx *BlockItemListContext) {}

// ExitBlockItemList is called when production blockItemList is exited.
func (s *BasetinycListener) ExitBlockItemList(ctx *BlockItemListContext) {}

// EnterBlockItem is called when production blockItem is entered.
func (s *BasetinycListener) EnterBlockItem(ctx *BlockItemContext) {}

// ExitBlockItem is called when production blockItem is exited.
func (s *BasetinycListener) ExitBlockItem(ctx *BlockItemContext) {}

// EnterExpressionStatement is called when production expressionStatement is entered.
func (s *BasetinycListener) EnterExpressionStatement(ctx *ExpressionStatementContext) {}

// ExitExpressionStatement is called when production expressionStatement is exited.
func (s *BasetinycListener) ExitExpressionStatement(ctx *ExpressionStatementContext) {}

// EnterSelectionStatement is called when production selectionStatement is entered.
func (s *BasetinycListener) EnterSelectionStatement(ctx *SelectionStatementContext) {}

// ExitSelectionStatement is called when production selectionStatement is exited.
func (s *BasetinycListener) ExitSelectionStatement(ctx *SelectionStatementContext) {}

// EnterIterationStatement is called when production iterationStatement is entered.
func (s *BasetinycListener) EnterIterationStatement(ctx *IterationStatementContext) {}

// ExitIterationStatement is called when production iterationStatement is exited.
func (s *BasetinycListener) ExitIterationStatement(ctx *IterationStatementContext) {}

// EnterForCondition is called when production forCondition is entered.
func (s *BasetinycListener) EnterForCondition(ctx *ForConditionContext) {}

// ExitForCondition is called when production forCondition is exited.
func (s *BasetinycListener) ExitForCondition(ctx *ForConditionContext) {}

// EnterForDeclaration is called when production forDeclaration is entered.
func (s *BasetinycListener) EnterForDeclaration(ctx *ForDeclarationContext) {}

// ExitForDeclaration is called when production forDeclaration is exited.
func (s *BasetinycListener) ExitForDeclaration(ctx *ForDeclarationContext) {}

// EnterForExpression is called when production forExpression is entered.
func (s *BasetinycListener) EnterForExpression(ctx *ForExpressionContext) {}

// ExitForExpression is called when production forExpression is exited.
func (s *BasetinycListener) ExitForExpression(ctx *ForExpressionContext) {}

// EnterJumpStatement is called when production jumpStatement is entered.
func (s *BasetinycListener) EnterJumpStatement(ctx *JumpStatementContext) {}

// ExitJumpStatement is called when production jumpStatement is exited.
func (s *BasetinycListener) ExitJumpStatement(ctx *JumpStatementContext) {}

// EnterCompilationUnit is called when production compilationUnit is entered.
func (s *BasetinycListener) EnterCompilationUnit(ctx *CompilationUnitContext) {}

// ExitCompilationUnit is called when production compilationUnit is exited.
func (s *BasetinycListener) ExitCompilationUnit(ctx *CompilationUnitContext) {}

// EnterExternalDeclaration is called when production externalDeclaration is entered.
func (s *BasetinycListener) EnterExternalDeclaration(ctx *ExternalDeclarationContext) {}

// ExitExternalDeclaration is called when production externalDeclaration is exited.
func (s *BasetinycListener) ExitExternalDeclaration(ctx *ExternalDeclarationContext) {}

// EnterFunctionDefinition is called when production functionDefinition is entered.
func (s *BasetinycListener) EnterFunctionDefinition(ctx *FunctionDefinitionContext) {}

// ExitFunctionDefinition is called when production functionDefinition is exited.
func (s *BasetinycListener) ExitFunctionDefinition(ctx *FunctionDefinitionContext) {}

// EnterDeclarationList is called when production declarationList is entered.
func (s *BasetinycListener) EnterDeclarationList(ctx *DeclarationListContext) {}

// ExitDeclarationList is called when production declarationList is exited.
func (s *BasetinycListener) ExitDeclarationList(ctx *DeclarationListContext) {}
