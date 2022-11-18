// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // SimpleQueryLanguage

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseSimpleQueryLanguageListener is a complete listener for a parse tree produced by SimpleQueryLanguageParser.
type BaseSimpleQueryLanguageListener struct{}

var _ SimpleQueryLanguageListener = &BaseSimpleQueryLanguageListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSimpleQueryLanguageListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSimpleQueryLanguageListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSimpleQueryLanguageListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSimpleQueryLanguageListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSimpleQueryLanguageListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSimpleQueryLanguageListener) ExitExpression(ctx *ExpressionContext) {}

// EnterFunction_param is called when production function_param is entered.
func (s *BaseSimpleQueryLanguageListener) EnterFunction_param(ctx *Function_paramContext) {}

// ExitFunction_param is called when production function_param is exited.
func (s *BaseSimpleQueryLanguageListener) ExitFunction_param(ctx *Function_paramContext) {}

// EnterField is called when production field is entered.
func (s *BaseSimpleQueryLanguageListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseSimpleQueryLanguageListener) ExitField(ctx *FieldContext) {}

// EnterArray is called when production array is entered.
func (s *BaseSimpleQueryLanguageListener) EnterArray(ctx *ArrayContext) {}

// ExitArray is called when production array is exited.
func (s *BaseSimpleQueryLanguageListener) ExitArray(ctx *ArrayContext) {}

// EnterObject is called when production object is entered.
func (s *BaseSimpleQueryLanguageListener) EnterObject(ctx *ObjectContext) {}

// ExitObject is called when production object is exited.
func (s *BaseSimpleQueryLanguageListener) ExitObject(ctx *ObjectContext) {}

// EnterObject_field is called when production object_field is entered.
func (s *BaseSimpleQueryLanguageListener) EnterObject_field(ctx *Object_fieldContext) {}

// ExitObject_field is called when production object_field is exited.
func (s *BaseSimpleQueryLanguageListener) ExitObject_field(ctx *Object_fieldContext) {}

// EnterName is called when production name is entered.
func (s *BaseSimpleQueryLanguageListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BaseSimpleQueryLanguageListener) ExitName(ctx *NameContext) {}

// EnterIndex is called when production index is entered.
func (s *BaseSimpleQueryLanguageListener) EnterIndex(ctx *IndexContext) {}

// ExitIndex is called when production index is exited.
func (s *BaseSimpleQueryLanguageListener) ExitIndex(ctx *IndexContext) {}
