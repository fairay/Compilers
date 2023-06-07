// Code generated from tinyc.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type tinycLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var TinycLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func tinyclexerLexerInit() {
	staticData := &TinycLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'if'", "'else'", "'while'", "'do'", "';'", "'{'", "'}'", "'('",
		"')'", "'='", "'<'", "'+'", "'-'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "STRING", "INT",
		"WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "STRING", "INT", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 16, 82, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6,
		1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12,
		1, 12, 1, 13, 4, 13, 70, 8, 13, 11, 13, 12, 13, 71, 1, 14, 4, 14, 75, 8,
		14, 11, 14, 12, 14, 76, 1, 15, 1, 15, 1, 15, 1, 15, 0, 0, 16, 1, 1, 3,
		2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12,
		25, 13, 27, 14, 29, 15, 31, 16, 1, 0, 3, 1, 0, 97, 122, 1, 0, 48, 57, 3,
		0, 9, 10, 13, 13, 32, 32, 83, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5,
		1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13,
		1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0,
		21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0,
		0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 1, 33, 1, 0, 0, 0, 3, 36, 1, 0, 0,
		0, 5, 41, 1, 0, 0, 0, 7, 47, 1, 0, 0, 0, 9, 50, 1, 0, 0, 0, 11, 52, 1,
		0, 0, 0, 13, 54, 1, 0, 0, 0, 15, 56, 1, 0, 0, 0, 17, 58, 1, 0, 0, 0, 19,
		60, 1, 0, 0, 0, 21, 62, 1, 0, 0, 0, 23, 64, 1, 0, 0, 0, 25, 66, 1, 0, 0,
		0, 27, 69, 1, 0, 0, 0, 29, 74, 1, 0, 0, 0, 31, 78, 1, 0, 0, 0, 33, 34,
		5, 105, 0, 0, 34, 35, 5, 102, 0, 0, 35, 2, 1, 0, 0, 0, 36, 37, 5, 101,
		0, 0, 37, 38, 5, 108, 0, 0, 38, 39, 5, 115, 0, 0, 39, 40, 5, 101, 0, 0,
		40, 4, 1, 0, 0, 0, 41, 42, 5, 119, 0, 0, 42, 43, 5, 104, 0, 0, 43, 44,
		5, 105, 0, 0, 44, 45, 5, 108, 0, 0, 45, 46, 5, 101, 0, 0, 46, 6, 1, 0,
		0, 0, 47, 48, 5, 100, 0, 0, 48, 49, 5, 111, 0, 0, 49, 8, 1, 0, 0, 0, 50,
		51, 5, 59, 0, 0, 51, 10, 1, 0, 0, 0, 52, 53, 5, 123, 0, 0, 53, 12, 1, 0,
		0, 0, 54, 55, 5, 125, 0, 0, 55, 14, 1, 0, 0, 0, 56, 57, 5, 40, 0, 0, 57,
		16, 1, 0, 0, 0, 58, 59, 5, 41, 0, 0, 59, 18, 1, 0, 0, 0, 60, 61, 5, 61,
		0, 0, 61, 20, 1, 0, 0, 0, 62, 63, 5, 60, 0, 0, 63, 22, 1, 0, 0, 0, 64,
		65, 5, 43, 0, 0, 65, 24, 1, 0, 0, 0, 66, 67, 5, 45, 0, 0, 67, 26, 1, 0,
		0, 0, 68, 70, 7, 0, 0, 0, 69, 68, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 69,
		1, 0, 0, 0, 71, 72, 1, 0, 0, 0, 72, 28, 1, 0, 0, 0, 73, 75, 7, 1, 0, 0,
		74, 73, 1, 0, 0, 0, 75, 76, 1, 0, 0, 0, 76, 74, 1, 0, 0, 0, 76, 77, 1,
		0, 0, 0, 77, 30, 1, 0, 0, 0, 78, 79, 7, 2, 0, 0, 79, 80, 1, 0, 0, 0, 80,
		81, 6, 15, 0, 0, 81, 32, 1, 0, 0, 0, 3, 0, 71, 76, 1, 6, 0, 0,
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

// tinycLexerInit initializes any static state used to implement tinycLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewtinycLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func TinycLexerInit() {
	staticData := &TinycLexerLexerStaticData
	staticData.once.Do(tinyclexerLexerInit)
}

// NewtinycLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewtinycLexer(input antlr.CharStream) *tinycLexer {
	TinycLexerInit()
	l := new(tinycLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &TinycLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "tinyc.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// tinycLexer tokens.
const (
	tinycLexerT__0   = 1
	tinycLexerT__1   = 2
	tinycLexerT__2   = 3
	tinycLexerT__3   = 4
	tinycLexerT__4   = 5
	tinycLexerT__5   = 6
	tinycLexerT__6   = 7
	tinycLexerT__7   = 8
	tinycLexerT__8   = 9
	tinycLexerT__9   = 10
	tinycLexerT__10  = 11
	tinycLexerT__11  = 12
	tinycLexerT__12  = 13
	tinycLexerSTRING = 14
	tinycLexerINT    = 15
	tinycLexerWS     = 16
)
