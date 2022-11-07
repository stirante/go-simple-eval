// Code generated from ../grammar/SimpleQueryLanguage.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // SimpleQueryLanguage

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
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
		"", "'.'", "':'", "'<'", "'<='", "'=='", "'>'", "'>='", "'!='", "'&&'",
		"'||'", "'!'", "'+'", "'-'", "'*'", "'/'", "'('", "')'", "'['", "']'",
		"'?'", "'??'", "'..'", "','", "'{'", "'}'", "'null'", "'false'", "'true'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "Less", "LessOrEqual", "Equal", "Greater", "GreaterOrEqual",
		"NotEqual", "And", "Or", "Not", "Add", "Subtract", "Multiply", "Divide",
		"LeftParen", "RightParen", "LeftBracket", "RightBracket", "Question",
		"NullCoalescing", "Range", "Comma", "LeftBrace", "RightBrace", "Null",
		"False", "True", "ESCAPED_STRING", "STRING", "NUMBER", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "function_param", "field", "array", "object", "object_field",
		"name", "index",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 32, 157, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 3, 2, 38, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 78, 8, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 3, 2, 84, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 5, 2, 95, 8, 2, 10, 2, 12, 2, 98, 9, 2, 3, 2, 100, 8, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 108, 8, 2, 5, 2, 110, 8, 2, 10,
		2, 12, 2, 113, 9, 2, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 119, 8, 3, 10, 3, 12,
		3, 122, 9, 3, 3, 3, 124, 8, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4,
		132, 8, 4, 10, 4, 12, 4, 135, 9, 4, 3, 4, 137, 8, 4, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 148, 8, 5, 1, 6, 1, 6, 1, 7,
		1, 7, 1, 7, 3, 7, 155, 8, 7, 1, 7, 0, 1, 4, 8, 0, 2, 4, 6, 8, 10, 12, 14,
		0, 2, 1, 0, 14, 15, 1, 0, 12, 13, 186, 0, 16, 1, 0, 0, 0, 2, 18, 1, 0,
		0, 0, 4, 37, 1, 0, 0, 0, 6, 114, 1, 0, 0, 0, 8, 127, 1, 0, 0, 0, 10, 147,
		1, 0, 0, 0, 12, 149, 1, 0, 0, 0, 14, 154, 1, 0, 0, 0, 16, 17, 3, 4, 2,
		0, 17, 1, 1, 0, 0, 0, 18, 19, 3, 4, 2, 0, 19, 3, 1, 0, 0, 0, 20, 21, 6,
		2, -1, 0, 21, 22, 5, 16, 0, 0, 22, 23, 3, 4, 2, 0, 23, 24, 5, 17, 0, 0,
		24, 38, 1, 0, 0, 0, 25, 38, 5, 28, 0, 0, 26, 38, 5, 27, 0, 0, 27, 38, 5,
		26, 0, 0, 28, 38, 5, 31, 0, 0, 29, 38, 5, 29, 0, 0, 30, 38, 3, 6, 3, 0,
		31, 38, 3, 8, 4, 0, 32, 38, 3, 12, 6, 0, 33, 34, 5, 13, 0, 0, 34, 38, 3,
		4, 2, 15, 35, 36, 5, 11, 0, 0, 36, 38, 3, 4, 2, 2, 37, 20, 1, 0, 0, 0,
		37, 25, 1, 0, 0, 0, 37, 26, 1, 0, 0, 0, 37, 27, 1, 0, 0, 0, 37, 28, 1,
		0, 0, 0, 37, 29, 1, 0, 0, 0, 37, 30, 1, 0, 0, 0, 37, 31, 1, 0, 0, 0, 37,
		32, 1, 0, 0, 0, 37, 33, 1, 0, 0, 0, 37, 35, 1, 0, 0, 0, 38, 111, 1, 0,
		0, 0, 39, 40, 10, 14, 0, 0, 40, 41, 7, 0, 0, 0, 41, 110, 3, 4, 2, 15, 42,
		43, 10, 13, 0, 0, 43, 44, 7, 1, 0, 0, 44, 110, 3, 4, 2, 14, 45, 46, 10,
		12, 0, 0, 46, 47, 5, 22, 0, 0, 47, 110, 3, 4, 2, 13, 48, 49, 10, 11, 0,
		0, 49, 50, 5, 21, 0, 0, 50, 110, 3, 4, 2, 12, 51, 52, 10, 10, 0, 0, 52,
		53, 5, 5, 0, 0, 53, 110, 3, 4, 2, 11, 54, 55, 10, 9, 0, 0, 55, 56, 5, 3,
		0, 0, 56, 110, 3, 4, 2, 10, 57, 58, 10, 8, 0, 0, 58, 59, 5, 4, 0, 0, 59,
		110, 3, 4, 2, 9, 60, 61, 10, 7, 0, 0, 61, 62, 5, 6, 0, 0, 62, 110, 3, 4,
		2, 8, 63, 64, 10, 6, 0, 0, 64, 65, 5, 7, 0, 0, 65, 110, 3, 4, 2, 7, 66,
		67, 10, 5, 0, 0, 67, 68, 5, 8, 0, 0, 68, 110, 3, 4, 2, 6, 69, 70, 10, 4,
		0, 0, 70, 71, 5, 9, 0, 0, 71, 110, 3, 4, 2, 5, 72, 73, 10, 3, 0, 0, 73,
		74, 5, 10, 0, 0, 74, 110, 3, 4, 2, 4, 75, 77, 10, 18, 0, 0, 76, 78, 5,
		20, 0, 0, 77, 76, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79,
		80, 5, 1, 0, 0, 80, 110, 3, 12, 6, 0, 81, 83, 10, 17, 0, 0, 82, 84, 5,
		20, 0, 0, 83, 82, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85,
		86, 5, 18, 0, 0, 86, 87, 3, 14, 7, 0, 87, 88, 5, 19, 0, 0, 88, 110, 1,
		0, 0, 0, 89, 90, 10, 16, 0, 0, 90, 99, 5, 16, 0, 0, 91, 96, 3, 2, 1, 0,
		92, 93, 5, 23, 0, 0, 93, 95, 3, 2, 1, 0, 94, 92, 1, 0, 0, 0, 95, 98, 1,
		0, 0, 0, 96, 94, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 100, 1, 0, 0, 0, 98,
		96, 1, 0, 0, 0, 99, 91, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 101, 1, 0,
		0, 0, 101, 110, 5, 17, 0, 0, 102, 103, 10, 1, 0, 0, 103, 104, 5, 20, 0,
		0, 104, 107, 3, 4, 2, 0, 105, 106, 5, 2, 0, 0, 106, 108, 3, 4, 2, 0, 107,
		105, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 110, 1, 0, 0, 0, 109, 39, 1,
		0, 0, 0, 109, 42, 1, 0, 0, 0, 109, 45, 1, 0, 0, 0, 109, 48, 1, 0, 0, 0,
		109, 51, 1, 0, 0, 0, 109, 54, 1, 0, 0, 0, 109, 57, 1, 0, 0, 0, 109, 60,
		1, 0, 0, 0, 109, 63, 1, 0, 0, 0, 109, 66, 1, 0, 0, 0, 109, 69, 1, 0, 0,
		0, 109, 72, 1, 0, 0, 0, 109, 75, 1, 0, 0, 0, 109, 81, 1, 0, 0, 0, 109,
		89, 1, 0, 0, 0, 109, 102, 1, 0, 0, 0, 110, 113, 1, 0, 0, 0, 111, 109, 1,
		0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 5, 1, 0, 0, 0, 113, 111, 1, 0, 0, 0,
		114, 123, 5, 18, 0, 0, 115, 120, 3, 4, 2, 0, 116, 117, 5, 23, 0, 0, 117,
		119, 3, 4, 2, 0, 118, 116, 1, 0, 0, 0, 119, 122, 1, 0, 0, 0, 120, 118,
		1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121, 124, 1, 0, 0, 0, 122, 120, 1, 0,
		0, 0, 123, 115, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0,
		125, 126, 5, 19, 0, 0, 126, 7, 1, 0, 0, 0, 127, 136, 5, 24, 0, 0, 128,
		133, 3, 10, 5, 0, 129, 130, 5, 23, 0, 0, 130, 132, 3, 10, 5, 0, 131, 129,
		1, 0, 0, 0, 132, 135, 1, 0, 0, 0, 133, 131, 1, 0, 0, 0, 133, 134, 1, 0,
		0, 0, 134, 137, 1, 0, 0, 0, 135, 133, 1, 0, 0, 0, 136, 128, 1, 0, 0, 0,
		136, 137, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 139, 5, 25, 0, 0, 139,
		9, 1, 0, 0, 0, 140, 141, 3, 12, 6, 0, 141, 142, 5, 2, 0, 0, 142, 143, 3,
		4, 2, 0, 143, 148, 1, 0, 0, 0, 144, 145, 5, 29, 0, 0, 145, 146, 5, 2, 0,
		0, 146, 148, 3, 4, 2, 0, 147, 140, 1, 0, 0, 0, 147, 144, 1, 0, 0, 0, 148,
		11, 1, 0, 0, 0, 149, 150, 5, 30, 0, 0, 150, 13, 1, 0, 0, 0, 151, 155, 3,
		4, 2, 0, 152, 155, 5, 31, 0, 0, 153, 155, 5, 29, 0, 0, 154, 151, 1, 0,
		0, 0, 154, 152, 1, 0, 0, 0, 154, 153, 1, 0, 0, 0, 155, 15, 1, 0, 0, 0,
		14, 37, 77, 83, 96, 99, 107, 109, 111, 120, 123, 133, 136, 147, 154,
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
	this.GrammarFileName = "SimpleQueryLanguage.g4"

	return this
}

// SimpleQueryLanguageParser tokens.
const (
	SimpleQueryLanguageParserEOF            = antlr.TokenEOF
	SimpleQueryLanguageParserT__0           = 1
	SimpleQueryLanguageParserT__1           = 2
	SimpleQueryLanguageParserLess           = 3
	SimpleQueryLanguageParserLessOrEqual    = 4
	SimpleQueryLanguageParserEqual          = 5
	SimpleQueryLanguageParserGreater        = 6
	SimpleQueryLanguageParserGreaterOrEqual = 7
	SimpleQueryLanguageParserNotEqual       = 8
	SimpleQueryLanguageParserAnd            = 9
	SimpleQueryLanguageParserOr             = 10
	SimpleQueryLanguageParserNot            = 11
	SimpleQueryLanguageParserAdd            = 12
	SimpleQueryLanguageParserSubtract       = 13
	SimpleQueryLanguageParserMultiply       = 14
	SimpleQueryLanguageParserDivide         = 15
	SimpleQueryLanguageParserLeftParen      = 16
	SimpleQueryLanguageParserRightParen     = 17
	SimpleQueryLanguageParserLeftBracket    = 18
	SimpleQueryLanguageParserRightBracket   = 19
	SimpleQueryLanguageParserQuestion       = 20
	SimpleQueryLanguageParserNullCoalescing = 21
	SimpleQueryLanguageParserRange          = 22
	SimpleQueryLanguageParserComma          = 23
	SimpleQueryLanguageParserLeftBrace      = 24
	SimpleQueryLanguageParserRightBrace     = 25
	SimpleQueryLanguageParserNull           = 26
	SimpleQueryLanguageParserFalse          = 27
	SimpleQueryLanguageParserTrue           = 28
	SimpleQueryLanguageParserESCAPED_STRING = 29
	SimpleQueryLanguageParserSTRING         = 30
	SimpleQueryLanguageParserNUMBER         = 31
	SimpleQueryLanguageParserWS             = 32
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

func (s *FieldContext) Range() antlr.TerminalNode {
	return s.GetToken(SimpleQueryLanguageParserRange, 0)
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
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(109)
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
					p.Match(SimpleQueryLanguageParserRange)
				}
				{
					p.SetState(47)
					p.field(13)
				}

			case 4:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SimpleQueryLanguageParserRULE_field)
				p.SetState(48)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(49)
					p.Match(SimpleQueryLanguageParserNullCoalescing)
				}
				{
					p.SetState(50)
					p.field(12)
				}

			case 5:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SimpleQueryLanguageParserRULE_field)
				p.SetState(51)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(52)
					p.Match(SimpleQueryLanguageParserEqual)
				}
				{
					p.SetState(53)
					p.field(11)
				}

			case 6:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SimpleQueryLanguageParserRULE_field)
				p.SetState(54)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(55)
					p.Match(SimpleQueryLanguageParserLess)
				}
				{
					p.SetState(56)
					p.field(10)
				}

			case 7:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SimpleQueryLanguageParserRULE_field)
				p.SetState(57)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(58)
					p.Match(SimpleQueryLanguageParserLessOrEqual)
				}
				{
					p.SetState(59)
					p.field(9)
				}

			case 8:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SimpleQueryLanguageParserRULE_field)
				p.SetState(60)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(61)
					p.Match(SimpleQueryLanguageParserGreater)
				}
				{
					p.SetState(62)
					p.field(8)
				}

			case 9:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SimpleQueryLanguageParserRULE_field)
				p.SetState(63)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(64)
					p.Match(SimpleQueryLanguageParserGreaterOrEqual)
				}
				{
					p.SetState(65)
					p.field(7)
				}

			case 10:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SimpleQueryLanguageParserRULE_field)
				p.SetState(66)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(67)
					p.Match(SimpleQueryLanguageParserNotEqual)
				}
				{
					p.SetState(68)
					p.field(6)
				}

			case 11:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SimpleQueryLanguageParserRULE_field)
				p.SetState(69)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(70)
					p.Match(SimpleQueryLanguageParserAnd)
				}
				{
					p.SetState(71)
					p.field(5)
				}

			case 12:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SimpleQueryLanguageParserRULE_field)
				p.SetState(72)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(73)
					p.Match(SimpleQueryLanguageParserOr)
				}
				{
					p.SetState(74)
					p.field(4)
				}

			case 13:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SimpleQueryLanguageParserRULE_field)
				p.SetState(75)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
				}

				p.SetState(77)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == SimpleQueryLanguageParserQuestion {
					{
						p.SetState(76)
						p.Match(SimpleQueryLanguageParserQuestion)
					}

				}
				{
					p.SetState(79)
					p.Match(SimpleQueryLanguageParserT__0)
				}
				{
					p.SetState(80)
					p.Name()
				}

			case 14:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SimpleQueryLanguageParserRULE_field)
				p.SetState(81)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
				}

				p.SetState(83)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == SimpleQueryLanguageParserQuestion {
					{
						p.SetState(82)
						p.Match(SimpleQueryLanguageParserQuestion)
					}

				}
				{
					p.SetState(85)
					p.Match(SimpleQueryLanguageParserLeftBracket)
				}
				{
					p.SetState(86)
					p.Index()
				}
				{
					p.SetState(87)
					p.Match(SimpleQueryLanguageParserRightBracket)
				}

			case 15:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SimpleQueryLanguageParserRULE_field)
				p.SetState(89)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
				}
				{
					p.SetState(90)
					p.Match(SimpleQueryLanguageParserLeftParen)
				}
				p.SetState(99)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SimpleQueryLanguageParserNot)|(1<<SimpleQueryLanguageParserSubtract)|(1<<SimpleQueryLanguageParserLeftParen)|(1<<SimpleQueryLanguageParserLeftBracket)|(1<<SimpleQueryLanguageParserLeftBrace)|(1<<SimpleQueryLanguageParserNull)|(1<<SimpleQueryLanguageParserFalse)|(1<<SimpleQueryLanguageParserTrue)|(1<<SimpleQueryLanguageParserESCAPED_STRING)|(1<<SimpleQueryLanguageParserSTRING)|(1<<SimpleQueryLanguageParserNUMBER))) != 0 {
					{
						p.SetState(91)
						p.Function_param()
					}
					p.SetState(96)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)

					for _la == SimpleQueryLanguageParserComma {
						{
							p.SetState(92)
							p.Match(SimpleQueryLanguageParserComma)
						}
						{
							p.SetState(93)
							p.Function_param()
						}

						p.SetState(98)
						p.GetErrorHandler().Sync(p)
						_la = p.GetTokenStream().LA(1)
					}

				}
				{
					p.SetState(101)
					p.Match(SimpleQueryLanguageParserRightParen)
				}

			case 16:
				localctx = NewFieldContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SimpleQueryLanguageParserRULE_field)
				p.SetState(102)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(103)
					p.Match(SimpleQueryLanguageParserQuestion)
				}
				{
					p.SetState(104)
					p.field(0)
				}
				p.SetState(107)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
					{
						p.SetState(105)
						p.Match(SimpleQueryLanguageParserT__1)
					}
					{
						p.SetState(106)
						p.field(0)
					}

				}

			}

		}
		p.SetState(113)
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
		p.SetState(114)
		p.Match(SimpleQueryLanguageParserLeftBracket)
	}
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SimpleQueryLanguageParserNot)|(1<<SimpleQueryLanguageParserSubtract)|(1<<SimpleQueryLanguageParserLeftParen)|(1<<SimpleQueryLanguageParserLeftBracket)|(1<<SimpleQueryLanguageParserLeftBrace)|(1<<SimpleQueryLanguageParserNull)|(1<<SimpleQueryLanguageParserFalse)|(1<<SimpleQueryLanguageParserTrue)|(1<<SimpleQueryLanguageParserESCAPED_STRING)|(1<<SimpleQueryLanguageParserSTRING)|(1<<SimpleQueryLanguageParserNUMBER))) != 0 {
		{
			p.SetState(115)
			p.field(0)
		}
		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SimpleQueryLanguageParserComma {
			{
				p.SetState(116)
				p.Match(SimpleQueryLanguageParserComma)
			}
			{
				p.SetState(117)
				p.field(0)
			}

			p.SetState(122)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(125)
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
		p.SetState(127)
		p.Match(SimpleQueryLanguageParserLeftBrace)
	}
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SimpleQueryLanguageParserESCAPED_STRING || _la == SimpleQueryLanguageParserSTRING {
		{
			p.SetState(128)
			p.Object_field()
		}
		p.SetState(133)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SimpleQueryLanguageParserComma {
			{
				p.SetState(129)
				p.Match(SimpleQueryLanguageParserComma)
			}
			{
				p.SetState(130)
				p.Object_field()
			}

			p.SetState(135)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(138)
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

	p.SetState(147)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SimpleQueryLanguageParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(140)
			p.Name()
		}
		{
			p.SetState(141)
			p.Match(SimpleQueryLanguageParserT__1)
		}
		{
			p.SetState(142)
			p.field(0)
		}

	case SimpleQueryLanguageParserESCAPED_STRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(144)
			p.Match(SimpleQueryLanguageParserESCAPED_STRING)
		}
		{
			p.SetState(145)
			p.Match(SimpleQueryLanguageParserT__1)
		}
		{
			p.SetState(146)
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
		p.SetState(149)
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

	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(151)
			p.field(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(152)
			p.Match(SimpleQueryLanguageParserNUMBER)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(153)
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
