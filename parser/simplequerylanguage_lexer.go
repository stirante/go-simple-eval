// Code generated from ../grammar/SimpleQueryLanguage.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type SimpleQueryLanguageLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var simplequerylanguagelexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func simplequerylanguagelexerLexerInit() {
	staticData := &simplequerylanguagelexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
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
		"T__0", "T__1", "Less", "LessOrEqual", "Equal", "Greater", "GreaterOrEqual",
		"NotEqual", "And", "Or", "Not", "Add", "Subtract", "Multiply", "Divide",
		"LeftParen", "RightParen", "LeftBracket", "RightBracket", "Question",
		"NullCoalescing", "Range", "Comma", "LeftBrace", "RightBrace", "Null",
		"False", "True", "ESCAPED_STRING", "STRING", "NUMBER", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 32, 185, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4,
		1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8,
		1, 8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13,
		1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1,
		18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22,
		1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28,
		1, 28, 1, 28, 1, 28, 5, 28, 144, 8, 28, 10, 28, 12, 28, 147, 9, 28, 1,
		28, 1, 28, 1, 28, 1, 28, 1, 28, 5, 28, 154, 8, 28, 10, 28, 12, 28, 157,
		9, 28, 1, 28, 3, 28, 160, 8, 28, 1, 29, 1, 29, 5, 29, 164, 8, 29, 10, 29,
		12, 29, 167, 9, 29, 1, 30, 4, 30, 170, 8, 30, 11, 30, 12, 30, 171, 1, 30,
		1, 30, 4, 30, 176, 8, 30, 11, 30, 12, 30, 177, 3, 30, 180, 8, 30, 1, 31,
		1, 31, 1, 31, 1, 31, 0, 0, 32, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13,
		7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16,
		33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25,
		51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 1, 0, 6, 2, 0,
		34, 34, 92, 92, 2, 0, 39, 39, 92, 92, 4, 0, 36, 36, 65, 90, 95, 95, 97,
		122, 5, 0, 36, 36, 48, 57, 65, 90, 95, 95, 97, 122, 1, 0, 48, 57, 3, 0,
		9, 10, 13, 13, 32, 32, 193, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1,
		0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13,
		1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0,
		21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0,
		0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0,
		0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0,
		0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1,
		0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59,
		1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 1, 65, 1, 0, 0, 0, 3,
		67, 1, 0, 0, 0, 5, 69, 1, 0, 0, 0, 7, 71, 1, 0, 0, 0, 9, 74, 1, 0, 0, 0,
		11, 77, 1, 0, 0, 0, 13, 79, 1, 0, 0, 0, 15, 82, 1, 0, 0, 0, 17, 85, 1,
		0, 0, 0, 19, 88, 1, 0, 0, 0, 21, 91, 1, 0, 0, 0, 23, 93, 1, 0, 0, 0, 25,
		95, 1, 0, 0, 0, 27, 97, 1, 0, 0, 0, 29, 99, 1, 0, 0, 0, 31, 101, 1, 0,
		0, 0, 33, 103, 1, 0, 0, 0, 35, 105, 1, 0, 0, 0, 37, 107, 1, 0, 0, 0, 39,
		109, 1, 0, 0, 0, 41, 111, 1, 0, 0, 0, 43, 114, 1, 0, 0, 0, 45, 117, 1,
		0, 0, 0, 47, 119, 1, 0, 0, 0, 49, 121, 1, 0, 0, 0, 51, 123, 1, 0, 0, 0,
		53, 128, 1, 0, 0, 0, 55, 134, 1, 0, 0, 0, 57, 159, 1, 0, 0, 0, 59, 161,
		1, 0, 0, 0, 61, 169, 1, 0, 0, 0, 63, 181, 1, 0, 0, 0, 65, 66, 5, 46, 0,
		0, 66, 2, 1, 0, 0, 0, 67, 68, 5, 58, 0, 0, 68, 4, 1, 0, 0, 0, 69, 70, 5,
		60, 0, 0, 70, 6, 1, 0, 0, 0, 71, 72, 5, 60, 0, 0, 72, 73, 5, 61, 0, 0,
		73, 8, 1, 0, 0, 0, 74, 75, 5, 61, 0, 0, 75, 76, 5, 61, 0, 0, 76, 10, 1,
		0, 0, 0, 77, 78, 5, 62, 0, 0, 78, 12, 1, 0, 0, 0, 79, 80, 5, 62, 0, 0,
		80, 81, 5, 61, 0, 0, 81, 14, 1, 0, 0, 0, 82, 83, 5, 33, 0, 0, 83, 84, 5,
		61, 0, 0, 84, 16, 1, 0, 0, 0, 85, 86, 5, 38, 0, 0, 86, 87, 5, 38, 0, 0,
		87, 18, 1, 0, 0, 0, 88, 89, 5, 124, 0, 0, 89, 90, 5, 124, 0, 0, 90, 20,
		1, 0, 0, 0, 91, 92, 5, 33, 0, 0, 92, 22, 1, 0, 0, 0, 93, 94, 5, 43, 0,
		0, 94, 24, 1, 0, 0, 0, 95, 96, 5, 45, 0, 0, 96, 26, 1, 0, 0, 0, 97, 98,
		5, 42, 0, 0, 98, 28, 1, 0, 0, 0, 99, 100, 5, 47, 0, 0, 100, 30, 1, 0, 0,
		0, 101, 102, 5, 40, 0, 0, 102, 32, 1, 0, 0, 0, 103, 104, 5, 41, 0, 0, 104,
		34, 1, 0, 0, 0, 105, 106, 5, 91, 0, 0, 106, 36, 1, 0, 0, 0, 107, 108, 5,
		93, 0, 0, 108, 38, 1, 0, 0, 0, 109, 110, 5, 63, 0, 0, 110, 40, 1, 0, 0,
		0, 111, 112, 5, 63, 0, 0, 112, 113, 5, 63, 0, 0, 113, 42, 1, 0, 0, 0, 114,
		115, 5, 46, 0, 0, 115, 116, 5, 46, 0, 0, 116, 44, 1, 0, 0, 0, 117, 118,
		5, 44, 0, 0, 118, 46, 1, 0, 0, 0, 119, 120, 5, 123, 0, 0, 120, 48, 1, 0,
		0, 0, 121, 122, 5, 125, 0, 0, 122, 50, 1, 0, 0, 0, 123, 124, 5, 110, 0,
		0, 124, 125, 5, 117, 0, 0, 125, 126, 5, 108, 0, 0, 126, 127, 5, 108, 0,
		0, 127, 52, 1, 0, 0, 0, 128, 129, 5, 102, 0, 0, 129, 130, 5, 97, 0, 0,
		130, 131, 5, 108, 0, 0, 131, 132, 5, 115, 0, 0, 132, 133, 5, 101, 0, 0,
		133, 54, 1, 0, 0, 0, 134, 135, 5, 116, 0, 0, 135, 136, 5, 114, 0, 0, 136,
		137, 5, 117, 0, 0, 137, 138, 5, 101, 0, 0, 138, 56, 1, 0, 0, 0, 139, 145,
		5, 34, 0, 0, 140, 141, 5, 92, 0, 0, 141, 144, 9, 0, 0, 0, 142, 144, 8,
		0, 0, 0, 143, 140, 1, 0, 0, 0, 143, 142, 1, 0, 0, 0, 144, 147, 1, 0, 0,
		0, 145, 143, 1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146, 148, 1, 0, 0, 0, 147,
		145, 1, 0, 0, 0, 148, 160, 5, 34, 0, 0, 149, 155, 5, 39, 0, 0, 150, 151,
		5, 92, 0, 0, 151, 154, 9, 0, 0, 0, 152, 154, 8, 1, 0, 0, 153, 150, 1, 0,
		0, 0, 153, 152, 1, 0, 0, 0, 154, 157, 1, 0, 0, 0, 155, 153, 1, 0, 0, 0,
		155, 156, 1, 0, 0, 0, 156, 158, 1, 0, 0, 0, 157, 155, 1, 0, 0, 0, 158,
		160, 5, 39, 0, 0, 159, 139, 1, 0, 0, 0, 159, 149, 1, 0, 0, 0, 160, 58,
		1, 0, 0, 0, 161, 165, 7, 2, 0, 0, 162, 164, 7, 3, 0, 0, 163, 162, 1, 0,
		0, 0, 164, 167, 1, 0, 0, 0, 165, 163, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0,
		166, 60, 1, 0, 0, 0, 167, 165, 1, 0, 0, 0, 168, 170, 7, 4, 0, 0, 169, 168,
		1, 0, 0, 0, 170, 171, 1, 0, 0, 0, 171, 169, 1, 0, 0, 0, 171, 172, 1, 0,
		0, 0, 172, 179, 1, 0, 0, 0, 173, 175, 5, 46, 0, 0, 174, 176, 7, 4, 0, 0,
		175, 174, 1, 0, 0, 0, 176, 177, 1, 0, 0, 0, 177, 175, 1, 0, 0, 0, 177,
		178, 1, 0, 0, 0, 178, 180, 1, 0, 0, 0, 179, 173, 1, 0, 0, 0, 179, 180,
		1, 0, 0, 0, 180, 62, 1, 0, 0, 0, 181, 182, 7, 5, 0, 0, 182, 183, 1, 0,
		0, 0, 183, 184, 6, 31, 0, 0, 184, 64, 1, 0, 0, 0, 10, 0, 143, 145, 153,
		155, 159, 165, 171, 177, 179, 1, 0, 1, 0,
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

// SimpleQueryLanguageLexerInit initializes any static state used to implement SimpleQueryLanguageLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewSimpleQueryLanguageLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func SimpleQueryLanguageLexerInit() {
	staticData := &simplequerylanguagelexerLexerStaticData
	staticData.once.Do(simplequerylanguagelexerLexerInit)
}

// NewSimpleQueryLanguageLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewSimpleQueryLanguageLexer(input antlr.CharStream) *SimpleQueryLanguageLexer {
	SimpleQueryLanguageLexerInit()
	l := new(SimpleQueryLanguageLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &simplequerylanguagelexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "SimpleQueryLanguage.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SimpleQueryLanguageLexer tokens.
const (
	SimpleQueryLanguageLexerT__0           = 1
	SimpleQueryLanguageLexerT__1           = 2
	SimpleQueryLanguageLexerLess           = 3
	SimpleQueryLanguageLexerLessOrEqual    = 4
	SimpleQueryLanguageLexerEqual          = 5
	SimpleQueryLanguageLexerGreater        = 6
	SimpleQueryLanguageLexerGreaterOrEqual = 7
	SimpleQueryLanguageLexerNotEqual       = 8
	SimpleQueryLanguageLexerAnd            = 9
	SimpleQueryLanguageLexerOr             = 10
	SimpleQueryLanguageLexerNot            = 11
	SimpleQueryLanguageLexerAdd            = 12
	SimpleQueryLanguageLexerSubtract       = 13
	SimpleQueryLanguageLexerMultiply       = 14
	SimpleQueryLanguageLexerDivide         = 15
	SimpleQueryLanguageLexerLeftParen      = 16
	SimpleQueryLanguageLexerRightParen     = 17
	SimpleQueryLanguageLexerLeftBracket    = 18
	SimpleQueryLanguageLexerRightBracket   = 19
	SimpleQueryLanguageLexerQuestion       = 20
	SimpleQueryLanguageLexerNullCoalescing = 21
	SimpleQueryLanguageLexerRange          = 22
	SimpleQueryLanguageLexerComma          = 23
	SimpleQueryLanguageLexerLeftBrace      = 24
	SimpleQueryLanguageLexerRightBrace     = 25
	SimpleQueryLanguageLexerNull           = 26
	SimpleQueryLanguageLexerFalse          = 27
	SimpleQueryLanguageLexerTrue           = 28
	SimpleQueryLanguageLexerESCAPED_STRING = 29
	SimpleQueryLanguageLexerSTRING         = 30
	SimpleQueryLanguageLexerNUMBER         = 31
	SimpleQueryLanguageLexerWS             = 32
)
