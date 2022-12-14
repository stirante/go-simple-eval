package eval

import (
	"github.com/Bedrock-OSS/go-burrito/burrito"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/gammazero/deque"
	"github.com/stirante/go-simple-eval/eval/functions"
	"github.com/stirante/go-simple-eval/eval/utils"
	"github.com/stirante/go-simple-eval/parser"
)

// CollectingErrorListener is an error listener that collects all errors by appending them to the Error field
type CollectingErrorListener struct {
	*antlr.DefaultErrorListener
	Error error
}

func (l *CollectingErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.Error = burrito.WrapErrorf(l.Error, "column: %d %s", column, msg)
}

// Eval evaluates the given expression and returns the result
func Eval(text string, scope map[string]interface{}) (interface{}, error) {
	listener := CollectingErrorListener{DefaultErrorListener: antlr.NewDefaultErrorListener()}
	is := antlr.NewInputStream(text)
	lexer := parser.NewSimpleQueryLanguageLexer(is)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(&listener)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewSimpleQueryLanguageParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(&listener)
	p.BuildParseTrees = true
	tree := p.Expression()
	if listener.Error != nil {
		return nil, burrito.WrapErrorf(listener.Error, "Failed to parse expression %s", text)
	}
	s := deque.Deque[interface{}]{}
	s.PushBack(scope)
	visitor := ExpressionVisitor{
		scope: s,
	}
	r := visitor.Visit(tree)
	var err error
	if isError(r) {
		err = getError(r)
	}
	return utils.UnwrapContainers(r), err
}

func Init() {
	// Init functions
	functions.Init()
}
