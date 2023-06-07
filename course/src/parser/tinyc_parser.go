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
		"", "'if'", "'else'", "'while'", "'do'", "';'", "'{'", "'}'", "'('",
		"')'", "'='", "'<'", "'+'", "'-'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "STRING", "INT",
		"WS",
	}
	staticData.RuleNames = []string{
		"program", "statement", "paren_expr", "expr", "test", "sum_", "term",
		"id_", "integer",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 16, 101, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 4, 0, 20, 8, 0,
		11, 0, 12, 0, 21, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 5, 1, 48, 8, 1, 10, 1, 12, 1, 51, 9, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 3, 1, 58, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 3, 3, 69, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 76, 8,
		4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 87, 8, 5,
		10, 5, 12, 5, 90, 9, 5, 1, 6, 1, 6, 1, 6, 3, 6, 95, 8, 6, 1, 7, 1, 7, 1,
		8, 1, 8, 1, 8, 0, 1, 10, 9, 0, 2, 4, 6, 8, 10, 12, 14, 16, 0, 0, 105, 0,
		19, 1, 0, 0, 0, 2, 57, 1, 0, 0, 0, 4, 59, 1, 0, 0, 0, 6, 68, 1, 0, 0, 0,
		8, 75, 1, 0, 0, 0, 10, 77, 1, 0, 0, 0, 12, 94, 1, 0, 0, 0, 14, 96, 1, 0,
		0, 0, 16, 98, 1, 0, 0, 0, 18, 20, 3, 2, 1, 0, 19, 18, 1, 0, 0, 0, 20, 21,
		1, 0, 0, 0, 21, 19, 1, 0, 0, 0, 21, 22, 1, 0, 0, 0, 22, 23, 1, 0, 0, 0,
		23, 24, 5, 0, 0, 1, 24, 1, 1, 0, 0, 0, 25, 26, 5, 1, 0, 0, 26, 27, 3, 4,
		2, 0, 27, 28, 3, 2, 1, 0, 28, 58, 1, 0, 0, 0, 29, 30, 5, 1, 0, 0, 30, 31,
		3, 4, 2, 0, 31, 32, 3, 2, 1, 0, 32, 33, 5, 2, 0, 0, 33, 34, 3, 2, 1, 0,
		34, 58, 1, 0, 0, 0, 35, 36, 5, 3, 0, 0, 36, 37, 3, 4, 2, 0, 37, 38, 3,
		2, 1, 0, 38, 58, 1, 0, 0, 0, 39, 40, 5, 4, 0, 0, 40, 41, 3, 2, 1, 0, 41,
		42, 5, 3, 0, 0, 42, 43, 3, 4, 2, 0, 43, 44, 5, 5, 0, 0, 44, 58, 1, 0, 0,
		0, 45, 49, 5, 6, 0, 0, 46, 48, 3, 2, 1, 0, 47, 46, 1, 0, 0, 0, 48, 51,
		1, 0, 0, 0, 49, 47, 1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 52, 1, 0, 0, 0,
		51, 49, 1, 0, 0, 0, 52, 58, 5, 7, 0, 0, 53, 54, 3, 6, 3, 0, 54, 55, 5,
		5, 0, 0, 55, 58, 1, 0, 0, 0, 56, 58, 5, 5, 0, 0, 57, 25, 1, 0, 0, 0, 57,
		29, 1, 0, 0, 0, 57, 35, 1, 0, 0, 0, 57, 39, 1, 0, 0, 0, 57, 45, 1, 0, 0,
		0, 57, 53, 1, 0, 0, 0, 57, 56, 1, 0, 0, 0, 58, 3, 1, 0, 0, 0, 59, 60, 5,
		8, 0, 0, 60, 61, 3, 6, 3, 0, 61, 62, 5, 9, 0, 0, 62, 5, 1, 0, 0, 0, 63,
		69, 3, 8, 4, 0, 64, 65, 3, 14, 7, 0, 65, 66, 5, 10, 0, 0, 66, 67, 3, 6,
		3, 0, 67, 69, 1, 0, 0, 0, 68, 63, 1, 0, 0, 0, 68, 64, 1, 0, 0, 0, 69, 7,
		1, 0, 0, 0, 70, 76, 3, 10, 5, 0, 71, 72, 3, 10, 5, 0, 72, 73, 5, 11, 0,
		0, 73, 74, 3, 10, 5, 0, 74, 76, 1, 0, 0, 0, 75, 70, 1, 0, 0, 0, 75, 71,
		1, 0, 0, 0, 76, 9, 1, 0, 0, 0, 77, 78, 6, 5, -1, 0, 78, 79, 3, 12, 6, 0,
		79, 88, 1, 0, 0, 0, 80, 81, 10, 2, 0, 0, 81, 82, 5, 12, 0, 0, 82, 87, 3,
		12, 6, 0, 83, 84, 10, 1, 0, 0, 84, 85, 5, 13, 0, 0, 85, 87, 3, 12, 6, 0,
		86, 80, 1, 0, 0, 0, 86, 83, 1, 0, 0, 0, 87, 90, 1, 0, 0, 0, 88, 86, 1,
		0, 0, 0, 88, 89, 1, 0, 0, 0, 89, 11, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0, 91,
		95, 3, 14, 7, 0, 92, 95, 3, 16, 8, 0, 93, 95, 3, 4, 2, 0, 94, 91, 1, 0,
		0, 0, 94, 92, 1, 0, 0, 0, 94, 93, 1, 0, 0, 0, 95, 13, 1, 0, 0, 0, 96, 97,
		5, 14, 0, 0, 97, 15, 1, 0, 0, 0, 98, 99, 5, 15, 0, 0, 99, 17, 1, 0, 0,
		0, 8, 21, 49, 57, 68, 75, 86, 88, 94,
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
	tinycParserEOF    = antlr.TokenEOF
	tinycParserT__0   = 1
	tinycParserT__1   = 2
	tinycParserT__2   = 3
	tinycParserT__3   = 4
	tinycParserT__4   = 5
	tinycParserT__5   = 6
	tinycParserT__6   = 7
	tinycParserT__7   = 8
	tinycParserT__8   = 9
	tinycParserT__9   = 10
	tinycParserT__10  = 11
	tinycParserT__11  = 12
	tinycParserT__12  = 13
	tinycParserSTRING = 14
	tinycParserINT    = 15
	tinycParserWS     = 16
)

// tinycParser rules.
const (
	tinycParserRULE_program    = 0
	tinycParserRULE_statement  = 1
	tinycParserRULE_paren_expr = 2
	tinycParserRULE_expr       = 3
	tinycParserRULE_test       = 4
	tinycParserRULE_sum_       = 5
	tinycParserRULE_term       = 6
	tinycParserRULE_id_        = 7
	tinycParserRULE_integer    = 8
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(tinycParserEOF, 0)
}

func (s *ProgramContext) AllStatement() []IStatementContext {
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

func (s *ProgramContext) Statement(i int) IStatementContext {
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

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, tinycParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&49530) != 0) {
		{
			p.SetState(18)
			p.Statement()
		}

		p.SetState(21)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(23)
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

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Paren_expr() IParen_exprContext
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext
	Expr() IExprContext

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

func (s *StatementContext) Paren_expr() IParen_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParen_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParen_exprContext)
}

func (s *StatementContext) AllStatement() []IStatementContext {
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

func (s *StatementContext) Statement(i int) IStatementContext {
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

func (s *StatementContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
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
	p.EnterRule(localctx, 2, tinycParserRULE_statement)
	var _la int

	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(25)
			p.Match(tinycParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(26)
			p.Paren_expr()
		}
		{
			p.SetState(27)
			p.Statement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(29)
			p.Match(tinycParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(30)
			p.Paren_expr()
		}
		{
			p.SetState(31)
			p.Statement()
		}
		{
			p.SetState(32)
			p.Match(tinycParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(33)
			p.Statement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(35)
			p.Match(tinycParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(36)
			p.Paren_expr()
		}
		{
			p.SetState(37)
			p.Statement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(39)
			p.Match(tinycParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(40)
			p.Statement()
		}
		{
			p.SetState(41)
			p.Match(tinycParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(42)
			p.Paren_expr()
		}
		{
			p.SetState(43)
			p.Match(tinycParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(45)
			p.Match(tinycParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&49530) != 0 {
			{
				p.SetState(46)
				p.Statement()
			}

			p.SetState(51)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(52)
			p.Match(tinycParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(53)
			p.Expr()
		}
		{
			p.SetState(54)
			p.Match(tinycParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(56)
			p.Match(tinycParserT__4)
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

// IParen_exprContext is an interface to support dynamic dispatch.
type IParen_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext

	// IsParen_exprContext differentiates from other interfaces.
	IsParen_exprContext()
}

type Paren_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParen_exprContext() *Paren_exprContext {
	var p = new(Paren_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_paren_expr
	return p
}

func InitEmptyParen_exprContext(p *Paren_exprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_paren_expr
}

func (*Paren_exprContext) IsParen_exprContext() {}

func NewParen_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Paren_exprContext {
	var p = new(Paren_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_paren_expr

	return p
}

func (s *Paren_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Paren_exprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Paren_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Paren_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Paren_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterParen_expr(s)
	}
}

func (s *Paren_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitParen_expr(s)
	}
}

func (s *Paren_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitParen_expr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) Paren_expr() (localctx IParen_exprContext) {
	localctx = NewParen_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, tinycParserRULE_paren_expr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(59)
		p.Match(tinycParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(60)
		p.Expr()
	}
	{
		p.SetState(61)
		p.Match(tinycParserT__8)
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

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Test() ITestContext
	Id_() IId_Context
	Expr() IExprContext

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) Test() ITestContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITestContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITestContext)
}

func (s *ExprContext) Id_() IId_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IId_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IId_Context)
}

func (s *ExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (s *ExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, tinycParserRULE_expr)
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(63)
			p.Test()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(64)
			p.Id_()
		}
		{
			p.SetState(65)
			p.Match(tinycParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(66)
			p.Expr()
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

// ITestContext is an interface to support dynamic dispatch.
type ITestContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSum_() []ISum_Context
	Sum_(i int) ISum_Context

	// IsTestContext differentiates from other interfaces.
	IsTestContext()
}

type TestContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTestContext() *TestContext {
	var p = new(TestContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_test
	return p
}

func InitEmptyTestContext(p *TestContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_test
}

func (*TestContext) IsTestContext() {}

func NewTestContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TestContext {
	var p = new(TestContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_test

	return p
}

func (s *TestContext) GetParser() antlr.Parser { return s.parser }

func (s *TestContext) AllSum_() []ISum_Context {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISum_Context); ok {
			len++
		}
	}

	tst := make([]ISum_Context, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISum_Context); ok {
			tst[i] = t.(ISum_Context)
			i++
		}
	}

	return tst
}

func (s *TestContext) Sum_(i int) ISum_Context {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISum_Context); ok {
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

	return t.(ISum_Context)
}

func (s *TestContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TestContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TestContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterTest(s)
	}
}

func (s *TestContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitTest(s)
	}
}

func (s *TestContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitTest(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) Test() (localctx ITestContext) {
	localctx = NewTestContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, tinycParserRULE_test)
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(70)
			p.sum_(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(71)
			p.sum_(0)
		}
		{
			p.SetState(72)
			p.Match(tinycParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(73)
			p.sum_(0)
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

// ISum_Context is an interface to support dynamic dispatch.
type ISum_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Term() ITermContext
	Sum_() ISum_Context

	// IsSum_Context differentiates from other interfaces.
	IsSum_Context()
}

type Sum_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySum_Context() *Sum_Context {
	var p = new(Sum_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_sum_
	return p
}

func InitEmptySum_Context(p *Sum_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_sum_
}

func (*Sum_Context) IsSum_Context() {}

func NewSum_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sum_Context {
	var p = new(Sum_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_sum_

	return p
}

func (s *Sum_Context) GetParser() antlr.Parser { return s.parser }

func (s *Sum_Context) Term() ITermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *Sum_Context) Sum_() ISum_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISum_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISum_Context)
}

func (s *Sum_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sum_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Sum_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterSum_(s)
	}
}

func (s *Sum_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitSum_(s)
	}
}

func (s *Sum_Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitSum_(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) Sum_() (localctx ISum_Context) {
	return p.sum_(0)
}

func (p *tinycParser) sum_(_p int) (localctx ISum_Context) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewSum_Context(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ISum_Context = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, tinycParserRULE_sum_, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(78)
		p.Term()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(86)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
			case 1:
				localctx = NewSum_Context(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, tinycParserRULE_sum_)
				p.SetState(80)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(81)
					p.Match(tinycParserT__11)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(82)
					p.Term()
				}

			case 2:
				localctx = NewSum_Context(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, tinycParserRULE_sum_)
				p.SetState(83)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(84)
					p.Match(tinycParserT__12)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(85)
					p.Term()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(90)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
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

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Id_() IId_Context
	Integer() IIntegerContext
	Paren_expr() IParen_exprContext

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_term
	return p
}

func InitEmptyTermContext(p *TermContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_term
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) Id_() IId_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IId_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IId_Context)
}

func (s *TermContext) Integer() IIntegerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntegerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *TermContext) Paren_expr() IParen_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParen_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParen_exprContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (s *TermContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitTerm(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) Term() (localctx ITermContext) {
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, tinycParserRULE_term)
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case tinycParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(91)
			p.Id_()
		}

	case tinycParserINT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(92)
			p.Integer()
		}

	case tinycParserT__7:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(93)
			p.Paren_expr()
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

// IId_Context is an interface to support dynamic dispatch.
type IId_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode

	// IsId_Context differentiates from other interfaces.
	IsId_Context()
}

type Id_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyId_Context() *Id_Context {
	var p = new(Id_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_id_
	return p
}

func InitEmptyId_Context(p *Id_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_id_
}

func (*Id_Context) IsId_Context() {}

func NewId_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Id_Context {
	var p = new(Id_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_id_

	return p
}

func (s *Id_Context) GetParser() antlr.Parser { return s.parser }

func (s *Id_Context) STRING() antlr.TerminalNode {
	return s.GetToken(tinycParserSTRING, 0)
}

func (s *Id_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Id_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Id_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterId_(s)
	}
}

func (s *Id_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitId_(s)
	}
}

func (s *Id_Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitId_(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) Id_() (localctx IId_Context) {
	localctx = NewId_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, tinycParserRULE_id_)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)
		p.Match(tinycParserSTRING)
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

// IIntegerContext is an interface to support dynamic dispatch.
type IIntegerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode

	// IsIntegerContext differentiates from other interfaces.
	IsIntegerContext()
}

type IntegerContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerContext() *IntegerContext {
	var p = new(IntegerContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_integer
	return p
}

func InitEmptyIntegerContext(p *IntegerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tinycParserRULE_integer
}

func (*IntegerContext) IsIntegerContext() {}

func NewIntegerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerContext {
	var p = new(IntegerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_integer

	return p
}

func (s *IntegerContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerContext) INT() antlr.TerminalNode {
	return s.GetToken(tinycParserINT, 0)
}

func (s *IntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterInteger(s)
	}
}

func (s *IntegerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitInteger(s)
	}
}

func (s *IntegerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitInteger(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) Integer() (localctx IIntegerContext) {
	localctx = NewIntegerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, tinycParserRULE_integer)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(98)
		p.Match(tinycParserINT)
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

func (p *tinycParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 5:
		var t *Sum_Context = nil
		if localctx != nil {
			t = localctx.(*Sum_Context)
		}
		return p.Sum__Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *tinycParser) Sum__Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
