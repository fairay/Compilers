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
		"", "'auto'", "'break'", "'char'", "'const'", "'continue'", "'else'",
		"'extern'", "'float'", "'goto'", "'if'", "'inline'", "'int'", "'register'",
		"'restrict'", "'return'", "'sizeof'", "'struct'", "'switch'", "'typedef'",
		"'union'", "'while'", "'_Noreturn'", "'('", "')'", "'['", "']'", "'{'",
		"'}'", "'<'", "'<='", "'>'", "'>='", "'<<'", "'>>'", "'+'", "'++'",
		"'-'", "'--'", "'*'", "'/'", "'%'", "'&'", "'|'", "'&&'", "'||'", "'^'",
		"'!'", "'~'", "'?'", "':'", "';'", "','", "'='", "'=='", "'!='", "'->'",
		"'.'", "'...'",
	}
	staticData.SymbolicNames = []string{
		"", "Auto", "Break", "Char", "Const", "Continue", "Else", "Extern",
		"Float", "Goto", "If", "Inline", "Int", "Register", "Restrict", "Return",
		"Sizeof", "Struct", "Switch", "Typedef", "Union", "While", "Noreturn",
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
		"funcCall", "unaryExpression", "unaryOperator", "castExpression", "multiplicativeExpression",
		"additiveExpression", "relationalExpression", "equalityExpression",
		"logicalAndExpression", "logicalOrExpression", "assignmentExpression",
		"assignmentOperator", "expression", "declaration", "declarationSpecifiers",
		"initDeclaratorList", "initDeclarator", "typeSpecifier", "specifierQualifierList",
		"structDeclaratorList", "structDeclarator", "declarator", "parameterTypeList",
		"parameterList", "parameterDeclaration", "identifierList", "typeName",
		"abstractDeclarator", "initializer", "initializerList", "designation",
		"designatorList", "designator", "statement", "labeledStatement", "compoundStatement",
		"blockItem", "expressionStatement", "selectionStatement", "iterationStatement",
		"jumpStatement", "compilationUnit", "externalDeclaration", "functionDefinition",
		"declarationList",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 72, 500, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 1, 0, 1, 0, 1, 0, 4, 0, 102, 8, 0, 11, 0, 12, 0, 103,
		1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 110, 8, 0, 1, 1, 1, 1, 1, 1, 5, 1, 115, 8,
		1, 10, 1, 12, 1, 118, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 3, 3, 131, 8, 3, 1, 3, 1, 3, 3, 3, 135, 8, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 142, 8, 3, 10, 3, 12, 3, 145, 9, 3, 1, 4,
		1, 4, 1, 4, 1, 4, 5, 4, 151, 8, 4, 10, 4, 12, 4, 154, 9, 4, 3, 4, 156,
		8, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 3, 7, 172, 8, 7, 1, 8, 1, 8, 1, 8, 5, 8, 177, 8, 8, 10,
		8, 12, 8, 180, 9, 8, 1, 9, 1, 9, 1, 9, 5, 9, 185, 8, 9, 10, 9, 12, 9, 188,
		9, 9, 1, 10, 1, 10, 1, 10, 5, 10, 193, 8, 10, 10, 10, 12, 10, 196, 9, 10,
		1, 11, 1, 11, 1, 11, 5, 11, 201, 8, 11, 10, 11, 12, 11, 204, 9, 11, 1,
		12, 1, 12, 1, 12, 5, 12, 209, 8, 12, 10, 12, 12, 12, 212, 9, 12, 1, 13,
		1, 13, 1, 13, 5, 13, 217, 8, 13, 10, 13, 12, 13, 220, 9, 13, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 3, 14, 227, 8, 14, 1, 15, 1, 15, 1, 16, 1, 16,
		1, 16, 5, 16, 234, 8, 16, 10, 16, 12, 16, 237, 9, 16, 1, 17, 1, 17, 3,
		17, 241, 8, 17, 1, 17, 1, 17, 1, 18, 4, 18, 246, 8, 18, 11, 18, 12, 18,
		247, 1, 19, 1, 19, 1, 19, 5, 19, 253, 8, 19, 10, 19, 12, 19, 256, 9, 19,
		1, 20, 1, 20, 1, 20, 3, 20, 261, 8, 20, 1, 21, 1, 21, 1, 22, 1, 22, 3,
		22, 267, 8, 22, 1, 23, 1, 23, 1, 23, 5, 23, 272, 8, 23, 10, 23, 12, 23,
		275, 9, 23, 1, 24, 1, 24, 3, 24, 279, 8, 24, 1, 24, 1, 24, 3, 24, 283,
		8, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 3, 25, 291, 8, 25, 1,
		25, 1, 25, 1, 25, 3, 25, 296, 8, 25, 1, 25, 1, 25, 1, 25, 1, 25, 3, 25,
		302, 8, 25, 1, 25, 5, 25, 305, 8, 25, 10, 25, 12, 25, 308, 9, 25, 1, 26,
		1, 26, 1, 26, 3, 26, 313, 8, 26, 1, 27, 1, 27, 1, 27, 5, 27, 318, 8, 27,
		10, 27, 12, 27, 321, 9, 27, 1, 27, 3, 27, 324, 8, 27, 1, 28, 1, 28, 1,
		28, 1, 29, 1, 29, 1, 29, 5, 29, 332, 8, 29, 10, 29, 12, 29, 335, 9, 29,
		1, 30, 1, 30, 3, 30, 339, 8, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 3, 31, 348, 8, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31,
		3, 31, 356, 8, 31, 1, 31, 3, 31, 359, 8, 31, 1, 31, 1, 31, 1, 31, 3, 31,
		364, 8, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 3,
		31, 374, 8, 31, 1, 31, 5, 31, 377, 8, 31, 10, 31, 12, 31, 380, 9, 31, 1,
		32, 1, 32, 1, 32, 1, 32, 3, 32, 386, 8, 32, 1, 32, 1, 32, 3, 32, 390, 8,
		32, 1, 33, 3, 33, 393, 8, 33, 1, 33, 1, 33, 1, 33, 3, 33, 398, 8, 33, 1,
		33, 5, 33, 401, 8, 33, 10, 33, 12, 33, 404, 9, 33, 1, 34, 1, 34, 1, 34,
		1, 35, 4, 35, 410, 8, 35, 11, 35, 12, 35, 411, 1, 36, 1, 36, 1, 36, 1,
		36, 1, 36, 1, 36, 3, 36, 420, 8, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37,
		1, 37, 3, 37, 428, 8, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 5,
		39, 436, 8, 39, 10, 39, 12, 39, 439, 9, 39, 1, 39, 1, 39, 1, 40, 1, 40,
		3, 40, 445, 8, 40, 1, 41, 3, 41, 448, 8, 41, 1, 41, 1, 41, 1, 42, 1, 42,
		1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 3, 42, 459, 8, 42, 1, 43, 1, 43, 1,
		43, 1, 43, 1, 43, 1, 43, 1, 44, 1, 44, 1, 44, 1, 44, 1, 45, 4, 45, 472,
		8, 45, 11, 45, 12, 45, 473, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 3, 46, 481,
		8, 46, 1, 47, 3, 47, 484, 8, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 3,
		47, 491, 8, 47, 1, 47, 1, 47, 1, 48, 4, 48, 496, 8, 48, 11, 48, 12, 48,
		497, 1, 48, 0, 2, 50, 62, 49, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
		24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58,
		60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94,
		96, 0, 6, 3, 0, 35, 35, 37, 37, 47, 47, 1, 0, 39, 41, 2, 0, 35, 35, 37,
		37, 1, 0, 29, 32, 1, 0, 54, 55, 3, 0, 3, 3, 8, 8, 12, 12, 521, 0, 109,
		1, 0, 0, 0, 2, 111, 1, 0, 0, 0, 4, 119, 1, 0, 0, 0, 6, 134, 1, 0, 0, 0,
		8, 146, 1, 0, 0, 0, 10, 159, 1, 0, 0, 0, 12, 162, 1, 0, 0, 0, 14, 171,
		1, 0, 0, 0, 16, 173, 1, 0, 0, 0, 18, 181, 1, 0, 0, 0, 20, 189, 1, 0, 0,
		0, 22, 197, 1, 0, 0, 0, 24, 205, 1, 0, 0, 0, 26, 213, 1, 0, 0, 0, 28, 226,
		1, 0, 0, 0, 30, 228, 1, 0, 0, 0, 32, 230, 1, 0, 0, 0, 34, 238, 1, 0, 0,
		0, 36, 245, 1, 0, 0, 0, 38, 249, 1, 0, 0, 0, 40, 257, 1, 0, 0, 0, 42, 262,
		1, 0, 0, 0, 44, 264, 1, 0, 0, 0, 46, 268, 1, 0, 0, 0, 48, 282, 1, 0, 0,
		0, 50, 290, 1, 0, 0, 0, 52, 309, 1, 0, 0, 0, 54, 323, 1, 0, 0, 0, 56, 325,
		1, 0, 0, 0, 58, 328, 1, 0, 0, 0, 60, 336, 1, 0, 0, 0, 62, 358, 1, 0, 0,
		0, 64, 389, 1, 0, 0, 0, 66, 392, 1, 0, 0, 0, 68, 405, 1, 0, 0, 0, 70, 409,
		1, 0, 0, 0, 72, 419, 1, 0, 0, 0, 74, 427, 1, 0, 0, 0, 76, 429, 1, 0, 0,
		0, 78, 433, 1, 0, 0, 0, 80, 444, 1, 0, 0, 0, 82, 447, 1, 0, 0, 0, 84, 451,
		1, 0, 0, 0, 86, 460, 1, 0, 0, 0, 88, 466, 1, 0, 0, 0, 90, 471, 1, 0, 0,
		0, 92, 480, 1, 0, 0, 0, 94, 483, 1, 0, 0, 0, 96, 495, 1, 0, 0, 0, 98, 110,
		5, 59, 0, 0, 99, 110, 5, 60, 0, 0, 100, 102, 5, 62, 0, 0, 101, 100, 1,
		0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 101, 1, 0, 0, 0, 103, 104, 1, 0, 0,
		0, 104, 110, 1, 0, 0, 0, 105, 106, 5, 23, 0, 0, 106, 107, 3, 32, 16, 0,
		107, 108, 5, 24, 0, 0, 108, 110, 1, 0, 0, 0, 109, 98, 1, 0, 0, 0, 109,
		99, 1, 0, 0, 0, 109, 101, 1, 0, 0, 0, 109, 105, 1, 0, 0, 0, 110, 1, 1,
		0, 0, 0, 111, 116, 3, 4, 2, 0, 112, 113, 5, 52, 0, 0, 113, 115, 3, 4, 2,
		0, 114, 112, 1, 0, 0, 0, 115, 118, 1, 0, 0, 0, 116, 114, 1, 0, 0, 0, 116,
		117, 1, 0, 0, 0, 117, 3, 1, 0, 0, 0, 118, 116, 1, 0, 0, 0, 119, 120, 3,
		60, 30, 0, 120, 121, 5, 50, 0, 0, 121, 122, 3, 28, 14, 0, 122, 5, 1, 0,
		0, 0, 123, 135, 3, 0, 0, 0, 124, 125, 5, 23, 0, 0, 125, 126, 3, 60, 30,
		0, 126, 127, 5, 24, 0, 0, 127, 128, 5, 27, 0, 0, 128, 130, 3, 66, 33, 0,
		129, 131, 5, 52, 0, 0, 130, 129, 1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131,
		132, 1, 0, 0, 0, 132, 133, 5, 28, 0, 0, 133, 135, 1, 0, 0, 0, 134, 123,
		1, 0, 0, 0, 134, 124, 1, 0, 0, 0, 135, 143, 1, 0, 0, 0, 136, 137, 5, 25,
		0, 0, 137, 138, 3, 32, 16, 0, 138, 139, 5, 26, 0, 0, 139, 142, 1, 0, 0,
		0, 140, 142, 3, 8, 4, 0, 141, 136, 1, 0, 0, 0, 141, 140, 1, 0, 0, 0, 142,
		145, 1, 0, 0, 0, 143, 141, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144, 7, 1,
		0, 0, 0, 145, 143, 1, 0, 0, 0, 146, 155, 5, 23, 0, 0, 147, 152, 3, 28,
		14, 0, 148, 149, 5, 52, 0, 0, 149, 151, 3, 28, 14, 0, 150, 148, 1, 0, 0,
		0, 151, 154, 1, 0, 0, 0, 152, 150, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153,
		156, 1, 0, 0, 0, 154, 152, 1, 0, 0, 0, 155, 147, 1, 0, 0, 0, 155, 156,
		1, 0, 0, 0, 156, 157, 1, 0, 0, 0, 157, 158, 5, 24, 0, 0, 158, 9, 1, 0,
		0, 0, 159, 160, 3, 12, 6, 0, 160, 161, 3, 14, 7, 0, 161, 11, 1, 0, 0, 0,
		162, 163, 7, 0, 0, 0, 163, 13, 1, 0, 0, 0, 164, 165, 5, 23, 0, 0, 165,
		166, 3, 60, 30, 0, 166, 167, 5, 24, 0, 0, 167, 168, 3, 14, 7, 0, 168, 172,
		1, 0, 0, 0, 169, 172, 3, 6, 3, 0, 170, 172, 3, 10, 5, 0, 171, 164, 1, 0,
		0, 0, 171, 169, 1, 0, 0, 0, 171, 170, 1, 0, 0, 0, 172, 15, 1, 0, 0, 0,
		173, 178, 3, 14, 7, 0, 174, 175, 7, 1, 0, 0, 175, 177, 3, 14, 7, 0, 176,
		174, 1, 0, 0, 0, 177, 180, 1, 0, 0, 0, 178, 176, 1, 0, 0, 0, 178, 179,
		1, 0, 0, 0, 179, 17, 1, 0, 0, 0, 180, 178, 1, 0, 0, 0, 181, 186, 3, 16,
		8, 0, 182, 183, 7, 2, 0, 0, 183, 185, 3, 16, 8, 0, 184, 182, 1, 0, 0, 0,
		185, 188, 1, 0, 0, 0, 186, 184, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187,
		19, 1, 0, 0, 0, 188, 186, 1, 0, 0, 0, 189, 194, 3, 18, 9, 0, 190, 191,
		7, 3, 0, 0, 191, 193, 3, 18, 9, 0, 192, 190, 1, 0, 0, 0, 193, 196, 1, 0,
		0, 0, 194, 192, 1, 0, 0, 0, 194, 195, 1, 0, 0, 0, 195, 21, 1, 0, 0, 0,
		196, 194, 1, 0, 0, 0, 197, 202, 3, 20, 10, 0, 198, 199, 7, 4, 0, 0, 199,
		201, 3, 20, 10, 0, 200, 198, 1, 0, 0, 0, 201, 204, 1, 0, 0, 0, 202, 200,
		1, 0, 0, 0, 202, 203, 1, 0, 0, 0, 203, 23, 1, 0, 0, 0, 204, 202, 1, 0,
		0, 0, 205, 210, 3, 22, 11, 0, 206, 207, 5, 44, 0, 0, 207, 209, 3, 22, 11,
		0, 208, 206, 1, 0, 0, 0, 209, 212, 1, 0, 0, 0, 210, 208, 1, 0, 0, 0, 210,
		211, 1, 0, 0, 0, 211, 25, 1, 0, 0, 0, 212, 210, 1, 0, 0, 0, 213, 218, 3,
		24, 12, 0, 214, 215, 5, 45, 0, 0, 215, 217, 3, 24, 12, 0, 216, 214, 1,
		0, 0, 0, 217, 220, 1, 0, 0, 0, 218, 216, 1, 0, 0, 0, 218, 219, 1, 0, 0,
		0, 219, 27, 1, 0, 0, 0, 220, 218, 1, 0, 0, 0, 221, 227, 3, 26, 13, 0, 222,
		223, 3, 6, 3, 0, 223, 224, 3, 30, 15, 0, 224, 225, 3, 28, 14, 0, 225, 227,
		1, 0, 0, 0, 226, 221, 1, 0, 0, 0, 226, 222, 1, 0, 0, 0, 227, 29, 1, 0,
		0, 0, 228, 229, 5, 53, 0, 0, 229, 31, 1, 0, 0, 0, 230, 235, 3, 28, 14,
		0, 231, 232, 5, 52, 0, 0, 232, 234, 3, 28, 14, 0, 233, 231, 1, 0, 0, 0,
		234, 237, 1, 0, 0, 0, 235, 233, 1, 0, 0, 0, 235, 236, 1, 0, 0, 0, 236,
		33, 1, 0, 0, 0, 237, 235, 1, 0, 0, 0, 238, 240, 3, 36, 18, 0, 239, 241,
		3, 38, 19, 0, 240, 239, 1, 0, 0, 0, 240, 241, 1, 0, 0, 0, 241, 242, 1,
		0, 0, 0, 242, 243, 5, 51, 0, 0, 243, 35, 1, 0, 0, 0, 244, 246, 3, 42, 21,
		0, 245, 244, 1, 0, 0, 0, 246, 247, 1, 0, 0, 0, 247, 245, 1, 0, 0, 0, 247,
		248, 1, 0, 0, 0, 248, 37, 1, 0, 0, 0, 249, 254, 3, 40, 20, 0, 250, 251,
		5, 52, 0, 0, 251, 253, 3, 40, 20, 0, 252, 250, 1, 0, 0, 0, 253, 256, 1,
		0, 0, 0, 254, 252, 1, 0, 0, 0, 254, 255, 1, 0, 0, 0, 255, 39, 1, 0, 0,
		0, 256, 254, 1, 0, 0, 0, 257, 260, 3, 50, 25, 0, 258, 259, 5, 53, 0, 0,
		259, 261, 3, 64, 32, 0, 260, 258, 1, 0, 0, 0, 260, 261, 1, 0, 0, 0, 261,
		41, 1, 0, 0, 0, 262, 263, 7, 5, 0, 0, 263, 43, 1, 0, 0, 0, 264, 266, 3,
		42, 21, 0, 265, 267, 3, 44, 22, 0, 266, 265, 1, 0, 0, 0, 266, 267, 1, 0,
		0, 0, 267, 45, 1, 0, 0, 0, 268, 273, 3, 48, 24, 0, 269, 270, 5, 52, 0,
		0, 270, 272, 3, 48, 24, 0, 271, 269, 1, 0, 0, 0, 272, 275, 1, 0, 0, 0,
		273, 271, 1, 0, 0, 0, 273, 274, 1, 0, 0, 0, 274, 47, 1, 0, 0, 0, 275, 273,
		1, 0, 0, 0, 276, 283, 3, 50, 25, 0, 277, 279, 3, 50, 25, 0, 278, 277, 1,
		0, 0, 0, 278, 279, 1, 0, 0, 0, 279, 280, 1, 0, 0, 0, 280, 281, 5, 50, 0,
		0, 281, 283, 3, 26, 13, 0, 282, 276, 1, 0, 0, 0, 282, 278, 1, 0, 0, 0,
		283, 49, 1, 0, 0, 0, 284, 285, 6, 25, -1, 0, 285, 291, 5, 59, 0, 0, 286,
		287, 5, 23, 0, 0, 287, 288, 3, 50, 25, 0, 288, 289, 5, 24, 0, 0, 289, 291,
		1, 0, 0, 0, 290, 284, 1, 0, 0, 0, 290, 286, 1, 0, 0, 0, 291, 306, 1, 0,
		0, 0, 292, 293, 10, 2, 0, 0, 293, 295, 5, 25, 0, 0, 294, 296, 3, 28, 14,
		0, 295, 294, 1, 0, 0, 0, 295, 296, 1, 0, 0, 0, 296, 297, 1, 0, 0, 0, 297,
		305, 5, 26, 0, 0, 298, 299, 10, 1, 0, 0, 299, 301, 5, 23, 0, 0, 300, 302,
		3, 58, 29, 0, 301, 300, 1, 0, 0, 0, 301, 302, 1, 0, 0, 0, 302, 303, 1,
		0, 0, 0, 303, 305, 5, 24, 0, 0, 304, 292, 1, 0, 0, 0, 304, 298, 1, 0, 0,
		0, 305, 308, 1, 0, 0, 0, 306, 304, 1, 0, 0, 0, 306, 307, 1, 0, 0, 0, 307,
		51, 1, 0, 0, 0, 308, 306, 1, 0, 0, 0, 309, 312, 3, 54, 27, 0, 310, 311,
		5, 52, 0, 0, 311, 313, 5, 58, 0, 0, 312, 310, 1, 0, 0, 0, 312, 313, 1,
		0, 0, 0, 313, 53, 1, 0, 0, 0, 314, 319, 3, 56, 28, 0, 315, 316, 5, 52,
		0, 0, 316, 318, 3, 56, 28, 0, 317, 315, 1, 0, 0, 0, 318, 321, 1, 0, 0,
		0, 319, 317, 1, 0, 0, 0, 319, 320, 1, 0, 0, 0, 320, 324, 1, 0, 0, 0, 321,
		319, 1, 0, 0, 0, 322, 324, 1, 0, 0, 0, 323, 314, 1, 0, 0, 0, 323, 322,
		1, 0, 0, 0, 324, 55, 1, 0, 0, 0, 325, 326, 3, 36, 18, 0, 326, 327, 3, 50,
		25, 0, 327, 57, 1, 0, 0, 0, 328, 333, 5, 59, 0, 0, 329, 330, 5, 52, 0,
		0, 330, 332, 5, 59, 0, 0, 331, 329, 1, 0, 0, 0, 332, 335, 1, 0, 0, 0, 333,
		331, 1, 0, 0, 0, 333, 334, 1, 0, 0, 0, 334, 59, 1, 0, 0, 0, 335, 333, 1,
		0, 0, 0, 336, 338, 3, 44, 22, 0, 337, 339, 3, 62, 31, 0, 338, 337, 1, 0,
		0, 0, 338, 339, 1, 0, 0, 0, 339, 61, 1, 0, 0, 0, 340, 341, 6, 31, -1, 0,
		341, 342, 5, 23, 0, 0, 342, 343, 3, 62, 31, 0, 343, 344, 5, 24, 0, 0, 344,
		359, 1, 0, 0, 0, 345, 347, 5, 25, 0, 0, 346, 348, 3, 28, 14, 0, 347, 346,
		1, 0, 0, 0, 347, 348, 1, 0, 0, 0, 348, 349, 1, 0, 0, 0, 349, 359, 5, 26,
		0, 0, 350, 351, 5, 25, 0, 0, 351, 352, 5, 39, 0, 0, 352, 359, 5, 26, 0,
		0, 353, 355, 5, 23, 0, 0, 354, 356, 3, 52, 26, 0, 355, 354, 1, 0, 0, 0,
		355, 356, 1, 0, 0, 0, 356, 357, 1, 0, 0, 0, 357, 359, 5, 24, 0, 0, 358,
		340, 1, 0, 0, 0, 358, 345, 1, 0, 0, 0, 358, 350, 1, 0, 0, 0, 358, 353,
		1, 0, 0, 0, 359, 378, 1, 0, 0, 0, 360, 361, 10, 3, 0, 0, 361, 363, 5, 25,
		0, 0, 362, 364, 3, 28, 14, 0, 363, 362, 1, 0, 0, 0, 363, 364, 1, 0, 0,
		0, 364, 365, 1, 0, 0, 0, 365, 377, 5, 26, 0, 0, 366, 367, 10, 2, 0, 0,
		367, 368, 5, 25, 0, 0, 368, 369, 5, 39, 0, 0, 369, 377, 5, 26, 0, 0, 370,
		371, 10, 1, 0, 0, 371, 373, 5, 23, 0, 0, 372, 374, 3, 52, 26, 0, 373, 372,
		1, 0, 0, 0, 373, 374, 1, 0, 0, 0, 374, 375, 1, 0, 0, 0, 375, 377, 5, 24,
		0, 0, 376, 360, 1, 0, 0, 0, 376, 366, 1, 0, 0, 0, 376, 370, 1, 0, 0, 0,
		377, 380, 1, 0, 0, 0, 378, 376, 1, 0, 0, 0, 378, 379, 1, 0, 0, 0, 379,
		63, 1, 0, 0, 0, 380, 378, 1, 0, 0, 0, 381, 390, 3, 28, 14, 0, 382, 383,
		5, 27, 0, 0, 383, 385, 3, 66, 33, 0, 384, 386, 5, 52, 0, 0, 385, 384, 1,
		0, 0, 0, 385, 386, 1, 0, 0, 0, 386, 387, 1, 0, 0, 0, 387, 388, 5, 28, 0,
		0, 388, 390, 1, 0, 0, 0, 389, 381, 1, 0, 0, 0, 389, 382, 1, 0, 0, 0, 390,
		65, 1, 0, 0, 0, 391, 393, 3, 68, 34, 0, 392, 391, 1, 0, 0, 0, 392, 393,
		1, 0, 0, 0, 393, 394, 1, 0, 0, 0, 394, 402, 3, 64, 32, 0, 395, 397, 5,
		52, 0, 0, 396, 398, 3, 68, 34, 0, 397, 396, 1, 0, 0, 0, 397, 398, 1, 0,
		0, 0, 398, 399, 1, 0, 0, 0, 399, 401, 3, 64, 32, 0, 400, 395, 1, 0, 0,
		0, 401, 404, 1, 0, 0, 0, 402, 400, 1, 0, 0, 0, 402, 403, 1, 0, 0, 0, 403,
		67, 1, 0, 0, 0, 404, 402, 1, 0, 0, 0, 405, 406, 3, 70, 35, 0, 406, 407,
		5, 53, 0, 0, 407, 69, 1, 0, 0, 0, 408, 410, 3, 72, 36, 0, 409, 408, 1,
		0, 0, 0, 410, 411, 1, 0, 0, 0, 411, 409, 1, 0, 0, 0, 411, 412, 1, 0, 0,
		0, 412, 71, 1, 0, 0, 0, 413, 414, 5, 25, 0, 0, 414, 415, 3, 26, 13, 0,
		415, 416, 5, 26, 0, 0, 416, 420, 1, 0, 0, 0, 417, 418, 5, 57, 0, 0, 418,
		420, 5, 59, 0, 0, 419, 413, 1, 0, 0, 0, 419, 417, 1, 0, 0, 0, 420, 73,
		1, 0, 0, 0, 421, 428, 3, 76, 38, 0, 422, 428, 3, 78, 39, 0, 423, 428, 3,
		82, 41, 0, 424, 428, 3, 84, 42, 0, 425, 428, 3, 86, 43, 0, 426, 428, 3,
		88, 44, 0, 427, 421, 1, 0, 0, 0, 427, 422, 1, 0, 0, 0, 427, 423, 1, 0,
		0, 0, 427, 424, 1, 0, 0, 0, 427, 425, 1, 0, 0, 0, 427, 426, 1, 0, 0, 0,
		428, 75, 1, 0, 0, 0, 429, 430, 5, 59, 0, 0, 430, 431, 5, 50, 0, 0, 431,
		432, 3, 74, 37, 0, 432, 77, 1, 0, 0, 0, 433, 437, 5, 27, 0, 0, 434, 436,
		3, 80, 40, 0, 435, 434, 1, 0, 0, 0, 436, 439, 1, 0, 0, 0, 437, 435, 1,
		0, 0, 0, 437, 438, 1, 0, 0, 0, 438, 440, 1, 0, 0, 0, 439, 437, 1, 0, 0,
		0, 440, 441, 5, 28, 0, 0, 441, 79, 1, 0, 0, 0, 442, 445, 3, 74, 37, 0,
		443, 445, 3, 34, 17, 0, 444, 442, 1, 0, 0, 0, 444, 443, 1, 0, 0, 0, 445,
		81, 1, 0, 0, 0, 446, 448, 3, 32, 16, 0, 447, 446, 1, 0, 0, 0, 447, 448,
		1, 0, 0, 0, 448, 449, 1, 0, 0, 0, 449, 450, 5, 51, 0, 0, 450, 83, 1, 0,
		0, 0, 451, 452, 5, 10, 0, 0, 452, 453, 5, 23, 0, 0, 453, 454, 3, 32, 16,
		0, 454, 455, 5, 24, 0, 0, 455, 458, 3, 74, 37, 0, 456, 457, 5, 6, 0, 0,
		457, 459, 3, 74, 37, 0, 458, 456, 1, 0, 0, 0, 458, 459, 1, 0, 0, 0, 459,
		85, 1, 0, 0, 0, 460, 461, 5, 21, 0, 0, 461, 462, 5, 23, 0, 0, 462, 463,
		3, 32, 16, 0, 463, 464, 5, 24, 0, 0, 464, 465, 3, 74, 37, 0, 465, 87, 1,
		0, 0, 0, 466, 467, 5, 15, 0, 0, 467, 468, 3, 32, 16, 0, 468, 469, 5, 51,
		0, 0, 469, 89, 1, 0, 0, 0, 470, 472, 3, 92, 46, 0, 471, 470, 1, 0, 0, 0,
		472, 473, 1, 0, 0, 0, 473, 471, 1, 0, 0, 0, 473, 474, 1, 0, 0, 0, 474,
		475, 1, 0, 0, 0, 475, 476, 5, 0, 0, 1, 476, 91, 1, 0, 0, 0, 477, 481, 3,
		94, 47, 0, 478, 481, 3, 34, 17, 0, 479, 481, 5, 51, 0, 0, 480, 477, 1,
		0, 0, 0, 480, 478, 1, 0, 0, 0, 480, 479, 1, 0, 0, 0, 481, 93, 1, 0, 0,
		0, 482, 484, 3, 36, 18, 0, 483, 482, 1, 0, 0, 0, 483, 484, 1, 0, 0, 0,
		484, 485, 1, 0, 0, 0, 485, 486, 5, 59, 0, 0, 486, 487, 5, 23, 0, 0, 487,
		488, 3, 54, 27, 0, 488, 490, 5, 24, 0, 0, 489, 491, 3, 96, 48, 0, 490,
		489, 1, 0, 0, 0, 490, 491, 1, 0, 0, 0, 491, 492, 1, 0, 0, 0, 492, 493,
		3, 78, 39, 0, 493, 95, 1, 0, 0, 0, 494, 496, 3, 34, 17, 0, 495, 494, 1,
		0, 0, 0, 496, 497, 1, 0, 0, 0, 497, 495, 1, 0, 0, 0, 497, 498, 1, 0, 0,
		0, 498, 97, 1, 0, 0, 0, 60, 103, 109, 116, 130, 134, 141, 143, 152, 155,
		171, 178, 186, 194, 202, 210, 218, 226, 235, 240, 247, 254, 260, 266, 273,
		278, 282, 290, 295, 301, 304, 306, 312, 319, 323, 333, 338, 347, 355, 358,
		363, 373, 376, 378, 385, 389, 392, 397, 402, 411, 419, 427, 437, 444, 447,
		458, 473, 480, 483, 490, 497,
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
	tinycParserChar                   = 3
	tinycParserConst                  = 4
	tinycParserContinue               = 5
	tinycParserElse                   = 6
	tinycParserExtern                 = 7
	tinycParserFloat                  = 8
	tinycParserGoto                   = 9
	tinycParserIf                     = 10
	tinycParserInline                 = 11
	tinycParserInt                    = 12
	tinycParserRegister               = 13
	tinycParserRestrict               = 14
	tinycParserReturn                 = 15
	tinycParserSizeof                 = 16
	tinycParserStruct                 = 17
	tinycParserSwitch                 = 18
	tinycParserTypedef                = 19
	tinycParserUnion                  = 20
	tinycParserWhile                  = 21
	tinycParserNoreturn               = 22
	tinycParserLeftParen              = 23
	tinycParserRightParen             = 24
	tinycParserLeftBracket            = 25
	tinycParserRightBracket           = 26
	tinycParserLeftBrace              = 27
	tinycParserRightBrace             = 28
	tinycParserLess                   = 29
	tinycParserLessEqual              = 30
	tinycParserGreater                = 31
	tinycParserGreaterEqual           = 32
	tinycParserLeftShift              = 33
	tinycParserRightShift             = 34
	tinycParserPlus                   = 35
	tinycParserPlusPlus               = 36
	tinycParserMinus                  = 37
	tinycParserMinusMinus             = 38
	tinycParserStar                   = 39
	tinycParserDiv                    = 40
	tinycParserMod                    = 41
	tinycParserAnd                    = 42
	tinycParserOr                     = 43
	tinycParserAndAnd                 = 44
	tinycParserOrOr                   = 45
	tinycParserCaret                  = 46
	tinycParserNot                    = 47
	tinycParserTilde                  = 48
	tinycParserQuestion               = 49
	tinycParserColon                  = 50
	tinycParserSemi                   = 51
	tinycParserComma                  = 52
	tinycParserAssign                 = 53
	tinycParserEqual                  = 54
	tinycParserNotEqual               = 55
	tinycParserArrow                  = 56
	tinycParserDot                    = 57
	tinycParserEllipsis               = 58
	tinycParserIdentifier             = 59
	tinycParserConstant               = 60
	tinycParserDigitSequence          = 61
	tinycParserStringLiteral          = 62
	tinycParserComplexDefine          = 63
	tinycParserIncludeDirective       = 64
	tinycParserAsmBlock               = 65
	tinycParserLineAfterPreprocessing = 66
	tinycParserLineDirective          = 67
	tinycParserPragmaDirective        = 68
	tinycParserWhitespace             = 69
	tinycParserNewline                = 70
	tinycParserBlockComment           = 71
	tinycParserLineComment            = 72
)

// tinycParser rules.
const (
	tinycParserRULE_primaryExpression        = 0
	tinycParserRULE_genericAssocList         = 1
	tinycParserRULE_genericAssociation       = 2
	tinycParserRULE_postfixExpression        = 3
	tinycParserRULE_funcCall                 = 4
	tinycParserRULE_unaryExpression          = 5
	tinycParserRULE_unaryOperator            = 6
	tinycParserRULE_castExpression           = 7
	tinycParserRULE_multiplicativeExpression = 8
	tinycParserRULE_additiveExpression       = 9
	tinycParserRULE_relationalExpression     = 10
	tinycParserRULE_equalityExpression       = 11
	tinycParserRULE_logicalAndExpression     = 12
	tinycParserRULE_logicalOrExpression      = 13
	tinycParserRULE_assignmentExpression     = 14
	tinycParserRULE_assignmentOperator       = 15
	tinycParserRULE_expression               = 16
	tinycParserRULE_declaration              = 17
	tinycParserRULE_declarationSpecifiers    = 18
	tinycParserRULE_initDeclaratorList       = 19
	tinycParserRULE_initDeclarator           = 20
	tinycParserRULE_typeSpecifier            = 21
	tinycParserRULE_specifierQualifierList   = 22
	tinycParserRULE_structDeclaratorList     = 23
	tinycParserRULE_structDeclarator         = 24
	tinycParserRULE_declarator               = 25
	tinycParserRULE_parameterTypeList        = 26
	tinycParserRULE_parameterList            = 27
	tinycParserRULE_parameterDeclaration     = 28
	tinycParserRULE_identifierList           = 29
	tinycParserRULE_typeName                 = 30
	tinycParserRULE_abstractDeclarator       = 31
	tinycParserRULE_initializer              = 32
	tinycParserRULE_initializerList          = 33
	tinycParserRULE_designation              = 34
	tinycParserRULE_designatorList           = 35
	tinycParserRULE_designator               = 36
	tinycParserRULE_statement                = 37
	tinycParserRULE_labeledStatement         = 38
	tinycParserRULE_compoundStatement        = 39
	tinycParserRULE_blockItem                = 40
	tinycParserRULE_expressionStatement      = 41
	tinycParserRULE_selectionStatement       = 42
	tinycParserRULE_iterationStatement       = 43
	tinycParserRULE_jumpStatement            = 44
	tinycParserRULE_compilationUnit          = 45
	tinycParserRULE_externalDeclaration      = 46
	tinycParserRULE_functionDefinition       = 47
	tinycParserRULE_declarationList          = 48
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

	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case tinycParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(98)
			p.Match(tinycParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case tinycParserConstant:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(99)
			p.Match(tinycParserConstant)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case tinycParserStringLiteral:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(101)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == tinycParserStringLiteral {
			{
				p.SetState(100)
				p.Match(tinycParserStringLiteral)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(103)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case tinycParserLeftParen:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(105)
			p.Match(tinycParserLeftParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(106)
			p.Expression()
		}
		{
			p.SetState(107)
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
		p.SetState(111)
		p.GenericAssociation()
	}
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == tinycParserComma {
		{
			p.SetState(112)
			p.Match(tinycParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(113)
			p.GenericAssociation()
		}

		p.SetState(118)
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
	TypeName() ITypeNameContext
	Colon() antlr.TerminalNode
	AssignmentExpression() IAssignmentExpressionContext

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
	{
		p.SetState(119)
		p.TypeName()
	}
	{
		p.SetState(120)
		p.Match(tinycParserColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(121)
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
	LeftParen() antlr.TerminalNode
	TypeName() ITypeNameContext
	RightParen() antlr.TerminalNode
	LeftBrace() antlr.TerminalNode
	InitializerList() IInitializerListContext
	RightBrace() antlr.TerminalNode
	AllLeftBracket() []antlr.TerminalNode
	LeftBracket(i int) antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllRightBracket() []antlr.TerminalNode
	RightBracket(i int) antlr.TerminalNode
	AllFuncCall() []IFuncCallContext
	FuncCall(i int) IFuncCallContext
	Comma() antlr.TerminalNode

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

func (s *PostfixExpressionContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(tinycParserLeftParen, 0)
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

func (s *PostfixExpressionContext) RightParen() antlr.TerminalNode {
	return s.GetToken(tinycParserRightParen, 0)
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

func (s *PostfixExpressionContext) AllFuncCall() []IFuncCallContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFuncCallContext); ok {
			len++
		}
	}

	tst := make([]IFuncCallContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFuncCallContext); ok {
			tst[i] = t.(IFuncCallContext)
			i++
		}
	}

	return tst
}

func (s *PostfixExpressionContext) FuncCall(i int) IFuncCallContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncCallContext); ok {
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

	return t.(IFuncCallContext)
}

func (s *PostfixExpressionContext) Comma() antlr.TerminalNode {
	return s.GetToken(tinycParserComma, 0)
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
	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(123)
			p.PrimaryExpression()
		}

	case 2:
		{
			p.SetState(124)
			p.Match(tinycParserLeftParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(125)
			p.TypeName()
		}
		{
			p.SetState(126)
			p.Match(tinycParserRightParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(127)
			p.Match(tinycParserLeftBrace)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(128)
			p.InitializerList()
		}
		p.SetState(130)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == tinycParserComma {
			{
				p.SetState(129)
				p.Match(tinycParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(132)
			p.Match(tinycParserRightBrace)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == tinycParserLeftParen || _la == tinycParserLeftBracket {
		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case tinycParserLeftBracket:
			{
				p.SetState(136)
				p.Match(tinycParserLeftBracket)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(137)
				p.Expression()
			}
			{
				p.SetState(138)
				p.Match(tinycParserRightBracket)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case tinycParserLeftParen:
			{
				p.SetState(140)
				p.FuncCall()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(145)
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

// IFuncCallContext is an interface to support dynamic dispatch.
type IFuncCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LeftParen() antlr.TerminalNode
	RightParen() antlr.TerminalNode
	AllAssignmentExpression() []IAssignmentExpressionContext
	AssignmentExpression(i int) IAssignmentExpressionContext
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

	// IsFuncCallContext differentiates from other interfaces.
	IsFuncCallContext()
}

type FuncCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncCallContext() *FuncCallContext {
	var p = new(FuncCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_funcCall
	return p
}

func InitEmptyFuncCallContext(p *FuncCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_funcCall
}

func (*FuncCallContext) IsFuncCallContext() {}

func NewFuncCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncCallContext {
	var p = new(FuncCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_funcCall

	return p
}

func (s *FuncCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncCallContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(tinycParserLeftParen, 0)
}

func (s *FuncCallContext) RightParen() antlr.TerminalNode {
	return s.GetToken(tinycParserRightParen, 0)
}

func (s *FuncCallContext) AllAssignmentExpression() []IAssignmentExpressionContext {
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

func (s *FuncCallContext) AssignmentExpression(i int) IAssignmentExpressionContext {
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

func (s *FuncCallContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(tinycParserComma)
}

func (s *FuncCallContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(tinycParserComma, i)
}

func (s *FuncCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterFuncCall(s)
	}
}

func (s *FuncCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitFuncCall(s)
	}
}

func (s *FuncCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitFuncCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) FuncCall() (localctx IFuncCallContext) {
	localctx = NewFuncCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, tinycParserRULE_funcCall)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(146)
		p.Match(tinycParserLeftParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&6341209184633094144) != 0 {
		{
			p.SetState(147)
			p.AssignmentExpression()
		}
		p.SetState(152)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == tinycParserComma {
			{
				p.SetState(148)
				p.Match(tinycParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(149)
				p.AssignmentExpression()
			}

			p.SetState(154)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(157)
		p.Match(tinycParserRightParen)
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

// IUnaryExpressionContext is an interface to support dynamic dispatch.
type IUnaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UnaryOperator() IUnaryOperatorContext
	CastExpression() ICastExpressionContext

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
	{
		p.SetState(159)
		p.UnaryOperator()
	}
	{
		p.SetState(160)
		p.CastExpression()
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
	Plus() antlr.TerminalNode
	Minus() antlr.TerminalNode
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

func (s *UnaryOperatorContext) Plus() antlr.TerminalNode {
	return s.GetToken(tinycParserPlus, 0)
}

func (s *UnaryOperatorContext) Minus() antlr.TerminalNode {
	return s.GetToken(tinycParserMinus, 0)
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
		p.SetState(162)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&140909287047168) != 0) {
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
	PostfixExpression() IPostfixExpressionContext
	UnaryExpression() IUnaryExpressionContext

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

func (s *CastExpressionContext) PostfixExpression() IPostfixExpressionContext {
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
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(164)
			p.Match(tinycParserLeftParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(165)
			p.TypeName()
		}
		{
			p.SetState(166)
			p.Match(tinycParserRightParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(167)
			p.CastExpression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(169)
			p.PostfixExpression()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(170)
			p.UnaryExpression()
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
		p.SetState(173)
		p.CastExpression()
	}
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3848290697216) != 0 {
		{
			p.SetState(174)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3848290697216) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(175)
			p.CastExpression()
		}

		p.SetState(180)
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
		p.SetState(181)
		p.MultiplicativeExpression()
	}
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == tinycParserPlus || _la == tinycParserMinus {
		{
			p.SetState(182)
			_la = p.GetTokenStream().LA(1)

			if !(_la == tinycParserPlus || _la == tinycParserMinus) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(183)
			p.MultiplicativeExpression()
		}

		p.SetState(188)
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
		p.SetState(189)
		p.AdditiveExpression()
	}
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8053063680) != 0 {
		{
			p.SetState(190)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8053063680) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(191)
			p.AdditiveExpression()
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
		p.SetState(197)
		p.RelationalExpression()
	}
	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == tinycParserEqual || _la == tinycParserNotEqual {
		{
			p.SetState(198)
			_la = p.GetTokenStream().LA(1)

			if !(_la == tinycParserEqual || _la == tinycParserNotEqual) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(199)
			p.RelationalExpression()
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

// ILogicalAndExpressionContext is an interface to support dynamic dispatch.
type ILogicalAndExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllEqualityExpression() []IEqualityExpressionContext
	EqualityExpression(i int) IEqualityExpressionContext
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

func (s *LogicalAndExpressionContext) AllEqualityExpression() []IEqualityExpressionContext {
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

func (s *LogicalAndExpressionContext) EqualityExpression(i int) IEqualityExpressionContext {
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
	p.EnterRule(localctx, 24, tinycParserRULE_logicalAndExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(205)
		p.EqualityExpression()
	}
	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == tinycParserAndAnd {
		{
			p.SetState(206)
			p.Match(tinycParserAndAnd)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(207)
			p.EqualityExpression()
		}

		p.SetState(212)
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
	p.EnterRule(localctx, 26, tinycParserRULE_logicalOrExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(213)
		p.LogicalAndExpression()
	}
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == tinycParserOrOr {
		{
			p.SetState(214)
			p.Match(tinycParserOrOr)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(215)
			p.LogicalAndExpression()
		}

		p.SetState(220)
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
	PostfixExpression() IPostfixExpressionContext
	AssignmentOperator() IAssignmentOperatorContext
	AssignmentExpression() IAssignmentExpressionContext

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

func (s *AssignmentExpressionContext) PostfixExpression() IPostfixExpressionContext {
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
	p.EnterRule(localctx, 28, tinycParserRULE_assignmentExpression)
	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(221)
			p.LogicalOrExpression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(222)
			p.PostfixExpression()
		}
		{
			p.SetState(223)
			p.AssignmentOperator()
		}
		{
			p.SetState(224)
			p.AssignmentExpression()
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
	p.EnterRule(localctx, 30, tinycParserRULE_assignmentOperator)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(228)
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
	p.EnterRule(localctx, 32, tinycParserRULE_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(230)
		p.AssignmentExpression()
	}
	p.SetState(235)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == tinycParserComma {
		{
			p.SetState(231)
			p.Match(tinycParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(232)
			p.AssignmentExpression()
		}

		p.SetState(237)
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
	p.EnterRule(localctx, 34, tinycParserRULE_declaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(238)
		p.DeclarationSpecifiers()
	}
	p.SetState(240)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == tinycParserLeftParen || _la == tinycParserIdentifier {
		{
			p.SetState(239)
			p.InitDeclaratorList()
		}

	}
	{
		p.SetState(242)
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
	AllTypeSpecifier() []ITypeSpecifierContext
	TypeSpecifier(i int) ITypeSpecifierContext

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

func (s *DeclarationSpecifiersContext) AllTypeSpecifier() []ITypeSpecifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeSpecifierContext); ok {
			len++
		}
	}

	tst := make([]ITypeSpecifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeSpecifierContext); ok {
			tst[i] = t.(ITypeSpecifierContext)
			i++
		}
	}

	return tst
}

func (s *DeclarationSpecifiersContext) TypeSpecifier(i int) ITypeSpecifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeSpecifierContext); ok {
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

	return t.(ITypeSpecifierContext)
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
	p.EnterRule(localctx, 36, tinycParserRULE_declarationSpecifiers)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4360) != 0) {
		{
			p.SetState(244)
			p.TypeSpecifier()
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
	p.EnterRule(localctx, 38, tinycParserRULE_initDeclaratorList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(249)
		p.InitDeclarator()
	}
	p.SetState(254)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == tinycParserComma {
		{
			p.SetState(250)
			p.Match(tinycParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(251)
			p.InitDeclarator()
		}

		p.SetState(256)
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
	p.EnterRule(localctx, 40, tinycParserRULE_initDeclarator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(257)
		p.declarator(0)
	}
	p.SetState(260)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == tinycParserAssign {
		{
			p.SetState(258)
			p.Match(tinycParserAssign)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(259)
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
	Char() antlr.TerminalNode
	Int() antlr.TerminalNode
	Float() antlr.TerminalNode

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

func (s *TypeSpecifierContext) Char() antlr.TerminalNode {
	return s.GetToken(tinycParserChar, 0)
}

func (s *TypeSpecifierContext) Int() antlr.TerminalNode {
	return s.GetToken(tinycParserInt, 0)
}

func (s *TypeSpecifierContext) Float() antlr.TerminalNode {
	return s.GetToken(tinycParserFloat, 0)
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
	p.EnterRule(localctx, 42, tinycParserRULE_typeSpecifier)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(262)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4360) != 0) {
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
	p.EnterRule(localctx, 44, tinycParserRULE_specifierQualifierList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(264)
		p.TypeSpecifier()
	}
	p.SetState(266)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4360) != 0 {
		{
			p.SetState(265)
			p.SpecifierQualifierList()
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
	p.EnterRule(localctx, 46, tinycParserRULE_structDeclaratorList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(268)
		p.StructDeclarator()
	}
	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == tinycParserComma {
		{
			p.SetState(269)
			p.Match(tinycParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(270)
			p.StructDeclarator()
		}

		p.SetState(275)
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
	LogicalOrExpression() ILogicalOrExpressionContext

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

func (s *StructDeclaratorContext) LogicalOrExpression() ILogicalOrExpressionContext {
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
	p.EnterRule(localctx, 48, tinycParserRULE_structDeclarator)
	var _la int

	p.SetState(282)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(276)
			p.declarator(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(278)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == tinycParserLeftParen || _la == tinycParserIdentifier {
			{
				p.SetState(277)
				p.declarator(0)
			}

		}
		{
			p.SetState(280)
			p.Match(tinycParserColon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(281)
			p.LogicalOrExpression()
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
	Identifier() antlr.TerminalNode
	LeftParen() antlr.TerminalNode
	Declarator() IDeclaratorContext
	RightParen() antlr.TerminalNode
	LeftBracket() antlr.TerminalNode
	RightBracket() antlr.TerminalNode
	AssignmentExpression() IAssignmentExpressionContext
	IdentifierList() IIdentifierListContext

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

func (s *DeclaratorContext) Identifier() antlr.TerminalNode {
	return s.GetToken(tinycParserIdentifier, 0)
}

func (s *DeclaratorContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(tinycParserLeftParen, 0)
}

func (s *DeclaratorContext) Declarator() IDeclaratorContext {
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

func (s *DeclaratorContext) RightParen() antlr.TerminalNode {
	return s.GetToken(tinycParserRightParen, 0)
}

func (s *DeclaratorContext) LeftBracket() antlr.TerminalNode {
	return s.GetToken(tinycParserLeftBracket, 0)
}

func (s *DeclaratorContext) RightBracket() antlr.TerminalNode {
	return s.GetToken(tinycParserRightBracket, 0)
}

func (s *DeclaratorContext) AssignmentExpression() IAssignmentExpressionContext {
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

func (s *DeclaratorContext) IdentifierList() IIdentifierListContext {
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
	return p.declarator(0)
}

func (p *tinycParser) declarator(_p int) (localctx IDeclaratorContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewDeclaratorContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IDeclaratorContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 50
	p.EnterRecursionRule(localctx, 50, tinycParserRULE_declarator, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(290)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case tinycParserIdentifier:
		{
			p.SetState(285)
			p.Match(tinycParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case tinycParserLeftParen:
		{
			p.SetState(286)
			p.Match(tinycParserLeftParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(287)
			p.declarator(0)
		}
		{
			p.SetState(288)
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
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(306)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(304)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) {
			case 1:
				localctx = NewDeclaratorContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, tinycParserRULE_declarator)
				p.SetState(292)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(293)
					p.Match(tinycParserLeftBracket)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(295)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&6341209184633094144) != 0 {
					{
						p.SetState(294)
						p.AssignmentExpression()
					}

				}
				{
					p.SetState(297)
					p.Match(tinycParserRightBracket)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 2:
				localctx = NewDeclaratorContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, tinycParserRULE_declarator)
				p.SetState(298)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(299)
					p.Match(tinycParserLeftParen)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(301)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == tinycParserIdentifier {
					{
						p.SetState(300)
						p.IdentifierList()
					}

				}
				{
					p.SetState(303)
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
		p.SetState(308)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 52, tinycParserRULE_parameterTypeList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(309)
		p.ParameterList()
	}
	p.SetState(312)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == tinycParserComma {
		{
			p.SetState(310)
			p.Match(tinycParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(311)
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
	p.EnterRule(localctx, 54, tinycParserRULE_parameterList)
	var _alt int

	p.SetState(323)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case tinycParserChar, tinycParserFloat, tinycParserInt:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(314)
			p.ParameterDeclaration()
		}
		p.SetState(319)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(315)
					p.Match(tinycParserComma)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(316)
					p.ParameterDeclaration()
				}

			}
			p.SetState(321)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	case tinycParserRightParen, tinycParserComma:
		p.EnterOuterAlt(localctx, 2)

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

// IParameterDeclarationContext is an interface to support dynamic dispatch.
type IParameterDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DeclarationSpecifiers() IDeclarationSpecifiersContext
	Declarator() IDeclaratorContext

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
	p.EnterRule(localctx, 56, tinycParserRULE_parameterDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(325)
		p.DeclarationSpecifiers()
	}
	{
		p.SetState(326)
		p.declarator(0)
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
	p.EnterRule(localctx, 58, tinycParserRULE_identifierList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(328)
		p.Match(tinycParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(333)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == tinycParserComma {
		{
			p.SetState(329)
			p.Match(tinycParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(330)
			p.Match(tinycParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(335)
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
	p.EnterRule(localctx, 60, tinycParserRULE_typeName)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(336)
		p.SpecifierQualifierList()
	}
	p.SetState(338)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == tinycParserLeftParen || _la == tinycParserLeftBracket {
		{
			p.SetState(337)
			p.abstractDeclarator(0)
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
	LeftParen() antlr.TerminalNode
	AbstractDeclarator() IAbstractDeclaratorContext
	RightParen() antlr.TerminalNode
	LeftBracket() antlr.TerminalNode
	RightBracket() antlr.TerminalNode
	AssignmentExpression() IAssignmentExpressionContext
	Star() antlr.TerminalNode
	ParameterTypeList() IParameterTypeListContext

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

func (s *AbstractDeclaratorContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(tinycParserLeftParen, 0)
}

func (s *AbstractDeclaratorContext) AbstractDeclarator() IAbstractDeclaratorContext {
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

func (s *AbstractDeclaratorContext) RightParen() antlr.TerminalNode {
	return s.GetToken(tinycParserRightParen, 0)
}

func (s *AbstractDeclaratorContext) LeftBracket() antlr.TerminalNode {
	return s.GetToken(tinycParserLeftBracket, 0)
}

func (s *AbstractDeclaratorContext) RightBracket() antlr.TerminalNode {
	return s.GetToken(tinycParserRightBracket, 0)
}

func (s *AbstractDeclaratorContext) AssignmentExpression() IAssignmentExpressionContext {
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

func (s *AbstractDeclaratorContext) Star() antlr.TerminalNode {
	return s.GetToken(tinycParserStar, 0)
}

func (s *AbstractDeclaratorContext) ParameterTypeList() IParameterTypeListContext {
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
	return p.abstractDeclarator(0)
}

func (p *tinycParser) abstractDeclarator(_p int) (localctx IAbstractDeclaratorContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewAbstractDeclaratorContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IAbstractDeclaratorContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 62
	p.EnterRecursionRule(localctx, 62, tinycParserRULE_abstractDeclarator, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(358)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(341)
			p.Match(tinycParserLeftParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(342)
			p.abstractDeclarator(0)
		}
		{
			p.SetState(343)
			p.Match(tinycParserRightParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(345)
			p.Match(tinycParserLeftBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(347)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&6341209184633094144) != 0 {
			{
				p.SetState(346)
				p.AssignmentExpression()
			}

		}
		{
			p.SetState(349)
			p.Match(tinycParserRightBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		{
			p.SetState(350)
			p.Match(tinycParserLeftBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(351)
			p.Match(tinycParserStar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(352)
			p.Match(tinycParserRightBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		{
			p.SetState(353)
			p.Match(tinycParserLeftParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(355)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(354)
				p.ParameterTypeList()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		{
			p.SetState(357)
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
	p.SetState(378)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(376)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAbstractDeclaratorContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, tinycParserRULE_abstractDeclarator)
				p.SetState(360)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(361)
					p.Match(tinycParserLeftBracket)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(363)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&6341209184633094144) != 0 {
					{
						p.SetState(362)
						p.AssignmentExpression()
					}

				}
				{
					p.SetState(365)
					p.Match(tinycParserRightBracket)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 2:
				localctx = NewAbstractDeclaratorContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, tinycParserRULE_abstractDeclarator)
				p.SetState(366)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(367)
					p.Match(tinycParserLeftBracket)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(368)
					p.Match(tinycParserStar)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(369)
					p.Match(tinycParserRightBracket)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 3:
				localctx = NewAbstractDeclaratorContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, tinycParserRULE_abstractDeclarator)
				p.SetState(370)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(371)
					p.Match(tinycParserLeftParen)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(373)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext()) == 1 {
					{
						p.SetState(372)
						p.ParameterTypeList()
					}

				} else if p.HasError() { // JIM
					goto errorExit
				}
				{
					p.SetState(375)
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
		p.SetState(380)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 64, tinycParserRULE_initializer)
	var _la int

	p.SetState(389)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case tinycParserLeftParen, tinycParserPlus, tinycParserMinus, tinycParserNot, tinycParserIdentifier, tinycParserConstant, tinycParserStringLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(381)
			p.AssignmentExpression()
		}

	case tinycParserLeftBrace:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(382)
			p.Match(tinycParserLeftBrace)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(383)
			p.InitializerList()
		}
		p.SetState(385)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == tinycParserComma {
			{
				p.SetState(384)
				p.Match(tinycParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(387)
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
	p.EnterRule(localctx, 66, tinycParserRULE_initializerList)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(392)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == tinycParserLeftBracket || _la == tinycParserDot {
		{
			p.SetState(391)
			p.Designation()
		}

	}
	{
		p.SetState(394)
		p.Initializer()
	}
	p.SetState(402)
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
				p.SetState(395)
				p.Match(tinycParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(397)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == tinycParserLeftBracket || _la == tinycParserDot {
				{
					p.SetState(396)
					p.Designation()
				}

			}
			{
				p.SetState(399)
				p.Initializer()
			}

		}
		p.SetState(404)
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
	p.EnterRule(localctx, 68, tinycParserRULE_designation)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(405)
		p.DesignatorList()
	}
	{
		p.SetState(406)
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
	p.EnterRule(localctx, 70, tinycParserRULE_designatorList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(409)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == tinycParserLeftBracket || _la == tinycParserDot {
		{
			p.SetState(408)
			p.Designator()
		}

		p.SetState(411)
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
	LogicalOrExpression() ILogicalOrExpressionContext
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

func (s *DesignatorContext) LogicalOrExpression() ILogicalOrExpressionContext {
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
	p.EnterRule(localctx, 72, tinycParserRULE_designator)
	p.SetState(419)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case tinycParserLeftBracket:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(413)
			p.Match(tinycParserLeftBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(414)
			p.LogicalOrExpression()
		}
		{
			p.SetState(415)
			p.Match(tinycParserRightBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case tinycParserDot:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(417)
			p.Match(tinycParserDot)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(418)
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
	p.EnterRule(localctx, 74, tinycParserRULE_statement)
	p.SetState(427)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 50, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(421)
			p.LabeledStatement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(422)
			p.CompoundStatement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(423)
			p.ExpressionStatement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(424)
			p.SelectionStatement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(425)
			p.IterationStatement()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(426)
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
	p.EnterRule(localctx, 76, tinycParserRULE_labeledStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(429)
		p.Match(tinycParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(430)
		p.Match(tinycParserColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(431)
		p.Statement()
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
	AllBlockItem() []IBlockItemContext
	BlockItem(i int) IBlockItemContext

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

func (s *CompoundStatementContext) AllBlockItem() []IBlockItemContext {
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

func (s *CompoundStatementContext) BlockItem(i int) IBlockItemContext {
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
	p.EnterRule(localctx, 78, tinycParserRULE_compoundStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(433)
		p.Match(tinycParserLeftBrace)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(437)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&6343460984583132424) != 0 {
		{
			p.SetState(434)
			p.BlockItem()
		}

		p.SetState(439)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(440)
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
	p.EnterRule(localctx, 80, tinycParserRULE_blockItem)
	p.SetState(444)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case tinycParserIf, tinycParserReturn, tinycParserWhile, tinycParserLeftParen, tinycParserLeftBrace, tinycParserPlus, tinycParserMinus, tinycParserNot, tinycParserSemi, tinycParserIdentifier, tinycParserConstant, tinycParserStringLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(442)
			p.Statement()
		}

	case tinycParserChar, tinycParserFloat, tinycParserInt:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(443)
			p.Declaration()
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
	p.EnterRule(localctx, 82, tinycParserRULE_expressionStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(447)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&6341209184633094144) != 0 {
		{
			p.SetState(446)
			p.Expression()
		}

	}
	{
		p.SetState(449)
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
	p.EnterRule(localctx, 84, tinycParserRULE_selectionStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(451)
		p.Match(tinycParserIf)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(452)
		p.Match(tinycParserLeftParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(453)
		p.Expression()
	}
	{
		p.SetState(454)
		p.Match(tinycParserRightParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(455)
		p.Statement()
	}
	p.SetState(458)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 54, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(456)
			p.Match(tinycParserElse)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(457)
			p.Statement()
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
	p.EnterRule(localctx, 86, tinycParserRULE_iterationStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(460)
		p.Match(tinycParserWhile)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(461)
		p.Match(tinycParserLeftParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(462)
		p.Expression()
	}
	{
		p.SetState(463)
		p.Match(tinycParserRightParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(464)
		p.Statement()
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
	Return() antlr.TerminalNode
	Expression() IExpressionContext
	Semi() antlr.TerminalNode

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

func (s *JumpStatementContext) Return() antlr.TerminalNode {
	return s.GetToken(tinycParserReturn, 0)
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

func (s *JumpStatementContext) Semi() antlr.TerminalNode {
	return s.GetToken(tinycParserSemi, 0)
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
	p.EnterRule(localctx, 88, tinycParserRULE_jumpStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(466)
		p.Match(tinycParserReturn)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(467)
		p.Expression()
	}
	{
		p.SetState(468)
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
	p.EnterRule(localctx, 90, tinycParserRULE_compilationUnit)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(471)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&578712552117113096) != 0) {
		{
			p.SetState(470)
			p.ExternalDeclaration()
		}

		p.SetState(473)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(475)
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
	p.EnterRule(localctx, 92, tinycParserRULE_externalDeclaration)
	p.SetState(480)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 56, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(477)
			p.FunctionDefinition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(478)
			p.Declaration()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(479)
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
	Identifier() antlr.TerminalNode
	LeftParen() antlr.TerminalNode
	ParameterList() IParameterListContext
	RightParen() antlr.TerminalNode
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

func (s *FunctionDefinitionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(tinycParserIdentifier, 0)
}

func (s *FunctionDefinitionContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(tinycParserLeftParen, 0)
}

func (s *FunctionDefinitionContext) ParameterList() IParameterListContext {
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

func (s *FunctionDefinitionContext) RightParen() antlr.TerminalNode {
	return s.GetToken(tinycParserRightParen, 0)
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
	p.EnterRule(localctx, 94, tinycParserRULE_functionDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(483)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4360) != 0 {
		{
			p.SetState(482)
			p.DeclarationSpecifiers()
		}

	}
	{
		p.SetState(485)
		p.Match(tinycParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(486)
		p.Match(tinycParserLeftParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(487)
		p.ParameterList()
	}
	{
		p.SetState(488)
		p.Match(tinycParserRightParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(490)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4360) != 0 {
		{
			p.SetState(489)
			p.DeclarationList()
		}

	}
	{
		p.SetState(492)
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
	p.EnterRule(localctx, 96, tinycParserRULE_declarationList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(495)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4360) != 0) {
		{
			p.SetState(494)
			p.Declaration()
		}

		p.SetState(497)
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
	case 25:
		var t *DeclaratorContext = nil
		if localctx != nil {
			t = localctx.(*DeclaratorContext)
		}
		return p.Declarator_Sempred(t, predIndex)

	case 31:
		var t *AbstractDeclaratorContext = nil
		if localctx != nil {
			t = localctx.(*AbstractDeclaratorContext)
		}
		return p.AbstractDeclarator_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *tinycParser) Declarator_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *tinycParser) AbstractDeclarator_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
