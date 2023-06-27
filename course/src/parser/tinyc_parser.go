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
		"", "'.'", "'char'", "'const'", "'else'", "'float'", "'if'", "'int'",
		"'return'", "'while'", "'('", "')'", "'['", "']'", "'{'", "'}'", "'<'",
		"'<='", "'>'", "'>='", "'+'", "'-'", "'*'", "'/'", "'%'", "'&'", "'|'",
		"'&&'", "'||'", "'^'", "'!'", "'~'", "'?'", "':'", "';'", "','", "'='",
		"'=='", "'!='",
	}
	staticData.SymbolicNames = []string{
		"", "", "Char", "Const", "Else", "Float", "If", "Int", "Return", "While",
		"LeftParen", "RightParen", "LeftBracket", "RightBracket", "LeftBrace",
		"RightBrace", "Less", "LessEqual", "Greater", "GreaterEqual", "Plus",
		"Minus", "Star", "Div", "Mod", "And", "Or", "AndAnd", "OrOr", "Caret",
		"Not", "Tilde", "Question", "Colon", "Semi", "Comma", "Assign", "Equal",
		"NotEqual", "Identifier", "Constant", "DigitSequence", "StringLiteral",
		"LineDirective", "Whitespace", "Newline", "BlockComment", "LineComment",
	}
	staticData.RuleNames = []string{
		"primaryExpression", "postfixExpression", "funcCall", "unaryExpression",
		"unaryOperator", "castExpression", "multiplicativeExpression", "additiveExpression",
		"relationalExpression", "equalityExpression", "logicalAndExpression",
		"logicalOrExpression", "assignmentExpression", "assignmentOperator",
		"expression", "declaration", "declarationSpecifiers", "initDeclaratorList",
		"initDeclarator", "typeSpecifier", "structDeclaratorList", "structDeclarator",
		"declarator", "parameterList", "parameterDeclaration", "identifierList",
		"initializer", "initializerList", "designation", "designatorList", "designator",
		"statement", "labeledStatement", "compoundStatement", "blockItem", "expressionStatement",
		"selectionStatement", "iterationStatement", "jumpStatement", "compilationUnit",
		"externalDeclaration", "functionDefinition", "declarationList",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 47, 405, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 1, 0, 1, 0, 1, 0, 4, 0, 90, 8, 0, 11, 0, 12, 0, 91, 1, 0, 1,
		0, 1, 0, 1, 0, 3, 0, 98, 8, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1,
		106, 8, 1, 10, 1, 12, 1, 109, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 115,
		8, 2, 10, 2, 12, 2, 118, 9, 2, 3, 2, 120, 8, 2, 1, 2, 1, 2, 1, 3, 1, 3,
		1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 3, 5, 131, 8, 5, 1, 6, 1, 6, 1, 6, 5, 6,
		136, 8, 6, 10, 6, 12, 6, 139, 9, 6, 1, 7, 1, 7, 1, 7, 5, 7, 144, 8, 7,
		10, 7, 12, 7, 147, 9, 7, 1, 8, 1, 8, 1, 8, 5, 8, 152, 8, 8, 10, 8, 12,
		8, 155, 9, 8, 1, 9, 1, 9, 1, 9, 5, 9, 160, 8, 9, 10, 9, 12, 9, 163, 9,
		9, 1, 10, 1, 10, 1, 10, 5, 10, 168, 8, 10, 10, 10, 12, 10, 171, 9, 10,
		1, 11, 1, 11, 1, 11, 5, 11, 176, 8, 11, 10, 11, 12, 11, 179, 9, 11, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 186, 8, 12, 1, 13, 1, 13, 1, 14,
		1, 14, 1, 14, 5, 14, 193, 8, 14, 10, 14, 12, 14, 196, 9, 14, 1, 15, 1,
		15, 3, 15, 200, 8, 15, 1, 15, 1, 15, 1, 16, 4, 16, 205, 8, 16, 11, 16,
		12, 16, 206, 1, 17, 1, 17, 1, 17, 5, 17, 212, 8, 17, 10, 17, 12, 17, 215,
		9, 17, 1, 18, 1, 18, 1, 18, 3, 18, 220, 8, 18, 1, 19, 1, 19, 1, 20, 1,
		20, 1, 20, 5, 20, 227, 8, 20, 10, 20, 12, 20, 230, 9, 20, 1, 21, 1, 21,
		3, 21, 234, 8, 21, 1, 21, 1, 21, 3, 21, 238, 8, 21, 1, 22, 1, 22, 1, 22,
		1, 22, 1, 22, 1, 22, 3, 22, 246, 8, 22, 1, 22, 1, 22, 1, 22, 3, 22, 251,
		8, 22, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 257, 8, 22, 1, 22, 5, 22, 260,
		8, 22, 10, 22, 12, 22, 263, 9, 22, 1, 23, 1, 23, 1, 23, 5, 23, 268, 8,
		23, 10, 23, 12, 23, 271, 9, 23, 1, 23, 3, 23, 274, 8, 23, 1, 24, 1, 24,
		1, 24, 1, 25, 1, 25, 1, 25, 5, 25, 282, 8, 25, 10, 25, 12, 25, 285, 9,
		25, 1, 26, 1, 26, 1, 26, 1, 26, 3, 26, 291, 8, 26, 1, 26, 1, 26, 3, 26,
		295, 8, 26, 1, 27, 3, 27, 298, 8, 27, 1, 27, 1, 27, 1, 27, 3, 27, 303,
		8, 27, 1, 27, 5, 27, 306, 8, 27, 10, 27, 12, 27, 309, 9, 27, 1, 28, 1,
		28, 1, 28, 1, 29, 4, 29, 315, 8, 29, 11, 29, 12, 29, 316, 1, 30, 1, 30,
		1, 30, 1, 30, 1, 30, 1, 30, 3, 30, 325, 8, 30, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 3, 31, 333, 8, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33,
		1, 33, 5, 33, 341, 8, 33, 10, 33, 12, 33, 344, 9, 33, 1, 33, 1, 33, 1,
		34, 1, 34, 3, 34, 350, 8, 34, 1, 35, 3, 35, 353, 8, 35, 1, 35, 1, 35, 1,
		36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 3, 36, 364, 8, 36, 1, 37,
		1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 4,
		39, 377, 8, 39, 11, 39, 12, 39, 378, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40,
		3, 40, 386, 8, 40, 1, 41, 3, 41, 389, 8, 41, 1, 41, 1, 41, 1, 41, 1, 41,
		1, 41, 3, 41, 396, 8, 41, 1, 41, 1, 41, 1, 42, 4, 42, 401, 8, 42, 11, 42,
		12, 42, 402, 1, 42, 0, 1, 44, 43, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
		22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56,
		58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 0, 6, 2, 0, 20,
		21, 30, 30, 1, 0, 22, 24, 1, 0, 20, 21, 1, 0, 16, 19, 1, 0, 37, 38, 3,
		0, 2, 2, 5, 5, 7, 7, 415, 0, 97, 1, 0, 0, 0, 2, 99, 1, 0, 0, 0, 4, 110,
		1, 0, 0, 0, 6, 123, 1, 0, 0, 0, 8, 126, 1, 0, 0, 0, 10, 130, 1, 0, 0, 0,
		12, 132, 1, 0, 0, 0, 14, 140, 1, 0, 0, 0, 16, 148, 1, 0, 0, 0, 18, 156,
		1, 0, 0, 0, 20, 164, 1, 0, 0, 0, 22, 172, 1, 0, 0, 0, 24, 185, 1, 0, 0,
		0, 26, 187, 1, 0, 0, 0, 28, 189, 1, 0, 0, 0, 30, 197, 1, 0, 0, 0, 32, 204,
		1, 0, 0, 0, 34, 208, 1, 0, 0, 0, 36, 216, 1, 0, 0, 0, 38, 221, 1, 0, 0,
		0, 40, 223, 1, 0, 0, 0, 42, 237, 1, 0, 0, 0, 44, 245, 1, 0, 0, 0, 46, 273,
		1, 0, 0, 0, 48, 275, 1, 0, 0, 0, 50, 278, 1, 0, 0, 0, 52, 294, 1, 0, 0,
		0, 54, 297, 1, 0, 0, 0, 56, 310, 1, 0, 0, 0, 58, 314, 1, 0, 0, 0, 60, 324,
		1, 0, 0, 0, 62, 332, 1, 0, 0, 0, 64, 334, 1, 0, 0, 0, 66, 338, 1, 0, 0,
		0, 68, 349, 1, 0, 0, 0, 70, 352, 1, 0, 0, 0, 72, 356, 1, 0, 0, 0, 74, 365,
		1, 0, 0, 0, 76, 371, 1, 0, 0, 0, 78, 376, 1, 0, 0, 0, 80, 385, 1, 0, 0,
		0, 82, 388, 1, 0, 0, 0, 84, 400, 1, 0, 0, 0, 86, 98, 5, 39, 0, 0, 87, 98,
		5, 40, 0, 0, 88, 90, 5, 42, 0, 0, 89, 88, 1, 0, 0, 0, 90, 91, 1, 0, 0,
		0, 91, 89, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 98, 1, 0, 0, 0, 93, 94,
		5, 10, 0, 0, 94, 95, 3, 28, 14, 0, 95, 96, 5, 11, 0, 0, 96, 98, 1, 0, 0,
		0, 97, 86, 1, 0, 0, 0, 97, 87, 1, 0, 0, 0, 97, 89, 1, 0, 0, 0, 97, 93,
		1, 0, 0, 0, 98, 1, 1, 0, 0, 0, 99, 107, 3, 0, 0, 0, 100, 101, 5, 12, 0,
		0, 101, 102, 3, 28, 14, 0, 102, 103, 5, 13, 0, 0, 103, 106, 1, 0, 0, 0,
		104, 106, 3, 4, 2, 0, 105, 100, 1, 0, 0, 0, 105, 104, 1, 0, 0, 0, 106,
		109, 1, 0, 0, 0, 107, 105, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 3, 1,
		0, 0, 0, 109, 107, 1, 0, 0, 0, 110, 119, 5, 10, 0, 0, 111, 116, 3, 24,
		12, 0, 112, 113, 5, 35, 0, 0, 113, 115, 3, 24, 12, 0, 114, 112, 1, 0, 0,
		0, 115, 118, 1, 0, 0, 0, 116, 114, 1, 0, 0, 0, 116, 117, 1, 0, 0, 0, 117,
		120, 1, 0, 0, 0, 118, 116, 1, 0, 0, 0, 119, 111, 1, 0, 0, 0, 119, 120,
		1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121, 122, 5, 11, 0, 0, 122, 5, 1, 0,
		0, 0, 123, 124, 3, 8, 4, 0, 124, 125, 3, 10, 5, 0, 125, 7, 1, 0, 0, 0,
		126, 127, 7, 0, 0, 0, 127, 9, 1, 0, 0, 0, 128, 131, 3, 2, 1, 0, 129, 131,
		3, 6, 3, 0, 130, 128, 1, 0, 0, 0, 130, 129, 1, 0, 0, 0, 131, 11, 1, 0,
		0, 0, 132, 137, 3, 10, 5, 0, 133, 134, 7, 1, 0, 0, 134, 136, 3, 10, 5,
		0, 135, 133, 1, 0, 0, 0, 136, 139, 1, 0, 0, 0, 137, 135, 1, 0, 0, 0, 137,
		138, 1, 0, 0, 0, 138, 13, 1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 140, 145, 3,
		12, 6, 0, 141, 142, 7, 2, 0, 0, 142, 144, 3, 12, 6, 0, 143, 141, 1, 0,
		0, 0, 144, 147, 1, 0, 0, 0, 145, 143, 1, 0, 0, 0, 145, 146, 1, 0, 0, 0,
		146, 15, 1, 0, 0, 0, 147, 145, 1, 0, 0, 0, 148, 153, 3, 14, 7, 0, 149,
		150, 7, 3, 0, 0, 150, 152, 3, 14, 7, 0, 151, 149, 1, 0, 0, 0, 152, 155,
		1, 0, 0, 0, 153, 151, 1, 0, 0, 0, 153, 154, 1, 0, 0, 0, 154, 17, 1, 0,
		0, 0, 155, 153, 1, 0, 0, 0, 156, 161, 3, 16, 8, 0, 157, 158, 7, 4, 0, 0,
		158, 160, 3, 16, 8, 0, 159, 157, 1, 0, 0, 0, 160, 163, 1, 0, 0, 0, 161,
		159, 1, 0, 0, 0, 161, 162, 1, 0, 0, 0, 162, 19, 1, 0, 0, 0, 163, 161, 1,
		0, 0, 0, 164, 169, 3, 18, 9, 0, 165, 166, 5, 27, 0, 0, 166, 168, 3, 18,
		9, 0, 167, 165, 1, 0, 0, 0, 168, 171, 1, 0, 0, 0, 169, 167, 1, 0, 0, 0,
		169, 170, 1, 0, 0, 0, 170, 21, 1, 0, 0, 0, 171, 169, 1, 0, 0, 0, 172, 177,
		3, 20, 10, 0, 173, 174, 5, 28, 0, 0, 174, 176, 3, 20, 10, 0, 175, 173,
		1, 0, 0, 0, 176, 179, 1, 0, 0, 0, 177, 175, 1, 0, 0, 0, 177, 178, 1, 0,
		0, 0, 178, 23, 1, 0, 0, 0, 179, 177, 1, 0, 0, 0, 180, 186, 3, 22, 11, 0,
		181, 182, 3, 2, 1, 0, 182, 183, 3, 26, 13, 0, 183, 184, 3, 24, 12, 0, 184,
		186, 1, 0, 0, 0, 185, 180, 1, 0, 0, 0, 185, 181, 1, 0, 0, 0, 186, 25, 1,
		0, 0, 0, 187, 188, 5, 36, 0, 0, 188, 27, 1, 0, 0, 0, 189, 194, 3, 24, 12,
		0, 190, 191, 5, 35, 0, 0, 191, 193, 3, 24, 12, 0, 192, 190, 1, 0, 0, 0,
		193, 196, 1, 0, 0, 0, 194, 192, 1, 0, 0, 0, 194, 195, 1, 0, 0, 0, 195,
		29, 1, 0, 0, 0, 196, 194, 1, 0, 0, 0, 197, 199, 3, 32, 16, 0, 198, 200,
		3, 34, 17, 0, 199, 198, 1, 0, 0, 0, 199, 200, 1, 0, 0, 0, 200, 201, 1,
		0, 0, 0, 201, 202, 5, 34, 0, 0, 202, 31, 1, 0, 0, 0, 203, 205, 3, 38, 19,
		0, 204, 203, 1, 0, 0, 0, 205, 206, 1, 0, 0, 0, 206, 204, 1, 0, 0, 0, 206,
		207, 1, 0, 0, 0, 207, 33, 1, 0, 0, 0, 208, 213, 3, 36, 18, 0, 209, 210,
		5, 35, 0, 0, 210, 212, 3, 36, 18, 0, 211, 209, 1, 0, 0, 0, 212, 215, 1,
		0, 0, 0, 213, 211, 1, 0, 0, 0, 213, 214, 1, 0, 0, 0, 214, 35, 1, 0, 0,
		0, 215, 213, 1, 0, 0, 0, 216, 219, 3, 44, 22, 0, 217, 218, 5, 36, 0, 0,
		218, 220, 3, 52, 26, 0, 219, 217, 1, 0, 0, 0, 219, 220, 1, 0, 0, 0, 220,
		37, 1, 0, 0, 0, 221, 222, 7, 5, 0, 0, 222, 39, 1, 0, 0, 0, 223, 228, 3,
		42, 21, 0, 224, 225, 5, 35, 0, 0, 225, 227, 3, 42, 21, 0, 226, 224, 1,
		0, 0, 0, 227, 230, 1, 0, 0, 0, 228, 226, 1, 0, 0, 0, 228, 229, 1, 0, 0,
		0, 229, 41, 1, 0, 0, 0, 230, 228, 1, 0, 0, 0, 231, 238, 3, 44, 22, 0, 232,
		234, 3, 44, 22, 0, 233, 232, 1, 0, 0, 0, 233, 234, 1, 0, 0, 0, 234, 235,
		1, 0, 0, 0, 235, 236, 5, 33, 0, 0, 236, 238, 3, 22, 11, 0, 237, 231, 1,
		0, 0, 0, 237, 233, 1, 0, 0, 0, 238, 43, 1, 0, 0, 0, 239, 240, 6, 22, -1,
		0, 240, 246, 5, 39, 0, 0, 241, 242, 5, 10, 0, 0, 242, 243, 3, 44, 22, 0,
		243, 244, 5, 11, 0, 0, 244, 246, 1, 0, 0, 0, 245, 239, 1, 0, 0, 0, 245,
		241, 1, 0, 0, 0, 246, 261, 1, 0, 0, 0, 247, 248, 10, 2, 0, 0, 248, 250,
		5, 12, 0, 0, 249, 251, 3, 24, 12, 0, 250, 249, 1, 0, 0, 0, 250, 251, 1,
		0, 0, 0, 251, 252, 1, 0, 0, 0, 252, 260, 5, 13, 0, 0, 253, 254, 10, 1,
		0, 0, 254, 256, 5, 10, 0, 0, 255, 257, 3, 50, 25, 0, 256, 255, 1, 0, 0,
		0, 256, 257, 1, 0, 0, 0, 257, 258, 1, 0, 0, 0, 258, 260, 5, 11, 0, 0, 259,
		247, 1, 0, 0, 0, 259, 253, 1, 0, 0, 0, 260, 263, 1, 0, 0, 0, 261, 259,
		1, 0, 0, 0, 261, 262, 1, 0, 0, 0, 262, 45, 1, 0, 0, 0, 263, 261, 1, 0,
		0, 0, 264, 269, 3, 48, 24, 0, 265, 266, 5, 35, 0, 0, 266, 268, 3, 48, 24,
		0, 267, 265, 1, 0, 0, 0, 268, 271, 1, 0, 0, 0, 269, 267, 1, 0, 0, 0, 269,
		270, 1, 0, 0, 0, 270, 274, 1, 0, 0, 0, 271, 269, 1, 0, 0, 0, 272, 274,
		1, 0, 0, 0, 273, 264, 1, 0, 0, 0, 273, 272, 1, 0, 0, 0, 274, 47, 1, 0,
		0, 0, 275, 276, 3, 32, 16, 0, 276, 277, 3, 44, 22, 0, 277, 49, 1, 0, 0,
		0, 278, 283, 5, 39, 0, 0, 279, 280, 5, 35, 0, 0, 280, 282, 5, 39, 0, 0,
		281, 279, 1, 0, 0, 0, 282, 285, 1, 0, 0, 0, 283, 281, 1, 0, 0, 0, 283,
		284, 1, 0, 0, 0, 284, 51, 1, 0, 0, 0, 285, 283, 1, 0, 0, 0, 286, 295, 3,
		24, 12, 0, 287, 288, 5, 14, 0, 0, 288, 290, 3, 54, 27, 0, 289, 291, 5,
		35, 0, 0, 290, 289, 1, 0, 0, 0, 290, 291, 1, 0, 0, 0, 291, 292, 1, 0, 0,
		0, 292, 293, 5, 15, 0, 0, 293, 295, 1, 0, 0, 0, 294, 286, 1, 0, 0, 0, 294,
		287, 1, 0, 0, 0, 295, 53, 1, 0, 0, 0, 296, 298, 3, 56, 28, 0, 297, 296,
		1, 0, 0, 0, 297, 298, 1, 0, 0, 0, 298, 299, 1, 0, 0, 0, 299, 307, 3, 52,
		26, 0, 300, 302, 5, 35, 0, 0, 301, 303, 3, 56, 28, 0, 302, 301, 1, 0, 0,
		0, 302, 303, 1, 0, 0, 0, 303, 304, 1, 0, 0, 0, 304, 306, 3, 52, 26, 0,
		305, 300, 1, 0, 0, 0, 306, 309, 1, 0, 0, 0, 307, 305, 1, 0, 0, 0, 307,
		308, 1, 0, 0, 0, 308, 55, 1, 0, 0, 0, 309, 307, 1, 0, 0, 0, 310, 311, 3,
		58, 29, 0, 311, 312, 5, 36, 0, 0, 312, 57, 1, 0, 0, 0, 313, 315, 3, 60,
		30, 0, 314, 313, 1, 0, 0, 0, 315, 316, 1, 0, 0, 0, 316, 314, 1, 0, 0, 0,
		316, 317, 1, 0, 0, 0, 317, 59, 1, 0, 0, 0, 318, 319, 5, 12, 0, 0, 319,
		320, 3, 22, 11, 0, 320, 321, 5, 13, 0, 0, 321, 325, 1, 0, 0, 0, 322, 323,
		5, 1, 0, 0, 323, 325, 5, 39, 0, 0, 324, 318, 1, 0, 0, 0, 324, 322, 1, 0,
		0, 0, 325, 61, 1, 0, 0, 0, 326, 333, 3, 64, 32, 0, 327, 333, 3, 66, 33,
		0, 328, 333, 3, 70, 35, 0, 329, 333, 3, 72, 36, 0, 330, 333, 3, 74, 37,
		0, 331, 333, 3, 76, 38, 0, 332, 326, 1, 0, 0, 0, 332, 327, 1, 0, 0, 0,
		332, 328, 1, 0, 0, 0, 332, 329, 1, 0, 0, 0, 332, 330, 1, 0, 0, 0, 332,
		331, 1, 0, 0, 0, 333, 63, 1, 0, 0, 0, 334, 335, 5, 39, 0, 0, 335, 336,
		5, 33, 0, 0, 336, 337, 3, 62, 31, 0, 337, 65, 1, 0, 0, 0, 338, 342, 5,
		14, 0, 0, 339, 341, 3, 68, 34, 0, 340, 339, 1, 0, 0, 0, 341, 344, 1, 0,
		0, 0, 342, 340, 1, 0, 0, 0, 342, 343, 1, 0, 0, 0, 343, 345, 1, 0, 0, 0,
		344, 342, 1, 0, 0, 0, 345, 346, 5, 15, 0, 0, 346, 67, 1, 0, 0, 0, 347,
		350, 3, 62, 31, 0, 348, 350, 3, 30, 15, 0, 349, 347, 1, 0, 0, 0, 349, 348,
		1, 0, 0, 0, 350, 69, 1, 0, 0, 0, 351, 353, 3, 28, 14, 0, 352, 351, 1, 0,
		0, 0, 352, 353, 1, 0, 0, 0, 353, 354, 1, 0, 0, 0, 354, 355, 5, 34, 0, 0,
		355, 71, 1, 0, 0, 0, 356, 357, 5, 6, 0, 0, 357, 358, 5, 10, 0, 0, 358,
		359, 3, 28, 14, 0, 359, 360, 5, 11, 0, 0, 360, 363, 3, 62, 31, 0, 361,
		362, 5, 4, 0, 0, 362, 364, 3, 62, 31, 0, 363, 361, 1, 0, 0, 0, 363, 364,
		1, 0, 0, 0, 364, 73, 1, 0, 0, 0, 365, 366, 5, 9, 0, 0, 366, 367, 5, 10,
		0, 0, 367, 368, 3, 28, 14, 0, 368, 369, 5, 11, 0, 0, 369, 370, 3, 62, 31,
		0, 370, 75, 1, 0, 0, 0, 371, 372, 5, 8, 0, 0, 372, 373, 3, 28, 14, 0, 373,
		374, 5, 34, 0, 0, 374, 77, 1, 0, 0, 0, 375, 377, 3, 80, 40, 0, 376, 375,
		1, 0, 0, 0, 377, 378, 1, 0, 0, 0, 378, 376, 1, 0, 0, 0, 378, 379, 1, 0,
		0, 0, 379, 380, 1, 0, 0, 0, 380, 381, 5, 0, 0, 1, 381, 79, 1, 0, 0, 0,
		382, 386, 3, 82, 41, 0, 383, 386, 3, 30, 15, 0, 384, 386, 5, 34, 0, 0,
		385, 382, 1, 0, 0, 0, 385, 383, 1, 0, 0, 0, 385, 384, 1, 0, 0, 0, 386,
		81, 1, 0, 0, 0, 387, 389, 3, 32, 16, 0, 388, 387, 1, 0, 0, 0, 388, 389,
		1, 0, 0, 0, 389, 390, 1, 0, 0, 0, 390, 391, 5, 39, 0, 0, 391, 392, 5, 10,
		0, 0, 392, 393, 3, 46, 23, 0, 393, 395, 5, 11, 0, 0, 394, 396, 3, 84, 42,
		0, 395, 394, 1, 0, 0, 0, 395, 396, 1, 0, 0, 0, 396, 397, 1, 0, 0, 0, 397,
		398, 3, 66, 33, 0, 398, 83, 1, 0, 0, 0, 399, 401, 3, 30, 15, 0, 400, 399,
		1, 0, 0, 0, 401, 402, 1, 0, 0, 0, 402, 400, 1, 0, 0, 0, 402, 403, 1, 0,
		0, 0, 403, 85, 1, 0, 0, 0, 47, 91, 97, 105, 107, 116, 119, 130, 137, 145,
		153, 161, 169, 177, 185, 194, 199, 206, 213, 219, 228, 233, 237, 245, 250,
		256, 259, 261, 269, 273, 283, 290, 294, 297, 302, 307, 316, 324, 332, 342,
		349, 352, 363, 378, 385, 388, 395, 402,
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
	tinycParserEOF           = antlr.TokenEOF
	tinycParserT__0          = 1
	tinycParserChar          = 2
	tinycParserConst         = 3
	tinycParserElse          = 4
	tinycParserFloat         = 5
	tinycParserIf            = 6
	tinycParserInt           = 7
	tinycParserReturn        = 8
	tinycParserWhile         = 9
	tinycParserLeftParen     = 10
	tinycParserRightParen    = 11
	tinycParserLeftBracket   = 12
	tinycParserRightBracket  = 13
	tinycParserLeftBrace     = 14
	tinycParserRightBrace    = 15
	tinycParserLess          = 16
	tinycParserLessEqual     = 17
	tinycParserGreater       = 18
	tinycParserGreaterEqual  = 19
	tinycParserPlus          = 20
	tinycParserMinus         = 21
	tinycParserStar          = 22
	tinycParserDiv           = 23
	tinycParserMod           = 24
	tinycParserAnd           = 25
	tinycParserOr            = 26
	tinycParserAndAnd        = 27
	tinycParserOrOr          = 28
	tinycParserCaret         = 29
	tinycParserNot           = 30
	tinycParserTilde         = 31
	tinycParserQuestion      = 32
	tinycParserColon         = 33
	tinycParserSemi          = 34
	tinycParserComma         = 35
	tinycParserAssign        = 36
	tinycParserEqual         = 37
	tinycParserNotEqual      = 38
	tinycParserIdentifier    = 39
	tinycParserConstant      = 40
	tinycParserDigitSequence = 41
	tinycParserStringLiteral = 42
	tinycParserLineDirective = 43
	tinycParserWhitespace    = 44
	tinycParserNewline       = 45
	tinycParserBlockComment  = 46
	tinycParserLineComment   = 47
)

// tinycParser rules.
const (
	tinycParserRULE_primaryExpression        = 0
	tinycParserRULE_postfixExpression        = 1
	tinycParserRULE_funcCall                 = 2
	tinycParserRULE_unaryExpression          = 3
	tinycParserRULE_unaryOperator            = 4
	tinycParserRULE_castExpression           = 5
	tinycParserRULE_multiplicativeExpression = 6
	tinycParserRULE_additiveExpression       = 7
	tinycParserRULE_relationalExpression     = 8
	tinycParserRULE_equalityExpression       = 9
	tinycParserRULE_logicalAndExpression     = 10
	tinycParserRULE_logicalOrExpression      = 11
	tinycParserRULE_assignmentExpression     = 12
	tinycParserRULE_assignmentOperator       = 13
	tinycParserRULE_expression               = 14
	tinycParserRULE_declaration              = 15
	tinycParserRULE_declarationSpecifiers    = 16
	tinycParserRULE_initDeclaratorList       = 17
	tinycParserRULE_initDeclarator           = 18
	tinycParserRULE_typeSpecifier            = 19
	tinycParserRULE_structDeclaratorList     = 20
	tinycParserRULE_structDeclarator         = 21
	tinycParserRULE_declarator               = 22
	tinycParserRULE_parameterList            = 23
	tinycParserRULE_parameterDeclaration     = 24
	tinycParserRULE_identifierList           = 25
	tinycParserRULE_initializer              = 26
	tinycParserRULE_initializerList          = 27
	tinycParserRULE_designation              = 28
	tinycParserRULE_designatorList           = 29
	tinycParserRULE_designator               = 30
	tinycParserRULE_statement                = 31
	tinycParserRULE_labeledStatement         = 32
	tinycParserRULE_compoundStatement        = 33
	tinycParserRULE_blockItem                = 34
	tinycParserRULE_expressionStatement      = 35
	tinycParserRULE_selectionStatement       = 36
	tinycParserRULE_iterationStatement       = 37
	tinycParserRULE_jumpStatement            = 38
	tinycParserRULE_compilationUnit          = 39
	tinycParserRULE_externalDeclaration      = 40
	tinycParserRULE_functionDefinition       = 41
	tinycParserRULE_declarationList          = 42
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

	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case tinycParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(86)
			p.Match(tinycParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case tinycParserConstant:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(87)
			p.Match(tinycParserConstant)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case tinycParserStringLiteral:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == tinycParserStringLiteral {
			{
				p.SetState(88)
				p.Match(tinycParserStringLiteral)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(91)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case tinycParserLeftParen:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(93)
			p.Match(tinycParserLeftParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(94)
			p.Expression()
		}
		{
			p.SetState(95)
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

// IPostfixExpressionContext is an interface to support dynamic dispatch.
type IPostfixExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PrimaryExpression() IPrimaryExpressionContext
	AllLeftBracket() []antlr.TerminalNode
	LeftBracket(i int) antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllRightBracket() []antlr.TerminalNode
	RightBracket(i int) antlr.TerminalNode
	AllFuncCall() []IFuncCallContext
	FuncCall(i int) IFuncCallContext

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
	p.EnterRule(localctx, 2, tinycParserRULE_postfixExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(99)
		p.PrimaryExpression()
	}
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == tinycParserLeftParen || _la == tinycParserLeftBracket {
		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case tinycParserLeftBracket:
			{
				p.SetState(100)
				p.Match(tinycParserLeftBracket)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(101)
				p.Expression()
			}
			{
				p.SetState(102)
				p.Match(tinycParserRightBracket)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case tinycParserLeftParen:
			{
				p.SetState(104)
				p.FuncCall()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(109)
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
	p.EnterRule(localctx, 4, tinycParserRULE_funcCall)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(110)
		p.Match(tinycParserLeftParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&6048390841344) != 0 {
		{
			p.SetState(111)
			p.AssignmentExpression()
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
				p.AssignmentExpression()
			}

			p.SetState(118)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(121)
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
	p.EnterRule(localctx, 6, tinycParserRULE_unaryExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(123)
		p.UnaryOperator()
	}
	{
		p.SetState(124)
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
	p.EnterRule(localctx, 8, tinycParserRULE_unaryOperator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(126)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1076887552) != 0) {
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
	p.EnterRule(localctx, 10, tinycParserRULE_castExpression)
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case tinycParserLeftParen, tinycParserIdentifier, tinycParserConstant, tinycParserStringLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(128)
			p.PostfixExpression()
		}

	case tinycParserPlus, tinycParserMinus, tinycParserNot:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(129)
			p.UnaryExpression()
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
	p.EnterRule(localctx, 12, tinycParserRULE_multiplicativeExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(132)
		p.CastExpression()
	}
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&29360128) != 0 {
		{
			p.SetState(133)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&29360128) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(134)
			p.CastExpression()
		}

		p.SetState(139)
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
	p.EnterRule(localctx, 14, tinycParserRULE_additiveExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(140)
		p.MultiplicativeExpression()
	}
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == tinycParserPlus || _la == tinycParserMinus {
		{
			p.SetState(141)
			_la = p.GetTokenStream().LA(1)

			if !(_la == tinycParserPlus || _la == tinycParserMinus) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(142)
			p.MultiplicativeExpression()
		}

		p.SetState(147)
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
	p.EnterRule(localctx, 16, tinycParserRULE_relationalExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(148)
		p.AdditiveExpression()
	}
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&983040) != 0 {
		{
			p.SetState(149)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&983040) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(150)
			p.AdditiveExpression()
		}

		p.SetState(155)
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
	p.EnterRule(localctx, 18, tinycParserRULE_equalityExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(156)
		p.RelationalExpression()
	}
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == tinycParserEqual || _la == tinycParserNotEqual {
		{
			p.SetState(157)
			_la = p.GetTokenStream().LA(1)

			if !(_la == tinycParserEqual || _la == tinycParserNotEqual) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(158)
			p.RelationalExpression()
		}

		p.SetState(163)
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
	p.EnterRule(localctx, 20, tinycParserRULE_logicalAndExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(164)
		p.EqualityExpression()
	}
	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == tinycParserAndAnd {
		{
			p.SetState(165)
			p.Match(tinycParserAndAnd)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(166)
			p.EqualityExpression()
		}

		p.SetState(171)
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
	p.EnterRule(localctx, 22, tinycParserRULE_logicalOrExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(172)
		p.LogicalAndExpression()
	}
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == tinycParserOrOr {
		{
			p.SetState(173)
			p.Match(tinycParserOrOr)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(174)
			p.LogicalAndExpression()
		}

		p.SetState(179)
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
	p.EnterRule(localctx, 24, tinycParserRULE_assignmentExpression)
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(180)
			p.LogicalOrExpression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(181)
			p.PostfixExpression()
		}
		{
			p.SetState(182)
			p.AssignmentOperator()
		}
		{
			p.SetState(183)
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
	p.EnterRule(localctx, 26, tinycParserRULE_assignmentOperator)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(187)
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
	p.EnterRule(localctx, 28, tinycParserRULE_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(189)
		p.AssignmentExpression()
	}
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == tinycParserComma {
		{
			p.SetState(190)
			p.Match(tinycParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(191)
			p.AssignmentExpression()
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
	p.EnterRule(localctx, 30, tinycParserRULE_declaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(197)
		p.DeclarationSpecifiers()
	}
	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == tinycParserLeftParen || _la == tinycParserIdentifier {
		{
			p.SetState(198)
			p.InitDeclaratorList()
		}

	}
	{
		p.SetState(201)
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
	p.EnterRule(localctx, 32, tinycParserRULE_declarationSpecifiers)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(204)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&164) != 0) {
		{
			p.SetState(203)
			p.TypeSpecifier()
		}

		p.SetState(206)
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
	p.EnterRule(localctx, 34, tinycParserRULE_initDeclaratorList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(208)
		p.InitDeclarator()
	}
	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == tinycParserComma {
		{
			p.SetState(209)
			p.Match(tinycParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(210)
			p.InitDeclarator()
		}

		p.SetState(215)
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
	p.EnterRule(localctx, 36, tinycParserRULE_initDeclarator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(216)
		p.declarator(0)
	}
	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == tinycParserAssign {
		{
			p.SetState(217)
			p.Match(tinycParserAssign)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(218)
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
	p.EnterRule(localctx, 38, tinycParserRULE_typeSpecifier)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(221)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&164) != 0) {
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
	p.EnterRule(localctx, 40, tinycParserRULE_structDeclaratorList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(223)
		p.StructDeclarator()
	}
	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == tinycParserComma {
		{
			p.SetState(224)
			p.Match(tinycParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(225)
			p.StructDeclarator()
		}

		p.SetState(230)
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
	p.EnterRule(localctx, 42, tinycParserRULE_structDeclarator)
	var _la int

	p.SetState(237)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(231)
			p.declarator(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(233)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == tinycParserLeftParen || _la == tinycParserIdentifier {
			{
				p.SetState(232)
				p.declarator(0)
			}

		}
		{
			p.SetState(235)
			p.Match(tinycParserColon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(236)
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
	_startState := 44
	p.EnterRecursionRule(localctx, 44, tinycParserRULE_declarator, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case tinycParserIdentifier:
		{
			p.SetState(240)
			p.Match(tinycParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case tinycParserLeftParen:
		{
			p.SetState(241)
			p.Match(tinycParserLeftParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(242)
			p.declarator(0)
		}
		{
			p.SetState(243)
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
	p.SetState(261)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(259)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) {
			case 1:
				localctx = NewDeclaratorContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, tinycParserRULE_declarator)
				p.SetState(247)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(248)
					p.Match(tinycParserLeftBracket)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(250)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&6048390841344) != 0 {
					{
						p.SetState(249)
						p.AssignmentExpression()
					}

				}
				{
					p.SetState(252)
					p.Match(tinycParserRightBracket)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 2:
				localctx = NewDeclaratorContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, tinycParserRULE_declarator)
				p.SetState(253)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(254)
					p.Match(tinycParserLeftParen)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(256)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == tinycParserIdentifier {
					{
						p.SetState(255)
						p.IdentifierList()
					}

				}
				{
					p.SetState(258)
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
		p.SetState(263)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 46, tinycParserRULE_parameterList)
	var _la int

	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case tinycParserChar, tinycParserFloat, tinycParserInt:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(264)
			p.ParameterDeclaration()
		}
		p.SetState(269)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == tinycParserComma {
			{
				p.SetState(265)
				p.Match(tinycParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(266)
				p.ParameterDeclaration()
			}

			p.SetState(271)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case tinycParserRightParen:
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
	p.EnterRule(localctx, 48, tinycParserRULE_parameterDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(275)
		p.DeclarationSpecifiers()
	}
	{
		p.SetState(276)
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
	p.EnterRule(localctx, 50, tinycParserRULE_identifierList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(278)
		p.Match(tinycParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == tinycParserComma {
		{
			p.SetState(279)
			p.Match(tinycParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(280)
			p.Match(tinycParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(285)
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
	p.EnterRule(localctx, 52, tinycParserRULE_initializer)
	var _la int

	p.SetState(294)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case tinycParserLeftParen, tinycParserPlus, tinycParserMinus, tinycParserNot, tinycParserIdentifier, tinycParserConstant, tinycParserStringLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(286)
			p.AssignmentExpression()
		}

	case tinycParserLeftBrace:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(287)
			p.Match(tinycParserLeftBrace)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(288)
			p.InitializerList()
		}
		p.SetState(290)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == tinycParserComma {
			{
				p.SetState(289)
				p.Match(tinycParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(292)
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
	p.EnterRule(localctx, 54, tinycParserRULE_initializerList)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(297)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == tinycParserT__0 || _la == tinycParserLeftBracket {
		{
			p.SetState(296)
			p.Designation()
		}

	}
	{
		p.SetState(299)
		p.Initializer()
	}
	p.SetState(307)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(300)
				p.Match(tinycParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(302)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == tinycParserT__0 || _la == tinycParserLeftBracket {
				{
					p.SetState(301)
					p.Designation()
				}

			}
			{
				p.SetState(304)
				p.Initializer()
			}

		}
		p.SetState(309)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 56, tinycParserRULE_designation)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(310)
		p.DesignatorList()
	}
	{
		p.SetState(311)
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
	p.EnterRule(localctx, 58, tinycParserRULE_designatorList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(314)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == tinycParserT__0 || _la == tinycParserLeftBracket {
		{
			p.SetState(313)
			p.Designator()
		}

		p.SetState(316)
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
	p.EnterRule(localctx, 60, tinycParserRULE_designator)
	p.SetState(324)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case tinycParserLeftBracket:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(318)
			p.Match(tinycParserLeftBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(319)
			p.LogicalOrExpression()
		}
		{
			p.SetState(320)
			p.Match(tinycParserRightBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case tinycParserT__0:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(322)
			p.Match(tinycParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(323)
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
	p.EnterRule(localctx, 62, tinycParserRULE_statement)
	p.SetState(332)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(326)
			p.LabeledStatement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(327)
			p.CompoundStatement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(328)
			p.ExpressionStatement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(329)
			p.SelectionStatement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(330)
			p.IterationStatement()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(331)
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
	p.EnterRule(localctx, 64, tinycParserRULE_labeledStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(334)
		p.Match(tinycParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(335)
		p.Match(tinycParserColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(336)
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
	p.EnterRule(localctx, 66, tinycParserRULE_compoundStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(338)
		p.Match(tinycParserLeftBrace)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(342)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&6065570727908) != 0 {
		{
			p.SetState(339)
			p.BlockItem()
		}

		p.SetState(344)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(345)
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
	p.EnterRule(localctx, 68, tinycParserRULE_blockItem)
	p.SetState(349)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case tinycParserIf, tinycParserReturn, tinycParserWhile, tinycParserLeftParen, tinycParserLeftBrace, tinycParserPlus, tinycParserMinus, tinycParserNot, tinycParserSemi, tinycParserIdentifier, tinycParserConstant, tinycParserStringLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(347)
			p.Statement()
		}

	case tinycParserChar, tinycParserFloat, tinycParserInt:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(348)
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
	p.EnterRule(localctx, 70, tinycParserRULE_expressionStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(352)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&6048390841344) != 0 {
		{
			p.SetState(351)
			p.Expression()
		}

	}
	{
		p.SetState(354)
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
	p.EnterRule(localctx, 72, tinycParserRULE_selectionStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(356)
		p.Match(tinycParserIf)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(357)
		p.Match(tinycParserLeftParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(358)
		p.Expression()
	}
	{
		p.SetState(359)
		p.Match(tinycParserRightParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(360)
		p.Statement()
	}
	p.SetState(363)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(361)
			p.Match(tinycParserElse)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(362)
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
	p.EnterRule(localctx, 74, tinycParserRULE_iterationStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(365)
		p.Match(tinycParserWhile)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(366)
		p.Match(tinycParserLeftParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(367)
		p.Expression()
	}
	{
		p.SetState(368)
		p.Match(tinycParserRightParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(369)
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
	p.EnterRule(localctx, 76, tinycParserRULE_jumpStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(371)
		p.Match(tinycParserReturn)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(372)
		p.Expression()
	}
	{
		p.SetState(373)
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
	p.EnterRule(localctx, 78, tinycParserRULE_compilationUnit)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(376)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&566935683236) != 0) {
		{
			p.SetState(375)
			p.ExternalDeclaration()
		}

		p.SetState(378)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(380)
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
	p.EnterRule(localctx, 80, tinycParserRULE_externalDeclaration)
	p.SetState(385)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(382)
			p.FunctionDefinition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(383)
			p.Declaration()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(384)
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
	p.EnterRule(localctx, 82, tinycParserRULE_functionDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(388)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&164) != 0 {
		{
			p.SetState(387)
			p.DeclarationSpecifiers()
		}

	}
	{
		p.SetState(390)
		p.Match(tinycParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(391)
		p.Match(tinycParserLeftParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(392)
		p.ParameterList()
	}
	{
		p.SetState(393)
		p.Match(tinycParserRightParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(395)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&164) != 0 {
		{
			p.SetState(394)
			p.DeclarationList()
		}

	}
	{
		p.SetState(397)
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
	p.EnterRule(localctx, 84, tinycParserRULE_declarationList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(400)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&164) != 0) {
		{
			p.SetState(399)
			p.Declaration()
		}

		p.SetState(402)
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
	case 22:
		var t *DeclaratorContext = nil
		if localctx != nil {
			t = localctx.(*DeclaratorContext)
		}
		return p.Declarator_Sempred(t, predIndex)

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
