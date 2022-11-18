// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // SimpleQueryLanguage

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// A complete Visitor for a parse tree produced by SimpleQueryLanguageParser.
type SimpleQueryLanguageVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SimpleQueryLanguageParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by SimpleQueryLanguageParser#function_param.
	VisitFunction_param(ctx *Function_paramContext) interface{}

	// Visit a parse tree produced by SimpleQueryLanguageParser#field.
	VisitField(ctx *FieldContext) interface{}

	// Visit a parse tree produced by SimpleQueryLanguageParser#array.
	VisitArray(ctx *ArrayContext) interface{}

	// Visit a parse tree produced by SimpleQueryLanguageParser#object.
	VisitObject(ctx *ObjectContext) interface{}

	// Visit a parse tree produced by SimpleQueryLanguageParser#object_field.
	VisitObject_field(ctx *Object_fieldContext) interface{}

	// Visit a parse tree produced by SimpleQueryLanguageParser#name.
	VisitName(ctx *NameContext) interface{}

	// Visit a parse tree produced by SimpleQueryLanguageParser#index.
	VisitIndex(ctx *IndexContext) interface{}
}
