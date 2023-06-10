// Code generated from tinyc.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // tinyc
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type tinycParser struct {
	*antlr.BaseParser
}

var TinycParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func tinycParserInit() {
	staticData := &TinycParserStaticData
	staticData.LiteralNames = []string{
		"", "'auto'", "'break'", "'case'", "'char'", "'const'", "'continue'",
		"'default'", "'do'", "'double'", "'else'", "'extern'", "'float'", "'for'",
		"'goto'", "'if'", "'inline'", "'int'", "'register'", "'restrict'", "'return'",
		"'short'", "'signed'", "'sizeof'", "'struct'", "'switch'", "'typedef'",
		"'union'", "'void'", "'while'", "'_Noreturn'", "'('", "')'", "'['",
		"']'", "'{'", "'}'", "'<'", "'<='", "'>'", "'>='", "'<<'", "'>>'", "'+'",
		"'++'", "'-'", "'--'", "'*'", "'/'", "'%'", "'&'", "'|'", "'&&'", "'||'",
		"'^'", "'!'", "'~'", "'?'", "':'", "';'", "','", "'='", "'=='", "'!='",
		"'->'", "'.'", "'...'",
	}
	staticData.SymbolicNames = []string{
		"", "Auto", "Break", "Case", "Char", "Const", "Continue", "Default",
		"Do", "Double", "Else", "Extern", "Float", "For", "Goto", "If", "Inline",
		"Int", "Register", "Restrict", "Return", "Short", "Signed", "Sizeof",
		"Struct", "Switch", "Typedef", "Union", "Void", "While", "Noreturn",
		"LeftParen", "RightParen", "LeftBracket", "RightBracket", "LeftBrace",
		"RightBrace", "Less", "LessEqual", "Greater", "GreaterEqual", "LeftShift",
		"RightShift", "Plus", "PlusPlus", "Minus", "MinusMinus", "Star", "Div",
		"Mod", "And", "Or", "AndAnd", "OrOr", "Caret", "Not", "Tilde", "Question",
		"Colon", "Semi", "Comma", "Assign", "Equal", "NotEqual", "Arrow", "Dot",
		"Ellipsis", "Identifier", "Constant", "DigitSequence", "StringLiteral",
		"ComplexDefine", "IncludeDirective", "AsmBlock", "LineAfterPreprocessing",
		"LineDirective", "PragmaDirective", "Whitespace", "Newline", "BlockComment",
		"LineComment",
	}
	staticData.RuleNames = []string{
		"primaryExpression", "genericAssocList", "genericAssociation", "postfixExpression",
		"argumentExpressionList", "unaryExpression", "unaryOperator", "castExpression",
		"multiplicativeExpression", "additiveExpression", "relationalExpression",
		"equalityExpression", "inclusiveOrExpression", "logicalAndExpression",
		"logicalOrExpression", "assignmentExpression", "assignmentOperator",
		"expression", "constantExpression", "declaration", "declarationSpecifiers",
		"declarationSpecifiers2", "declarationSpecifier", "initDeclaratorList",
		"initDeclarator", "typeSpecifier", "structOrUnionSpecifier", "structOrUnion",
		"structDeclarationList", "structDeclaration", "specifierQualifierList",
		"structDeclaratorList", "structDeclarator", "declarator", "directDeclarator",
		"gccAttribute", "nestedParenthesesBlock", "pointer", "parameterTypeList",
		"parameterList", "parameterDeclaration", "identifierList", "typeName",
		"abstractDeclarator", "directAbstractDeclarator", "typedefName", "initializer",
		"initializerList", "designation", "designatorList", "designator", "statement",
		"labeledStatement", "compoundStatement", "blockItemList", "blockItem",
		"expressionStatement", "selectionStatement", "iterationStatement", "forCondition",
		"forDeclaration", "forExpression", "jumpStatement", "compilationUnit",
		"externalDeclaration", "functionDefinition", "declarationList",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 80, 723, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2, 52, 7,
		52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57, 7, 57,
		2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7, 62, 2,
		63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 1, 0, 1, 0, 1, 0,
		4, 0, 138, 8, 0, 11, 0, 12, 0, 139, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 3, 0, 150, 8, 0, 1, 1, 1, 1, 1, 1, 5, 1, 155, 8, 1, 10, 1,
		12, 1, 158, 9, 1, 1, 2, 1, 2, 3, 2, 162, 8, 2, 1, 2, 1, 2, 1, 2, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 174, 8, 3, 1, 3, 1, 3, 3, 3,
		178, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 186, 8, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 5, 3, 193, 8, 3, 10, 3, 12, 3, 196, 9, 3, 1, 4, 1,
		4, 1, 4, 5, 4, 201, 8, 4, 10, 4, 12, 4, 204, 9, 4, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 3, 5, 212, 8, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 3, 7, 223, 8, 7, 1, 8, 1, 8, 1, 8, 5, 8, 228, 8, 8, 10,
		8, 12, 8, 231, 9, 8, 1, 9, 1, 9, 1, 9, 5, 9, 236, 8, 9, 10, 9, 12, 9, 239,
		9, 9, 1, 10, 1, 10, 1, 10, 5, 10, 244, 8, 10, 10, 10, 12, 10, 247, 9, 10,
		1, 11, 1, 11, 1, 11, 5, 11, 252, 8, 11, 10, 11, 12, 11, 255, 9, 11, 1,
		12, 1, 12, 1, 12, 5, 12, 260, 8, 12, 10, 12, 12, 12, 263, 9, 12, 1, 13,
		1, 13, 1, 13, 5, 13, 268, 8, 13, 10, 13, 12, 13, 271, 9, 13, 1, 14, 1,
		14, 1, 14, 5, 14, 276, 8, 14, 10, 14, 12, 14, 279, 9, 14, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 287, 8, 15, 1, 16, 1, 16, 1, 17, 1,
		17, 1, 17, 5, 17, 294, 8, 17, 10, 17, 12, 17, 297, 9, 17, 1, 18, 1, 18,
		1, 19, 1, 19, 3, 19, 303, 8, 19, 1, 19, 1, 19, 1, 20, 4, 20, 308, 8, 20,
		11, 20, 12, 20, 309, 1, 21, 4, 21, 313, 8, 21, 11, 21, 12, 21, 314, 1,
		22, 1, 22, 1, 23, 1, 23, 1, 23, 5, 23, 322, 8, 23, 10, 23, 12, 23, 325,
		9, 23, 1, 24, 1, 24, 1, 24, 3, 24, 330, 8, 24, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 3, 25, 341, 8, 25, 1, 26, 1, 26,
		3, 26, 345, 8, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 3,
		26, 354, 8, 26, 1, 27, 1, 27, 1, 28, 4, 28, 359, 8, 28, 11, 28, 12, 28,
		360, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 3, 29, 370, 8, 29,
		1, 30, 1, 30, 3, 30, 374, 8, 30, 1, 31, 1, 31, 1, 31, 5, 31, 379, 8, 31,
		10, 31, 12, 31, 382, 9, 31, 1, 32, 1, 32, 3, 32, 386, 8, 32, 1, 32, 1,
		32, 3, 32, 390, 8, 32, 1, 33, 3, 33, 393, 8, 33, 1, 33, 1, 33, 1, 34, 1,
		34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 3, 34, 406, 8, 34,
		1, 34, 1, 34, 1, 34, 3, 34, 411, 8, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1,
		34, 1, 34, 1, 34, 1, 34, 1, 34, 3, 34, 422, 8, 34, 1, 34, 5, 34, 425, 8,
		34, 10, 34, 12, 34, 428, 9, 34, 1, 35, 1, 35, 1, 35, 3, 35, 433, 8, 35,
		1, 35, 3, 35, 436, 8, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 5, 36, 443,
		8, 36, 10, 36, 12, 36, 446, 9, 36, 1, 37, 4, 37, 449, 8, 37, 11, 37, 12,
		37, 450, 1, 38, 1, 38, 1, 38, 3, 38, 456, 8, 38, 1, 39, 1, 39, 1, 39, 5,
		39, 461, 8, 39, 10, 39, 12, 39, 464, 9, 39, 1, 40, 1, 40, 1, 40, 1, 40,
		1, 40, 3, 40, 471, 8, 40, 3, 40, 473, 8, 40, 1, 41, 1, 41, 1, 41, 5, 41,
		478, 8, 41, 10, 41, 12, 41, 481, 9, 41, 1, 42, 1, 42, 3, 42, 485, 8, 42,
		1, 43, 1, 43, 3, 43, 489, 8, 43, 1, 43, 3, 43, 492, 8, 43, 1, 44, 1, 44,
		1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 3, 44, 501, 8, 44, 1, 44, 1, 44, 1,
		44, 1, 44, 1, 44, 1, 44, 3, 44, 509, 8, 44, 1, 44, 3, 44, 512, 8, 44, 1,
		44, 1, 44, 1, 44, 3, 44, 517, 8, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44,
		1, 44, 1, 44, 1, 44, 3, 44, 527, 8, 44, 1, 44, 5, 44, 530, 8, 44, 10, 44,
		12, 44, 533, 9, 44, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 1, 46, 3, 46, 541,
		8, 46, 1, 46, 1, 46, 3, 46, 545, 8, 46, 1, 47, 3, 47, 548, 8, 47, 1, 47,
		1, 47, 1, 47, 3, 47, 553, 8, 47, 1, 47, 5, 47, 556, 8, 47, 10, 47, 12,
		47, 559, 9, 47, 1, 48, 1, 48, 1, 48, 1, 49, 4, 49, 565, 8, 49, 11, 49,
		12, 49, 566, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 3, 50, 575, 8, 50,
		1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 3, 51, 583, 8, 51, 1, 52, 1,
		52, 1, 52, 1, 52, 1, 52, 1, 52, 1, 52, 1, 52, 1, 52, 1, 52, 1, 52, 3, 52,
		596, 8, 52, 1, 53, 1, 53, 3, 53, 600, 8, 53, 1, 53, 1, 53, 1, 54, 4, 54,
		605, 8, 54, 11, 54, 12, 54, 606, 1, 55, 1, 55, 3, 55, 611, 8, 55, 1, 56,
		3, 56, 614, 8, 56, 1, 56, 1, 56, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1,
		57, 1, 57, 3, 57, 625, 8, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57,
		3, 57, 633, 8, 57, 1, 58, 1, 58, 1, 58, 1, 58, 1, 58, 1, 58, 1, 58, 1,
		58, 1, 58, 1, 58, 1, 58, 1, 58, 1, 58, 1, 58, 1, 58, 1, 58, 1, 58, 1, 58,
		1, 58, 1, 58, 3, 58, 655, 8, 58, 1, 59, 1, 59, 3, 59, 659, 8, 59, 3, 59,
		661, 8, 59, 1, 59, 1, 59, 3, 59, 665, 8, 59, 1, 59, 1, 59, 3, 59, 669,
		8, 59, 1, 60, 1, 60, 3, 60, 673, 8, 60, 1, 61, 1, 61, 1, 61, 5, 61, 678,
		8, 61, 10, 61, 12, 61, 681, 9, 61, 1, 62, 1, 62, 1, 62, 1, 62, 1, 62, 1,
		62, 3, 62, 689, 8, 62, 1, 62, 1, 62, 3, 62, 693, 8, 62, 1, 62, 1, 62, 1,
		63, 4, 63, 698, 8, 63, 11, 63, 12, 63, 699, 1, 63, 1, 63, 1, 64, 1, 64,
		1, 64, 3, 64, 707, 8, 64, 1, 65, 3, 65, 710, 8, 65, 1, 65, 1, 65, 3, 65,
		714, 8, 65, 1, 65, 1, 65, 1, 66, 4, 66, 719, 8, 66, 11, 66, 12, 66, 720,
		1, 66, 0, 2, 68, 88, 67, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
		26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60,
		62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96,
		98, 100, 102, 104, 106, 108, 110, 112, 114, 116, 118, 120, 122, 124, 126,
		128, 130, 132, 0, 10, 1, 0, 64, 65, 5, 0, 43, 43, 45, 45, 47, 47, 50, 50,
		55, 56, 1, 0, 47, 49, 2, 0, 43, 43, 45, 45, 1, 0, 37, 40, 1, 0, 62, 63,
		2, 0, 24, 24, 27, 27, 2, 0, 31, 32, 60, 60, 1, 0, 31, 32, 2, 0, 47, 47,
		54, 54, 776, 0, 149, 1, 0, 0, 0, 2, 151, 1, 0, 0, 0, 4, 161, 1, 0, 0, 0,
		6, 177, 1, 0, 0, 0, 8, 197, 1, 0, 0, 0, 10, 211, 1, 0, 0, 0, 12, 213, 1,
		0, 0, 0, 14, 222, 1, 0, 0, 0, 16, 224, 1, 0, 0, 0, 18, 232, 1, 0, 0, 0,
		20, 240, 1, 0, 0, 0, 22, 248, 1, 0, 0, 0, 24, 256, 1, 0, 0, 0, 26, 264,
		1, 0, 0, 0, 28, 272, 1, 0, 0, 0, 30, 286, 1, 0, 0, 0, 32, 288, 1, 0, 0,
		0, 34, 290, 1, 0, 0, 0, 36, 298, 1, 0, 0, 0, 38, 300, 1, 0, 0, 0, 40, 307,
		1, 0, 0, 0, 42, 312, 1, 0, 0, 0, 44, 316, 1, 0, 0, 0, 46, 318, 1, 0, 0,
		0, 48, 326, 1, 0, 0, 0, 50, 340, 1, 0, 0, 0, 52, 353, 1, 0, 0, 0, 54, 355,
		1, 0, 0, 0, 56, 358, 1, 0, 0, 0, 58, 369, 1, 0, 0, 0, 60, 371, 1, 0, 0,
		0, 62, 375, 1, 0, 0, 0, 64, 389, 1, 0, 0, 0, 66, 392, 1, 0, 0, 0, 68, 405,
		1, 0, 0, 0, 70, 429, 1, 0, 0, 0, 72, 444, 1, 0, 0, 0, 74, 448, 1, 0, 0,
		0, 76, 452, 1, 0, 0, 0, 78, 457, 1, 0, 0, 0, 80, 472, 1, 0, 0, 0, 82, 474,
		1, 0, 0, 0, 84, 482, 1, 0, 0, 0, 86, 491, 1, 0, 0, 0, 88, 511, 1, 0, 0,
		0, 90, 534, 1, 0, 0, 0, 92, 544, 1, 0, 0, 0, 94, 547, 1, 0, 0, 0, 96, 560,
		1, 0, 0, 0, 98, 564, 1, 0, 0, 0, 100, 574, 1, 0, 0, 0, 102, 582, 1, 0,
		0, 0, 104, 595, 1, 0, 0, 0, 106, 597, 1, 0, 0, 0, 108, 604, 1, 0, 0, 0,
		110, 610, 1, 0, 0, 0, 112, 613, 1, 0, 0, 0, 114, 632, 1, 0, 0, 0, 116,
		654, 1, 0, 0, 0, 118, 660, 1, 0, 0, 0, 120, 670, 1, 0, 0, 0, 122, 674,
		1, 0, 0, 0, 124, 692, 1, 0, 0, 0, 126, 697, 1, 0, 0, 0, 128, 706, 1, 0,
		0, 0, 130, 709, 1, 0, 0, 0, 132, 718, 1, 0, 0, 0, 134, 150, 5, 67, 0, 0,
		135, 150, 5, 68, 0, 0, 136, 138, 5, 70, 0, 0, 137, 136, 1, 0, 0, 0, 138,
		139, 1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0, 140, 150,
		1, 0, 0, 0, 141, 142, 5, 31, 0, 0, 142, 143, 3, 34, 17, 0, 143, 144, 5,
		32, 0, 0, 144, 150, 1, 0, 0, 0, 145, 146, 5, 31, 0, 0, 146, 147, 3, 106,
		53, 0, 147, 148, 5, 32, 0, 0, 148, 150, 1, 0, 0, 0, 149, 134, 1, 0, 0,
		0, 149, 135, 1, 0, 0, 0, 149, 137, 1, 0, 0, 0, 149, 141, 1, 0, 0, 0, 149,
		145, 1, 0, 0, 0, 150, 1, 1, 0, 0, 0, 151, 156, 3, 4, 2, 0, 152, 153, 5,
		60, 0, 0, 153, 155, 3, 4, 2, 0, 154, 152, 1, 0, 0, 0, 155, 158, 1, 0, 0,
		0, 156, 154, 1, 0, 0, 0, 156, 157, 1, 0, 0, 0, 157, 3, 1, 0, 0, 0, 158,
		156, 1, 0, 0, 0, 159, 162, 3, 84, 42, 0, 160, 162, 5, 7, 0, 0, 161, 159,
		1, 0, 0, 0, 161, 160, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 164, 5, 58,
		0, 0, 164, 165, 3, 30, 15, 0, 165, 5, 1, 0, 0, 0, 166, 178, 3, 0, 0, 0,
		167, 168, 5, 31, 0, 0, 168, 169, 3, 84, 42, 0, 169, 170, 5, 32, 0, 0, 170,
		171, 5, 35, 0, 0, 171, 173, 3, 94, 47, 0, 172, 174, 5, 60, 0, 0, 173, 172,
		1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 175, 176, 5, 36,
		0, 0, 176, 178, 1, 0, 0, 0, 177, 166, 1, 0, 0, 0, 177, 167, 1, 0, 0, 0,
		178, 194, 1, 0, 0, 0, 179, 180, 5, 33, 0, 0, 180, 181, 3, 34, 17, 0, 181,
		182, 5, 34, 0, 0, 182, 193, 1, 0, 0, 0, 183, 185, 5, 31, 0, 0, 184, 186,
		3, 8, 4, 0, 185, 184, 1, 0, 0, 0, 185, 186, 1, 0, 0, 0, 186, 187, 1, 0,
		0, 0, 187, 193, 5, 32, 0, 0, 188, 189, 7, 0, 0, 0, 189, 193, 5, 67, 0,
		0, 190, 193, 5, 44, 0, 0, 191, 193, 5, 46, 0, 0, 192, 179, 1, 0, 0, 0,
		192, 183, 1, 0, 0, 0, 192, 188, 1, 0, 0, 0, 192, 190, 1, 0, 0, 0, 192,
		191, 1, 0, 0, 0, 193, 196, 1, 0, 0, 0, 194, 192, 1, 0, 0, 0, 194, 195,
		1, 0, 0, 0, 195, 7, 1, 0, 0, 0, 196, 194, 1, 0, 0, 0, 197, 202, 3, 30,
		15, 0, 198, 199, 5, 60, 0, 0, 199, 201, 3, 30, 15, 0, 200, 198, 1, 0, 0,
		0, 201, 204, 1, 0, 0, 0, 202, 200, 1, 0, 0, 0, 202, 203, 1, 0, 0, 0, 203,
		9, 1, 0, 0, 0, 204, 202, 1, 0, 0, 0, 205, 212, 3, 6, 3, 0, 206, 207, 3,
		12, 6, 0, 207, 208, 3, 14, 7, 0, 208, 212, 1, 0, 0, 0, 209, 210, 5, 52,
		0, 0, 210, 212, 5, 67, 0, 0, 211, 205, 1, 0, 0, 0, 211, 206, 1, 0, 0, 0,
		211, 209, 1, 0, 0, 0, 212, 11, 1, 0, 0, 0, 213, 214, 7, 1, 0, 0, 214, 13,
		1, 0, 0, 0, 215, 216, 5, 31, 0, 0, 216, 217, 3, 84, 42, 0, 217, 218, 5,
		32, 0, 0, 218, 219, 3, 14, 7, 0, 219, 223, 1, 0, 0, 0, 220, 223, 3, 10,
		5, 0, 221, 223, 5, 69, 0, 0, 222, 215, 1, 0, 0, 0, 222, 220, 1, 0, 0, 0,
		222, 221, 1, 0, 0, 0, 223, 15, 1, 0, 0, 0, 224, 229, 3, 14, 7, 0, 225,
		226, 7, 2, 0, 0, 226, 228, 3, 14, 7, 0, 227, 225, 1, 0, 0, 0, 228, 231,
		1, 0, 0, 0, 229, 227, 1, 0, 0, 0, 229, 230, 1, 0, 0, 0, 230, 17, 1, 0,
		0, 0, 231, 229, 1, 0, 0, 0, 232, 237, 3, 16, 8, 0, 233, 234, 7, 3, 0, 0,
		234, 236, 3, 16, 8, 0, 235, 233, 1, 0, 0, 0, 236, 239, 1, 0, 0, 0, 237,
		235, 1, 0, 0, 0, 237, 238, 1, 0, 0, 0, 238, 19, 1, 0, 0, 0, 239, 237, 1,
		0, 0, 0, 240, 245, 3, 18, 9, 0, 241, 242, 7, 4, 0, 0, 242, 244, 3, 18,
		9, 0, 243, 241, 1, 0, 0, 0, 244, 247, 1, 0, 0, 0, 245, 243, 1, 0, 0, 0,
		245, 246, 1, 0, 0, 0, 246, 21, 1, 0, 0, 0, 247, 245, 1, 0, 0, 0, 248, 253,
		3, 20, 10, 0, 249, 250, 7, 5, 0, 0, 250, 252, 3, 20, 10, 0, 251, 249, 1,
		0, 0, 0, 252, 255, 1, 0, 0, 0, 253, 251, 1, 0, 0, 0, 253, 254, 1, 0, 0,
		0, 254, 23, 1, 0, 0, 0, 255, 253, 1, 0, 0, 0, 256, 261, 3, 22, 11, 0, 257,
		258, 5, 51, 0, 0, 258, 260, 3, 22, 11, 0, 259, 257, 1, 0, 0, 0, 260, 263,
		1, 0, 0, 0, 261, 259, 1, 0, 0, 0, 261, 262, 1, 0, 0, 0, 262, 25, 1, 0,
		0, 0, 263, 261, 1, 0, 0, 0, 264, 269, 3, 24, 12, 0, 265, 266, 5, 52, 0,
		0, 266, 268, 3, 24, 12, 0, 267, 265, 1, 0, 0, 0, 268, 271, 1, 0, 0, 0,
		269, 267, 1, 0, 0, 0, 269, 270, 1, 0, 0, 0, 270, 27, 1, 0, 0, 0, 271, 269,
		1, 0, 0, 0, 272, 277, 3, 26, 13, 0, 273, 274, 5, 53, 0, 0, 274, 276, 3,
		26, 13, 0, 275, 273, 1, 0, 0, 0, 276, 279, 1, 0, 0, 0, 277, 275, 1, 0,
		0, 0, 277, 278, 1, 0, 0, 0, 278, 29, 1, 0, 0, 0, 279, 277, 1, 0, 0, 0,
		280, 287, 3, 28, 14, 0, 281, 282, 3, 10, 5, 0, 282, 283, 3, 32, 16, 0,
		283, 284, 3, 30, 15, 0, 284, 287, 1, 0, 0, 0, 285, 287, 5, 69, 0, 0, 286,
		280, 1, 0, 0, 0, 286, 281, 1, 0, 0, 0, 286, 285, 1, 0, 0, 0, 287, 31, 1,
		0, 0, 0, 288, 289, 5, 61, 0, 0, 289, 33, 1, 0, 0, 0, 290, 295, 3, 30, 15,
		0, 291, 292, 5, 60, 0, 0, 292, 294, 3, 30, 15, 0, 293, 291, 1, 0, 0, 0,
		294, 297, 1, 0, 0, 0, 295, 293, 1, 0, 0, 0, 295, 296, 1, 0, 0, 0, 296,
		35, 1, 0, 0, 0, 297, 295, 1, 0, 0, 0, 298, 299, 3, 28, 14, 0, 299, 37,
		1, 0, 0, 0, 300, 302, 3, 40, 20, 0, 301, 303, 3, 46, 23, 0, 302, 301, 1,
		0, 0, 0, 302, 303, 1, 0, 0, 0, 303, 304, 1, 0, 0, 0, 304, 305, 5, 59, 0,
		0, 305, 39, 1, 0, 0, 0, 306, 308, 3, 44, 22, 0, 307, 306, 1, 0, 0, 0, 308,
		309, 1, 0, 0, 0, 309, 307, 1, 0, 0, 0, 309, 310, 1, 0, 0, 0, 310, 41, 1,
		0, 0, 0, 311, 313, 3, 44, 22, 0, 312, 311, 1, 0, 0, 0, 313, 314, 1, 0,
		0, 0, 314, 312, 1, 0, 0, 0, 314, 315, 1, 0, 0, 0, 315, 43, 1, 0, 0, 0,
		316, 317, 3, 50, 25, 0, 317, 45, 1, 0, 0, 0, 318, 323, 3, 48, 24, 0, 319,
		320, 5, 60, 0, 0, 320, 322, 3, 48, 24, 0, 321, 319, 1, 0, 0, 0, 322, 325,
		1, 0, 0, 0, 323, 321, 1, 0, 0, 0, 323, 324, 1, 0, 0, 0, 324, 47, 1, 0,
		0, 0, 325, 323, 1, 0, 0, 0, 326, 329, 3, 66, 33, 0, 327, 328, 5, 61, 0,
		0, 328, 330, 3, 92, 46, 0, 329, 327, 1, 0, 0, 0, 329, 330, 1, 0, 0, 0,
		330, 49, 1, 0, 0, 0, 331, 341, 5, 28, 0, 0, 332, 341, 5, 4, 0, 0, 333,
		341, 5, 21, 0, 0, 334, 341, 5, 17, 0, 0, 335, 341, 5, 12, 0, 0, 336, 341,
		5, 9, 0, 0, 337, 341, 5, 22, 0, 0, 338, 341, 3, 52, 26, 0, 339, 341, 3,
		90, 45, 0, 340, 331, 1, 0, 0, 0, 340, 332, 1, 0, 0, 0, 340, 333, 1, 0,
		0, 0, 340, 334, 1, 0, 0, 0, 340, 335, 1, 0, 0, 0, 340, 336, 1, 0, 0, 0,
		340, 337, 1, 0, 0, 0, 340, 338, 1, 0, 0, 0, 340, 339, 1, 0, 0, 0, 341,
		51, 1, 0, 0, 0, 342, 344, 3, 54, 27, 0, 343, 345, 5, 67, 0, 0, 344, 343,
		1, 0, 0, 0, 344, 345, 1, 0, 0, 0, 345, 346, 1, 0, 0, 0, 346, 347, 5, 35,
		0, 0, 347, 348, 3, 56, 28, 0, 348, 349, 5, 36, 0, 0, 349, 354, 1, 0, 0,
		0, 350, 351, 3, 54, 27, 0, 351, 352, 5, 67, 0, 0, 352, 354, 1, 0, 0, 0,
		353, 342, 1, 0, 0, 0, 353, 350, 1, 0, 0, 0, 354, 53, 1, 0, 0, 0, 355, 356,
		7, 6, 0, 0, 356, 55, 1, 0, 0, 0, 357, 359, 3, 58, 29, 0, 358, 357, 1, 0,
		0, 0, 359, 360, 1, 0, 0, 0, 360, 358, 1, 0, 0, 0, 360, 361, 1, 0, 0, 0,
		361, 57, 1, 0, 0, 0, 362, 363, 3, 60, 30, 0, 363, 364, 3, 62, 31, 0, 364,
		365, 5, 59, 0, 0, 365, 370, 1, 0, 0, 0, 366, 367, 3, 60, 30, 0, 367, 368,
		5, 59, 0, 0, 368, 370, 1, 0, 0, 0, 369, 362, 1, 0, 0, 0, 369, 366, 1, 0,
		0, 0, 370, 59, 1, 0, 0, 0, 371, 373, 3, 50, 25, 0, 372, 374, 3, 60, 30,
		0, 373, 372, 1, 0, 0, 0, 373, 374, 1, 0, 0, 0, 374, 61, 1, 0, 0, 0, 375,
		380, 3, 64, 32, 0, 376, 377, 5, 60, 0, 0, 377, 379, 3, 64, 32, 0, 378,
		376, 1, 0, 0, 0, 379, 382, 1, 0, 0, 0, 380, 378, 1, 0, 0, 0, 380, 381,
		1, 0, 0, 0, 381, 63, 1, 0, 0, 0, 382, 380, 1, 0, 0, 0, 383, 390, 3, 66,
		33, 0, 384, 386, 3, 66, 33, 0, 385, 384, 1, 0, 0, 0, 385, 386, 1, 0, 0,
		0, 386, 387, 1, 0, 0, 0, 387, 388, 5, 58, 0, 0, 388, 390, 3, 36, 18, 0,
		389, 383, 1, 0, 0, 0, 389, 385, 1, 0, 0, 0, 390, 65, 1, 0, 0, 0, 391, 393,
		3, 74, 37, 0, 392, 391, 1, 0, 0, 0, 392, 393, 1, 0, 0, 0, 393, 394, 1,
		0, 0, 0, 394, 395, 3, 68, 34, 0, 395, 67, 1, 0, 0, 0, 396, 397, 6, 34,
		-1, 0, 397, 406, 5, 67, 0, 0, 398, 399, 5, 31, 0, 0, 399, 400, 3, 66, 33,
		0, 400, 401, 5, 32, 0, 0, 401, 406, 1, 0, 0, 0, 402, 403, 5, 67, 0, 0,
		403, 404, 5, 58, 0, 0, 404, 406, 5, 69, 0, 0, 405, 396, 1, 0, 0, 0, 405,
		398, 1, 0, 0, 0, 405, 402, 1, 0, 0, 0, 406, 426, 1, 0, 0, 0, 407, 408,
		10, 4, 0, 0, 408, 410, 5, 33, 0, 0, 409, 411, 3, 30, 15, 0, 410, 409, 1,
		0, 0, 0, 410, 411, 1, 0, 0, 0, 411, 412, 1, 0, 0, 0, 412, 425, 5, 34, 0,
		0, 413, 414, 10, 3, 0, 0, 414, 415, 5, 31, 0, 0, 415, 416, 3, 76, 38, 0,
		416, 417, 5, 32, 0, 0, 417, 425, 1, 0, 0, 0, 418, 419, 10, 2, 0, 0, 419,
		421, 5, 31, 0, 0, 420, 422, 3, 82, 41, 0, 421, 420, 1, 0, 0, 0, 421, 422,
		1, 0, 0, 0, 422, 423, 1, 0, 0, 0, 423, 425, 5, 32, 0, 0, 424, 407, 1, 0,
		0, 0, 424, 413, 1, 0, 0, 0, 424, 418, 1, 0, 0, 0, 425, 428, 1, 0, 0, 0,
		426, 424, 1, 0, 0, 0, 426, 427, 1, 0, 0, 0, 427, 69, 1, 0, 0, 0, 428, 426,
		1, 0, 0, 0, 429, 435, 8, 7, 0, 0, 430, 432, 5, 31, 0, 0, 431, 433, 3, 8,
		4, 0, 432, 431, 1, 0, 0, 0, 432, 433, 1, 0, 0, 0, 433, 434, 1, 0, 0, 0,
		434, 436, 5, 32, 0, 0, 435, 430, 1, 0, 0, 0, 435, 436, 1, 0, 0, 0, 436,
		71, 1, 0, 0, 0, 437, 443, 8, 8, 0, 0, 438, 439, 5, 31, 0, 0, 439, 440,
		3, 72, 36, 0, 440, 441, 5, 32, 0, 0, 441, 443, 1, 0, 0, 0, 442, 437, 1,
		0, 0, 0, 442, 438, 1, 0, 0, 0, 443, 446, 1, 0, 0, 0, 444, 442, 1, 0, 0,
		0, 444, 445, 1, 0, 0, 0, 445, 73, 1, 0, 0, 0, 446, 444, 1, 0, 0, 0, 447,
		449, 7, 9, 0, 0, 448, 447, 1, 0, 0, 0, 449, 450, 1, 0, 0, 0, 450, 448,
		1, 0, 0, 0, 450, 451, 1, 0, 0, 0, 451, 75, 1, 0, 0, 0, 452, 455, 3, 78,
		39, 0, 453, 454, 5, 60, 0, 0, 454, 456, 5, 66, 0, 0, 455, 453, 1, 0, 0,
		0, 455, 456, 1, 0, 0, 0, 456, 77, 1, 0, 0, 0, 457, 462, 3, 80, 40, 0, 458,
		459, 5, 60, 0, 0, 459, 461, 3, 80, 40, 0, 460, 458, 1, 0, 0, 0, 461, 464,
		1, 0, 0, 0, 462, 460, 1, 0, 0, 0, 462, 463, 1, 0, 0, 0, 463, 79, 1, 0,
		0, 0, 464, 462, 1, 0, 0, 0, 465, 466, 3, 40, 20, 0, 466, 467, 3, 66, 33,
		0, 467, 473, 1, 0, 0, 0, 468, 470, 3, 42, 21, 0, 469, 471, 3, 86, 43, 0,
		470, 469, 1, 0, 0, 0, 470, 471, 1, 0, 0, 0, 471, 473, 1, 0, 0, 0, 472,
		465, 1, 0, 0, 0, 472, 468, 1, 0, 0, 0, 473, 81, 1, 0, 0, 0, 474, 479, 5,
		67, 0, 0, 475, 476, 5, 60, 0, 0, 476, 478, 5, 67, 0, 0, 477, 475, 1, 0,
		0, 0, 478, 481, 1, 0, 0, 0, 479, 477, 1, 0, 0, 0, 479, 480, 1, 0, 0, 0,
		480, 83, 1, 0, 0, 0, 481, 479, 1, 0, 0, 0, 482, 484, 3, 60, 30, 0, 483,
		485, 3, 86, 43, 0, 484, 483, 1, 0, 0, 0, 484, 485, 1, 0, 0, 0, 485, 85,
		1, 0, 0, 0, 486, 492, 3, 74, 37, 0, 487, 489, 3, 74, 37, 0, 488, 487, 1,
		0, 0, 0, 488, 489, 1, 0, 0, 0, 489, 490, 1, 0, 0, 0, 490, 492, 3, 88, 44,
		0, 491, 486, 1, 0, 0, 0, 491, 488, 1, 0, 0, 0, 492, 87, 1, 0, 0, 0, 493,
		494, 6, 44, -1, 0, 494, 495, 5, 31, 0, 0, 495, 496, 3, 86, 43, 0, 496,
		497, 5, 32, 0, 0, 497, 512, 1, 0, 0, 0, 498, 500, 5, 33, 0, 0, 499, 501,
		3, 30, 15, 0, 500, 499, 1, 0, 0, 0, 500, 501, 1, 0, 0, 0, 501, 502, 1,
		0, 0, 0, 502, 512, 5, 34, 0, 0, 503, 504, 5, 33, 0, 0, 504, 505, 5, 47,
		0, 0, 505, 512, 5, 34, 0, 0, 506, 508, 5, 31, 0, 0, 507, 509, 3, 76, 38,
		0, 508, 507, 1, 0, 0, 0, 508, 509, 1, 0, 0, 0, 509, 510, 1, 0, 0, 0, 510,
		512, 5, 32, 0, 0, 511, 493, 1, 0, 0, 0, 511, 498, 1, 0, 0, 0, 511, 503,
		1, 0, 0, 0, 511, 506, 1, 0, 0, 0, 512, 531, 1, 0, 0, 0, 513, 514, 10, 3,
		0, 0, 514, 516, 5, 33, 0, 0, 515, 517, 3, 30, 15, 0, 516, 515, 1, 0, 0,
		0, 516, 517, 1, 0, 0, 0, 517, 518, 1, 0, 0, 0, 518, 530, 5, 34, 0, 0, 519,
		520, 10, 2, 0, 0, 520, 521, 5, 33, 0, 0, 521, 522, 5, 47, 0, 0, 522, 530,
		5, 34, 0, 0, 523, 524, 10, 1, 0, 0, 524, 526, 5, 31, 0, 0, 525, 527, 3,
		76, 38, 0, 526, 525, 1, 0, 0, 0, 526, 527, 1, 0, 0, 0, 527, 528, 1, 0,
		0, 0, 528, 530, 5, 32, 0, 0, 529, 513, 1, 0, 0, 0, 529, 519, 1, 0, 0, 0,
		529, 523, 1, 0, 0, 0, 530, 533, 1, 0, 0, 0, 531, 529, 1, 0, 0, 0, 531,
		532, 1, 0, 0, 0, 532, 89, 1, 0, 0, 0, 533, 531, 1, 0, 0, 0, 534, 535, 5,
		67, 0, 0, 535, 91, 1, 0, 0, 0, 536, 545, 3, 30, 15, 0, 537, 538, 5, 35,
		0, 0, 538, 540, 3, 94, 47, 0, 539, 541, 5, 60, 0, 0, 540, 539, 1, 0, 0,
		0, 540, 541, 1, 0, 0, 0, 541, 542, 1, 0, 0, 0, 542, 543, 5, 36, 0, 0, 543,
		545, 1, 0, 0, 0, 544, 536, 1, 0, 0, 0, 544, 537, 1, 0, 0, 0, 545, 93, 1,
		0, 0, 0, 546, 548, 3, 96, 48, 0, 547, 546, 1, 0, 0, 0, 547, 548, 1, 0,
		0, 0, 548, 549, 1, 0, 0, 0, 549, 557, 3, 92, 46, 0, 550, 552, 5, 60, 0,
		0, 551, 553, 3, 96, 48, 0, 552, 551, 1, 0, 0, 0, 552, 553, 1, 0, 0, 0,
		553, 554, 1, 0, 0, 0, 554, 556, 3, 92, 46, 0, 555, 550, 1, 0, 0, 0, 556,
		559, 1, 0, 0, 0, 557, 555, 1, 0, 0, 0, 557, 558, 1, 0, 0, 0, 558, 95, 1,
		0, 0, 0, 559, 557, 1, 0, 0, 0, 560, 561, 3, 98, 49, 0, 561, 562, 5, 61,
		0, 0, 562, 97, 1, 0, 0, 0, 563, 565, 3, 100, 50, 0, 564, 563, 1, 0, 0,
		0, 565, 566, 1, 0, 0, 0, 566, 564, 1, 0, 0, 0, 566, 567, 1, 0, 0, 0, 567,
		99, 1, 0, 0, 0, 568, 569, 5, 33, 0, 0, 569, 570, 3, 36, 18, 0, 570, 571,
		5, 34, 0, 0, 571, 575, 1, 0, 0, 0, 572, 573, 5, 65, 0, 0, 573, 575, 5,
		67, 0, 0, 574, 568, 1, 0, 0, 0, 574, 572, 1, 0, 0, 0, 575, 101, 1, 0, 0,
		0, 576, 583, 3, 104, 52, 0, 577, 583, 3, 106, 53, 0, 578, 583, 3, 112,
		56, 0, 579, 583, 3, 114, 57, 0, 580, 583, 3, 116, 58, 0, 581, 583, 3, 124,
		62, 0, 582, 576, 1, 0, 0, 0, 582, 577, 1, 0, 0, 0, 582, 578, 1, 0, 0, 0,
		582, 579, 1, 0, 0, 0, 582, 580, 1, 0, 0, 0, 582, 581, 1, 0, 0, 0, 583,
		103, 1, 0, 0, 0, 584, 585, 5, 67, 0, 0, 585, 586, 5, 58, 0, 0, 586, 596,
		3, 102, 51, 0, 587, 588, 5, 3, 0, 0, 588, 589, 3, 36, 18, 0, 589, 590,
		5, 58, 0, 0, 590, 591, 3, 102, 51, 0, 591, 596, 1, 0, 0, 0, 592, 593, 5,
		7, 0, 0, 593, 594, 5, 58, 0, 0, 594, 596, 3, 102, 51, 0, 595, 584, 1, 0,
		0, 0, 595, 587, 1, 0, 0, 0, 595, 592, 1, 0, 0, 0, 596, 105, 1, 0, 0, 0,
		597, 599, 5, 35, 0, 0, 598, 600, 3, 108, 54, 0, 599, 598, 1, 0, 0, 0, 599,
		600, 1, 0, 0, 0, 600, 601, 1, 0, 0, 0, 601, 602, 5, 36, 0, 0, 602, 107,
		1, 0, 0, 0, 603, 605, 3, 110, 55, 0, 604, 603, 1, 0, 0, 0, 605, 606, 1,
		0, 0, 0, 606, 604, 1, 0, 0, 0, 606, 607, 1, 0, 0, 0, 607, 109, 1, 0, 0,
		0, 608, 611, 3, 102, 51, 0, 609, 611, 3, 38, 19, 0, 610, 608, 1, 0, 0,
		0, 610, 609, 1, 0, 0, 0, 611, 111, 1, 0, 0, 0, 612, 614, 3, 34, 17, 0,
		613, 612, 1, 0, 0, 0, 613, 614, 1, 0, 0, 0, 614, 615, 1, 0, 0, 0, 615,
		616, 5, 59, 0, 0, 616, 113, 1, 0, 0, 0, 617, 618, 5, 15, 0, 0, 618, 619,
		5, 31, 0, 0, 619, 620, 3, 34, 17, 0, 620, 621, 5, 32, 0, 0, 621, 624, 3,
		102, 51, 0, 622, 623, 5, 10, 0, 0, 623, 625, 3, 102, 51, 0, 624, 622, 1,
		0, 0, 0, 624, 625, 1, 0, 0, 0, 625, 633, 1, 0, 0, 0, 626, 627, 5, 25, 0,
		0, 627, 628, 5, 31, 0, 0, 628, 629, 3, 34, 17, 0, 629, 630, 5, 32, 0, 0,
		630, 631, 3, 102, 51, 0, 631, 633, 1, 0, 0, 0, 632, 617, 1, 0, 0, 0, 632,
		626, 1, 0, 0, 0, 633, 115, 1, 0, 0, 0, 634, 635, 5, 29, 0, 0, 635, 636,
		5, 31, 0, 0, 636, 637, 3, 34, 17, 0, 637, 638, 5, 32, 0, 0, 638, 639, 3,
		102, 51, 0, 639, 655, 1, 0, 0, 0, 640, 641, 5, 8, 0, 0, 641, 642, 3, 102,
		51, 0, 642, 643, 5, 29, 0, 0, 643, 644, 5, 31, 0, 0, 644, 645, 3, 34, 17,
		0, 645, 646, 5, 32, 0, 0, 646, 647, 5, 59, 0, 0, 647, 655, 1, 0, 0, 0,
		648, 649, 5, 13, 0, 0, 649, 650, 5, 31, 0, 0, 650, 651, 3, 118, 59, 0,
		651, 652, 5, 32, 0, 0, 652, 653, 3, 102, 51, 0, 653, 655, 1, 0, 0, 0, 654,
		634, 1, 0, 0, 0, 654, 640, 1, 0, 0, 0, 654, 648, 1, 0, 0, 0, 655, 117,
		1, 0, 0, 0, 656, 661, 3, 120, 60, 0, 657, 659, 3, 34, 17, 0, 658, 657,
		1, 0, 0, 0, 658, 659, 1, 0, 0, 0, 659, 661, 1, 0, 0, 0, 660, 656, 1, 0,
		0, 0, 660, 658, 1, 0, 0, 0, 661, 662, 1, 0, 0, 0, 662, 664, 5, 59, 0, 0,
		663, 665, 3, 122, 61, 0, 664, 663, 1, 0, 0, 0, 664, 665, 1, 0, 0, 0, 665,
		666, 1, 0, 0, 0, 666, 668, 5, 59, 0, 0, 667, 669, 3, 122, 61, 0, 668, 667,
		1, 0, 0, 0, 668, 669, 1, 0, 0, 0, 669, 119, 1, 0, 0, 0, 670, 672, 3, 40,
		20, 0, 671, 673, 3, 46, 23, 0, 672, 671, 1, 0, 0, 0, 672, 673, 1, 0, 0,
		0, 673, 121, 1, 0, 0, 0, 674, 679, 3, 30, 15, 0, 675, 676, 5, 60, 0, 0,
		676, 678, 3, 30, 15, 0, 677, 675, 1, 0, 0, 0, 678, 681, 1, 0, 0, 0, 679,
		677, 1, 0, 0, 0, 679, 680, 1, 0, 0, 0, 680, 123, 1, 0, 0, 0, 681, 679,
		1, 0, 0, 0, 682, 683, 5, 14, 0, 0, 683, 693, 5, 67, 0, 0, 684, 693, 5,
		6, 0, 0, 685, 693, 5, 2, 0, 0, 686, 688, 5, 20, 0, 0, 687, 689, 3, 34,
		17, 0, 688, 687, 1, 0, 0, 0, 688, 689, 1, 0, 0, 0, 689, 693, 1, 0, 0, 0,
		690, 691, 5, 14, 0, 0, 691, 693, 3, 10, 5, 0, 692, 682, 1, 0, 0, 0, 692,
		684, 1, 0, 0, 0, 692, 685, 1, 0, 0, 0, 692, 686, 1, 0, 0, 0, 692, 690,
		1, 0, 0, 0, 693, 694, 1, 0, 0, 0, 694, 695, 5, 59, 0, 0, 695, 125, 1, 0,
		0, 0, 696, 698, 3, 128, 64, 0, 697, 696, 1, 0, 0, 0, 698, 699, 1, 0, 0,
		0, 699, 697, 1, 0, 0, 0, 699, 700, 1, 0, 0, 0, 700, 701, 1, 0, 0, 0, 701,
		702, 5, 0, 0, 1, 702, 127, 1, 0, 0, 0, 703, 707, 3, 130, 65, 0, 704, 707,
		3, 38, 19, 0, 705, 707, 5, 59, 0, 0, 706, 703, 1, 0, 0, 0, 706, 704, 1,
		0, 0, 0, 706, 705, 1, 0, 0, 0, 707, 129, 1, 0, 0, 0, 708, 710, 3, 40, 20,
		0, 709, 708, 1, 0, 0, 0, 709, 710, 1, 0, 0, 0, 710, 711, 1, 0, 0, 0, 711,
		713, 3, 66, 33, 0, 712, 714, 3, 132, 66, 0, 713, 712, 1, 0, 0, 0, 713,
		714, 1, 0, 0, 0, 714, 715, 1, 0, 0, 0, 715, 716, 3, 106, 53, 0, 716, 131,
		1, 0, 0, 0, 717, 719, 3, 38, 19, 0, 718, 717, 1, 0, 0, 0, 719, 720, 1,
		0, 0, 0, 720, 718, 1, 0, 0, 0, 720, 721, 1, 0, 0, 0, 721, 133, 1, 0, 0,
		0, 90, 139, 149, 156, 161, 173, 177, 185, 192, 194, 202, 211, 222, 229,
		237, 245, 253, 261, 269, 277, 286, 295, 302, 309, 314, 323, 329, 340, 344,
		353, 360, 369, 373, 380, 385, 389, 392, 405, 410, 421, 424, 426, 432, 435,
		442, 444, 450, 455, 462, 470, 472, 479, 484, 488, 491, 500, 508, 511, 516,
		526, 529, 531, 540, 544, 547, 552, 557, 566, 574, 582, 595, 599, 606, 610,
		613, 624, 632, 654, 658, 660, 664, 668, 672, 679, 688, 692, 699, 706, 709,
		713, 720,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// tinycParserInit initializes any static state used to implement tinycParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewtinycParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func TinycParserInit() {
	staticData := &TinycParserStaticData
	staticData.once.Do(tinycParserInit)
}

// NewtinycParser produces a new parser instance for the optional input antlr.TokenStream.
func NewtinycParser(input antlr.TokenStream) *tinycParser {
	TinycParserInit()
	this := new(tinycParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &TinycParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "tinyc.g4"

	return this
}

// tinycParser tokens.
const (
	tinycParserEOF                    = antlr.TokenEOF
	tinycParserAuto                   = 1
	tinycParserBreak                  = 2
	tinycParserCase                   = 3
	tinycParserChar                   = 4
	tinycParserConst                  = 5
	tinycParserContinue               = 6
	tinycParserDefault                = 7
	tinycParserDo                     = 8
	tinycParserDouble                 = 9
	tinycParserElse                   = 10
	tinycParserExtern                 = 11
	tinycParserFloat                  = 12
	tinycParserFor                    = 13
	tinycParserGoto                   = 14
	tinycParserIf                     = 15
	tinycParserInline                 = 16
	tinycParserInt                    = 17
	tinycParserRegister               = 18
	tinycParserRestrict               = 19
	tinycParserReturn                 = 20
	tinycParserShort                  = 21
	tinycParserSigned                 = 22
	tinycParserSizeof                 = 23
	tinycParserStruct                 = 24
	tinycParserSwitch                 = 25
	tinycParserTypedef                = 26
	tinycParserUnion                  = 27
	tinycParserVoid                   = 28
	tinycParserWhile                  = 29
	tinycParserNoreturn               = 30
	tinycParserLeftParen              = 31
	tinycParserRightParen             = 32
	tinycParserLeftBracket            = 33
	tinycParserRightBracket           = 34
	tinycParserLeftBrace              = 35
	tinycParserRightBrace             = 36
	tinycParserLess                   = 37
	tinycParserLessEqual              = 38
	tinycParserGreater                = 39
	tinycParserGreaterEqual           = 40
	tinycParserLeftShift              = 41
	tinycParserRightShift             = 42
	tinycParserPlus                   = 43
	tinycParserPlusPlus               = 44
	tinycParserMinus                  = 45
	tinycParserMinusMinus             = 46
	tinycParserStar                   = 47
	tinycParserDiv                    = 48
	tinycParserMod                    = 49
	tinycParserAnd                    = 50
	tinycParserOr                     = 51
	tinycParserAndAnd                 = 52
	tinycParserOrOr                   = 53
	tinycParserCaret                  = 54
	tinycParserNot                    = 55
	tinycParserTilde                  = 56
	tinycParserQuestion               = 57
	tinycParserColon                  = 58
	tinycParserSemi                   = 59
	tinycParserComma                  = 60
	tinycParserAssign                 = 61
	tinycParserEqual                  = 62
	tinycParserNotEqual               = 63
	tinycParserArrow                  = 64
	tinycParserDot                    = 65
	tinycParserEllipsis               = 66
	tinycParserIdentifier             = 67
	tinycParserConstant               = 68
	tinycParserDigitSequence          = 69
	tinycParserStringLiteral          = 70
	tinycParserComplexDefine          = 71
	tinycParserIncludeDirective       = 72
	tinycParserAsmBlock               = 73
	tinycParserLineAfterPreprocessing = 74
	tinycParserLineDirective          = 75
	tinycParserPragmaDirective        = 76
	tinycParserWhitespace             = 77
	tinycParserNewline                = 78
	tinycParserBlockComment           = 79
	tinycParserLineComment            = 80
)

// tinycParser rules.
const (
	tinycParserRULE_primaryExpression        = 0
	tinycParserRULE_genericAssocList         = 1
	tinycParserRULE_genericAssociation       = 2
	tinycParserRULE_postfixExpression        = 3
	tinycParserRULE_argumentExpressionList   = 4
	tinycParserRULE_unaryExpression          = 5
	tinycParserRULE_unaryOperator            = 6
	tinycParserRULE_castExpression           = 7
	tinycParserRULE_multiplicativeExpression = 8
	tinycParserRULE_additiveExpression       = 9
	tinycParserRULE_relationalExpression     = 10
	tinycParserRULE_equalityExpression       = 11
	tinycParserRULE_inclusiveOrExpression    = 12
	tinycParserRULE_logicalAndExpression     = 13
	tinycParserRULE_logicalOrExpression      = 14
	tinycParserRULE_assignmentExpression     = 15
	tinycParserRULE_assignmentOperator       = 16
	tinycParserRULE_expression               = 17
	tinycParserRULE_constantExpression       = 18
	tinycParserRULE_declaration              = 19
	tinycParserRULE_declarationSpecifiers    = 20
	tinycParserRULE_declarationSpecifiers2   = 21
	tinycParserRULE_declarationSpecifier     = 22
	tinycParserRULE_initDeclaratorList       = 23
	tinycParserRULE_initDeclarator           = 24
	tinycParserRULE_typeSpecifier            = 25
	tinycParserRULE_structOrUnionSpecifier   = 26
	tinycParserRULE_structOrUnion            = 27
	tinycParserRULE_structDeclarationList    = 28
	tinycParserRULE_structDeclaration        = 29
	tinycParserRULE_specifierQualifierList   = 30
	tinycParserRULE_structDeclaratorList     = 31
	tinycParserRULE_structDeclarator         = 32
	tinycParserRULE_declarator               = 33
	tinycParserRULE_directDeclarator         = 34
	tinycParserRULE_gccAttribute             = 35
	tinycParserRULE_nestedParenthesesBlock   = 36
	tinycParserRULE_pointer                  = 37
	tinycParserRULE_parameterTypeList        = 38
	tinycParserRULE_parameterList            = 39
	tinycParserRULE_parameterDeclaration     = 40
	tinycParserRULE_identifierList           = 41
	tinycParserRULE_typeName                 = 42
	tinycParserRULE_abstractDeclarator       = 43
	tinycParserRULE_directAbstractDeclarator = 44
	tinycParserRULE_typedefName              = 45
	tinycParserRULE_initializer              = 46
	tinycParserRULE_initializerList          = 47
	tinycParserRULE_designation              = 48
	tinycParserRULE_designatorList           = 49
	tinycParserRULE_designator               = 50
	tinycParserRULE_statement                = 51
	tinycParserRULE_labeledStatement         = 52
	tinycParserRULE_compoundStatement        = 53
	tinycParserRULE_blockItemList            = 54
	tinycParserRULE_blockItem                = 55
	tinycParserRULE_expressionStatement      = 56
	tinycParserRULE_selectionStatement       = 57
	tinycParserRULE_iterationStatement       = 58
	tinycParserRULE_forCondition             = 59
	tinycParserRULE_forDeclaration           = 60
	tinycParserRULE_forExpression            = 61
	tinycParserRULE_jumpStatement            = 62
	tinycParserRULE_compilationUnit          = 63
	tinycParserRULE_externalDeclaration      = 64
	tinycParserRULE_functionDefinition       = 65
	tinycParserRULE_declarationList          = 66
)

// IPrimaryExpressionContext is an interface to support dynamic dispatch.
type IPrimaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	Constant() antlr.TerminalNode
	AllStringLiteral() []antlr.TerminalNode
	StringLiteral(i int) antlr.TerminalNode
	LeftParen() antlr.TerminalNode
	Expression() IExpressionContext
	RightParen() antlr.TerminalNode
	CompoundStatement() ICompoundStatementContext

	// IsPrimaryExpressionContext differentiates from other interfaces.
	IsPrimaryExpressionContext()
}

type PrimaryExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExpressionContext() *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_primaryExpression
	return p
}

func InitEmptyPrimaryExpressionContext(p *PrimaryExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_primaryExpression
}

func (*PrimaryExpressionContext) IsPrimaryExpressionContext() {}

func NewPrimaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_primaryExpression

	return p
}

func (s *PrimaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExpressionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(tinycParserIdentifier, 0)
}

func (s *PrimaryExpressionContext) Constant() antlr.TerminalNode {
	return s.GetToken(tinycParserConstant, 0)
}

func (s *PrimaryExpressionContext) AllStringLiteral() []antlr.TerminalNode {
	return s.GetTokens(tinycParserStringLiteral)
}

func (s *PrimaryExpressionContext) StringLiteral(i int) antlr.TerminalNode {
	return s.GetToken(tinycParserStringLiteral, i)
}

func (s *PrimaryExpressionContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(tinycParserLeftParen, 0)
}

func (s *PrimaryExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PrimaryExpressionContext) RightParen() antlr.TerminalNode {
	return s.GetToken(tinycParserRightParen, 0)
}

func (s *PrimaryExpressionContext) CompoundStatement() ICompoundStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICompoundStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICompoundStatementContext)
}

func (s *PrimaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterPrimaryExpression(s)
	}
}

func (s *PrimaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitPrimaryExpression(s)
	}
}

func (s *PrimaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitPrimaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) PrimaryExpression() (localctx IPrimaryExpressionContext) {
	localctx = NewPrimaryExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, tinycParserRULE_primaryExpression)
	var _la int

	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(134)
			p.Match(tinycParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(135)
			p.Match(tinycParserConstant)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(137)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == tinycParserStringLiteral {
			{
				p.SetState(136)
				p.Match(tinycParserStringLiteral)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(139)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(141)
			p.Match(tinycParserLeftParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(142)
			p.Expression()
		}
		{
			p.SetState(143)
			p.Match(tinycParserRightParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(145)
			p.Match(tinycParserLeftParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(146)
			p.CompoundStatement()
		}
		{
			p.SetState(147)
			p.Match(tinycParserRightParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGenericAssocListContext is an interface to support dynamic dispatch.
type IGenericAssocListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllGenericAssociation() []IGenericAssociationContext
	GenericAssociation(i int) IGenericAssociationContext
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

	// IsGenericAssocListContext differentiates from other interfaces.
	IsGenericAssocListContext()
}

type GenericAssocListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGenericAssocListContext() *GenericAssocListContext {
	var p = new(GenericAssocListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_genericAssocList
	return p
}

func InitEmptyGenericAssocListContext(p *GenericAssocListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_genericAssocList
}

func (*GenericAssocListContext) IsGenericAssocListContext() {}

func NewGenericAssocListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GenericAssocListContext {
	var p = new(GenericAssocListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_genericAssocList

	return p
}

func (s *GenericAssocListContext) GetParser() antlr.Parser { return s.parser }

func (s *GenericAssocListContext) AllGenericAssociation() []IGenericAssociationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IGenericAssociationContext); ok {
			len++
		}
	}

	tst := make([]IGenericAssociationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IGenericAssociationContext); ok {
			tst[i] = t.(IGenericAssociationContext)
			i++
		}
	}

	return tst
}

func (s *GenericAssocListContext) GenericAssociation(i int) IGenericAssociationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGenericAssociationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGenericAssociationContext)
}

func (s *GenericAssocListContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(tinycParserComma)
}

func (s *GenericAssocListContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(tinycParserComma, i)
}

func (s *GenericAssocListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GenericAssocListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GenericAssocListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterGenericAssocList(s)
	}
}

func (s *GenericAssocListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitGenericAssocList(s)
	}
}

func (s *GenericAssocListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitGenericAssocList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) GenericAssocList() (localctx IGenericAssocListContext) {
	localctx = NewGenericAssocListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, tinycParserRULE_genericAssocList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(151)
		p.GenericAssociation()
	}
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == tinycParserComma {
		{
			p.SetState(152)
			p.Match(tinycParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(153)
			p.GenericAssociation()
		}

		p.SetState(158)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGenericAssociationContext is an interface to support dynamic dispatch.
type IGenericAssociationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Colon() antlr.TerminalNode
	AssignmentExpression() IAssignmentExpressionContext
	TypeName() ITypeNameContext
	Default() antlr.TerminalNode

	// IsGenericAssociationContext differentiates from other interfaces.
	IsGenericAssociationContext()
}

type GenericAssociationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGenericAssociationContext() *GenericAssociationContext {
	var p = new(GenericAssociationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_genericAssociation
	return p
}

func InitEmptyGenericAssociationContext(p *GenericAssociationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_genericAssociation
}

func (*GenericAssociationContext) IsGenericAssociationContext() {}

func NewGenericAssociationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GenericAssociationContext {
	var p = new(GenericAssociationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_genericAssociation

	return p
}

func (s *GenericAssociationContext) GetParser() antlr.Parser { return s.parser }

func (s *GenericAssociationContext) Colon() antlr.TerminalNode {
	return s.GetToken(tinycParserColon, 0)
}

func (s *GenericAssociationContext) AssignmentExpression() IAssignmentExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentExpressionContext)
}

func (s *GenericAssociationContext) TypeName() ITypeNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *GenericAssociationContext) Default() antlr.TerminalNode {
	return s.GetToken(tinycParserDefault, 0)
}

func (s *GenericAssociationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GenericAssociationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GenericAssociationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterGenericAssociation(s)
	}
}

func (s *GenericAssociationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitGenericAssociation(s)
	}
}

func (s *GenericAssociationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitGenericAssociation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) GenericAssociation() (localctx IGenericAssociationContext) {
	localctx = NewGenericAssociationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, tinycParserRULE_genericAssociation)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case tinycParserChar, tinycParserDouble, tinycParserFloat, tinycParserInt, tinycParserShort, tinycParserSigned, tinycParserStruct, tinycParserUnion, tinycParserVoid, tinycParserIdentifier:
		{
			p.SetState(159)
			p.TypeName()
		}

	case tinycParserDefault:
		{
			p.SetState(160)
			p.Match(tinycParserDefault)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(163)
		p.Match(tinycParserColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(164)
		p.AssignmentExpression()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPostfixExpressionContext is an interface to support dynamic dispatch.
type IPostfixExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PrimaryExpression() IPrimaryExpressionContext
	AllLeftParen() []antlr.TerminalNode
	LeftParen(i int) antlr.TerminalNode
	TypeName() ITypeNameContext
	AllRightParen() []antlr.TerminalNode
	RightParen(i int) antlr.TerminalNode
	LeftBrace() antlr.TerminalNode
	InitializerList() IInitializerListContext
	RightBrace() antlr.TerminalNode
	AllLeftBracket() []antlr.TerminalNode
	LeftBracket(i int) antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllRightBracket() []antlr.TerminalNode
	RightBracket(i int) antlr.TerminalNode
	AllIdentifier() []antlr.TerminalNode
	Identifier(i int) antlr.TerminalNode
	AllPlusPlus() []antlr.TerminalNode
	PlusPlus(i int) antlr.TerminalNode
	AllMinusMinus() []antlr.TerminalNode
	MinusMinus(i int) antlr.TerminalNode
	AllDot() []antlr.TerminalNode
	Dot(i int) antlr.TerminalNode
	AllArrow() []antlr.TerminalNode
	Arrow(i int) antlr.TerminalNode
	Comma() antlr.TerminalNode
	AllArgumentExpressionList() []IArgumentExpressionListContext
	ArgumentExpressionList(i int) IArgumentExpressionListContext

	// IsPostfixExpressionContext differentiates from other interfaces.
	IsPostfixExpressionContext()
}

type PostfixExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPostfixExpressionContext() *PostfixExpressionContext {
	var p = new(PostfixExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_postfixExpression
	return p
}

func InitEmptyPostfixExpressionContext(p *PostfixExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_postfixExpression
}

func (*PostfixExpressionContext) IsPostfixExpressionContext() {}

func NewPostfixExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PostfixExpressionContext {
	var p = new(PostfixExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_postfixExpression

	return p
}

func (s *PostfixExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *PostfixExpressionContext) PrimaryExpression() IPrimaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *PostfixExpressionContext) AllLeftParen() []antlr.TerminalNode {
	return s.GetTokens(tinycParserLeftParen)
}

func (s *PostfixExpressionContext) LeftParen(i int) antlr.TerminalNode {
	return s.GetToken(tinycParserLeftParen, i)
}

func (s *PostfixExpressionContext) TypeName() ITypeNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *PostfixExpressionContext) AllRightParen() []antlr.TerminalNode {
	return s.GetTokens(tinycParserRightParen)
}

func (s *PostfixExpressionContext) RightParen(i int) antlr.TerminalNode {
	return s.GetToken(tinycParserRightParen, i)
}

func (s *PostfixExpressionContext) LeftBrace() antlr.TerminalNode {
	return s.GetToken(tinycParserLeftBrace, 0)
}

func (s *PostfixExpressionContext) InitializerList() IInitializerListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInitializerListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInitializerListContext)
}

func (s *PostfixExpressionContext) RightBrace() antlr.TerminalNode {
	return s.GetToken(tinycParserRightBrace, 0)
}

func (s *PostfixExpressionContext) AllLeftBracket() []antlr.TerminalNode {
	return s.GetTokens(tinycParserLeftBracket)
}

func (s *PostfixExpressionContext) LeftBracket(i int) antlr.TerminalNode {
	return s.GetToken(tinycParserLeftBracket, i)
}

func (s *PostfixExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *PostfixExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PostfixExpressionContext) AllRightBracket() []antlr.TerminalNode {
	return s.GetTokens(tinycParserRightBracket)
}

func (s *PostfixExpressionContext) RightBracket(i int) antlr.TerminalNode {
	return s.GetToken(tinycParserRightBracket, i)
}

func (s *PostfixExpressionContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(tinycParserIdentifier)
}

func (s *PostfixExpressionContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(tinycParserIdentifier, i)
}

func (s *PostfixExpressionContext) AllPlusPlus() []antlr.TerminalNode {
	return s.GetTokens(tinycParserPlusPlus)
}

func (s *PostfixExpressionContext) PlusPlus(i int) antlr.TerminalNode {
	return s.GetToken(tinycParserPlusPlus, i)
}

func (s *PostfixExpressionContext) AllMinusMinus() []antlr.TerminalNode {
	return s.GetTokens(tinycParserMinusMinus)
}

func (s *PostfixExpressionContext) MinusMinus(i int) antlr.TerminalNode {
	return s.GetToken(tinycParserMinusMinus, i)
}

func (s *PostfixExpressionContext) AllDot() []antlr.TerminalNode {
	return s.GetTokens(tinycParserDot)
}

func (s *PostfixExpressionContext) Dot(i int) antlr.TerminalNode {
	return s.GetToken(tinycParserDot, i)
}

func (s *PostfixExpressionContext) AllArrow() []antlr.TerminalNode {
	return s.GetTokens(tinycParserArrow)
}

func (s *PostfixExpressionContext) Arrow(i int) antlr.TerminalNode {
	return s.GetToken(tinycParserArrow, i)
}

func (s *PostfixExpressionContext) Comma() antlr.TerminalNode {
	return s.GetToken(tinycParserComma, 0)
}

func (s *PostfixExpressionContext) AllArgumentExpressionList() []IArgumentExpressionListContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArgumentExpressionListContext); ok {
			len++
		}
	}

	tst := make([]IArgumentExpressionListContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArgumentExpressionListContext); ok {
			tst[i] = t.(IArgumentExpressionListContext)
			i++
		}
	}

	return tst
}

func (s *PostfixExpressionContext) ArgumentExpressionList(i int) IArgumentExpressionListContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentExpressionListContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentExpressionListContext)
}

func (s *PostfixExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostfixExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PostfixExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterPostfixExpression(s)
	}
}

func (s *PostfixExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitPostfixExpression(s)
	}
}

func (s *PostfixExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitPostfixExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) PostfixExpression() (localctx IPostfixExpressionContext) {
	localctx = NewPostfixExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, tinycParserRULE_postfixExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(166)
			p.PrimaryExpression()
		}

	case 2:
		{
			p.SetState(167)
			p.Match(tinycParserLeftParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(168)
			p.TypeName()
		}
		{
			p.SetState(169)
			p.Match(tinycParserRightParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(170)
			p.Match(tinycParserLeftBrace)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(171)
			p.InitializerList()
		}
		p.SetState(173)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == tinycParserComma {
			{
				p.SetState(172)
				p.Match(tinycParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(175)
			p.Match(tinycParserRightBrace)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-31)) & ^0x3f) == 0 && ((int64(1)<<(_la-31))&25769844741) != 0 {
		p.SetState(192)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case tinycParserLeftBracket:
			{
				p.SetState(179)
				p.Match(tinycParserLeftBracket)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(180)
				p.Expression()
			}
			{
				p.SetState(181)
				p.Match(tinycParserRightBracket)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case tinycParserLeftParen:
			{
				p.SetState(183)
				p.Match(tinycParserLeftParen)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(185)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if (int64((_la-31)) & ^0x3f) == 0 && ((int64(1)<<(_la-31))&1030845190145) != 0 {
				{
					p.SetState(184)
					p.ArgumentExpressionList()
				}

			}
			{
				p.SetState(187)
				p.Match(tinycParserRightParen)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case tinycParserArrow, tinycParserDot:
			{
				p.SetState(188)
				_la = p.GetTokenStream().LA(1)

				if !(_la == tinycParserArrow || _la == tinycParserDot) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(189)
				p.Match(tinycParserIdentifier)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case tinycParserPlusPlus:
			{
				p.SetState(190)
				p.Match(tinycParserPlusPlus)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case tinycParserMinusMinus:
			{
				p.SetState(191)
				p.Match(tinycParserMinusMinus)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(196)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgumentExpressionListContext is an interface to support dynamic dispatch.
type IArgumentExpressionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAssignmentExpression() []IAssignmentExpressionContext
	AssignmentExpression(i int) IAssignmentExpressionContext
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

	// IsArgumentExpressionListContext differentiates from other interfaces.
	IsArgumentExpressionListContext()
}

type ArgumentExpressionListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentExpressionListContext() *ArgumentExpressionListContext {
	var p = new(ArgumentExpressionListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_argumentExpressionList
	return p
}

func InitEmptyArgumentExpressionListContext(p *ArgumentExpressionListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_argumentExpressionList
}

func (*ArgumentExpressionListContext) IsArgumentExpressionListContext() {}

func NewArgumentExpressionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentExpressionListContext {
	var p = new(ArgumentExpressionListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_argumentExpressionList

	return p
}

func (s *ArgumentExpressionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentExpressionListContext) AllAssignmentExpression() []IAssignmentExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAssignmentExpressionContext); ok {
			len++
		}
	}

	tst := make([]IAssignmentExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAssignmentExpressionContext); ok {
			tst[i] = t.(IAssignmentExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ArgumentExpressionListContext) AssignmentExpression(i int) IAssignmentExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentExpressionContext)
}

func (s *ArgumentExpressionListContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(tinycParserComma)
}

func (s *ArgumentExpressionListContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(tinycParserComma, i)
}

func (s *ArgumentExpressionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentExpressionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentExpressionListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterArgumentExpressionList(s)
	}
}

func (s *ArgumentExpressionListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitArgumentExpressionList(s)
	}
}

func (s *ArgumentExpressionListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitArgumentExpressionList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) ArgumentExpressionList() (localctx IArgumentExpressionListContext) {
	localctx = NewArgumentExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, tinycParserRULE_argumentExpressionList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(197)
		p.AssignmentExpression()
	}
	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == tinycParserComma {
		{
			p.SetState(198)
			p.Match(tinycParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(199)
			p.AssignmentExpression()
		}

		p.SetState(204)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnaryExpressionContext is an interface to support dynamic dispatch.
type IUnaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PostfixExpression() IPostfixExpressionContext
	UnaryOperator() IUnaryOperatorContext
	CastExpression() ICastExpressionContext
	AndAnd() antlr.TerminalNode
	Identifier() antlr.TerminalNode

	// IsUnaryExpressionContext differentiates from other interfaces.
	IsUnaryExpressionContext()
}

type UnaryExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryExpressionContext() *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_unaryExpression
	return p
}

func InitEmptyUnaryExpressionContext(p *UnaryExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_unaryExpression
}

func (*UnaryExpressionContext) IsUnaryExpressionContext() {}

func NewUnaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_unaryExpression

	return p
}

func (s *UnaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryExpressionContext) PostfixExpression() IPostfixExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPostfixExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPostfixExpressionContext)
}

func (s *UnaryExpressionContext) UnaryOperator() IUnaryOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryOperatorContext)
}

func (s *UnaryExpressionContext) CastExpression() ICastExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICastExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICastExpressionContext)
}

func (s *UnaryExpressionContext) AndAnd() antlr.TerminalNode {
	return s.GetToken(tinycParserAndAnd, 0)
}

func (s *UnaryExpressionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(tinycParserIdentifier, 0)
}

func (s *UnaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterUnaryExpression(s)
	}
}

func (s *UnaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitUnaryExpression(s)
	}
}

func (s *UnaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitUnaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) UnaryExpression() (localctx IUnaryExpressionContext) {
	localctx = NewUnaryExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, tinycParserRULE_unaryExpression)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(211)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case tinycParserLeftParen, tinycParserIdentifier, tinycParserConstant, tinycParserStringLiteral:
		{
			p.SetState(205)
			p.PostfixExpression()
		}

	case tinycParserPlus, tinycParserMinus, tinycParserStar, tinycParserAnd, tinycParserNot, tinycParserTilde:
		{
			p.SetState(206)
			p.UnaryOperator()
		}
		{
			p.SetState(207)
			p.CastExpression()
		}

	case tinycParserAndAnd:
		{
			p.SetState(209)
			p.Match(tinycParserAndAnd)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(210)
			p.Match(tinycParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnaryOperatorContext is an interface to support dynamic dispatch.
type IUnaryOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	And() antlr.TerminalNode
	Star() antlr.TerminalNode
	Plus() antlr.TerminalNode
	Minus() antlr.TerminalNode
	Tilde() antlr.TerminalNode
	Not() antlr.TerminalNode

	// IsUnaryOperatorContext differentiates from other interfaces.
	IsUnaryOperatorContext()
}

type UnaryOperatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryOperatorContext() *UnaryOperatorContext {
	var p = new(UnaryOperatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_unaryOperator
	return p
}

func InitEmptyUnaryOperatorContext(p *UnaryOperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_unaryOperator
}

func (*UnaryOperatorContext) IsUnaryOperatorContext() {}

func NewUnaryOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryOperatorContext {
	var p = new(UnaryOperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_unaryOperator

	return p
}

func (s *UnaryOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryOperatorContext) And() antlr.TerminalNode {
	return s.GetToken(tinycParserAnd, 0)
}

func (s *UnaryOperatorContext) Star() antlr.TerminalNode {
	return s.GetToken(tinycParserStar, 0)
}

func (s *UnaryOperatorContext) Plus() antlr.TerminalNode {
	return s.GetToken(tinycParserPlus, 0)
}

func (s *UnaryOperatorContext) Minus() antlr.TerminalNode {
	return s.GetToken(tinycParserMinus, 0)
}

func (s *UnaryOperatorContext) Tilde() antlr.TerminalNode {
	return s.GetToken(tinycParserTilde, 0)
}

func (s *UnaryOperatorContext) Not() antlr.TerminalNode {
	return s.GetToken(tinycParserNot, 0)
}

func (s *UnaryOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterUnaryOperator(s)
	}
}

func (s *UnaryOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitUnaryOperator(s)
	}
}

func (s *UnaryOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitUnaryOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) UnaryOperator() (localctx IUnaryOperatorContext) {
	localctx = NewUnaryOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, tinycParserRULE_unaryOperator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(213)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&109397008917200896) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICastExpressionContext is an interface to support dynamic dispatch.
type ICastExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LeftParen() antlr.TerminalNode
	TypeName() ITypeNameContext
	RightParen() antlr.TerminalNode
	CastExpression() ICastExpressionContext
	UnaryExpression() IUnaryExpressionContext
	DigitSequence() antlr.TerminalNode

	// IsCastExpressionContext differentiates from other interfaces.
	IsCastExpressionContext()
}

type CastExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCastExpressionContext() *CastExpressionContext {
	var p = new(CastExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_castExpression
	return p
}

func InitEmptyCastExpressionContext(p *CastExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_castExpression
}

func (*CastExpressionContext) IsCastExpressionContext() {}

func NewCastExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CastExpressionContext {
	var p = new(CastExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_castExpression

	return p
}

func (s *CastExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *CastExpressionContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(tinycParserLeftParen, 0)
}

func (s *CastExpressionContext) TypeName() ITypeNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *CastExpressionContext) RightParen() antlr.TerminalNode {
	return s.GetToken(tinycParserRightParen, 0)
}

func (s *CastExpressionContext) CastExpression() ICastExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICastExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICastExpressionContext)
}

func (s *CastExpressionContext) UnaryExpression() IUnaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryExpressionContext)
}

func (s *CastExpressionContext) DigitSequence() antlr.TerminalNode {
	return s.GetToken(tinycParserDigitSequence, 0)
}

func (s *CastExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CastExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CastExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterCastExpression(s)
	}
}

func (s *CastExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitCastExpression(s)
	}
}

func (s *CastExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitCastExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) CastExpression() (localctx ICastExpressionContext) {
	localctx = NewCastExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, tinycParserRULE_castExpression)
	p.SetState(222)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(215)
			p.Match(tinycParserLeftParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(216)
			p.TypeName()
		}
		{
			p.SetState(217)
			p.Match(tinycParserRightParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(218)
			p.CastExpression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(220)
			p.UnaryExpression()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(221)
			p.Match(tinycParserDigitSequence)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMultiplicativeExpressionContext is an interface to support dynamic dispatch.
type IMultiplicativeExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllCastExpression() []ICastExpressionContext
	CastExpression(i int) ICastExpressionContext
	AllStar() []antlr.TerminalNode
	Star(i int) antlr.TerminalNode
	AllDiv() []antlr.TerminalNode
	Div(i int) antlr.TerminalNode
	AllMod() []antlr.TerminalNode
	Mod(i int) antlr.TerminalNode

	// IsMultiplicativeExpressionContext differentiates from other interfaces.
	IsMultiplicativeExpressionContext()
}

type MultiplicativeExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplicativeExpressionContext() *MultiplicativeExpressionContext {
	var p = new(MultiplicativeExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_multiplicativeExpression
	return p
}

func InitEmptyMultiplicativeExpressionContext(p *MultiplicativeExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_multiplicativeExpression
}

func (*MultiplicativeExpressionContext) IsMultiplicativeExpressionContext() {}

func NewMultiplicativeExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplicativeExpressionContext {
	var p = new(MultiplicativeExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_multiplicativeExpression

	return p
}

func (s *MultiplicativeExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplicativeExpressionContext) AllCastExpression() []ICastExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICastExpressionContext); ok {
			len++
		}
	}

	tst := make([]ICastExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICastExpressionContext); ok {
			tst[i] = t.(ICastExpressionContext)
			i++
		}
	}

	return tst
}

func (s *MultiplicativeExpressionContext) CastExpression(i int) ICastExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICastExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICastExpressionContext)
}

func (s *MultiplicativeExpressionContext) AllStar() []antlr.TerminalNode {
	return s.GetTokens(tinycParserStar)
}

func (s *MultiplicativeExpressionContext) Star(i int) antlr.TerminalNode {
	return s.GetToken(tinycParserStar, i)
}

func (s *MultiplicativeExpressionContext) AllDiv() []antlr.TerminalNode {
	return s.GetTokens(tinycParserDiv)
}

func (s *MultiplicativeExpressionContext) Div(i int) antlr.TerminalNode {
	return s.GetToken(tinycParserDiv, i)
}

func (s *MultiplicativeExpressionContext) AllMod() []antlr.TerminalNode {
	return s.GetTokens(tinycParserMod)
}

func (s *MultiplicativeExpressionContext) Mod(i int) antlr.TerminalNode {
	return s.GetToken(tinycParserMod, i)
}

func (s *MultiplicativeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiplicativeExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterMultiplicativeExpression(s)
	}
}

func (s *MultiplicativeExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitMultiplicativeExpression(s)
	}
}

func (s *MultiplicativeExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitMultiplicativeExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) MultiplicativeExpression() (localctx IMultiplicativeExpressionContext) {
	localctx = NewMultiplicativeExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, tinycParserRULE_multiplicativeExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(224)
		p.CastExpression()
	}
	p.SetState(229)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&985162418487296) != 0 {
		{
			p.SetState(225)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&985162418487296) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(226)
			p.CastExpression()
		}

		p.SetState(231)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAdditiveExpressionContext is an interface to support dynamic dispatch.
type IAdditiveExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllMultiplicativeExpression() []IMultiplicativeExpressionContext
	MultiplicativeExpression(i int) IMultiplicativeExpressionContext
	AllPlus() []antlr.TerminalNode
	Plus(i int) antlr.TerminalNode
	AllMinus() []antlr.TerminalNode
	Minus(i int) antlr.TerminalNode

	// IsAdditiveExpressionContext differentiates from other interfaces.
	IsAdditiveExpressionContext()
}

type AdditiveExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdditiveExpressionContext() *AdditiveExpressionContext {
	var p = new(AdditiveExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_additiveExpression
	return p
}

func InitEmptyAdditiveExpressionContext(p *AdditiveExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_additiveExpression
}

func (*AdditiveExpressionContext) IsAdditiveExpressionContext() {}

func NewAdditiveExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdditiveExpressionContext {
	var p = new(AdditiveExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_additiveExpression

	return p
}

func (s *AdditiveExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AdditiveExpressionContext) AllMultiplicativeExpression() []IMultiplicativeExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMultiplicativeExpressionContext); ok {
			len++
		}
	}

	tst := make([]IMultiplicativeExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMultiplicativeExpressionContext); ok {
			tst[i] = t.(IMultiplicativeExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AdditiveExpressionContext) MultiplicativeExpression(i int) IMultiplicativeExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiplicativeExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiplicativeExpressionContext)
}

func (s *AdditiveExpressionContext) AllPlus() []antlr.TerminalNode {
	return s.GetTokens(tinycParserPlus)
}

func (s *AdditiveExpressionContext) Plus(i int) antlr.TerminalNode {
	return s.GetToken(tinycParserPlus, i)
}

func (s *AdditiveExpressionContext) AllMinus() []antlr.TerminalNode {
	return s.GetTokens(tinycParserMinus)
}

func (s *AdditiveExpressionContext) Minus(i int) antlr.TerminalNode {
	return s.GetToken(tinycParserMinus, i)
}

func (s *AdditiveExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AdditiveExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterAdditiveExpression(s)
	}
}

func (s *AdditiveExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitAdditiveExpression(s)
	}
}

func (s *AdditiveExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitAdditiveExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) AdditiveExpression() (localctx IAdditiveExpressionContext) {
	localctx = NewAdditiveExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, tinycParserRULE_additiveExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(232)
		p.MultiplicativeExpression()
	}
	p.SetState(237)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == tinycParserPlus || _la == tinycParserMinus {
		{
			p.SetState(233)
			_la = p.GetTokenStream().LA(1)

			if !(_la == tinycParserPlus || _la == tinycParserMinus) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(234)
			p.MultiplicativeExpression()
		}

		p.SetState(239)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRelationalExpressionContext is an interface to support dynamic dispatch.
type IRelationalExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAdditiveExpression() []IAdditiveExpressionContext
	AdditiveExpression(i int) IAdditiveExpressionContext
	AllLess() []antlr.TerminalNode
	Less(i int) antlr.TerminalNode
	AllGreater() []antlr.TerminalNode
	Greater(i int) antlr.TerminalNode
	AllLessEqual() []antlr.TerminalNode
	LessEqual(i int) antlr.TerminalNode
	AllGreaterEqual() []antlr.TerminalNode
	GreaterEqual(i int) antlr.TerminalNode

	// IsRelationalExpressionContext differentiates from other interfaces.
	IsRelationalExpressionContext()
}

type RelationalExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationalExpressionContext() *RelationalExpressionContext {
	var p = new(RelationalExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_relationalExpression
	return p
}

func InitEmptyRelationalExpressionContext(p *RelationalExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_relationalExpression
}

func (*RelationalExpressionContext) IsRelationalExpressionContext() {}

func NewRelationalExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationalExpressionContext {
	var p = new(RelationalExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_relationalExpression

	return p
}

func (s *RelationalExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationalExpressionContext) AllAdditiveExpression() []IAdditiveExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAdditiveExpressionContext); ok {
			len++
		}
	}

	tst := make([]IAdditiveExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAdditiveExpressionContext); ok {
			tst[i] = t.(IAdditiveExpressionContext)
			i++
		}
	}

	return tst
}

func (s *RelationalExpressionContext) AdditiveExpression(i int) IAdditiveExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAdditiveExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAdditiveExpressionContext)
}

func (s *RelationalExpressionContext) AllLess() []antlr.TerminalNode {
	return s.GetTokens(tinycParserLess)
}

func (s *RelationalExpressionContext) Less(i int) antlr.TerminalNode {
	return s.GetToken(tinycParserLess, i)
}

func (s *RelationalExpressionContext) AllGreater() []antlr.TerminalNode {
	return s.GetTokens(tinycParserGreater)
}

func (s *RelationalExpressionContext) Greater(i int) antlr.TerminalNode {
	return s.GetToken(tinycParserGreater, i)
}

func (s *RelationalExpressionContext) AllLessEqual() []antlr.TerminalNode {
	return s.GetTokens(tinycParserLessEqual)
}

func (s *RelationalExpressionContext) LessEqual(i int) antlr.TerminalNode {
	return s.GetToken(tinycParserLessEqual, i)
}

func (s *RelationalExpressionContext) AllGreaterEqual() []antlr.TerminalNode {
	return s.GetTokens(tinycParserGreaterEqual)
}

func (s *RelationalExpressionContext) GreaterEqual(i int) antlr.TerminalNode {
	return s.GetToken(tinycParserGreaterEqual, i)
}

func (s *RelationalExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationalExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationalExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterRelationalExpression(s)
	}
}

func (s *RelationalExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitRelationalExpression(s)
	}
}

func (s *RelationalExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitRelationalExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) RelationalExpression() (localctx IRelationalExpressionContext) {
	localctx = NewRelationalExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, tinycParserRULE_relationalExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(240)
		p.AdditiveExpression()
	}
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2061584302080) != 0 {
		{
			p.SetState(241)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2061584302080) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(242)
			p.AdditiveExpression()
		}

		p.SetState(247)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEqualityExpressionContext is an interface to support dynamic dispatch.
type IEqualityExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllRelationalExpression() []IRelationalExpressionContext
	RelationalExpression(i int) IRelationalExpressionContext
	AllEqual() []antlr.TerminalNode
	Equal(i int) antlr.TerminalNode
	AllNotEqual() []antlr.TerminalNode
	NotEqual(i int) antlr.TerminalNode

	// IsEqualityExpressionContext differentiates from other interfaces.
	IsEqualityExpressionContext()
}

type EqualityExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqualityExpressionContext() *EqualityExpressionContext {
	var p = new(EqualityExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_equalityExpression
	return p
}

func InitEmptyEqualityExpressionContext(p *EqualityExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_equalityExpression
}

func (*EqualityExpressionContext) IsEqualityExpressionContext() {}

func NewEqualityExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualityExpressionContext {
	var p = new(EqualityExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_equalityExpression

	return p
}

func (s *EqualityExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualityExpressionContext) AllRelationalExpression() []IRelationalExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRelationalExpressionContext); ok {
			len++
		}
	}

	tst := make([]IRelationalExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRelationalExpressionContext); ok {
			tst[i] = t.(IRelationalExpressionContext)
			i++
		}
	}

	return tst
}

func (s *EqualityExpressionContext) RelationalExpression(i int) IRelationalExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationalExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationalExpressionContext)
}

func (s *EqualityExpressionContext) AllEqual() []antlr.TerminalNode {
	return s.GetTokens(tinycParserEqual)
}

func (s *EqualityExpressionContext) Equal(i int) antlr.TerminalNode {
	return s.GetToken(tinycParserEqual, i)
}

func (s *EqualityExpressionContext) AllNotEqual() []antlr.TerminalNode {
	return s.GetTokens(tinycParserNotEqual)
}

func (s *EqualityExpressionContext) NotEqual(i int) antlr.TerminalNode {
	return s.GetToken(tinycParserNotEqual, i)
}

func (s *EqualityExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqualityExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterEqualityExpression(s)
	}
}

func (s *EqualityExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitEqualityExpression(s)
	}
}

func (s *EqualityExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitEqualityExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) EqualityExpression() (localctx IEqualityExpressionContext) {
	localctx = NewEqualityExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, tinycParserRULE_equalityExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(248)
		p.RelationalExpression()
	}
	p.SetState(253)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == tinycParserEqual || _la == tinycParserNotEqual {
		{
			p.SetState(249)
			_la = p.GetTokenStream().LA(1)

			if !(_la == tinycParserEqual || _la == tinycParserNotEqual) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(250)
			p.RelationalExpression()
		}

		p.SetState(255)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInclusiveOrExpressionContext is an interface to support dynamic dispatch.
type IInclusiveOrExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllEqualityExpression() []IEqualityExpressionContext
	EqualityExpression(i int) IEqualityExpressionContext
	AllOr() []antlr.TerminalNode
	Or(i int) antlr.TerminalNode

	// IsInclusiveOrExpressionContext differentiates from other interfaces.
	IsInclusiveOrExpressionContext()
}

type InclusiveOrExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInclusiveOrExpressionContext() *InclusiveOrExpressionContext {
	var p = new(InclusiveOrExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_inclusiveOrExpression
	return p
}

func InitEmptyInclusiveOrExpressionContext(p *InclusiveOrExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_inclusiveOrExpression
}

func (*InclusiveOrExpressionContext) IsInclusiveOrExpressionContext() {}

func NewInclusiveOrExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InclusiveOrExpressionContext {
	var p = new(InclusiveOrExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_inclusiveOrExpression

	return p
}

func (s *InclusiveOrExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *InclusiveOrExpressionContext) AllEqualityExpression() []IEqualityExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEqualityExpressionContext); ok {
			len++
		}
	}

	tst := make([]IEqualityExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEqualityExpressionContext); ok {
			tst[i] = t.(IEqualityExpressionContext)
			i++
		}
	}

	return tst
}

func (s *InclusiveOrExpressionContext) EqualityExpression(i int) IEqualityExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEqualityExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEqualityExpressionContext)
}

func (s *InclusiveOrExpressionContext) AllOr() []antlr.TerminalNode {
	return s.GetTokens(tinycParserOr)
}

func (s *InclusiveOrExpressionContext) Or(i int) antlr.TerminalNode {
	return s.GetToken(tinycParserOr, i)
}

func (s *InclusiveOrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InclusiveOrExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InclusiveOrExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterInclusiveOrExpression(s)
	}
}

func (s *InclusiveOrExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitInclusiveOrExpression(s)
	}
}

func (s *InclusiveOrExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitInclusiveOrExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) InclusiveOrExpression() (localctx IInclusiveOrExpressionContext) {
	localctx = NewInclusiveOrExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, tinycParserRULE_inclusiveOrExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(256)
		p.EqualityExpression()
	}
	p.SetState(261)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == tinycParserOr {
		{
			p.SetState(257)
			p.Match(tinycParserOr)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(258)
			p.EqualityExpression()
		}

		p.SetState(263)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILogicalAndExpressionContext is an interface to support dynamic dispatch.
type ILogicalAndExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllInclusiveOrExpression() []IInclusiveOrExpressionContext
	InclusiveOrExpression(i int) IInclusiveOrExpressionContext
	AllAndAnd() []antlr.TerminalNode
	AndAnd(i int) antlr.TerminalNode

	// IsLogicalAndExpressionContext differentiates from other interfaces.
	IsLogicalAndExpressionContext()
}

type LogicalAndExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalAndExpressionContext() *LogicalAndExpressionContext {
	var p = new(LogicalAndExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_logicalAndExpression
	return p
}

func InitEmptyLogicalAndExpressionContext(p *LogicalAndExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_logicalAndExpression
}

func (*LogicalAndExpressionContext) IsLogicalAndExpressionContext() {}

func NewLogicalAndExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalAndExpressionContext {
	var p = new(LogicalAndExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_logicalAndExpression

	return p
}

func (s *LogicalAndExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalAndExpressionContext) AllInclusiveOrExpression() []IInclusiveOrExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInclusiveOrExpressionContext); ok {
			len++
		}
	}

	tst := make([]IInclusiveOrExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInclusiveOrExpressionContext); ok {
			tst[i] = t.(IInclusiveOrExpressionContext)
			i++
		}
	}

	return tst
}

func (s *LogicalAndExpressionContext) InclusiveOrExpression(i int) IInclusiveOrExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInclusiveOrExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInclusiveOrExpressionContext)
}

func (s *LogicalAndExpressionContext) AllAndAnd() []antlr.TerminalNode {
	return s.GetTokens(tinycParserAndAnd)
}

func (s *LogicalAndExpressionContext) AndAnd(i int) antlr.TerminalNode {
	return s.GetToken(tinycParserAndAnd, i)
}

func (s *LogicalAndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalAndExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicalAndExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterLogicalAndExpression(s)
	}
}

func (s *LogicalAndExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitLogicalAndExpression(s)
	}
}

func (s *LogicalAndExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitLogicalAndExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) LogicalAndExpression() (localctx ILogicalAndExpressionContext) {
	localctx = NewLogicalAndExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, tinycParserRULE_logicalAndExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(264)
		p.InclusiveOrExpression()
	}
	p.SetState(269)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == tinycParserAndAnd {
		{
			p.SetState(265)
			p.Match(tinycParserAndAnd)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(266)
			p.InclusiveOrExpression()
		}

		p.SetState(271)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILogicalOrExpressionContext is an interface to support dynamic dispatch.
type ILogicalOrExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllLogicalAndExpression() []ILogicalAndExpressionContext
	LogicalAndExpression(i int) ILogicalAndExpressionContext
	AllOrOr() []antlr.TerminalNode
	OrOr(i int) antlr.TerminalNode

	// IsLogicalOrExpressionContext differentiates from other interfaces.
	IsLogicalOrExpressionContext()
}

type LogicalOrExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalOrExpressionContext() *LogicalOrExpressionContext {
	var p = new(LogicalOrExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_logicalOrExpression
	return p
}

func InitEmptyLogicalOrExpressionContext(p *LogicalOrExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_logicalOrExpression
}

func (*LogicalOrExpressionContext) IsLogicalOrExpressionContext() {}

func NewLogicalOrExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalOrExpressionContext {
	var p = new(LogicalOrExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_logicalOrExpression

	return p
}

func (s *LogicalOrExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalOrExpressionContext) AllLogicalAndExpression() []ILogicalAndExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILogicalAndExpressionContext); ok {
			len++
		}
	}

	tst := make([]ILogicalAndExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILogicalAndExpressionContext); ok {
			tst[i] = t.(ILogicalAndExpressionContext)
			i++
		}
	}

	return tst
}

func (s *LogicalOrExpressionContext) LogicalAndExpression(i int) ILogicalAndExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicalAndExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogicalAndExpressionContext)
}

func (s *LogicalOrExpressionContext) AllOrOr() []antlr.TerminalNode {
	return s.GetTokens(tinycParserOrOr)
}

func (s *LogicalOrExpressionContext) OrOr(i int) antlr.TerminalNode {
	return s.GetToken(tinycParserOrOr, i)
}

func (s *LogicalOrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalOrExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicalOrExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterLogicalOrExpression(s)
	}
}

func (s *LogicalOrExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitLogicalOrExpression(s)
	}
}

func (s *LogicalOrExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitLogicalOrExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) LogicalOrExpression() (localctx ILogicalOrExpressionContext) {
	localctx = NewLogicalOrExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, tinycParserRULE_logicalOrExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(272)
		p.LogicalAndExpression()
	}
	p.SetState(277)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == tinycParserOrOr {
		{
			p.SetState(273)
			p.Match(tinycParserOrOr)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(274)
			p.LogicalAndExpression()
		}

		p.SetState(279)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignmentExpressionContext is an interface to support dynamic dispatch.
type IAssignmentExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LogicalOrExpression() ILogicalOrExpressionContext
	UnaryExpression() IUnaryExpressionContext
	AssignmentOperator() IAssignmentOperatorContext
	AssignmentExpression() IAssignmentExpressionContext
	DigitSequence() antlr.TerminalNode

	// IsAssignmentExpressionContext differentiates from other interfaces.
	IsAssignmentExpressionContext()
}

type AssignmentExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentExpressionContext() *AssignmentExpressionContext {
	var p = new(AssignmentExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_assignmentExpression
	return p
}

func InitEmptyAssignmentExpressionContext(p *AssignmentExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_assignmentExpression
}

func (*AssignmentExpressionContext) IsAssignmentExpressionContext() {}

func NewAssignmentExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentExpressionContext {
	var p = new(AssignmentExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_assignmentExpression

	return p
}

func (s *AssignmentExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentExpressionContext) LogicalOrExpression() ILogicalOrExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicalOrExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogicalOrExpressionContext)
}

func (s *AssignmentExpressionContext) UnaryExpression() IUnaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryExpressionContext)
}

func (s *AssignmentExpressionContext) AssignmentOperator() IAssignmentOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentOperatorContext)
}

func (s *AssignmentExpressionContext) AssignmentExpression() IAssignmentExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentExpressionContext)
}

func (s *AssignmentExpressionContext) DigitSequence() antlr.TerminalNode {
	return s.GetToken(tinycParserDigitSequence, 0)
}

func (s *AssignmentExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterAssignmentExpression(s)
	}
}

func (s *AssignmentExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitAssignmentExpression(s)
	}
}

func (s *AssignmentExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitAssignmentExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) AssignmentExpression() (localctx IAssignmentExpressionContext) {
	localctx = NewAssignmentExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, tinycParserRULE_assignmentExpression)
	p.SetState(286)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(280)
			p.LogicalOrExpression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(281)
			p.UnaryExpression()
		}
		{
			p.SetState(282)
			p.AssignmentOperator()
		}
		{
			p.SetState(283)
			p.AssignmentExpression()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(285)
			p.Match(tinycParserDigitSequence)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignmentOperatorContext is an interface to support dynamic dispatch.
type IAssignmentOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Assign() antlr.TerminalNode

	// IsAssignmentOperatorContext differentiates from other interfaces.
	IsAssignmentOperatorContext()
}

type AssignmentOperatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentOperatorContext() *AssignmentOperatorContext {
	var p = new(AssignmentOperatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_assignmentOperator
	return p
}

func InitEmptyAssignmentOperatorContext(p *AssignmentOperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_assignmentOperator
}

func (*AssignmentOperatorContext) IsAssignmentOperatorContext() {}

func NewAssignmentOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentOperatorContext {
	var p = new(AssignmentOperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_assignmentOperator

	return p
}

func (s *AssignmentOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentOperatorContext) Assign() antlr.TerminalNode {
	return s.GetToken(tinycParserAssign, 0)
}

func (s *AssignmentOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterAssignmentOperator(s)
	}
}

func (s *AssignmentOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitAssignmentOperator(s)
	}
}

func (s *AssignmentOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitAssignmentOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) AssignmentOperator() (localctx IAssignmentOperatorContext) {
	localctx = NewAssignmentOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, tinycParserRULE_assignmentOperator)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(288)
		p.Match(tinycParserAssign)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAssignmentExpression() []IAssignmentExpressionContext
	AssignmentExpression(i int) IAssignmentExpressionContext
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) AllAssignmentExpression() []IAssignmentExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAssignmentExpressionContext); ok {
			len++
		}
	}

	tst := make([]IAssignmentExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAssignmentExpressionContext); ok {
			tst[i] = t.(IAssignmentExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionContext) AssignmentExpression(i int) IAssignmentExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentExpressionContext)
}

func (s *ExpressionContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(tinycParserComma)
}

func (s *ExpressionContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(tinycParserComma, i)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, tinycParserRULE_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(290)
		p.AssignmentExpression()
	}
	p.SetState(295)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == tinycParserComma {
		{
			p.SetState(291)
			p.Match(tinycParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(292)
			p.AssignmentExpression()
		}

		p.SetState(297)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstantExpressionContext is an interface to support dynamic dispatch.
type IConstantExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LogicalOrExpression() ILogicalOrExpressionContext

	// IsConstantExpressionContext differentiates from other interfaces.
	IsConstantExpressionContext()
}

type ConstantExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantExpressionContext() *ConstantExpressionContext {
	var p = new(ConstantExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_constantExpression
	return p
}

func InitEmptyConstantExpressionContext(p *ConstantExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_constantExpression
}

func (*ConstantExpressionContext) IsConstantExpressionContext() {}

func NewConstantExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantExpressionContext {
	var p = new(ConstantExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_constantExpression

	return p
}

func (s *ConstantExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantExpressionContext) LogicalOrExpression() ILogicalOrExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicalOrExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogicalOrExpressionContext)
}

func (s *ConstantExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterConstantExpression(s)
	}
}

func (s *ConstantExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitConstantExpression(s)
	}
}

func (s *ConstantExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitConstantExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) ConstantExpression() (localctx IConstantExpressionContext) {
	localctx = NewConstantExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, tinycParserRULE_constantExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(298)
		p.LogicalOrExpression()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DeclarationSpecifiers() IDeclarationSpecifiersContext
	Semi() antlr.TerminalNode
	InitDeclaratorList() IInitDeclaratorListContext

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_declaration
	return p
}

func InitEmptyDeclarationContext(p *DeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_declaration
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) DeclarationSpecifiers() IDeclarationSpecifiersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationSpecifiersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationSpecifiersContext)
}

func (s *DeclarationContext) Semi() antlr.TerminalNode {
	return s.GetToken(tinycParserSemi, 0)
}

func (s *DeclarationContext) InitDeclaratorList() IInitDeclaratorListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInitDeclaratorListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInitDeclaratorListContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (s *DeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, tinycParserRULE_declaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(300)
		p.DeclarationSpecifiers()
	}
	p.SetState(302)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-31)) & ^0x3f) == 0 && ((int64(1)<<(_la-31))&68727930881) != 0 {
		{
			p.SetState(301)
			p.InitDeclaratorList()
		}

	}
	{
		p.SetState(304)
		p.Match(tinycParserSemi)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclarationSpecifiersContext is an interface to support dynamic dispatch.
type IDeclarationSpecifiersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDeclarationSpecifier() []IDeclarationSpecifierContext
	DeclarationSpecifier(i int) IDeclarationSpecifierContext

	// IsDeclarationSpecifiersContext differentiates from other interfaces.
	IsDeclarationSpecifiersContext()
}

type DeclarationSpecifiersContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationSpecifiersContext() *DeclarationSpecifiersContext {
	var p = new(DeclarationSpecifiersContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_declarationSpecifiers
	return p
}

func InitEmptyDeclarationSpecifiersContext(p *DeclarationSpecifiersContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_declarationSpecifiers
}

func (*DeclarationSpecifiersContext) IsDeclarationSpecifiersContext() {}

func NewDeclarationSpecifiersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationSpecifiersContext {
	var p = new(DeclarationSpecifiersContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_declarationSpecifiers

	return p
}

func (s *DeclarationSpecifiersContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationSpecifiersContext) AllDeclarationSpecifier() []IDeclarationSpecifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclarationSpecifierContext); ok {
			len++
		}
	}

	tst := make([]IDeclarationSpecifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclarationSpecifierContext); ok {
			tst[i] = t.(IDeclarationSpecifierContext)
			i++
		}
	}

	return tst
}

func (s *DeclarationSpecifiersContext) DeclarationSpecifier(i int) IDeclarationSpecifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationSpecifierContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationSpecifierContext)
}

func (s *DeclarationSpecifiersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationSpecifiersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationSpecifiersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterDeclarationSpecifiers(s)
	}
}

func (s *DeclarationSpecifiersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitDeclarationSpecifiers(s)
	}
}

func (s *DeclarationSpecifiersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitDeclarationSpecifiers(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) DeclarationSpecifiers() (localctx IDeclarationSpecifiersContext) {
	localctx = NewDeclarationSpecifiersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, tinycParserRULE_declarationSpecifiers)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(307)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(306)
				p.DeclarationSpecifier()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(309)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclarationSpecifiers2Context is an interface to support dynamic dispatch.
type IDeclarationSpecifiers2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDeclarationSpecifier() []IDeclarationSpecifierContext
	DeclarationSpecifier(i int) IDeclarationSpecifierContext

	// IsDeclarationSpecifiers2Context differentiates from other interfaces.
	IsDeclarationSpecifiers2Context()
}

type DeclarationSpecifiers2Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationSpecifiers2Context() *DeclarationSpecifiers2Context {
	var p = new(DeclarationSpecifiers2Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_declarationSpecifiers2
	return p
}

func InitEmptyDeclarationSpecifiers2Context(p *DeclarationSpecifiers2Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_declarationSpecifiers2
}

func (*DeclarationSpecifiers2Context) IsDeclarationSpecifiers2Context() {}

func NewDeclarationSpecifiers2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationSpecifiers2Context {
	var p = new(DeclarationSpecifiers2Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_declarationSpecifiers2

	return p
}

func (s *DeclarationSpecifiers2Context) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationSpecifiers2Context) AllDeclarationSpecifier() []IDeclarationSpecifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclarationSpecifierContext); ok {
			len++
		}
	}

	tst := make([]IDeclarationSpecifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclarationSpecifierContext); ok {
			tst[i] = t.(IDeclarationSpecifierContext)
			i++
		}
	}

	return tst
}

func (s *DeclarationSpecifiers2Context) DeclarationSpecifier(i int) IDeclarationSpecifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationSpecifierContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationSpecifierContext)
}

func (s *DeclarationSpecifiers2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationSpecifiers2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationSpecifiers2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterDeclarationSpecifiers2(s)
	}
}

func (s *DeclarationSpecifiers2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitDeclarationSpecifiers2(s)
	}
}

func (s *DeclarationSpecifiers2Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitDeclarationSpecifiers2(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) DeclarationSpecifiers2() (localctx IDeclarationSpecifiers2Context) {
	localctx = NewDeclarationSpecifiers2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, tinycParserRULE_declarationSpecifiers2)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(312)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64((_la-4)) & ^0x3f) == 0 && ((int64(1)<<(_la-4))&-9223372036828159711) != 0) {
		{
			p.SetState(311)
			p.DeclarationSpecifier()
		}

		p.SetState(314)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclarationSpecifierContext is an interface to support dynamic dispatch.
type IDeclarationSpecifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TypeSpecifier() ITypeSpecifierContext

	// IsDeclarationSpecifierContext differentiates from other interfaces.
	IsDeclarationSpecifierContext()
}

type DeclarationSpecifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationSpecifierContext() *DeclarationSpecifierContext {
	var p = new(DeclarationSpecifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_declarationSpecifier
	return p
}

func InitEmptyDeclarationSpecifierContext(p *DeclarationSpecifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_declarationSpecifier
}

func (*DeclarationSpecifierContext) IsDeclarationSpecifierContext() {}

func NewDeclarationSpecifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationSpecifierContext {
	var p = new(DeclarationSpecifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_declarationSpecifier

	return p
}

func (s *DeclarationSpecifierContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationSpecifierContext) TypeSpecifier() ITypeSpecifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeSpecifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierContext)
}

func (s *DeclarationSpecifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationSpecifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationSpecifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterDeclarationSpecifier(s)
	}
}

func (s *DeclarationSpecifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitDeclarationSpecifier(s)
	}
}

func (s *DeclarationSpecifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitDeclarationSpecifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) DeclarationSpecifier() (localctx IDeclarationSpecifierContext) {
	localctx = NewDeclarationSpecifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, tinycParserRULE_declarationSpecifier)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(316)
		p.TypeSpecifier()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInitDeclaratorListContext is an interface to support dynamic dispatch.
type IInitDeclaratorListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllInitDeclarator() []IInitDeclaratorContext
	InitDeclarator(i int) IInitDeclaratorContext
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

	// IsInitDeclaratorListContext differentiates from other interfaces.
	IsInitDeclaratorListContext()
}

type InitDeclaratorListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInitDeclaratorListContext() *InitDeclaratorListContext {
	var p = new(InitDeclaratorListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_initDeclaratorList
	return p
}

func InitEmptyInitDeclaratorListContext(p *InitDeclaratorListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_initDeclaratorList
}

func (*InitDeclaratorListContext) IsInitDeclaratorListContext() {}

func NewInitDeclaratorListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InitDeclaratorListContext {
	var p = new(InitDeclaratorListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_initDeclaratorList

	return p
}

func (s *InitDeclaratorListContext) GetParser() antlr.Parser { return s.parser }

func (s *InitDeclaratorListContext) AllInitDeclarator() []IInitDeclaratorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInitDeclaratorContext); ok {
			len++
		}
	}

	tst := make([]IInitDeclaratorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInitDeclaratorContext); ok {
			tst[i] = t.(IInitDeclaratorContext)
			i++
		}
	}

	return tst
}

func (s *InitDeclaratorListContext) InitDeclarator(i int) IInitDeclaratorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInitDeclaratorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInitDeclaratorContext)
}

func (s *InitDeclaratorListContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(tinycParserComma)
}

func (s *InitDeclaratorListContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(tinycParserComma, i)
}

func (s *InitDeclaratorListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitDeclaratorListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InitDeclaratorListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterInitDeclaratorList(s)
	}
}

func (s *InitDeclaratorListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitInitDeclaratorList(s)
	}
}

func (s *InitDeclaratorListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitInitDeclaratorList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) InitDeclaratorList() (localctx IInitDeclaratorListContext) {
	localctx = NewInitDeclaratorListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, tinycParserRULE_initDeclaratorList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(318)
		p.InitDeclarator()
	}
	p.SetState(323)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == tinycParserComma {
		{
			p.SetState(319)
			p.Match(tinycParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(320)
			p.InitDeclarator()
		}

		p.SetState(325)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInitDeclaratorContext is an interface to support dynamic dispatch.
type IInitDeclaratorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Declarator() IDeclaratorContext
	Assign() antlr.TerminalNode
	Initializer() IInitializerContext

	// IsInitDeclaratorContext differentiates from other interfaces.
	IsInitDeclaratorContext()
}

type InitDeclaratorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInitDeclaratorContext() *InitDeclaratorContext {
	var p = new(InitDeclaratorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_initDeclarator
	return p
}

func InitEmptyInitDeclaratorContext(p *InitDeclaratorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_initDeclarator
}

func (*InitDeclaratorContext) IsInitDeclaratorContext() {}

func NewInitDeclaratorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InitDeclaratorContext {
	var p = new(InitDeclaratorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_initDeclarator

	return p
}

func (s *InitDeclaratorContext) GetParser() antlr.Parser { return s.parser }

func (s *InitDeclaratorContext) Declarator() IDeclaratorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaratorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclaratorContext)
}

func (s *InitDeclaratorContext) Assign() antlr.TerminalNode {
	return s.GetToken(tinycParserAssign, 0)
}

func (s *InitDeclaratorContext) Initializer() IInitializerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInitializerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInitializerContext)
}

func (s *InitDeclaratorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitDeclaratorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InitDeclaratorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterInitDeclarator(s)
	}
}

func (s *InitDeclaratorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitInitDeclarator(s)
	}
}

func (s *InitDeclaratorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitInitDeclarator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) InitDeclarator() (localctx IInitDeclaratorContext) {
	localctx = NewInitDeclaratorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, tinycParserRULE_initDeclarator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(326)
		p.Declarator()
	}
	p.SetState(329)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == tinycParserAssign {
		{
			p.SetState(327)
			p.Match(tinycParserAssign)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(328)
			p.Initializer()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeSpecifierContext is an interface to support dynamic dispatch.
type ITypeSpecifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Void() antlr.TerminalNode
	Char() antlr.TerminalNode
	Short() antlr.TerminalNode
	Int() antlr.TerminalNode
	Float() antlr.TerminalNode
	Double() antlr.TerminalNode
	Signed() antlr.TerminalNode
	StructOrUnionSpecifier() IStructOrUnionSpecifierContext
	TypedefName() ITypedefNameContext

	// IsTypeSpecifierContext differentiates from other interfaces.
	IsTypeSpecifierContext()
}

type TypeSpecifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeSpecifierContext() *TypeSpecifierContext {
	var p = new(TypeSpecifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_typeSpecifier
	return p
}

func InitEmptyTypeSpecifierContext(p *TypeSpecifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_typeSpecifier
}

func (*TypeSpecifierContext) IsTypeSpecifierContext() {}

func NewTypeSpecifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeSpecifierContext {
	var p = new(TypeSpecifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_typeSpecifier

	return p
}

func (s *TypeSpecifierContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeSpecifierContext) Void() antlr.TerminalNode {
	return s.GetToken(tinycParserVoid, 0)
}

func (s *TypeSpecifierContext) Char() antlr.TerminalNode {
	return s.GetToken(tinycParserChar, 0)
}

func (s *TypeSpecifierContext) Short() antlr.TerminalNode {
	return s.GetToken(tinycParserShort, 0)
}

func (s *TypeSpecifierContext) Int() antlr.TerminalNode {
	return s.GetToken(tinycParserInt, 0)
}

func (s *TypeSpecifierContext) Float() antlr.TerminalNode {
	return s.GetToken(tinycParserFloat, 0)
}

func (s *TypeSpecifierContext) Double() antlr.TerminalNode {
	return s.GetToken(tinycParserDouble, 0)
}

func (s *TypeSpecifierContext) Signed() antlr.TerminalNode {
	return s.GetToken(tinycParserSigned, 0)
}

func (s *TypeSpecifierContext) StructOrUnionSpecifier() IStructOrUnionSpecifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructOrUnionSpecifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructOrUnionSpecifierContext)
}

func (s *TypeSpecifierContext) TypedefName() ITypedefNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypedefNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypedefNameContext)
}

func (s *TypeSpecifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSpecifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeSpecifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterTypeSpecifier(s)
	}
}

func (s *TypeSpecifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitTypeSpecifier(s)
	}
}

func (s *TypeSpecifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitTypeSpecifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) TypeSpecifier() (localctx ITypeSpecifierContext) {
	localctx = NewTypeSpecifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, tinycParserRULE_typeSpecifier)
	p.SetState(340)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case tinycParserVoid:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(331)
			p.Match(tinycParserVoid)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case tinycParserChar:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(332)
			p.Match(tinycParserChar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case tinycParserShort:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(333)
			p.Match(tinycParserShort)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case tinycParserInt:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(334)
			p.Match(tinycParserInt)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case tinycParserFloat:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(335)
			p.Match(tinycParserFloat)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case tinycParserDouble:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(336)
			p.Match(tinycParserDouble)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case tinycParserSigned:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(337)
			p.Match(tinycParserSigned)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case tinycParserStruct, tinycParserUnion:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(338)
			p.StructOrUnionSpecifier()
		}

	case tinycParserIdentifier:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(339)
			p.TypedefName()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructOrUnionSpecifierContext is an interface to support dynamic dispatch.
type IStructOrUnionSpecifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	StructOrUnion() IStructOrUnionContext
	LeftBrace() antlr.TerminalNode
	StructDeclarationList() IStructDeclarationListContext
	RightBrace() antlr.TerminalNode
	Identifier() antlr.TerminalNode

	// IsStructOrUnionSpecifierContext differentiates from other interfaces.
	IsStructOrUnionSpecifierContext()
}

type StructOrUnionSpecifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructOrUnionSpecifierContext() *StructOrUnionSpecifierContext {
	var p = new(StructOrUnionSpecifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_structOrUnionSpecifier
	return p
}

func InitEmptyStructOrUnionSpecifierContext(p *StructOrUnionSpecifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_structOrUnionSpecifier
}

func (*StructOrUnionSpecifierContext) IsStructOrUnionSpecifierContext() {}

func NewStructOrUnionSpecifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructOrUnionSpecifierContext {
	var p = new(StructOrUnionSpecifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_structOrUnionSpecifier

	return p
}

func (s *StructOrUnionSpecifierContext) GetParser() antlr.Parser { return s.parser }

func (s *StructOrUnionSpecifierContext) StructOrUnion() IStructOrUnionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructOrUnionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructOrUnionContext)
}

func (s *StructOrUnionSpecifierContext) LeftBrace() antlr.TerminalNode {
	return s.GetToken(tinycParserLeftBrace, 0)
}

func (s *StructOrUnionSpecifierContext) StructDeclarationList() IStructDeclarationListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructDeclarationListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructDeclarationListContext)
}

func (s *StructOrUnionSpecifierContext) RightBrace() antlr.TerminalNode {
	return s.GetToken(tinycParserRightBrace, 0)
}

func (s *StructOrUnionSpecifierContext) Identifier() antlr.TerminalNode {
	return s.GetToken(tinycParserIdentifier, 0)
}

func (s *StructOrUnionSpecifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructOrUnionSpecifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructOrUnionSpecifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterStructOrUnionSpecifier(s)
	}
}

func (s *StructOrUnionSpecifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitStructOrUnionSpecifier(s)
	}
}

func (s *StructOrUnionSpecifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitStructOrUnionSpecifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) StructOrUnionSpecifier() (localctx IStructOrUnionSpecifierContext) {
	localctx = NewStructOrUnionSpecifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, tinycParserRULE_structOrUnionSpecifier)
	var _la int

	p.SetState(353)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(342)
			p.StructOrUnion()
		}
		p.SetState(344)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == tinycParserIdentifier {
			{
				p.SetState(343)
				p.Match(tinycParserIdentifier)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(346)
			p.Match(tinycParserLeftBrace)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(347)
			p.StructDeclarationList()
		}
		{
			p.SetState(348)
			p.Match(tinycParserRightBrace)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(350)
			p.StructOrUnion()
		}
		{
			p.SetState(351)
			p.Match(tinycParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructOrUnionContext is an interface to support dynamic dispatch.
type IStructOrUnionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Struct() antlr.TerminalNode
	Union() antlr.TerminalNode

	// IsStructOrUnionContext differentiates from other interfaces.
	IsStructOrUnionContext()
}

type StructOrUnionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructOrUnionContext() *StructOrUnionContext {
	var p = new(StructOrUnionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_structOrUnion
	return p
}

func InitEmptyStructOrUnionContext(p *StructOrUnionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_structOrUnion
}

func (*StructOrUnionContext) IsStructOrUnionContext() {}

func NewStructOrUnionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructOrUnionContext {
	var p = new(StructOrUnionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_structOrUnion

	return p
}

func (s *StructOrUnionContext) GetParser() antlr.Parser { return s.parser }

func (s *StructOrUnionContext) Struct() antlr.TerminalNode {
	return s.GetToken(tinycParserStruct, 0)
}

func (s *StructOrUnionContext) Union() antlr.TerminalNode {
	return s.GetToken(tinycParserUnion, 0)
}

func (s *StructOrUnionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructOrUnionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructOrUnionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterStructOrUnion(s)
	}
}

func (s *StructOrUnionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitStructOrUnion(s)
	}
}

func (s *StructOrUnionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitStructOrUnion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) StructOrUnion() (localctx IStructOrUnionContext) {
	localctx = NewStructOrUnionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, tinycParserRULE_structOrUnion)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(355)
		_la = p.GetTokenStream().LA(1)

		if !(_la == tinycParserStruct || _la == tinycParserUnion) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructDeclarationListContext is an interface to support dynamic dispatch.
type IStructDeclarationListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStructDeclaration() []IStructDeclarationContext
	StructDeclaration(i int) IStructDeclarationContext

	// IsStructDeclarationListContext differentiates from other interfaces.
	IsStructDeclarationListContext()
}

type StructDeclarationListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructDeclarationListContext() *StructDeclarationListContext {
	var p = new(StructDeclarationListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_structDeclarationList
	return p
}

func InitEmptyStructDeclarationListContext(p *StructDeclarationListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_structDeclarationList
}

func (*StructDeclarationListContext) IsStructDeclarationListContext() {}

func NewStructDeclarationListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructDeclarationListContext {
	var p = new(StructDeclarationListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_structDeclarationList

	return p
}

func (s *StructDeclarationListContext) GetParser() antlr.Parser { return s.parser }

func (s *StructDeclarationListContext) AllStructDeclaration() []IStructDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStructDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IStructDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStructDeclarationContext); ok {
			tst[i] = t.(IStructDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *StructDeclarationListContext) StructDeclaration(i int) IStructDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructDeclarationContext)
}

func (s *StructDeclarationListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDeclarationListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructDeclarationListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterStructDeclarationList(s)
	}
}

func (s *StructDeclarationListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitStructDeclarationList(s)
	}
}

func (s *StructDeclarationListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitStructDeclarationList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) StructDeclarationList() (localctx IStructDeclarationListContext) {
	localctx = NewStructDeclarationListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, tinycParserRULE_structDeclarationList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(358)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64((_la-4)) & ^0x3f) == 0 && ((int64(1)<<(_la-4))&-9223372036828159711) != 0) {
		{
			p.SetState(357)
			p.StructDeclaration()
		}

		p.SetState(360)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructDeclarationContext is an interface to support dynamic dispatch.
type IStructDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SpecifierQualifierList() ISpecifierQualifierListContext
	StructDeclaratorList() IStructDeclaratorListContext
	Semi() antlr.TerminalNode

	// IsStructDeclarationContext differentiates from other interfaces.
	IsStructDeclarationContext()
}

type StructDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructDeclarationContext() *StructDeclarationContext {
	var p = new(StructDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_structDeclaration
	return p
}

func InitEmptyStructDeclarationContext(p *StructDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_structDeclaration
}

func (*StructDeclarationContext) IsStructDeclarationContext() {}

func NewStructDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructDeclarationContext {
	var p = new(StructDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_structDeclaration

	return p
}

func (s *StructDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *StructDeclarationContext) SpecifierQualifierList() ISpecifierQualifierListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISpecifierQualifierListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISpecifierQualifierListContext)
}

func (s *StructDeclarationContext) StructDeclaratorList() IStructDeclaratorListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructDeclaratorListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructDeclaratorListContext)
}

func (s *StructDeclarationContext) Semi() antlr.TerminalNode {
	return s.GetToken(tinycParserSemi, 0)
}

func (s *StructDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterStructDeclaration(s)
	}
}

func (s *StructDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitStructDeclaration(s)
	}
}

func (s *StructDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitStructDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) StructDeclaration() (localctx IStructDeclarationContext) {
	localctx = NewStructDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, tinycParserRULE_structDeclaration)
	p.SetState(369)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(362)
			p.SpecifierQualifierList()
		}
		{
			p.SetState(363)
			p.StructDeclaratorList()
		}
		{
			p.SetState(364)
			p.Match(tinycParserSemi)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(366)
			p.SpecifierQualifierList()
		}
		{
			p.SetState(367)
			p.Match(tinycParserSemi)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISpecifierQualifierListContext is an interface to support dynamic dispatch.
type ISpecifierQualifierListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TypeSpecifier() ITypeSpecifierContext
	SpecifierQualifierList() ISpecifierQualifierListContext

	// IsSpecifierQualifierListContext differentiates from other interfaces.
	IsSpecifierQualifierListContext()
}

type SpecifierQualifierListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpecifierQualifierListContext() *SpecifierQualifierListContext {
	var p = new(SpecifierQualifierListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_specifierQualifierList
	return p
}

func InitEmptySpecifierQualifierListContext(p *SpecifierQualifierListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_specifierQualifierList
}

func (*SpecifierQualifierListContext) IsSpecifierQualifierListContext() {}

func NewSpecifierQualifierListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpecifierQualifierListContext {
	var p = new(SpecifierQualifierListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_specifierQualifierList

	return p
}

func (s *SpecifierQualifierListContext) GetParser() antlr.Parser { return s.parser }

func (s *SpecifierQualifierListContext) TypeSpecifier() ITypeSpecifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeSpecifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierContext)
}

func (s *SpecifierQualifierListContext) SpecifierQualifierList() ISpecifierQualifierListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISpecifierQualifierListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISpecifierQualifierListContext)
}

func (s *SpecifierQualifierListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpecifierQualifierListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SpecifierQualifierListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterSpecifierQualifierList(s)
	}
}

func (s *SpecifierQualifierListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitSpecifierQualifierList(s)
	}
}

func (s *SpecifierQualifierListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitSpecifierQualifierList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) SpecifierQualifierList() (localctx ISpecifierQualifierListContext) {
	localctx = NewSpecifierQualifierListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, tinycParserRULE_specifierQualifierList)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(371)
		p.TypeSpecifier()
	}
	p.SetState(373)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(372)
			p.SpecifierQualifierList()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructDeclaratorListContext is an interface to support dynamic dispatch.
type IStructDeclaratorListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStructDeclarator() []IStructDeclaratorContext
	StructDeclarator(i int) IStructDeclaratorContext
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

	// IsStructDeclaratorListContext differentiates from other interfaces.
	IsStructDeclaratorListContext()
}

type StructDeclaratorListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructDeclaratorListContext() *StructDeclaratorListContext {
	var p = new(StructDeclaratorListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_structDeclaratorList
	return p
}

func InitEmptyStructDeclaratorListContext(p *StructDeclaratorListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_structDeclaratorList
}

func (*StructDeclaratorListContext) IsStructDeclaratorListContext() {}

func NewStructDeclaratorListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructDeclaratorListContext {
	var p = new(StructDeclaratorListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_structDeclaratorList

	return p
}

func (s *StructDeclaratorListContext) GetParser() antlr.Parser { return s.parser }

func (s *StructDeclaratorListContext) AllStructDeclarator() []IStructDeclaratorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStructDeclaratorContext); ok {
			len++
		}
	}

	tst := make([]IStructDeclaratorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStructDeclaratorContext); ok {
			tst[i] = t.(IStructDeclaratorContext)
			i++
		}
	}

	return tst
}

func (s *StructDeclaratorListContext) StructDeclarator(i int) IStructDeclaratorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructDeclaratorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructDeclaratorContext)
}

func (s *StructDeclaratorListContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(tinycParserComma)
}

func (s *StructDeclaratorListContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(tinycParserComma, i)
}

func (s *StructDeclaratorListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDeclaratorListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructDeclaratorListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterStructDeclaratorList(s)
	}
}

func (s *StructDeclaratorListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitStructDeclaratorList(s)
	}
}

func (s *StructDeclaratorListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitStructDeclaratorList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) StructDeclaratorList() (localctx IStructDeclaratorListContext) {
	localctx = NewStructDeclaratorListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, tinycParserRULE_structDeclaratorList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(375)
		p.StructDeclarator()
	}
	p.SetState(380)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == tinycParserComma {
		{
			p.SetState(376)
			p.Match(tinycParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(377)
			p.StructDeclarator()
		}

		p.SetState(382)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructDeclaratorContext is an interface to support dynamic dispatch.
type IStructDeclaratorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Declarator() IDeclaratorContext
	Colon() antlr.TerminalNode
	ConstantExpression() IConstantExpressionContext

	// IsStructDeclaratorContext differentiates from other interfaces.
	IsStructDeclaratorContext()
}

type StructDeclaratorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructDeclaratorContext() *StructDeclaratorContext {
	var p = new(StructDeclaratorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_structDeclarator
	return p
}

func InitEmptyStructDeclaratorContext(p *StructDeclaratorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_structDeclarator
}

func (*StructDeclaratorContext) IsStructDeclaratorContext() {}

func NewStructDeclaratorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructDeclaratorContext {
	var p = new(StructDeclaratorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_structDeclarator

	return p
}

func (s *StructDeclaratorContext) GetParser() antlr.Parser { return s.parser }

func (s *StructDeclaratorContext) Declarator() IDeclaratorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaratorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclaratorContext)
}

func (s *StructDeclaratorContext) Colon() antlr.TerminalNode {
	return s.GetToken(tinycParserColon, 0)
}

func (s *StructDeclaratorContext) ConstantExpression() IConstantExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantExpressionContext)
}

func (s *StructDeclaratorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDeclaratorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructDeclaratorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterStructDeclarator(s)
	}
}

func (s *StructDeclaratorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitStructDeclarator(s)
	}
}

func (s *StructDeclaratorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitStructDeclarator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) StructDeclarator() (localctx IStructDeclaratorContext) {
	localctx = NewStructDeclaratorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, tinycParserRULE_structDeclarator)
	var _la int

	p.SetState(389)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(383)
			p.Declarator()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(385)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64((_la-31)) & ^0x3f) == 0 && ((int64(1)<<(_la-31))&68727930881) != 0 {
			{
				p.SetState(384)
				p.Declarator()
			}

		}
		{
			p.SetState(387)
			p.Match(tinycParserColon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(388)
			p.ConstantExpression()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclaratorContext is an interface to support dynamic dispatch.
type IDeclaratorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DirectDeclarator() IDirectDeclaratorContext
	Pointer() IPointerContext

	// IsDeclaratorContext differentiates from other interfaces.
	IsDeclaratorContext()
}

type DeclaratorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclaratorContext() *DeclaratorContext {
	var p = new(DeclaratorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_declarator
	return p
}

func InitEmptyDeclaratorContext(p *DeclaratorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_declarator
}

func (*DeclaratorContext) IsDeclaratorContext() {}

func NewDeclaratorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclaratorContext {
	var p = new(DeclaratorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_declarator

	return p
}

func (s *DeclaratorContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclaratorContext) DirectDeclarator() IDirectDeclaratorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDirectDeclaratorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDirectDeclaratorContext)
}

func (s *DeclaratorContext) Pointer() IPointerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPointerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPointerContext)
}

func (s *DeclaratorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclaratorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclaratorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterDeclarator(s)
	}
}

func (s *DeclaratorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitDeclarator(s)
	}
}

func (s *DeclaratorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitDeclarator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) Declarator() (localctx IDeclaratorContext) {
	localctx = NewDeclaratorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, tinycParserRULE_declarator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(392)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == tinycParserStar || _la == tinycParserCaret {
		{
			p.SetState(391)
			p.Pointer()
		}

	}
	{
		p.SetState(394)
		p.directDeclarator(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDirectDeclaratorContext is an interface to support dynamic dispatch.
type IDirectDeclaratorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	LeftParen() antlr.TerminalNode
	Declarator() IDeclaratorContext
	RightParen() antlr.TerminalNode
	Colon() antlr.TerminalNode
	DigitSequence() antlr.TerminalNode
	DirectDeclarator() IDirectDeclaratorContext
	LeftBracket() antlr.TerminalNode
	RightBracket() antlr.TerminalNode
	AssignmentExpression() IAssignmentExpressionContext
	ParameterTypeList() IParameterTypeListContext
	IdentifierList() IIdentifierListContext

	// IsDirectDeclaratorContext differentiates from other interfaces.
	IsDirectDeclaratorContext()
}

type DirectDeclaratorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirectDeclaratorContext() *DirectDeclaratorContext {
	var p = new(DirectDeclaratorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_directDeclarator
	return p
}

func InitEmptyDirectDeclaratorContext(p *DirectDeclaratorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_directDeclarator
}

func (*DirectDeclaratorContext) IsDirectDeclaratorContext() {}

func NewDirectDeclaratorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectDeclaratorContext {
	var p = new(DirectDeclaratorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_directDeclarator

	return p
}

func (s *DirectDeclaratorContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectDeclaratorContext) Identifier() antlr.TerminalNode {
	return s.GetToken(tinycParserIdentifier, 0)
}

func (s *DirectDeclaratorContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(tinycParserLeftParen, 0)
}

func (s *DirectDeclaratorContext) Declarator() IDeclaratorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaratorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclaratorContext)
}

func (s *DirectDeclaratorContext) RightParen() antlr.TerminalNode {
	return s.GetToken(tinycParserRightParen, 0)
}

func (s *DirectDeclaratorContext) Colon() antlr.TerminalNode {
	return s.GetToken(tinycParserColon, 0)
}

func (s *DirectDeclaratorContext) DigitSequence() antlr.TerminalNode {
	return s.GetToken(tinycParserDigitSequence, 0)
}

func (s *DirectDeclaratorContext) DirectDeclarator() IDirectDeclaratorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDirectDeclaratorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDirectDeclaratorContext)
}

func (s *DirectDeclaratorContext) LeftBracket() antlr.TerminalNode {
	return s.GetToken(tinycParserLeftBracket, 0)
}

func (s *DirectDeclaratorContext) RightBracket() antlr.TerminalNode {
	return s.GetToken(tinycParserRightBracket, 0)
}

func (s *DirectDeclaratorContext) AssignmentExpression() IAssignmentExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentExpressionContext)
}

func (s *DirectDeclaratorContext) ParameterTypeList() IParameterTypeListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterTypeListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterTypeListContext)
}

func (s *DirectDeclaratorContext) IdentifierList() IIdentifierListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierListContext)
}

func (s *DirectDeclaratorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectDeclaratorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DirectDeclaratorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterDirectDeclarator(s)
	}
}

func (s *DirectDeclaratorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitDirectDeclarator(s)
	}
}

func (s *DirectDeclaratorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitDirectDeclarator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) DirectDeclarator() (localctx IDirectDeclaratorContext) {
	return p.directDeclarator(0)
}

func (p *tinycParser) directDeclarator(_p int) (localctx IDirectDeclaratorContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewDirectDeclaratorContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IDirectDeclaratorContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 68
	p.EnterRecursionRule(localctx, 68, tinycParserRULE_directDeclarator, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(405)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(397)
			p.Match(tinycParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(398)
			p.Match(tinycParserLeftParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(399)
			p.Declarator()
		}
		{
			p.SetState(400)
			p.Match(tinycParserRightParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		{
			p.SetState(402)
			p.Match(tinycParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(403)
			p.Match(tinycParserColon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(404)
			p.Match(tinycParserDigitSequence)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(426)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(424)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext()) {
			case 1:
				localctx = NewDirectDeclaratorContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, tinycParserRULE_directDeclarator)
				p.SetState(407)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(408)
					p.Match(tinycParserLeftBracket)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(410)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if (int64((_la-31)) & ^0x3f) == 0 && ((int64(1)<<(_la-31))&1030845190145) != 0 {
					{
						p.SetState(409)
						p.AssignmentExpression()
					}

				}
				{
					p.SetState(412)
					p.Match(tinycParserRightBracket)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 2:
				localctx = NewDirectDeclaratorContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, tinycParserRULE_directDeclarator)
				p.SetState(413)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(414)
					p.Match(tinycParserLeftParen)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(415)
					p.ParameterTypeList()
				}
				{
					p.SetState(416)
					p.Match(tinycParserRightParen)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 3:
				localctx = NewDirectDeclaratorContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, tinycParserRULE_directDeclarator)
				p.SetState(418)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(419)
					p.Match(tinycParserLeftParen)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(421)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == tinycParserIdentifier {
					{
						p.SetState(420)
						p.IdentifierList()
					}

				}
				{
					p.SetState(423)
					p.Match(tinycParserRightParen)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(428)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGccAttributeContext is an interface to support dynamic dispatch.
type IGccAttributeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Comma() antlr.TerminalNode
	AllLeftParen() []antlr.TerminalNode
	LeftParen(i int) antlr.TerminalNode
	AllRightParen() []antlr.TerminalNode
	RightParen(i int) antlr.TerminalNode
	ArgumentExpressionList() IArgumentExpressionListContext

	// IsGccAttributeContext differentiates from other interfaces.
	IsGccAttributeContext()
}

type GccAttributeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGccAttributeContext() *GccAttributeContext {
	var p = new(GccAttributeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_gccAttribute
	return p
}

func InitEmptyGccAttributeContext(p *GccAttributeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_gccAttribute
}

func (*GccAttributeContext) IsGccAttributeContext() {}

func NewGccAttributeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GccAttributeContext {
	var p = new(GccAttributeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_gccAttribute

	return p
}

func (s *GccAttributeContext) GetParser() antlr.Parser { return s.parser }

func (s *GccAttributeContext) Comma() antlr.TerminalNode {
	return s.GetToken(tinycParserComma, 0)
}

func (s *GccAttributeContext) AllLeftParen() []antlr.TerminalNode {
	return s.GetTokens(tinycParserLeftParen)
}

func (s *GccAttributeContext) LeftParen(i int) antlr.TerminalNode {
	return s.GetToken(tinycParserLeftParen, i)
}

func (s *GccAttributeContext) AllRightParen() []antlr.TerminalNode {
	return s.GetTokens(tinycParserRightParen)
}

func (s *GccAttributeContext) RightParen(i int) antlr.TerminalNode {
	return s.GetToken(tinycParserRightParen, i)
}

func (s *GccAttributeContext) ArgumentExpressionList() IArgumentExpressionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentExpressionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentExpressionListContext)
}

func (s *GccAttributeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GccAttributeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GccAttributeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterGccAttribute(s)
	}
}

func (s *GccAttributeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitGccAttribute(s)
	}
}

func (s *GccAttributeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitGccAttribute(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) GccAttribute() (localctx IGccAttributeContext) {
	localctx = NewGccAttributeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, tinycParserRULE_gccAttribute)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(429)
		_la = p.GetTokenStream().LA(1)

		if _la <= 0 || ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1152921511049297920) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(435)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == tinycParserLeftParen {
		{
			p.SetState(430)
			p.Match(tinycParserLeftParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(432)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64((_la-31)) & ^0x3f) == 0 && ((int64(1)<<(_la-31))&1030845190145) != 0 {
			{
				p.SetState(431)
				p.ArgumentExpressionList()
			}

		}
		{
			p.SetState(434)
			p.Match(tinycParserRightParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INestedParenthesesBlockContext is an interface to support dynamic dispatch.
type INestedParenthesesBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllLeftParen() []antlr.TerminalNode
	LeftParen(i int) antlr.TerminalNode
	AllNestedParenthesesBlock() []INestedParenthesesBlockContext
	NestedParenthesesBlock(i int) INestedParenthesesBlockContext
	AllRightParen() []antlr.TerminalNode
	RightParen(i int) antlr.TerminalNode

	// IsNestedParenthesesBlockContext differentiates from other interfaces.
	IsNestedParenthesesBlockContext()
}

type NestedParenthesesBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNestedParenthesesBlockContext() *NestedParenthesesBlockContext {
	var p = new(NestedParenthesesBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_nestedParenthesesBlock
	return p
}

func InitEmptyNestedParenthesesBlockContext(p *NestedParenthesesBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_nestedParenthesesBlock
}

func (*NestedParenthesesBlockContext) IsNestedParenthesesBlockContext() {}

func NewNestedParenthesesBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NestedParenthesesBlockContext {
	var p = new(NestedParenthesesBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_nestedParenthesesBlock

	return p
}

func (s *NestedParenthesesBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *NestedParenthesesBlockContext) AllLeftParen() []antlr.TerminalNode {
	return s.GetTokens(tinycParserLeftParen)
}

func (s *NestedParenthesesBlockContext) LeftParen(i int) antlr.TerminalNode {
	return s.GetToken(tinycParserLeftParen, i)
}

func (s *NestedParenthesesBlockContext) AllNestedParenthesesBlock() []INestedParenthesesBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INestedParenthesesBlockContext); ok {
			len++
		}
	}

	tst := make([]INestedParenthesesBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INestedParenthesesBlockContext); ok {
			tst[i] = t.(INestedParenthesesBlockContext)
			i++
		}
	}

	return tst
}

func (s *NestedParenthesesBlockContext) NestedParenthesesBlock(i int) INestedParenthesesBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INestedParenthesesBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INestedParenthesesBlockContext)
}

func (s *NestedParenthesesBlockContext) AllRightParen() []antlr.TerminalNode {
	return s.GetTokens(tinycParserRightParen)
}

func (s *NestedParenthesesBlockContext) RightParen(i int) antlr.TerminalNode {
	return s.GetToken(tinycParserRightParen, i)
}

func (s *NestedParenthesesBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NestedParenthesesBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NestedParenthesesBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterNestedParenthesesBlock(s)
	}
}

func (s *NestedParenthesesBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitNestedParenthesesBlock(s)
	}
}

func (s *NestedParenthesesBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitNestedParenthesesBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) NestedParenthesesBlock() (localctx INestedParenthesesBlockContext) {
	localctx = NewNestedParenthesesBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, tinycParserRULE_nestedParenthesesBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(444)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-4294967298) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&131071) != 0) {
		p.SetState(442)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case tinycParserAuto, tinycParserBreak, tinycParserCase, tinycParserChar, tinycParserConst, tinycParserContinue, tinycParserDefault, tinycParserDo, tinycParserDouble, tinycParserElse, tinycParserExtern, tinycParserFloat, tinycParserFor, tinycParserGoto, tinycParserIf, tinycParserInline, tinycParserInt, tinycParserRegister, tinycParserRestrict, tinycParserReturn, tinycParserShort, tinycParserSigned, tinycParserSizeof, tinycParserStruct, tinycParserSwitch, tinycParserTypedef, tinycParserUnion, tinycParserVoid, tinycParserWhile, tinycParserNoreturn, tinycParserLeftBracket, tinycParserRightBracket, tinycParserLeftBrace, tinycParserRightBrace, tinycParserLess, tinycParserLessEqual, tinycParserGreater, tinycParserGreaterEqual, tinycParserLeftShift, tinycParserRightShift, tinycParserPlus, tinycParserPlusPlus, tinycParserMinus, tinycParserMinusMinus, tinycParserStar, tinycParserDiv, tinycParserMod, tinycParserAnd, tinycParserOr, tinycParserAndAnd, tinycParserOrOr, tinycParserCaret, tinycParserNot, tinycParserTilde, tinycParserQuestion, tinycParserColon, tinycParserSemi, tinycParserComma, tinycParserAssign, tinycParserEqual, tinycParserNotEqual, tinycParserArrow, tinycParserDot, tinycParserEllipsis, tinycParserIdentifier, tinycParserConstant, tinycParserDigitSequence, tinycParserStringLiteral, tinycParserComplexDefine, tinycParserIncludeDirective, tinycParserAsmBlock, tinycParserLineAfterPreprocessing, tinycParserLineDirective, tinycParserPragmaDirective, tinycParserWhitespace, tinycParserNewline, tinycParserBlockComment, tinycParserLineComment:
			{
				p.SetState(437)
				_la = p.GetTokenStream().LA(1)

				if _la <= 0 || _la == tinycParserLeftParen || _la == tinycParserRightParen {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		case tinycParserLeftParen:
			{
				p.SetState(438)
				p.Match(tinycParserLeftParen)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(439)
				p.NestedParenthesesBlock()
			}
			{
				p.SetState(440)
				p.Match(tinycParserRightParen)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(446)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPointerContext is an interface to support dynamic dispatch.
type IPointerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStar() []antlr.TerminalNode
	Star(i int) antlr.TerminalNode
	AllCaret() []antlr.TerminalNode
	Caret(i int) antlr.TerminalNode

	// IsPointerContext differentiates from other interfaces.
	IsPointerContext()
}

type PointerContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPointerContext() *PointerContext {
	var p = new(PointerContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_pointer
	return p
}

func InitEmptyPointerContext(p *PointerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_pointer
}

func (*PointerContext) IsPointerContext() {}

func NewPointerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PointerContext {
	var p = new(PointerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_pointer

	return p
}

func (s *PointerContext) GetParser() antlr.Parser { return s.parser }

func (s *PointerContext) AllStar() []antlr.TerminalNode {
	return s.GetTokens(tinycParserStar)
}

func (s *PointerContext) Star(i int) antlr.TerminalNode {
	return s.GetToken(tinycParserStar, i)
}

func (s *PointerContext) AllCaret() []antlr.TerminalNode {
	return s.GetTokens(tinycParserCaret)
}

func (s *PointerContext) Caret(i int) antlr.TerminalNode {
	return s.GetToken(tinycParserCaret, i)
}

func (s *PointerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PointerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PointerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterPointer(s)
	}
}

func (s *PointerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitPointer(s)
	}
}

func (s *PointerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitPointer(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) Pointer() (localctx IPointerContext) {
	localctx = NewPointerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, tinycParserRULE_pointer)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(448)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == tinycParserStar || _la == tinycParserCaret {
		{
			p.SetState(447)
			_la = p.GetTokenStream().LA(1)

			if !(_la == tinycParserStar || _la == tinycParserCaret) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(450)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParameterTypeListContext is an interface to support dynamic dispatch.
type IParameterTypeListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ParameterList() IParameterListContext
	Comma() antlr.TerminalNode
	Ellipsis() antlr.TerminalNode

	// IsParameterTypeListContext differentiates from other interfaces.
	IsParameterTypeListContext()
}

type ParameterTypeListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterTypeListContext() *ParameterTypeListContext {
	var p = new(ParameterTypeListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_parameterTypeList
	return p
}

func InitEmptyParameterTypeListContext(p *ParameterTypeListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_parameterTypeList
}

func (*ParameterTypeListContext) IsParameterTypeListContext() {}

func NewParameterTypeListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterTypeListContext {
	var p = new(ParameterTypeListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_parameterTypeList

	return p
}

func (s *ParameterTypeListContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterTypeListContext) ParameterList() IParameterListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterListContext)
}

func (s *ParameterTypeListContext) Comma() antlr.TerminalNode {
	return s.GetToken(tinycParserComma, 0)
}

func (s *ParameterTypeListContext) Ellipsis() antlr.TerminalNode {
	return s.GetToken(tinycParserEllipsis, 0)
}

func (s *ParameterTypeListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterTypeListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterTypeListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterParameterTypeList(s)
	}
}

func (s *ParameterTypeListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitParameterTypeList(s)
	}
}

func (s *ParameterTypeListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitParameterTypeList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) ParameterTypeList() (localctx IParameterTypeListContext) {
	localctx = NewParameterTypeListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, tinycParserRULE_parameterTypeList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(452)
		p.ParameterList()
	}
	p.SetState(455)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == tinycParserComma {
		{
			p.SetState(453)
			p.Match(tinycParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(454)
			p.Match(tinycParserEllipsis)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParameterListContext is an interface to support dynamic dispatch.
type IParameterListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllParameterDeclaration() []IParameterDeclarationContext
	ParameterDeclaration(i int) IParameterDeclarationContext
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

	// IsParameterListContext differentiates from other interfaces.
	IsParameterListContext()
}

type ParameterListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterListContext() *ParameterListContext {
	var p = new(ParameterListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_parameterList
	return p
}

func InitEmptyParameterListContext(p *ParameterListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_parameterList
}

func (*ParameterListContext) IsParameterListContext() {}

func NewParameterListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterListContext {
	var p = new(ParameterListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_parameterList

	return p
}

func (s *ParameterListContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterListContext) AllParameterDeclaration() []IParameterDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParameterDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IParameterDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParameterDeclarationContext); ok {
			tst[i] = t.(IParameterDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *ParameterListContext) ParameterDeclaration(i int) IParameterDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterDeclarationContext)
}

func (s *ParameterListContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(tinycParserComma)
}

func (s *ParameterListContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(tinycParserComma, i)
}

func (s *ParameterListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterParameterList(s)
	}
}

func (s *ParameterListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitParameterList(s)
	}
}

func (s *ParameterListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitParameterList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) ParameterList() (localctx IParameterListContext) {
	localctx = NewParameterListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, tinycParserRULE_parameterList)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(457)
		p.ParameterDeclaration()
	}
	p.SetState(462)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 47, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(458)
				p.Match(tinycParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(459)
				p.ParameterDeclaration()
			}

		}
		p.SetState(464)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 47, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParameterDeclarationContext is an interface to support dynamic dispatch.
type IParameterDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DeclarationSpecifiers() IDeclarationSpecifiersContext
	Declarator() IDeclaratorContext
	DeclarationSpecifiers2() IDeclarationSpecifiers2Context
	AbstractDeclarator() IAbstractDeclaratorContext

	// IsParameterDeclarationContext differentiates from other interfaces.
	IsParameterDeclarationContext()
}

type ParameterDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterDeclarationContext() *ParameterDeclarationContext {
	var p = new(ParameterDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_parameterDeclaration
	return p
}

func InitEmptyParameterDeclarationContext(p *ParameterDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_parameterDeclaration
}

func (*ParameterDeclarationContext) IsParameterDeclarationContext() {}

func NewParameterDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterDeclarationContext {
	var p = new(ParameterDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_parameterDeclaration

	return p
}

func (s *ParameterDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterDeclarationContext) DeclarationSpecifiers() IDeclarationSpecifiersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationSpecifiersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationSpecifiersContext)
}

func (s *ParameterDeclarationContext) Declarator() IDeclaratorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaratorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclaratorContext)
}

func (s *ParameterDeclarationContext) DeclarationSpecifiers2() IDeclarationSpecifiers2Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationSpecifiers2Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationSpecifiers2Context)
}

func (s *ParameterDeclarationContext) AbstractDeclarator() IAbstractDeclaratorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAbstractDeclaratorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAbstractDeclaratorContext)
}

func (s *ParameterDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterParameterDeclaration(s)
	}
}

func (s *ParameterDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitParameterDeclaration(s)
	}
}

func (s *ParameterDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitParameterDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) ParameterDeclaration() (localctx IParameterDeclarationContext) {
	localctx = NewParameterDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, tinycParserRULE_parameterDeclaration)
	var _la int

	p.SetState(472)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 49, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(465)
			p.DeclarationSpecifiers()
		}
		{
			p.SetState(466)
			p.Declarator()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(468)
			p.DeclarationSpecifiers2()
		}
		p.SetState(470)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&18155146735255552) != 0 {
			{
				p.SetState(469)
				p.AbstractDeclarator()
			}

		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentifierListContext is an interface to support dynamic dispatch.
type IIdentifierListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentifier() []antlr.TerminalNode
	Identifier(i int) antlr.TerminalNode
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

	// IsIdentifierListContext differentiates from other interfaces.
	IsIdentifierListContext()
}

type IdentifierListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierListContext() *IdentifierListContext {
	var p = new(IdentifierListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_identifierList
	return p
}

func InitEmptyIdentifierListContext(p *IdentifierListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_identifierList
}

func (*IdentifierListContext) IsIdentifierListContext() {}

func NewIdentifierListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierListContext {
	var p = new(IdentifierListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_identifierList

	return p
}

func (s *IdentifierListContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierListContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(tinycParserIdentifier)
}

func (s *IdentifierListContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(tinycParserIdentifier, i)
}

func (s *IdentifierListContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(tinycParserComma)
}

func (s *IdentifierListContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(tinycParserComma, i)
}

func (s *IdentifierListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterIdentifierList(s)
	}
}

func (s *IdentifierListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitIdentifierList(s)
	}
}

func (s *IdentifierListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitIdentifierList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) IdentifierList() (localctx IIdentifierListContext) {
	localctx = NewIdentifierListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, tinycParserRULE_identifierList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(474)
		p.Match(tinycParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(479)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == tinycParserComma {
		{
			p.SetState(475)
			p.Match(tinycParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(476)
			p.Match(tinycParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(481)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeNameContext is an interface to support dynamic dispatch.
type ITypeNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SpecifierQualifierList() ISpecifierQualifierListContext
	AbstractDeclarator() IAbstractDeclaratorContext

	// IsTypeNameContext differentiates from other interfaces.
	IsTypeNameContext()
}

type TypeNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeNameContext() *TypeNameContext {
	var p = new(TypeNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_typeName
	return p
}

func InitEmptyTypeNameContext(p *TypeNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_typeName
}

func (*TypeNameContext) IsTypeNameContext() {}

func NewTypeNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeNameContext {
	var p = new(TypeNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_typeName

	return p
}

func (s *TypeNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeNameContext) SpecifierQualifierList() ISpecifierQualifierListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISpecifierQualifierListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISpecifierQualifierListContext)
}

func (s *TypeNameContext) AbstractDeclarator() IAbstractDeclaratorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAbstractDeclaratorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAbstractDeclaratorContext)
}

func (s *TypeNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterTypeName(s)
	}
}

func (s *TypeNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitTypeName(s)
	}
}

func (s *TypeNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitTypeName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) TypeName() (localctx ITypeNameContext) {
	localctx = NewTypeNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, tinycParserRULE_typeName)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(482)
		p.SpecifierQualifierList()
	}
	p.SetState(484)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&18155146735255552) != 0 {
		{
			p.SetState(483)
			p.AbstractDeclarator()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAbstractDeclaratorContext is an interface to support dynamic dispatch.
type IAbstractDeclaratorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Pointer() IPointerContext
	DirectAbstractDeclarator() IDirectAbstractDeclaratorContext

	// IsAbstractDeclaratorContext differentiates from other interfaces.
	IsAbstractDeclaratorContext()
}

type AbstractDeclaratorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAbstractDeclaratorContext() *AbstractDeclaratorContext {
	var p = new(AbstractDeclaratorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_abstractDeclarator
	return p
}

func InitEmptyAbstractDeclaratorContext(p *AbstractDeclaratorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_abstractDeclarator
}

func (*AbstractDeclaratorContext) IsAbstractDeclaratorContext() {}

func NewAbstractDeclaratorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AbstractDeclaratorContext {
	var p = new(AbstractDeclaratorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_abstractDeclarator

	return p
}

func (s *AbstractDeclaratorContext) GetParser() antlr.Parser { return s.parser }

func (s *AbstractDeclaratorContext) Pointer() IPointerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPointerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPointerContext)
}

func (s *AbstractDeclaratorContext) DirectAbstractDeclarator() IDirectAbstractDeclaratorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDirectAbstractDeclaratorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDirectAbstractDeclaratorContext)
}

func (s *AbstractDeclaratorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AbstractDeclaratorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AbstractDeclaratorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterAbstractDeclarator(s)
	}
}

func (s *AbstractDeclaratorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitAbstractDeclarator(s)
	}
}

func (s *AbstractDeclaratorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitAbstractDeclarator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) AbstractDeclarator() (localctx IAbstractDeclaratorContext) {
	localctx = NewAbstractDeclaratorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, tinycParserRULE_abstractDeclarator)
	var _la int

	p.SetState(491)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 53, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(486)
			p.Pointer()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(488)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == tinycParserStar || _la == tinycParserCaret {
			{
				p.SetState(487)
				p.Pointer()
			}

		}
		{
			p.SetState(490)
			p.directAbstractDeclarator(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDirectAbstractDeclaratorContext is an interface to support dynamic dispatch.
type IDirectAbstractDeclaratorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LeftParen() antlr.TerminalNode
	AbstractDeclarator() IAbstractDeclaratorContext
	RightParen() antlr.TerminalNode
	LeftBracket() antlr.TerminalNode
	RightBracket() antlr.TerminalNode
	AssignmentExpression() IAssignmentExpressionContext
	Star() antlr.TerminalNode
	ParameterTypeList() IParameterTypeListContext
	DirectAbstractDeclarator() IDirectAbstractDeclaratorContext

	// IsDirectAbstractDeclaratorContext differentiates from other interfaces.
	IsDirectAbstractDeclaratorContext()
}

type DirectAbstractDeclaratorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirectAbstractDeclaratorContext() *DirectAbstractDeclaratorContext {
	var p = new(DirectAbstractDeclaratorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_directAbstractDeclarator
	return p
}

func InitEmptyDirectAbstractDeclaratorContext(p *DirectAbstractDeclaratorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_directAbstractDeclarator
}

func (*DirectAbstractDeclaratorContext) IsDirectAbstractDeclaratorContext() {}

func NewDirectAbstractDeclaratorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectAbstractDeclaratorContext {
	var p = new(DirectAbstractDeclaratorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_directAbstractDeclarator

	return p
}

func (s *DirectAbstractDeclaratorContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectAbstractDeclaratorContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(tinycParserLeftParen, 0)
}

func (s *DirectAbstractDeclaratorContext) AbstractDeclarator() IAbstractDeclaratorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAbstractDeclaratorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAbstractDeclaratorContext)
}

func (s *DirectAbstractDeclaratorContext) RightParen() antlr.TerminalNode {
	return s.GetToken(tinycParserRightParen, 0)
}

func (s *DirectAbstractDeclaratorContext) LeftBracket() antlr.TerminalNode {
	return s.GetToken(tinycParserLeftBracket, 0)
}

func (s *DirectAbstractDeclaratorContext) RightBracket() antlr.TerminalNode {
	return s.GetToken(tinycParserRightBracket, 0)
}

func (s *DirectAbstractDeclaratorContext) AssignmentExpression() IAssignmentExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentExpressionContext)
}

func (s *DirectAbstractDeclaratorContext) Star() antlr.TerminalNode {
	return s.GetToken(tinycParserStar, 0)
}

func (s *DirectAbstractDeclaratorContext) ParameterTypeList() IParameterTypeListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterTypeListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterTypeListContext)
}

func (s *DirectAbstractDeclaratorContext) DirectAbstractDeclarator() IDirectAbstractDeclaratorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDirectAbstractDeclaratorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDirectAbstractDeclaratorContext)
}

func (s *DirectAbstractDeclaratorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectAbstractDeclaratorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DirectAbstractDeclaratorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterDirectAbstractDeclarator(s)
	}
}

func (s *DirectAbstractDeclaratorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitDirectAbstractDeclarator(s)
	}
}

func (s *DirectAbstractDeclaratorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitDirectAbstractDeclarator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) DirectAbstractDeclarator() (localctx IDirectAbstractDeclaratorContext) {
	return p.directAbstractDeclarator(0)
}

func (p *tinycParser) directAbstractDeclarator(_p int) (localctx IDirectAbstractDeclaratorContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewDirectAbstractDeclaratorContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IDirectAbstractDeclaratorContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 88
	p.EnterRecursionRule(localctx, 88, tinycParserRULE_directAbstractDeclarator, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(511)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 56, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(494)
			p.Match(tinycParserLeftParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(495)
			p.AbstractDeclarator()
		}
		{
			p.SetState(496)
			p.Match(tinycParserRightParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(498)
			p.Match(tinycParserLeftBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(500)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64((_la-31)) & ^0x3f) == 0 && ((int64(1)<<(_la-31))&1030845190145) != 0 {
			{
				p.SetState(499)
				p.AssignmentExpression()
			}

		}
		{
			p.SetState(502)
			p.Match(tinycParserRightBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		{
			p.SetState(503)
			p.Match(tinycParserLeftBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(504)
			p.Match(tinycParserStar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(505)
			p.Match(tinycParserRightBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		{
			p.SetState(506)
			p.Match(tinycParserLeftParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(508)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64((_la-4)) & ^0x3f) == 0 && ((int64(1)<<(_la-4))&-9223372036828159711) != 0 {
			{
				p.SetState(507)
				p.ParameterTypeList()
			}

		}
		{
			p.SetState(510)
			p.Match(tinycParserRightParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(531)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 60, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(529)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 59, p.GetParserRuleContext()) {
			case 1:
				localctx = NewDirectAbstractDeclaratorContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, tinycParserRULE_directAbstractDeclarator)
				p.SetState(513)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(514)
					p.Match(tinycParserLeftBracket)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(516)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if (int64((_la-31)) & ^0x3f) == 0 && ((int64(1)<<(_la-31))&1030845190145) != 0 {
					{
						p.SetState(515)
						p.AssignmentExpression()
					}

				}
				{
					p.SetState(518)
					p.Match(tinycParserRightBracket)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 2:
				localctx = NewDirectAbstractDeclaratorContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, tinycParserRULE_directAbstractDeclarator)
				p.SetState(519)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(520)
					p.Match(tinycParserLeftBracket)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(521)
					p.Match(tinycParserStar)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(522)
					p.Match(tinycParserRightBracket)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 3:
				localctx = NewDirectAbstractDeclaratorContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, tinycParserRULE_directAbstractDeclarator)
				p.SetState(523)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(524)
					p.Match(tinycParserLeftParen)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(526)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if (int64((_la-4)) & ^0x3f) == 0 && ((int64(1)<<(_la-4))&-9223372036828159711) != 0 {
					{
						p.SetState(525)
						p.ParameterTypeList()
					}

				}
				{
					p.SetState(528)
					p.Match(tinycParserRightParen)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(533)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 60, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypedefNameContext is an interface to support dynamic dispatch.
type ITypedefNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode

	// IsTypedefNameContext differentiates from other interfaces.
	IsTypedefNameContext()
}

type TypedefNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypedefNameContext() *TypedefNameContext {
	var p = new(TypedefNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_typedefName
	return p
}

func InitEmptyTypedefNameContext(p *TypedefNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_typedefName
}

func (*TypedefNameContext) IsTypedefNameContext() {}

func NewTypedefNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypedefNameContext {
	var p = new(TypedefNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_typedefName

	return p
}

func (s *TypedefNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TypedefNameContext) Identifier() antlr.TerminalNode {
	return s.GetToken(tinycParserIdentifier, 0)
}

func (s *TypedefNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypedefNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypedefNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterTypedefName(s)
	}
}

func (s *TypedefNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitTypedefName(s)
	}
}

func (s *TypedefNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitTypedefName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) TypedefName() (localctx ITypedefNameContext) {
	localctx = NewTypedefNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, tinycParserRULE_typedefName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(534)
		p.Match(tinycParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInitializerContext is an interface to support dynamic dispatch.
type IInitializerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AssignmentExpression() IAssignmentExpressionContext
	LeftBrace() antlr.TerminalNode
	InitializerList() IInitializerListContext
	RightBrace() antlr.TerminalNode
	Comma() antlr.TerminalNode

	// IsInitializerContext differentiates from other interfaces.
	IsInitializerContext()
}

type InitializerContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInitializerContext() *InitializerContext {
	var p = new(InitializerContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_initializer
	return p
}

func InitEmptyInitializerContext(p *InitializerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_initializer
}

func (*InitializerContext) IsInitializerContext() {}

func NewInitializerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InitializerContext {
	var p = new(InitializerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_initializer

	return p
}

func (s *InitializerContext) GetParser() antlr.Parser { return s.parser }

func (s *InitializerContext) AssignmentExpression() IAssignmentExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentExpressionContext)
}

func (s *InitializerContext) LeftBrace() antlr.TerminalNode {
	return s.GetToken(tinycParserLeftBrace, 0)
}

func (s *InitializerContext) InitializerList() IInitializerListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInitializerListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInitializerListContext)
}

func (s *InitializerContext) RightBrace() antlr.TerminalNode {
	return s.GetToken(tinycParserRightBrace, 0)
}

func (s *InitializerContext) Comma() antlr.TerminalNode {
	return s.GetToken(tinycParserComma, 0)
}

func (s *InitializerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitializerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InitializerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterInitializer(s)
	}
}

func (s *InitializerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitInitializer(s)
	}
}

func (s *InitializerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitInitializer(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) Initializer() (localctx IInitializerContext) {
	localctx = NewInitializerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, tinycParserRULE_initializer)
	var _la int

	p.SetState(544)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case tinycParserLeftParen, tinycParserPlus, tinycParserMinus, tinycParserStar, tinycParserAnd, tinycParserAndAnd, tinycParserNot, tinycParserTilde, tinycParserIdentifier, tinycParserConstant, tinycParserDigitSequence, tinycParserStringLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(536)
			p.AssignmentExpression()
		}

	case tinycParserLeftBrace:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(537)
			p.Match(tinycParserLeftBrace)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(538)
			p.InitializerList()
		}
		p.SetState(540)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == tinycParserComma {
			{
				p.SetState(539)
				p.Match(tinycParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(542)
			p.Match(tinycParserRightBrace)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInitializerListContext is an interface to support dynamic dispatch.
type IInitializerListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllInitializer() []IInitializerContext
	Initializer(i int) IInitializerContext
	AllDesignation() []IDesignationContext
	Designation(i int) IDesignationContext
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

	// IsInitializerListContext differentiates from other interfaces.
	IsInitializerListContext()
}

type InitializerListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInitializerListContext() *InitializerListContext {
	var p = new(InitializerListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_initializerList
	return p
}

func InitEmptyInitializerListContext(p *InitializerListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_initializerList
}

func (*InitializerListContext) IsInitializerListContext() {}

func NewInitializerListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InitializerListContext {
	var p = new(InitializerListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_initializerList

	return p
}

func (s *InitializerListContext) GetParser() antlr.Parser { return s.parser }

func (s *InitializerListContext) AllInitializer() []IInitializerContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInitializerContext); ok {
			len++
		}
	}

	tst := make([]IInitializerContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInitializerContext); ok {
			tst[i] = t.(IInitializerContext)
			i++
		}
	}

	return tst
}

func (s *InitializerListContext) Initializer(i int) IInitializerContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInitializerContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInitializerContext)
}

func (s *InitializerListContext) AllDesignation() []IDesignationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDesignationContext); ok {
			len++
		}
	}

	tst := make([]IDesignationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDesignationContext); ok {
			tst[i] = t.(IDesignationContext)
			i++
		}
	}

	return tst
}

func (s *InitializerListContext) Designation(i int) IDesignationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDesignationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDesignationContext)
}

func (s *InitializerListContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(tinycParserComma)
}

func (s *InitializerListContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(tinycParserComma, i)
}

func (s *InitializerListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitializerListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InitializerListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterInitializerList(s)
	}
}

func (s *InitializerListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitInitializerList(s)
	}
}

func (s *InitializerListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitInitializerList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) InitializerList() (localctx IInitializerListContext) {
	localctx = NewInitializerListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, tinycParserRULE_initializerList)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(547)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == tinycParserLeftBracket || _la == tinycParserDot {
		{
			p.SetState(546)
			p.Designation()
		}

	}
	{
		p.SetState(549)
		p.Initializer()
	}
	p.SetState(557)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 65, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(550)
				p.Match(tinycParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(552)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == tinycParserLeftBracket || _la == tinycParserDot {
				{
					p.SetState(551)
					p.Designation()
				}

			}
			{
				p.SetState(554)
				p.Initializer()
			}

		}
		p.SetState(559)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 65, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDesignationContext is an interface to support dynamic dispatch.
type IDesignationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DesignatorList() IDesignatorListContext
	Assign() antlr.TerminalNode

	// IsDesignationContext differentiates from other interfaces.
	IsDesignationContext()
}

type DesignationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDesignationContext() *DesignationContext {
	var p = new(DesignationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_designation
	return p
}

func InitEmptyDesignationContext(p *DesignationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_designation
}

func (*DesignationContext) IsDesignationContext() {}

func NewDesignationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DesignationContext {
	var p = new(DesignationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_designation

	return p
}

func (s *DesignationContext) GetParser() antlr.Parser { return s.parser }

func (s *DesignationContext) DesignatorList() IDesignatorListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDesignatorListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDesignatorListContext)
}

func (s *DesignationContext) Assign() antlr.TerminalNode {
	return s.GetToken(tinycParserAssign, 0)
}

func (s *DesignationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DesignationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DesignationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterDesignation(s)
	}
}

func (s *DesignationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitDesignation(s)
	}
}

func (s *DesignationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitDesignation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) Designation() (localctx IDesignationContext) {
	localctx = NewDesignationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, tinycParserRULE_designation)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(560)
		p.DesignatorList()
	}
	{
		p.SetState(561)
		p.Match(tinycParserAssign)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDesignatorListContext is an interface to support dynamic dispatch.
type IDesignatorListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDesignator() []IDesignatorContext
	Designator(i int) IDesignatorContext

	// IsDesignatorListContext differentiates from other interfaces.
	IsDesignatorListContext()
}

type DesignatorListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDesignatorListContext() *DesignatorListContext {
	var p = new(DesignatorListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_designatorList
	return p
}

func InitEmptyDesignatorListContext(p *DesignatorListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_designatorList
}

func (*DesignatorListContext) IsDesignatorListContext() {}

func NewDesignatorListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DesignatorListContext {
	var p = new(DesignatorListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_designatorList

	return p
}

func (s *DesignatorListContext) GetParser() antlr.Parser { return s.parser }

func (s *DesignatorListContext) AllDesignator() []IDesignatorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDesignatorContext); ok {
			len++
		}
	}

	tst := make([]IDesignatorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDesignatorContext); ok {
			tst[i] = t.(IDesignatorContext)
			i++
		}
	}

	return tst
}

func (s *DesignatorListContext) Designator(i int) IDesignatorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDesignatorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDesignatorContext)
}

func (s *DesignatorListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DesignatorListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DesignatorListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterDesignatorList(s)
	}
}

func (s *DesignatorListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitDesignatorList(s)
	}
}

func (s *DesignatorListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitDesignatorList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) DesignatorList() (localctx IDesignatorListContext) {
	localctx = NewDesignatorListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, tinycParserRULE_designatorList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(564)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == tinycParserLeftBracket || _la == tinycParserDot {
		{
			p.SetState(563)
			p.Designator()
		}

		p.SetState(566)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDesignatorContext is an interface to support dynamic dispatch.
type IDesignatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LeftBracket() antlr.TerminalNode
	ConstantExpression() IConstantExpressionContext
	RightBracket() antlr.TerminalNode
	Dot() antlr.TerminalNode
	Identifier() antlr.TerminalNode

	// IsDesignatorContext differentiates from other interfaces.
	IsDesignatorContext()
}

type DesignatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDesignatorContext() *DesignatorContext {
	var p = new(DesignatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_designator
	return p
}

func InitEmptyDesignatorContext(p *DesignatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_designator
}

func (*DesignatorContext) IsDesignatorContext() {}

func NewDesignatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DesignatorContext {
	var p = new(DesignatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_designator

	return p
}

func (s *DesignatorContext) GetParser() antlr.Parser { return s.parser }

func (s *DesignatorContext) LeftBracket() antlr.TerminalNode {
	return s.GetToken(tinycParserLeftBracket, 0)
}

func (s *DesignatorContext) ConstantExpression() IConstantExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantExpressionContext)
}

func (s *DesignatorContext) RightBracket() antlr.TerminalNode {
	return s.GetToken(tinycParserRightBracket, 0)
}

func (s *DesignatorContext) Dot() antlr.TerminalNode {
	return s.GetToken(tinycParserDot, 0)
}

func (s *DesignatorContext) Identifier() antlr.TerminalNode {
	return s.GetToken(tinycParserIdentifier, 0)
}

func (s *DesignatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DesignatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DesignatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterDesignator(s)
	}
}

func (s *DesignatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitDesignator(s)
	}
}

func (s *DesignatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitDesignator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) Designator() (localctx IDesignatorContext) {
	localctx = NewDesignatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, tinycParserRULE_designator)
	p.SetState(574)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case tinycParserLeftBracket:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(568)
			p.Match(tinycParserLeftBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(569)
			p.ConstantExpression()
		}
		{
			p.SetState(570)
			p.Match(tinycParserRightBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case tinycParserDot:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(572)
			p.Match(tinycParserDot)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(573)
			p.Match(tinycParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LabeledStatement() ILabeledStatementContext
	CompoundStatement() ICompoundStatementContext
	ExpressionStatement() IExpressionStatementContext
	SelectionStatement() ISelectionStatementContext
	IterationStatement() IIterationStatementContext
	JumpStatement() IJumpStatementContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) LabeledStatement() ILabeledStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILabeledStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILabeledStatementContext)
}

func (s *StatementContext) CompoundStatement() ICompoundStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICompoundStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICompoundStatementContext)
}

func (s *StatementContext) ExpressionStatement() IExpressionStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionStatementContext)
}

func (s *StatementContext) SelectionStatement() ISelectionStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectionStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectionStatementContext)
}

func (s *StatementContext) IterationStatement() IIterationStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIterationStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIterationStatementContext)
}

func (s *StatementContext) JumpStatement() IJumpStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJumpStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJumpStatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, tinycParserRULE_statement)
	p.SetState(582)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 68, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(576)
			p.LabeledStatement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(577)
			p.CompoundStatement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(578)
			p.ExpressionStatement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(579)
			p.SelectionStatement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(580)
			p.IterationStatement()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(581)
			p.JumpStatement()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILabeledStatementContext is an interface to support dynamic dispatch.
type ILabeledStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	Colon() antlr.TerminalNode
	Statement() IStatementContext
	Case() antlr.TerminalNode
	ConstantExpression() IConstantExpressionContext
	Default() antlr.TerminalNode

	// IsLabeledStatementContext differentiates from other interfaces.
	IsLabeledStatementContext()
}

type LabeledStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabeledStatementContext() *LabeledStatementContext {
	var p = new(LabeledStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_labeledStatement
	return p
}

func InitEmptyLabeledStatementContext(p *LabeledStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_labeledStatement
}

func (*LabeledStatementContext) IsLabeledStatementContext() {}

func NewLabeledStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabeledStatementContext {
	var p = new(LabeledStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_labeledStatement

	return p
}

func (s *LabeledStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *LabeledStatementContext) Identifier() antlr.TerminalNode {
	return s.GetToken(tinycParserIdentifier, 0)
}

func (s *LabeledStatementContext) Colon() antlr.TerminalNode {
	return s.GetToken(tinycParserColon, 0)
}

func (s *LabeledStatementContext) Statement() IStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *LabeledStatementContext) Case() antlr.TerminalNode {
	return s.GetToken(tinycParserCase, 0)
}

func (s *LabeledStatementContext) ConstantExpression() IConstantExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantExpressionContext)
}

func (s *LabeledStatementContext) Default() antlr.TerminalNode {
	return s.GetToken(tinycParserDefault, 0)
}

func (s *LabeledStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabeledStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabeledStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterLabeledStatement(s)
	}
}

func (s *LabeledStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitLabeledStatement(s)
	}
}

func (s *LabeledStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitLabeledStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) LabeledStatement() (localctx ILabeledStatementContext) {
	localctx = NewLabeledStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, tinycParserRULE_labeledStatement)
	p.SetState(595)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case tinycParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(584)
			p.Match(tinycParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(585)
			p.Match(tinycParserColon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(586)
			p.Statement()
		}

	case tinycParserCase:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(587)
			p.Match(tinycParserCase)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(588)
			p.ConstantExpression()
		}
		{
			p.SetState(589)
			p.Match(tinycParserColon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(590)
			p.Statement()
		}

	case tinycParserDefault:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(592)
			p.Match(tinycParserDefault)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(593)
			p.Match(tinycParserColon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(594)
			p.Statement()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICompoundStatementContext is an interface to support dynamic dispatch.
type ICompoundStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LeftBrace() antlr.TerminalNode
	RightBrace() antlr.TerminalNode
	BlockItemList() IBlockItemListContext

	// IsCompoundStatementContext differentiates from other interfaces.
	IsCompoundStatementContext()
}

type CompoundStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompoundStatementContext() *CompoundStatementContext {
	var p = new(CompoundStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_compoundStatement
	return p
}

func InitEmptyCompoundStatementContext(p *CompoundStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_compoundStatement
}

func (*CompoundStatementContext) IsCompoundStatementContext() {}

func NewCompoundStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompoundStatementContext {
	var p = new(CompoundStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_compoundStatement

	return p
}

func (s *CompoundStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *CompoundStatementContext) LeftBrace() antlr.TerminalNode {
	return s.GetToken(tinycParserLeftBrace, 0)
}

func (s *CompoundStatementContext) RightBrace() antlr.TerminalNode {
	return s.GetToken(tinycParserRightBrace, 0)
}

func (s *CompoundStatementContext) BlockItemList() IBlockItemListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockItemListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockItemListContext)
}

func (s *CompoundStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompoundStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompoundStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterCompoundStatement(s)
	}
}

func (s *CompoundStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitCompoundStatement(s)
	}
}

func (s *CompoundStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitCompoundStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) CompoundStatement() (localctx ICompoundStatementContext) {
	localctx = NewCompoundStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, tinycParserRULE_compoundStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(597)
		p.Match(tinycParserLeftBrace)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(599)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&690361398352606172) != 0) || ((int64((_la-67)) & ^0x3f) == 0 && ((int64(1)<<(_la-67))&15) != 0) {
		{
			p.SetState(598)
			p.BlockItemList()
		}

	}
	{
		p.SetState(601)
		p.Match(tinycParserRightBrace)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockItemListContext is an interface to support dynamic dispatch.
type IBlockItemListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllBlockItem() []IBlockItemContext
	BlockItem(i int) IBlockItemContext

	// IsBlockItemListContext differentiates from other interfaces.
	IsBlockItemListContext()
}

type BlockItemListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockItemListContext() *BlockItemListContext {
	var p = new(BlockItemListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_blockItemList
	return p
}

func InitEmptyBlockItemListContext(p *BlockItemListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_blockItemList
}

func (*BlockItemListContext) IsBlockItemListContext() {}

func NewBlockItemListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockItemListContext {
	var p = new(BlockItemListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_blockItemList

	return p
}

func (s *BlockItemListContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockItemListContext) AllBlockItem() []IBlockItemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockItemContext); ok {
			len++
		}
	}

	tst := make([]IBlockItemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockItemContext); ok {
			tst[i] = t.(IBlockItemContext)
			i++
		}
	}

	return tst
}

func (s *BlockItemListContext) BlockItem(i int) IBlockItemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockItemContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockItemContext)
}

func (s *BlockItemListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockItemListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockItemListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterBlockItemList(s)
	}
}

func (s *BlockItemListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitBlockItemList(s)
	}
}

func (s *BlockItemListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitBlockItemList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) BlockItemList() (localctx IBlockItemListContext) {
	localctx = NewBlockItemListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, tinycParserRULE_blockItemList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(604)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&690361398352606172) != 0) || ((int64((_la-67)) & ^0x3f) == 0 && ((int64(1)<<(_la-67))&15) != 0) {
		{
			p.SetState(603)
			p.BlockItem()
		}

		p.SetState(606)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockItemContext is an interface to support dynamic dispatch.
type IBlockItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Statement() IStatementContext
	Declaration() IDeclarationContext

	// IsBlockItemContext differentiates from other interfaces.
	IsBlockItemContext()
}

type BlockItemContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockItemContext() *BlockItemContext {
	var p = new(BlockItemContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_blockItem
	return p
}

func InitEmptyBlockItemContext(p *BlockItemContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_blockItem
}

func (*BlockItemContext) IsBlockItemContext() {}

func NewBlockItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockItemContext {
	var p = new(BlockItemContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_blockItem

	return p
}

func (s *BlockItemContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockItemContext) Statement() IStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockItemContext) Declaration() IDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *BlockItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterBlockItem(s)
	}
}

func (s *BlockItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitBlockItem(s)
	}
}

func (s *BlockItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitBlockItem(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) BlockItem() (localctx IBlockItemContext) {
	localctx = NewBlockItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, tinycParserRULE_blockItem)
	p.SetState(610)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 72, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(608)
			p.Statement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(609)
			p.Declaration()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionStatementContext is an interface to support dynamic dispatch.
type IExpressionStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Semi() antlr.TerminalNode
	Expression() IExpressionContext

	// IsExpressionStatementContext differentiates from other interfaces.
	IsExpressionStatementContext()
}

type ExpressionStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionStatementContext() *ExpressionStatementContext {
	var p = new(ExpressionStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_expressionStatement
	return p
}

func InitEmptyExpressionStatementContext(p *ExpressionStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_expressionStatement
}

func (*ExpressionStatementContext) IsExpressionStatementContext() {}

func NewExpressionStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionStatementContext {
	var p = new(ExpressionStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_expressionStatement

	return p
}

func (s *ExpressionStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionStatementContext) Semi() antlr.TerminalNode {
	return s.GetToken(tinycParserSemi, 0)
}

func (s *ExpressionStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterExpressionStatement(s)
	}
}

func (s *ExpressionStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitExpressionStatement(s)
	}
}

func (s *ExpressionStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitExpressionStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) ExpressionStatement() (localctx IExpressionStatementContext) {
	localctx = NewExpressionStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, tinycParserRULE_expressionStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(613)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-31)) & ^0x3f) == 0 && ((int64(1)<<(_la-31))&1030845190145) != 0 {
		{
			p.SetState(612)
			p.Expression()
		}

	}
	{
		p.SetState(615)
		p.Match(tinycParserSemi)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISelectionStatementContext is an interface to support dynamic dispatch.
type ISelectionStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	If() antlr.TerminalNode
	LeftParen() antlr.TerminalNode
	Expression() IExpressionContext
	RightParen() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext
	Else() antlr.TerminalNode
	Switch() antlr.TerminalNode

	// IsSelectionStatementContext differentiates from other interfaces.
	IsSelectionStatementContext()
}

type SelectionStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectionStatementContext() *SelectionStatementContext {
	var p = new(SelectionStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_selectionStatement
	return p
}

func InitEmptySelectionStatementContext(p *SelectionStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_selectionStatement
}

func (*SelectionStatementContext) IsSelectionStatementContext() {}

func NewSelectionStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectionStatementContext {
	var p = new(SelectionStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_selectionStatement

	return p
}

func (s *SelectionStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectionStatementContext) If() antlr.TerminalNode {
	return s.GetToken(tinycParserIf, 0)
}

func (s *SelectionStatementContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(tinycParserLeftParen, 0)
}

func (s *SelectionStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SelectionStatementContext) RightParen() antlr.TerminalNode {
	return s.GetToken(tinycParserRightParen, 0)
}

func (s *SelectionStatementContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *SelectionStatementContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *SelectionStatementContext) Else() antlr.TerminalNode {
	return s.GetToken(tinycParserElse, 0)
}

func (s *SelectionStatementContext) Switch() antlr.TerminalNode {
	return s.GetToken(tinycParserSwitch, 0)
}

func (s *SelectionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectionStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectionStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterSelectionStatement(s)
	}
}

func (s *SelectionStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitSelectionStatement(s)
	}
}

func (s *SelectionStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitSelectionStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) SelectionStatement() (localctx ISelectionStatementContext) {
	localctx = NewSelectionStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, tinycParserRULE_selectionStatement)
	p.SetState(632)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case tinycParserIf:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(617)
			p.Match(tinycParserIf)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(618)
			p.Match(tinycParserLeftParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(619)
			p.Expression()
		}
		{
			p.SetState(620)
			p.Match(tinycParserRightParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(621)
			p.Statement()
		}
		p.SetState(624)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 74, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(622)
				p.Match(tinycParserElse)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(623)
				p.Statement()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case tinycParserSwitch:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(626)
			p.Match(tinycParserSwitch)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(627)
			p.Match(tinycParserLeftParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(628)
			p.Expression()
		}
		{
			p.SetState(629)
			p.Match(tinycParserRightParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(630)
			p.Statement()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIterationStatementContext is an interface to support dynamic dispatch.
type IIterationStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	While() antlr.TerminalNode
	LeftParen() antlr.TerminalNode
	Expression() IExpressionContext
	RightParen() antlr.TerminalNode
	Statement() IStatementContext
	Do() antlr.TerminalNode
	Semi() antlr.TerminalNode
	For() antlr.TerminalNode
	ForCondition() IForConditionContext

	// IsIterationStatementContext differentiates from other interfaces.
	IsIterationStatementContext()
}

type IterationStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIterationStatementContext() *IterationStatementContext {
	var p = new(IterationStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_iterationStatement
	return p
}

func InitEmptyIterationStatementContext(p *IterationStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_iterationStatement
}

func (*IterationStatementContext) IsIterationStatementContext() {}

func NewIterationStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IterationStatementContext {
	var p = new(IterationStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_iterationStatement

	return p
}

func (s *IterationStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IterationStatementContext) While() antlr.TerminalNode {
	return s.GetToken(tinycParserWhile, 0)
}

func (s *IterationStatementContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(tinycParserLeftParen, 0)
}

func (s *IterationStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IterationStatementContext) RightParen() antlr.TerminalNode {
	return s.GetToken(tinycParserRightParen, 0)
}

func (s *IterationStatementContext) Statement() IStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *IterationStatementContext) Do() antlr.TerminalNode {
	return s.GetToken(tinycParserDo, 0)
}

func (s *IterationStatementContext) Semi() antlr.TerminalNode {
	return s.GetToken(tinycParserSemi, 0)
}

func (s *IterationStatementContext) For() antlr.TerminalNode {
	return s.GetToken(tinycParserFor, 0)
}

func (s *IterationStatementContext) ForCondition() IForConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForConditionContext)
}

func (s *IterationStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IterationStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IterationStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterIterationStatement(s)
	}
}

func (s *IterationStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitIterationStatement(s)
	}
}

func (s *IterationStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitIterationStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) IterationStatement() (localctx IIterationStatementContext) {
	localctx = NewIterationStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, tinycParserRULE_iterationStatement)
	p.SetState(654)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case tinycParserWhile:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(634)
			p.Match(tinycParserWhile)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(635)
			p.Match(tinycParserLeftParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(636)
			p.Expression()
		}
		{
			p.SetState(637)
			p.Match(tinycParserRightParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(638)
			p.Statement()
		}

	case tinycParserDo:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(640)
			p.Match(tinycParserDo)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(641)
			p.Statement()
		}
		{
			p.SetState(642)
			p.Match(tinycParserWhile)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(643)
			p.Match(tinycParserLeftParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(644)
			p.Expression()
		}
		{
			p.SetState(645)
			p.Match(tinycParserRightParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(646)
			p.Match(tinycParserSemi)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case tinycParserFor:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(648)
			p.Match(tinycParserFor)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(649)
			p.Match(tinycParserLeftParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(650)
			p.ForCondition()
		}
		{
			p.SetState(651)
			p.Match(tinycParserRightParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(652)
			p.Statement()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForConditionContext is an interface to support dynamic dispatch.
type IForConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSemi() []antlr.TerminalNode
	Semi(i int) antlr.TerminalNode
	ForDeclaration() IForDeclarationContext
	AllForExpression() []IForExpressionContext
	ForExpression(i int) IForExpressionContext
	Expression() IExpressionContext

	// IsForConditionContext differentiates from other interfaces.
	IsForConditionContext()
}

type ForConditionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForConditionContext() *ForConditionContext {
	var p = new(ForConditionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_forCondition
	return p
}

func InitEmptyForConditionContext(p *ForConditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_forCondition
}

func (*ForConditionContext) IsForConditionContext() {}

func NewForConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForConditionContext {
	var p = new(ForConditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_forCondition

	return p
}

func (s *ForConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ForConditionContext) AllSemi() []antlr.TerminalNode {
	return s.GetTokens(tinycParserSemi)
}

func (s *ForConditionContext) Semi(i int) antlr.TerminalNode {
	return s.GetToken(tinycParserSemi, i)
}

func (s *ForConditionContext) ForDeclaration() IForDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForDeclarationContext)
}

func (s *ForConditionContext) AllForExpression() []IForExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IForExpressionContext); ok {
			len++
		}
	}

	tst := make([]IForExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IForExpressionContext); ok {
			tst[i] = t.(IForExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ForConditionContext) ForExpression(i int) IForExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForExpressionContext)
}

func (s *ForConditionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ForConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterForCondition(s)
	}
}

func (s *ForConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitForCondition(s)
	}
}

func (s *ForConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitForCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) ForCondition() (localctx IForConditionContext) {
	localctx = NewForConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 118, tinycParserRULE_forCondition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(660)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 78, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(656)
			p.ForDeclaration()
		}

	case 2:
		p.SetState(658)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64((_la-31)) & ^0x3f) == 0 && ((int64(1)<<(_la-31))&1030845190145) != 0 {
			{
				p.SetState(657)
				p.Expression()
			}

		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(662)
		p.Match(tinycParserSemi)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(664)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-31)) & ^0x3f) == 0 && ((int64(1)<<(_la-31))&1030845190145) != 0 {
		{
			p.SetState(663)
			p.ForExpression()
		}

	}
	{
		p.SetState(666)
		p.Match(tinycParserSemi)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(668)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-31)) & ^0x3f) == 0 && ((int64(1)<<(_la-31))&1030845190145) != 0 {
		{
			p.SetState(667)
			p.ForExpression()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForDeclarationContext is an interface to support dynamic dispatch.
type IForDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DeclarationSpecifiers() IDeclarationSpecifiersContext
	InitDeclaratorList() IInitDeclaratorListContext

	// IsForDeclarationContext differentiates from other interfaces.
	IsForDeclarationContext()
}

type ForDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForDeclarationContext() *ForDeclarationContext {
	var p = new(ForDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_forDeclaration
	return p
}

func InitEmptyForDeclarationContext(p *ForDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_forDeclaration
}

func (*ForDeclarationContext) IsForDeclarationContext() {}

func NewForDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForDeclarationContext {
	var p = new(ForDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_forDeclaration

	return p
}

func (s *ForDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ForDeclarationContext) DeclarationSpecifiers() IDeclarationSpecifiersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationSpecifiersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationSpecifiersContext)
}

func (s *ForDeclarationContext) InitDeclaratorList() IInitDeclaratorListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInitDeclaratorListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInitDeclaratorListContext)
}

func (s *ForDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterForDeclaration(s)
	}
}

func (s *ForDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitForDeclaration(s)
	}
}

func (s *ForDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitForDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) ForDeclaration() (localctx IForDeclarationContext) {
	localctx = NewForDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 120, tinycParserRULE_forDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(670)
		p.DeclarationSpecifiers()
	}
	p.SetState(672)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-31)) & ^0x3f) == 0 && ((int64(1)<<(_la-31))&68727930881) != 0 {
		{
			p.SetState(671)
			p.InitDeclaratorList()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForExpressionContext is an interface to support dynamic dispatch.
type IForExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAssignmentExpression() []IAssignmentExpressionContext
	AssignmentExpression(i int) IAssignmentExpressionContext
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

	// IsForExpressionContext differentiates from other interfaces.
	IsForExpressionContext()
}

type ForExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForExpressionContext() *ForExpressionContext {
	var p = new(ForExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_forExpression
	return p
}

func InitEmptyForExpressionContext(p *ForExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_forExpression
}

func (*ForExpressionContext) IsForExpressionContext() {}

func NewForExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForExpressionContext {
	var p = new(ForExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_forExpression

	return p
}

func (s *ForExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ForExpressionContext) AllAssignmentExpression() []IAssignmentExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAssignmentExpressionContext); ok {
			len++
		}
	}

	tst := make([]IAssignmentExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAssignmentExpressionContext); ok {
			tst[i] = t.(IAssignmentExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ForExpressionContext) AssignmentExpression(i int) IAssignmentExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentExpressionContext)
}

func (s *ForExpressionContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(tinycParserComma)
}

func (s *ForExpressionContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(tinycParserComma, i)
}

func (s *ForExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterForExpression(s)
	}
}

func (s *ForExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitForExpression(s)
	}
}

func (s *ForExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitForExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) ForExpression() (localctx IForExpressionContext) {
	localctx = NewForExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 122, tinycParserRULE_forExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(674)
		p.AssignmentExpression()
	}
	p.SetState(679)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == tinycParserComma {
		{
			p.SetState(675)
			p.Match(tinycParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(676)
			p.AssignmentExpression()
		}

		p.SetState(681)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IJumpStatementContext is an interface to support dynamic dispatch.
type IJumpStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Semi() antlr.TerminalNode
	Goto() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	Continue() antlr.TerminalNode
	Break() antlr.TerminalNode
	Return() antlr.TerminalNode
	UnaryExpression() IUnaryExpressionContext
	Expression() IExpressionContext

	// IsJumpStatementContext differentiates from other interfaces.
	IsJumpStatementContext()
}

type JumpStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJumpStatementContext() *JumpStatementContext {
	var p = new(JumpStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_jumpStatement
	return p
}

func InitEmptyJumpStatementContext(p *JumpStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_jumpStatement
}

func (*JumpStatementContext) IsJumpStatementContext() {}

func NewJumpStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JumpStatementContext {
	var p = new(JumpStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_jumpStatement

	return p
}

func (s *JumpStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *JumpStatementContext) Semi() antlr.TerminalNode {
	return s.GetToken(tinycParserSemi, 0)
}

func (s *JumpStatementContext) Goto() antlr.TerminalNode {
	return s.GetToken(tinycParserGoto, 0)
}

func (s *JumpStatementContext) Identifier() antlr.TerminalNode {
	return s.GetToken(tinycParserIdentifier, 0)
}

func (s *JumpStatementContext) Continue() antlr.TerminalNode {
	return s.GetToken(tinycParserContinue, 0)
}

func (s *JumpStatementContext) Break() antlr.TerminalNode {
	return s.GetToken(tinycParserBreak, 0)
}

func (s *JumpStatementContext) Return() antlr.TerminalNode {
	return s.GetToken(tinycParserReturn, 0)
}

func (s *JumpStatementContext) UnaryExpression() IUnaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryExpressionContext)
}

func (s *JumpStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *JumpStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JumpStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JumpStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterJumpStatement(s)
	}
}

func (s *JumpStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitJumpStatement(s)
	}
}

func (s *JumpStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitJumpStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) JumpStatement() (localctx IJumpStatementContext) {
	localctx = NewJumpStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 124, tinycParserRULE_jumpStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(692)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 84, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(682)
			p.Match(tinycParserGoto)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(683)
			p.Match(tinycParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(684)
			p.Match(tinycParserContinue)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		{
			p.SetState(685)
			p.Match(tinycParserBreak)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		{
			p.SetState(686)
			p.Match(tinycParserReturn)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(688)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64((_la-31)) & ^0x3f) == 0 && ((int64(1)<<(_la-31))&1030845190145) != 0 {
			{
				p.SetState(687)
				p.Expression()
			}

		}

	case 5:
		{
			p.SetState(690)
			p.Match(tinycParserGoto)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(691)
			p.UnaryExpression()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(694)
		p.Match(tinycParserSemi)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICompilationUnitContext is an interface to support dynamic dispatch.
type ICompilationUnitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllExternalDeclaration() []IExternalDeclarationContext
	ExternalDeclaration(i int) IExternalDeclarationContext

	// IsCompilationUnitContext differentiates from other interfaces.
	IsCompilationUnitContext()
}

type CompilationUnitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompilationUnitContext() *CompilationUnitContext {
	var p = new(CompilationUnitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_compilationUnit
	return p
}

func InitEmptyCompilationUnitContext(p *CompilationUnitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_compilationUnit
}

func (*CompilationUnitContext) IsCompilationUnitContext() {}

func NewCompilationUnitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompilationUnitContext {
	var p = new(CompilationUnitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_compilationUnit

	return p
}

func (s *CompilationUnitContext) GetParser() antlr.Parser { return s.parser }

func (s *CompilationUnitContext) EOF() antlr.TerminalNode {
	return s.GetToken(tinycParserEOF, 0)
}

func (s *CompilationUnitContext) AllExternalDeclaration() []IExternalDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExternalDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IExternalDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExternalDeclarationContext); ok {
			tst[i] = t.(IExternalDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *CompilationUnitContext) ExternalDeclaration(i int) IExternalDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExternalDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExternalDeclarationContext)
}

func (s *CompilationUnitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompilationUnitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompilationUnitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterCompilationUnit(s)
	}
}

func (s *CompilationUnitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitCompilationUnit(s)
	}
}

func (s *CompilationUnitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitCompilationUnit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) CompilationUnit() (localctx ICompilationUnitContext) {
	localctx = NewCompilationUnitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 126, tinycParserRULE_compilationUnit)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(697)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64((_la-4)) & ^0x3f) == 0 && ((int64(1)<<(_la-4))&-9186208543675113183) != 0) {
		{
			p.SetState(696)
			p.ExternalDeclaration()
		}

		p.SetState(699)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(701)
		p.Match(tinycParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExternalDeclarationContext is an interface to support dynamic dispatch.
type IExternalDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FunctionDefinition() IFunctionDefinitionContext
	Declaration() IDeclarationContext
	Semi() antlr.TerminalNode

	// IsExternalDeclarationContext differentiates from other interfaces.
	IsExternalDeclarationContext()
}

type ExternalDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExternalDeclarationContext() *ExternalDeclarationContext {
	var p = new(ExternalDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_externalDeclaration
	return p
}

func InitEmptyExternalDeclarationContext(p *ExternalDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_externalDeclaration
}

func (*ExternalDeclarationContext) IsExternalDeclarationContext() {}

func NewExternalDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternalDeclarationContext {
	var p = new(ExternalDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_externalDeclaration

	return p
}

func (s *ExternalDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ExternalDeclarationContext) FunctionDefinition() IFunctionDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionDefinitionContext)
}

func (s *ExternalDeclarationContext) Declaration() IDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *ExternalDeclarationContext) Semi() antlr.TerminalNode {
	return s.GetToken(tinycParserSemi, 0)
}

func (s *ExternalDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExternalDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExternalDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterExternalDeclaration(s)
	}
}

func (s *ExternalDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitExternalDeclaration(s)
	}
}

func (s *ExternalDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitExternalDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) ExternalDeclaration() (localctx IExternalDeclarationContext) {
	localctx = NewExternalDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 128, tinycParserRULE_externalDeclaration)
	p.SetState(706)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 86, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(703)
			p.FunctionDefinition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(704)
			p.Declaration()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(705)
			p.Match(tinycParserSemi)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionDefinitionContext is an interface to support dynamic dispatch.
type IFunctionDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Declarator() IDeclaratorContext
	CompoundStatement() ICompoundStatementContext
	DeclarationSpecifiers() IDeclarationSpecifiersContext
	DeclarationList() IDeclarationListContext

	// IsFunctionDefinitionContext differentiates from other interfaces.
	IsFunctionDefinitionContext()
}

type FunctionDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDefinitionContext() *FunctionDefinitionContext {
	var p = new(FunctionDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_functionDefinition
	return p
}

func InitEmptyFunctionDefinitionContext(p *FunctionDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_functionDefinition
}

func (*FunctionDefinitionContext) IsFunctionDefinitionContext() {}

func NewFunctionDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDefinitionContext {
	var p = new(FunctionDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_functionDefinition

	return p
}

func (s *FunctionDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDefinitionContext) Declarator() IDeclaratorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaratorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclaratorContext)
}

func (s *FunctionDefinitionContext) CompoundStatement() ICompoundStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICompoundStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICompoundStatementContext)
}

func (s *FunctionDefinitionContext) DeclarationSpecifiers() IDeclarationSpecifiersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationSpecifiersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationSpecifiersContext)
}

func (s *FunctionDefinitionContext) DeclarationList() IDeclarationListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationListContext)
}

func (s *FunctionDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterFunctionDefinition(s)
	}
}

func (s *FunctionDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitFunctionDefinition(s)
	}
}

func (s *FunctionDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitFunctionDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) FunctionDefinition() (localctx IFunctionDefinitionContext) {
	localctx = NewFunctionDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 130, tinycParserRULE_functionDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(709)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 87, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(708)
			p.DeclarationSpecifiers()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(711)
		p.Declarator()
	}
	p.SetState(713)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-4)) & ^0x3f) == 0 && ((int64(1)<<(_la-4))&-9223372036828159711) != 0 {
		{
			p.SetState(712)
			p.DeclarationList()
		}

	}
	{
		p.SetState(715)
		p.CompoundStatement()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclarationListContext is an interface to support dynamic dispatch.
type IDeclarationListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDeclaration() []IDeclarationContext
	Declaration(i int) IDeclarationContext

	// IsDeclarationListContext differentiates from other interfaces.
	IsDeclarationListContext()
}

type DeclarationListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationListContext() *DeclarationListContext {
	var p = new(DeclarationListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_declarationList
	return p
}

func InitEmptyDeclarationListContext(p *DeclarationListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_declarationList
}

func (*DeclarationListContext) IsDeclarationListContext() {}

func NewDeclarationListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationListContext {
	var p = new(DeclarationListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_declarationList

	return p
}

func (s *DeclarationListContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationListContext) AllDeclaration() []IDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclarationContext); ok {
			tst[i] = t.(IDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *DeclarationListContext) Declaration(i int) IDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *DeclarationListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterDeclarationList(s)
	}
}

func (s *DeclarationListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitDeclarationList(s)
	}
}

func (s *DeclarationListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitDeclarationList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) DeclarationList() (localctx IDeclarationListContext) {
	localctx = NewDeclarationListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 132, tinycParserRULE_declarationList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(718)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64((_la-4)) & ^0x3f) == 0 && ((int64(1)<<(_la-4))&-9223372036828159711) != 0) {
		{
			p.SetState(717)
			p.Declaration()
		}

		p.SetState(720)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *tinycParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 34:
		var t *DirectDeclaratorContext = nil
		if localctx != nil {
			t = localctx.(*DirectDeclaratorContext)
		}
		return p.DirectDeclarator_Sempred(t, predIndex)

	case 44:
		var t *DirectAbstractDeclaratorContext = nil
		if localctx != nil {
			t = localctx.(*DirectAbstractDeclaratorContext)
		}
		return p.DirectAbstractDeclarator_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *tinycParser) DirectDeclarator_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *tinycParser) DirectAbstractDeclarator_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
