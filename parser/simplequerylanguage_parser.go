// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // SimpleQueryLanguage

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type SimpleQueryLanguageParser struct {
	*antlr.BaseParser
}

var simplequerylanguageParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func simplequerylanguageParserInit() {
	staticData := &simplequerylanguageParserStaticData
	staticData.literalNames = []string{
		"", "':'", "'<'", "'<='", "'=='", "'>'", "'>='", "'!='", "'&&'", "'||'",
		"'!'", "'+'", "'-'", "'*'", "'/'", "'('", "')'", "'['", "']'", "'?'",
		"'??'", "'.'", "','", "'{'", "'}'", "'null'", "'false'", "'true'",
	}
	staticData.symbolicNames = []string{
		"", "", "Less", "LessOrEqual", "Equal", "Greater", "GreaterOrEqual",
		"NotEqual", "And", "Or", "Not", "Add", "Subtract", "Multiply", "Divide",
		"LeftParen", "RightParen", "LeftBracket", "RightBracket", "Question",
		"NullCoalescing", "Dot", "Comma", "LeftBrace", "RightBrace", "Null",
		"False", "True", "ESCAPED_STRING", "STRING", "NUMBER", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "function_param", "field", "array", "object", "object_field",
		"name", "index",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 31, 158, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 3, 2, 38, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 79, 8, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 3, 2, 85, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 5, 2, 96, 8, 2, 10, 2, 12, 2, 99, 9, 2, 3, 2, 101, 8,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 109, 8, 2, 5, 2, 111, 8, 2,
		10, 2, 12, 2, 114, 9, 2, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 120, 8, 3, 10, 3,
		12, 3, 123, 9, 3, 3, 3, 125, 8, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4,
		5, 4, 133, 8, 4, 10, 4, 12, 4, 136, 9, 4, 3, 4, 138, 8, 4, 1, 4, 1, 4,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 149, 8, 5, 1, 6, 1, 6,
		1, 7, 1, 7, 1, 7, 3, 7, 156, 8, 7, 1, 7, 0, 1, 4, 8, 0, 2, 4, 6, 8, 10,
		12, 14, 0, 2, 1, 0, 13, 14, 1, 0, 11, 12, 187, 0, 16, 1, 0, 0, 0, 2, 18,
		1, 0, 0, 0, 4, 37, 1, 0, 0, 0, 6, 115, 1, 0, 0, 0, 8, 128, 1, 0, 0, 0,
		10, 148, 1, 0, 0, 0, 12, 150, 1, 0, 0, 0, 14, 155, 1, 0, 0, 0, 16, 17,
		3, 4, 2, 0, 17, 1, 1, 0, 0, 0, 18, 19, 3, 4, 2, 0, 19, 3, 1, 0, 0, 0, 20,
		21, 6, 2, -1, 0, 21, 22, 5, 15, 0, 0, 22, 23, 3, 4, 2, 0, 23, 24, 5, 16,
		0, 0, 24, 38, 1, 0, 0, 0, 25, 38, 5, 27, 0, 0, 26, 38, 5, 26, 0, 0, 27,
		38, 5, 25, 0, 0, 28, 38, 5, 30, 0, 0, 29, 38, 5, 28, 0, 0, 30, 38, 3, 6,
		3, 0, 31, 38, 3, 8, 4, 0, 32, 38, 3, 12, 6, 0, 33, 34, 5, 12, 0, 0, 34,
		38, 3, 4, 2, 15, 35, 36, 5, 10, 0, 0, 36, 38, 3, 4, 2, 2, 37, 20, 1, 0,
		0, 0, 37, 25, 1, 0, 0, 0, 37, 26, 1, 0, 0, 0, 37, 27, 1, 0, 0, 0, 37, 28,
		1, 0, 0, 0, 37, 29, 1, 0, 0, 0, 37, 30, 1, 0, 0, 0, 37, 31, 1, 0, 0, 0,
		37, 32, 1, 0, 0, 0, 37, 33, 1, 0, 0, 0, 37, 35, 1, 0, 0, 0, 38, 112, 1,
		0, 0, 0, 39, 40, 10, 14, 0, 0, 40, 41, 7, 0, 0, 0, 41, 111, 3, 4, 2, 15,
		42, 43, 10, 13, 0, 0, 43, 44, 7, 1, 0, 0, 44, 111, 3, 4, 2, 14, 45, 46,
		10, 12, 0, 0, 46, 47, 5, 21, 0, 0, 47, 48, 5, 21, 0, 0, 48, 111, 3, 4,
		2, 13, 49, 50, 10, 11, 0, 0, 50, 51, 5, 20, 0, 0, 51, 111, 3, 4, 2, 12,
		52, 53, 10, 10, 0, 0, 53, 54, 5, 4, 0, 0, 54, 111, 3, 4, 2, 11, 55, 56,
		10, 9, 0, 0, 56, 57, 5, 2, 0, 0, 57, 111, 3, 4, 2, 10, 58, 59, 10, 8, 0,
		0, 59, 60, 5, 3, 0, 0, 60, 111, 3, 4, 2, 9, 61, 62, 10, 7, 0, 0, 62, 63,
		5, 5, 0, 0, 63, 111, 3, 4, 2, 8, 64, 65, 10, 6, 0, 0, 65, 66, 5, 6, 0,
		0, 66, 111, 3, 4, 2, 7, 67, 68, 10, 5, 0, 0, 68, 69, 5, 7, 0, 0, 69, 111,
		3, 4, 2, 6, 70, 71, 10, 4, 0, 0, 71, 72, 5, 8, 0, 0, 72, 111, 3, 4, 2,
		5, 73, 74, 10, 3, 0, 0, 74, 75, 5, 9, 0, 0, 75, 111, 3, 4, 2, 4, 76, 78,
		10, 18, 0, 0, 77, 79, 5, 19, 0, 0, 78, 77, 1, 0, 0, 0, 78, 79, 1, 0, 0,
		0, 79, 80, 1, 0, 0, 0, 80, 81, 5, 21, 0, 0, 81, 111, 3, 12, 6, 0, 82, 84,
		10, 17, 0, 0, 83, 85, 5, 19, 0, 0, 84, 83, 1, 0, 0, 0, 84, 85, 1, 0, 0,
		0, 85, 86, 1, 0, 0, 0, 86, 87, 5, 17, 0, 0, 87, 88, 3, 14, 7, 0, 88, 89,
		5, 18, 0, 0, 89, 111, 1, 0, 0, 0, 90, 91, 10, 16, 0, 0, 91, 100, 5, 15,
		0, 0, 92, 97, 3, 2, 1, 0, 93, 94, 5, 22, 0, 0, 94, 96, 3, 2, 1, 0, 95,
		93, 1, 0, 0, 0, 96, 99, 1, 0, 0, 0, 97, 95, 1, 0, 0, 0, 97, 98, 1, 0, 0,
		0, 98, 101, 1, 0, 0, 0, 99, 97, 1, 0, 0, 0, 100, 92, 1, 0, 0, 0, 100, 101,
		1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 111, 5, 16, 0, 0, 103, 104, 10,
		1, 0, 0, 104, 105, 5, 19, 0, 0, 105, 108, 3, 4, 2, 0, 106, 107, 5, 1, 0,
		0, 107, 109, 3, 4, 2, 0, 108, 106, 1, 0, 0, 0, 108, 109, 1, 0, 0, 0, 109,
		111, 1, 0, 0, 0, 110, 39, 1, 0, 0, 0, 110, 42, 1, 0, 0, 0, 110, 45, 1,
		0, 0, 0, 110, 49, 1, 0, 0, 0, 110, 52, 1, 0, 0, 0, 110, 55, 1, 0, 0, 0,
		110, 58, 1, 0, 0, 0, 110, 61, 1, 0, 0, 0, 110, 64, 1, 0, 0, 0, 110, 67,
		1, 0, 0, 0, 110, 70, 1, 0, 0, 0, 110, 73, 1, 0, 0, 0, 110, 76, 1, 0, 0,
		0, 110, 82, 1, 0, 0, 0, 110, 90, 1, 0, 0, 0, 110, 103, 1, 0, 0, 0, 111,
		114, 1, 0, 0, 0, 112, 110, 1, 0, 0, 0, 112, 113, 1, 0, 0, 0, 113, 5, 1,
		0, 0, 0, 114, 112, 1, 0, 0, 0, 115, 124, 5, 17, 0, 0, 116, 121, 3, 4, 2,
		0, 117, 118, 5, 22, 0, 0, 118, 120, 3, 4, 2, 0, 119, 117, 1, 0, 0, 0, 120,
		123, 1, 0, 0, 0, 121, 119, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 125,
		1, 0, 0, 0, 123, 121, 1, 0, 0, 0, 124, 116, 1, 0, 0, 0, 124, 125, 1, 0,
		0, 0, 125, 126, 1, 0, 0, 0, 126, 127, 5, 18, 0, 0, 127, 7, 1, 0, 0, 0,
		128, 137, 5, 23, 0, 0, 129, 134, 3, 10, 5, 0, 130, 131, 5, 22, 0, 0, 131,
		133, 3, 10, 5, 0, 132, 130, 1, 0, 0, 0, 133, 136, 1, 0, 0, 0, 134, 132,
		1, 0, 0, 0, 134, 135, 1, 0, 0, 0, 135, 138, 1, 0, 0, 0, 136, 134, 1, 0,
		0, 0, 137, 129, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0,
		139, 140, 5, 24, 0, 0, 140, 9, 1, 0, 0, 0, 141, 142, 3, 12, 6, 0, 142,
		143, 5, 1, 0, 0, 143, 144, 3, 4, 2, 0, 144, 149, 1, 0, 0, 0, 145, 146,
		5, 28, 0, 0, 146, 147, 5, 1, 0, 0, 147, 149, 3, 4, 2, 0, 148, 141, 1, 0,
		0, 0, 148, 145, 1, 0, 0, 0, 149, 11, 1, 0, 0, 0, 150, 151, 5, 29, 0, 0,
		151, 13, 1, 0, 0, 0, 152, 156, 3, 4, 2, 0, 153, 156, 5, 30, 0, 0, 154,
		156, 5, 28, 0, 0, 155, 152, 1, 0, 0, 0, 155, 153, 1, 0, 0, 0, 155, 154,
		1, 0, 0, 0, 156, 15, 1, 0, 0, 0, 14, 37, 78, 84, 97, 100, 108, 110, 112,
		121, 124, 134, 137, 148, 155,
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

// SimpleQueryLanguageParserInit initializes any static state used to implement SimpleQueryLanguageParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSimpleQueryLanguageParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SimpleQueryLanguageParserInit() {
	staticData := &simplequerylanguageParserStaticData
	staticData.once.Do(simplequerylanguageParserInit)
}

// NewSimpleQueryLanguageParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSimpleQueryLanguageParser(input antlr.TokenStream) *SimpleQueryLanguageParser {
	SimpleQueryLanguageParserInit()
	this := new(SimpleQueryLanguageParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &simplequerylanguageParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

// SimpleQueryLanguageParser tokens.
const (
	SimpleQueryLanguageParserEOF            = antlr.TokenEOF
	SimpleQueryLanguageParserT__0           = 1
	SimpleQueryLanguageParserLess           = 2
	SimpleQueryLanguageParserLessOrEqual    = 3
	SimpleQueryLanguageParserEqual          = 4
	SimpleQueryLanguageParserGreater        = 5
	SimpleQueryLanguageParserGreaterOrEqual = 6
	SimpleQueryLanguageParserNotEqual       = 7
	SimpleQueryLanguageParserAnd            = 8
	SimpleQueryLanguageParserOr             = 9
	SimpleQueryLanguageParserNot            = 10
	SimpleQueryLanguageParserAdd            = 11
	SimpleQueryLanguageParserSubtract       = 12
	SimpleQueryLanguageParserMultiply       = 13
	SimpleQueryLanguageParserDivide         = 14
	SimpleQueryLanguageParserLeftParen      = 15
	SimpleQueryLanguageParserRightParen     = 16
	SimpleQueryLanguageParserLeftBracket    = 17
	SimpleQueryLanguageParserRightBracket   = 18
	SimpleQueryLanguageParserQuestion       = 19
	SimpleQueryLanguageParserNullCoalescing = 20
	SimpleQueryLanguageParserDot            = 21
	SimpleQueryLanguageParserComma          = 22
	SimpleQueryLanguageParserLeftBrace      = 23
	SimpleQueryLanguageParserRightBrace     = 24
	SimpleQueryLanguageParserNull           = 25
	SimpleQueryLanguageParserFalse          = 26
	SimpleQueryLanguageParserTrue           = 27
	SimpleQueryLanguageParserESCAPED_STRING = 28
	SimpleQueryLanguageParserSTRING         = 29
	SimpleQueryLanguageParserNUMBER         = 30
	SimpleQueryLanguageParserWS             = 31
)

// SimpleQueryLanguageParser rules.
const (
	SimpleQueryLanguageParserRULE_expression     = 0
	SimpleQueryLanguageParserRULE_function_param = 1
	SimpleQueryLanguageParserRULE_field          = 2
	SimpleQueryLanguageParserRULE_array          = 3
	SimpleQueryLanguageParserRULE_object         = 4
	SimpleQueryLanguageParserRULE_object_field   = 5
	SimpleQueryLanguageParserRULE_name           = 6
	SimpleQueryLanguageParserRULE_index          = 7
)

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleQueryLanguageParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleQueryLanguageParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Field() IFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleQueryLanguageListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleQueryLanguageListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleQueryLanguageVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleQueryLanguageParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SimpleQueryLanguageParserRULE_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(16)
		p.field(0)
	}

	return localctx
}

// IFunction_paramContext is an interface to support dynamic dispatch.
type IFunction_paramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunction_paramContext differentiates from other interfaces.
	IsFunction_paramContext()
}

type Function_paramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_paramContext() *Function_paramContext {
	var p = new(Function_paramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleQueryLanguageParserRULE_function_param
	return p
}

func (*Function_paramContext) IsFunction_paramContext() {}

func NewFunction_paramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_paramContext {
	var p = new(Function_paramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleQueryLanguageParserRULE_function_param

	return p
}

func (s *Function_paramContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_paramContext) Field() IFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *Function_paramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_paramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_paramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleQueryLanguageListener); ok {
		listenerT.EnterFunction_param(s)
	}
}

func (s *Function_paramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleQueryLanguageListener); ok {
		listenerT.ExitFunction_param(s)
	}
}

func (s *Function_paramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleQueryLanguageVisitor:
		return t.VisitFunction_param(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleQueryLanguageParser) Function_param() (localctx IFunction_paramContext) {
	this := p
	_ = this

	localctx = NewFunction_paramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SimpleQueryLanguageParserRULE_function_param)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(18)
		p.field(0)
	}

	return localctx
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleQueryLanguageParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleQueryLanguageParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(SimpleQueryLanguageParserLeftParen, 0)
}

func (s *FieldContext) AllField() []IFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldContext); ok {
			len++
		}
	}

	tst := make([]IFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldContext); ok {
			tst[i] = t.(IFieldContext)
			i++
		}
	}

	return tst
}

func (s *FieldContext) Field(i int) IFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
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

	return t.(IFieldContext)
}

func (s *FieldContext) RightParen() antlr.TerminalNode {
	return s.GetToken(SimpleQueryLanguageParserRightParen, 0)
}

func (s *FieldContext) True() antlr.TerminalNode {
	return s.GetToken(SimpleQueryLanguageParserTrue, 0)
}

func (s *FieldContext) False() antlr.TerminalNode {
	return s.GetToken(SimpleQueryLanguageParserFalse, 0)
}

func (s *FieldContext) Null() antlr.TerminalNode {
	return s.GetToken(SimpleQueryLanguageParserNull, 0)
}

func (s *FieldContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SimpleQueryLanguageParserNUMBER, 0)
}

func (s *FieldContext) ESCAPED_STRING() antlr.TerminalNode {
	return s.GetToken(SimpleQueryLanguageParserESCAPED_STRING, 0)
}

func (s *FieldContext) Array() IArrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}

func (s *FieldContext) Object() IObjectContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjectContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjectContext)
}

func (s *FieldContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *FieldContext) Subtract() antlr.TerminalNode {
	return s.GetToken(SimpleQueryLanguageParserSubtract, 0)
}

func (s *FieldContext) Not() antlr.TerminalNode {
	return s.GetToken(SimpleQueryLanguageParserNot, 0)
}

func (s *FieldContext) Divide() antlr.TerminalNode {
	return s.GetToken(SimpleQueryLanguageParserDivide, 0)
}

func (s *FieldContext) Multiply() antlr.TerminalNode {
	return s.GetToken(SimpleQueryLanguageParserMultiply, 0)
}

func (s *FieldContext) Add() antlr.TerminalNode {
	return s.GetToken(SimpleQueryLanguageParserAdd, 0)
}

func (s *FieldContext) AllDot() []antlr.TerminalNode {
	return s.GetTokens(SimpleQueryLanguageParserDot)
}

func (s *FieldContext) Dot(i int) antlr.TerminalNode {
	return s.GetToken(SimpleQueryLanguageParserDot, i)
}

func (s *FieldContext) NullCoalescing() antlr.TerminalNode {
	return s.GetToken(SimpleQueryLanguageParserNullCoalescing, 0)
}

func (s *FieldContext) Equal() antlr.TerminalNode {
	return s.GetToken(SimpleQueryLanguageParserEqual, 0)
}

func (s *FieldContext) Less() antlr.TerminalNode {
	return s.GetToken(SimpleQueryLanguageParserLess, 0)
}

func (s *FieldContext) LessOrEqual() antlr.TerminalNode {
	return s.GetToken(SimpleQueryLanguageParserLessOrEqual, 0)
}

func (s *FieldContext) Greater() antlr.TerminalNode {
	return s.GetToken(SimpleQueryLanguageParserGreater, 0)
}

func (s *FieldContext) GreaterOrEqual() antlr.TerminalNode {
	return s.GetToken(SimpleQueryLanguageParserGreaterOrEqual, 0)
}

func (s *FieldContext) NotEqual() antlr.TerminalNode {
	return s.GetToken(SimpleQueryLanguageParserNotEqual, 0)
}

func (s *FieldContext) And() antlr.TerminalNode {
	return s.GetToken(SimpleQueryLanguageParserAnd, 0)
}

func (s *FieldContext) Or() antlr.TerminalNode {
	return s.GetToken(SimpleQueryLanguageParserOr, 0)
}

func (s *FieldContext) Question() antlr.TerminalNode {
	return s.GetToken(SimpleQueryLanguageParserQuestion, 0)
}

func (s *FieldContext) LeftBracket() antlr.TerminalNode {
	return s.GetToken(SimpleQueryLanguageParserLeftBracket, 0)
}

func (s *FieldContext) Index() IIndexContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexContext)
}

func (s *FieldContext) RightBracket() antlr.TerminalNode {
	return s.GetToken(SimpleQueryLanguageParserRightBracket, 0)
}

func (s *FieldContext) AllFunction_param() []IFunction_paramContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunction_paramContext); ok {
			len++
		}
	}

	tst := make([]IFunction_paramContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunction_paramContext); ok {
			tst[i] = t.(IFunction_paramContext)
			i++
		}
	}

	return tst
}

func (s *FieldContext) Function_param(i int) IFunction_paramContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_paramContext); ok {
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

	return t.(IFunction_paramContext)
}

func (s *FieldContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(SimpleQueryLanguageParserComma)
}

func (s *FieldContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(SimpleQueryLanguageParserComma, i)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleQueryLanguageListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleQueryLanguageListener); ok {
		listenerT.ExitField(s)
	}
}

func (s *FieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleQueryLanguageVisitor:
		return t.VisitField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleQueryLanguageParser) Field() (localctx IFieldContext) {
	return p.field(0)
}

func (p *SimpleQueryLanguageParser) field(_p int) (localctx IFieldContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewFieldContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IFieldContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, SimpleQueryLanguageParserRULE_field, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(37)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SimpleQueryLanguageParserLeftParen:
		{
			p.SetState(21)
			p.Match(SimpleQueryLanguageParserLeftParen)
		}
		{
			p.SetState(22)
			p.field(0)
		}
		{
			p.SetState(23)
			p.Match(SimpleQueryLanguageParserRightParen)
		}

	case SimpleQueryLanguageParserTrue:
		{
			p.SetState(25)
			p.Match(SimpleQueryLanguageParserTrue)
		}

	case SimpleQueryLanguageParserFalse:
		{
			p.SetState(26)
			p.Match(SimpleQueryLanguageParserFalse)
		}

	case SimpleQueryLanguageParserNull:
		{
			p.SetState(27)
			p.Match(SimpleQueryLanguageParserNull)
		}

	case SimpleQueryLanguageParserNUMBER:
		{
			p.SetState(28)
			p.Match(SimpleQueryLanguageParserNUMBER)
		}

	case SimpleQueryLanguageParserESCAPED_STRING:
		{
			p.SetState(29)
			p.Match(SimpleQueryLanguageParserESCAPED_STRING)
		}

	case SimpleQueryLanguageParserLeftBracket:
		{
			p.SetState(30)
			p.Array()
		}

	case SimpleQueryLanguageParserLeftBrace:
		{
			p.SetState(31)
			p.Object()
		}

	case SimpleQueryLanguageParserSTRING:
		{
			p.SetState(32)
			p.Name()
		}

	case SimpleQueryLanguageParserSubtract:
		{
			p.SetState(33)
			p.Match(SimpleQueryLanguageParserSubtract)
		}
		{
			p.SetState(34)
			p.field(15)
		}

	case SimpleQueryLanguageParserNot:
		{
			p.SetState(35)
			p.Match(SimpleQueryLanguageParserNot)
		}
		{
			p.SetState(36)
			p.field(2)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(110)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
			case 1:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SimpleQueryLanguageParserRULE_field)
				p.SetState(39)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(40)
					_la = p.GetTokenStream().LA(1)

					if !(_la == SimpleQueryLanguageParserMultiply || _la == SimpleQueryLanguageParserDivide) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(41)
					p.field(15)
				}

			case 2:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SimpleQueryLanguageParserRULE_field)
				p.SetState(42)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(43)
					_la = p.GetTokenStream().LA(1)

					if !(_la == SimpleQueryLanguageParserAdd || _la == SimpleQueryLanguageParserSubtract) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(44)
					p.field(14)
				}

			case 3:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SimpleQueryLanguageParserRULE_field)
				p.SetState(45)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(46)
					p.Match(SimpleQueryLanguageParserDot)
				}
				{
					p.SetState(47)
					p.Match(SimpleQueryLanguageParserDot)
				}
				{
					p.SetState(48)
					p.field(13)
				}

			case 4:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SimpleQueryLanguageParserRULE_field)
				p.SetState(49)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(50)
					p.Match(SimpleQueryLanguageParserNullCoalescing)
				}
				{
					p.SetState(51)
					p.field(12)
				}

			case 5:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SimpleQueryLanguageParserRULE_field)
				p.SetState(52)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(53)
					p.Match(SimpleQueryLanguageParserEqual)
				}
				{
					p.SetState(54)
					p.field(11)
				}

			case 6:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SimpleQueryLanguageParserRULE_field)
				p.SetState(55)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(56)
					p.Match(SimpleQueryLanguageParserLess)
				}
				{
					p.SetState(57)
					p.field(10)
				}

			case 7:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SimpleQueryLanguageParserRULE_field)
				p.SetState(58)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(59)
					p.Match(SimpleQueryLanguageParserLessOrEqual)
				}
				{
					p.SetState(60)
					p.field(9)
				}

			case 8:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SimpleQueryLanguageParserRULE_field)
				p.SetState(61)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(62)
					p.Match(SimpleQueryLanguageParserGreater)
				}
				{
					p.SetState(63)
					p.field(8)
				}

			case 9:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SimpleQueryLanguageParserRULE_field)
				p.SetState(64)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(65)
					p.Match(SimpleQueryLanguageParserGreaterOrEqual)
				}
				{
					p.SetState(66)
					p.field(7)
				}

			case 10:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SimpleQueryLanguageParserRULE_field)
				p.SetState(67)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(68)
					p.Match(SimpleQueryLanguageParserNotEqual)
				}
				{
					p.SetState(69)
					p.field(6)
				}

			case 11:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SimpleQueryLanguageParserRULE_field)
				p.SetState(70)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(71)
					p.Match(SimpleQueryLanguageParserAnd)
				}
				{
					p.SetState(72)
					p.field(5)
				}

			case 12:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SimpleQueryLanguageParserRULE_field)
				p.SetState(73)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(74)
					p.Match(SimpleQueryLanguageParserOr)
				}
				{
					p.SetState(75)
					p.field(4)
				}

			case 13:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SimpleQueryLanguageParserRULE_field)
				p.SetState(76)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
				}

				p.SetState(78)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == SimpleQueryLanguageParserQuestion {
					{
						p.SetState(77)
						p.Match(SimpleQueryLanguageParserQuestion)
					}

				}
				{
					p.SetState(80)
					p.Match(SimpleQueryLanguageParserDot)
				}
				{
					p.SetState(81)
					p.Name()
				}

			case 14:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SimpleQueryLanguageParserRULE_field)
				p.SetState(82)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
				}

				p.SetState(84)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == SimpleQueryLanguageParserQuestion {
					{
						p.SetState(83)
						p.Match(SimpleQueryLanguageParserQuestion)
					}

				}
				{
					p.SetState(86)
					p.Match(SimpleQueryLanguageParserLeftBracket)
				}
				{
					p.SetState(87)
					p.Index()
				}
				{
					p.SetState(88)
					p.Match(SimpleQueryLanguageParserRightBracket)
				}

			case 15:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SimpleQueryLanguageParserRULE_field)
				p.SetState(90)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
				}
				{
					p.SetState(91)
					p.Match(SimpleQueryLanguageParserLeftParen)
				}
				p.SetState(100)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2122486784) != 0 {
					{
						p.SetState(92)
						p.Function_param()
					}
					p.SetState(97)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)

					for _la == SimpleQueryLanguageParserComma {
						{
							p.SetState(93)
							p.Match(SimpleQueryLanguageParserComma)
						}
						{
							p.SetState(94)
							p.Function_param()
						}

						p.SetState(99)
						p.GetErrorHandler().Sync(p)
						_la = p.GetTokenStream().LA(1)
					}

				}
				{
					p.SetState(102)
					p.Match(SimpleQueryLanguageParserRightParen)
				}

			case 16:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SimpleQueryLanguageParserRULE_field)
				p.SetState(103)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(104)
					p.Match(SimpleQueryLanguageParserQuestion)
				}
				{
					p.SetState(105)
					p.field(0)
				}
				p.SetState(108)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
					{
						p.SetState(106)
						p.Match(SimpleQueryLanguageParserT__0)
					}
					{
						p.SetState(107)
						p.field(0)
					}

				}

			}

		}
		p.SetState(114)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}

	return localctx
}

// IArrayContext is an interface to support dynamic dispatch.
type IArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayContext differentiates from other interfaces.
	IsArrayContext()
}

type ArrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayContext() *ArrayContext {
	var p = new(ArrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleQueryLanguageParserRULE_array
	return p
}

func (*ArrayContext) IsArrayContext() {}

func NewArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayContext {
	var p = new(ArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleQueryLanguageParserRULE_array

	return p
}

func (s *ArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayContext) LeftBracket() antlr.TerminalNode {
	return s.GetToken(SimpleQueryLanguageParserLeftBracket, 0)
}

func (s *ArrayContext) RightBracket() antlr.TerminalNode {
	return s.GetToken(SimpleQueryLanguageParserRightBracket, 0)
}

func (s *ArrayContext) AllField() []IFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldContext); ok {
			len++
		}
	}

	tst := make([]IFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldContext); ok {
			tst[i] = t.(IFieldContext)
			i++
		}
	}

	return tst
}

func (s *ArrayContext) Field(i int) IFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
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

	return t.(IFieldContext)
}

func (s *ArrayContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(SimpleQueryLanguageParserComma)
}

func (s *ArrayContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(SimpleQueryLanguageParserComma, i)
}

func (s *ArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleQueryLanguageListener); ok {
		listenerT.EnterArray(s)
	}
}

func (s *ArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleQueryLanguageListener); ok {
		listenerT.ExitArray(s)
	}
}

func (s *ArrayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleQueryLanguageVisitor:
		return t.VisitArray(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleQueryLanguageParser) Array() (localctx IArrayContext) {
	this := p
	_ = this

	localctx = NewArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SimpleQueryLanguageParserRULE_array)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)
		p.Match(SimpleQueryLanguageParserLeftBracket)
	}
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2122486784) != 0 {
		{
			p.SetState(116)
			p.field(0)
		}
		p.SetState(121)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SimpleQueryLanguageParserComma {
			{
				p.SetState(117)
				p.Match(SimpleQueryLanguageParserComma)
			}
			{
				p.SetState(118)
				p.field(0)
			}

			p.SetState(123)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(126)
		p.Match(SimpleQueryLanguageParserRightBracket)
	}

	return localctx
}

// IObjectContext is an interface to support dynamic dispatch.
type IObjectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjectContext differentiates from other interfaces.
	IsObjectContext()
}

type ObjectContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectContext() *ObjectContext {
	var p = new(ObjectContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleQueryLanguageParserRULE_object
	return p
}

func (*ObjectContext) IsObjectContext() {}

func NewObjectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectContext {
	var p = new(ObjectContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleQueryLanguageParserRULE_object

	return p
}

func (s *ObjectContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectContext) LeftBrace() antlr.TerminalNode {
	return s.GetToken(SimpleQueryLanguageParserLeftBrace, 0)
}

func (s *ObjectContext) RightBrace() antlr.TerminalNode {
	return s.GetToken(SimpleQueryLanguageParserRightBrace, 0)
}

func (s *ObjectContext) AllObject_field() []IObject_fieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IObject_fieldContext); ok {
			len++
		}
	}

	tst := make([]IObject_fieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IObject_fieldContext); ok {
			tst[i] = t.(IObject_fieldContext)
			i++
		}
	}

	return tst
}

func (s *ObjectContext) Object_field(i int) IObject_fieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObject_fieldContext); ok {
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

	return t.(IObject_fieldContext)
}

func (s *ObjectContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(SimpleQueryLanguageParserComma)
}

func (s *ObjectContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(SimpleQueryLanguageParserComma, i)
}

func (s *ObjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleQueryLanguageListener); ok {
		listenerT.EnterObject(s)
	}
}

func (s *ObjectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleQueryLanguageListener); ok {
		listenerT.ExitObject(s)
	}
}

func (s *ObjectContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleQueryLanguageVisitor:
		return t.VisitObject(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleQueryLanguageParser) Object() (localctx IObjectContext) {
	this := p
	_ = this

	localctx = NewObjectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SimpleQueryLanguageParserRULE_object)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(128)
		p.Match(SimpleQueryLanguageParserLeftBrace)
	}
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SimpleQueryLanguageParserESCAPED_STRING || _la == SimpleQueryLanguageParserSTRING {
		{
			p.SetState(129)
			p.Object_field()
		}
		p.SetState(134)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SimpleQueryLanguageParserComma {
			{
				p.SetState(130)
				p.Match(SimpleQueryLanguageParserComma)
			}
			{
				p.SetState(131)
				p.Object_field()
			}

			p.SetState(136)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(139)
		p.Match(SimpleQueryLanguageParserRightBrace)
	}

	return localctx
}

// IObject_fieldContext is an interface to support dynamic dispatch.
type IObject_fieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObject_fieldContext differentiates from other interfaces.
	IsObject_fieldContext()
}

type Object_fieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObject_fieldContext() *Object_fieldContext {
	var p = new(Object_fieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleQueryLanguageParserRULE_object_field
	return p
}

func (*Object_fieldContext) IsObject_fieldContext() {}

func NewObject_fieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Object_fieldContext {
	var p = new(Object_fieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleQueryLanguageParserRULE_object_field

	return p
}

func (s *Object_fieldContext) GetParser() antlr.Parser { return s.parser }

func (s *Object_fieldContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *Object_fieldContext) Field() IFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *Object_fieldContext) ESCAPED_STRING() antlr.TerminalNode {
	return s.GetToken(SimpleQueryLanguageParserESCAPED_STRING, 0)
}

func (s *Object_fieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Object_fieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Object_fieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleQueryLanguageListener); ok {
		listenerT.EnterObject_field(s)
	}
}

func (s *Object_fieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleQueryLanguageListener); ok {
		listenerT.ExitObject_field(s)
	}
}

func (s *Object_fieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleQueryLanguageVisitor:
		return t.VisitObject_field(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleQueryLanguageParser) Object_field() (localctx IObject_fieldContext) {
	this := p
	_ = this

	localctx = NewObject_fieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SimpleQueryLanguageParserRULE_object_field)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(148)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SimpleQueryLanguageParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(141)
			p.Name()
		}
		{
			p.SetState(142)
			p.Match(SimpleQueryLanguageParserT__0)
		}
		{
			p.SetState(143)
			p.field(0)
		}

	case SimpleQueryLanguageParserESCAPED_STRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(145)
			p.Match(SimpleQueryLanguageParserESCAPED_STRING)
		}
		{
			p.SetState(146)
			p.Match(SimpleQueryLanguageParserT__0)
		}
		{
			p.SetState(147)
			p.field(0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INameContext is an interface to support dynamic dispatch.
type INameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNameContext differentiates from other interfaces.
	IsNameContext()
}

type NameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameContext() *NameContext {
	var p = new(NameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleQueryLanguageParserRULE_name
	return p
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleQueryLanguageParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) STRING() antlr.TerminalNode {
	return s.GetToken(SimpleQueryLanguageParserSTRING, 0)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleQueryLanguageListener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleQueryLanguageListener); ok {
		listenerT.ExitName(s)
	}
}

func (s *NameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleQueryLanguageVisitor:
		return t.VisitName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleQueryLanguageParser) Name() (localctx INameContext) {
	this := p
	_ = this

	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SimpleQueryLanguageParserRULE_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(150)
		p.Match(SimpleQueryLanguageParserSTRING)
	}

	return localctx
}

// IIndexContext is an interface to support dynamic dispatch.
type IIndexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIndexContext differentiates from other interfaces.
	IsIndexContext()
}

type IndexContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexContext() *IndexContext {
	var p = new(IndexContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleQueryLanguageParserRULE_index
	return p
}

func (*IndexContext) IsIndexContext() {}

func NewIndexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexContext {
	var p = new(IndexContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleQueryLanguageParserRULE_index

	return p
}

func (s *IndexContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexContext) Field() IFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *IndexContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SimpleQueryLanguageParserNUMBER, 0)
}

func (s *IndexContext) ESCAPED_STRING() antlr.TerminalNode {
	return s.GetToken(SimpleQueryLanguageParserESCAPED_STRING, 0)
}

func (s *IndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleQueryLanguageListener); ok {
		listenerT.EnterIndex(s)
	}
}

func (s *IndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleQueryLanguageListener); ok {
		listenerT.ExitIndex(s)
	}
}

func (s *IndexContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleQueryLanguageVisitor:
		return t.VisitIndex(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleQueryLanguageParser) Index() (localctx IIndexContext) {
	this := p
	_ = this

	localctx = NewIndexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SimpleQueryLanguageParserRULE_index)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(152)
			p.field(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(153)
			p.Match(SimpleQueryLanguageParserNUMBER)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(154)
			p.Match(SimpleQueryLanguageParserESCAPED_STRING)
		}

	}

	return localctx
}

func (p *SimpleQueryLanguageParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *FieldContext = nil
		if localctx != nil {
			t = localctx.(*FieldContext)
		}
		return p.Field_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SimpleQueryLanguageParser) Field_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 13:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 14:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 15:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
