// Code generated from ../grammar/SimpleQueryLanguage.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // SimpleQueryLanguage

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SimpleQueryLanguageListener is a complete listener for a parse tree produced by SimpleQueryLanguageParser.
type SimpleQueryLanguageListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterFunction_param is called when entering the function_param production.
	EnterFunction_param(c *Function_paramContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterArray is called when entering the array production.
	EnterArray(c *ArrayContext)

	// EnterObject is called when entering the object production.
	EnterObject(c *ObjectContext)

	// EnterObject_field is called when entering the object_field production.
	EnterObject_field(c *Object_fieldContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterIndex is called when entering the index production.
	EnterIndex(c *IndexContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitFunction_param is called when exiting the function_param production.
	ExitFunction_param(c *Function_paramContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitArray is called when exiting the array production.
	ExitArray(c *ArrayContext)

	// ExitObject is called when exiting the object production.
	ExitObject(c *ObjectContext)

	// ExitObject_field is called when exiting the object_field production.
	ExitObject_field(c *Object_fieldContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitIndex is called when exiting the index production.
	ExitIndex(c *IndexContext)
}
