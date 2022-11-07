// Code generated from ../grammar/SimpleQueryLanguage.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // SimpleQueryLanguage

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseSimpleQueryLanguageVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSimpleQueryLanguageVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleQueryLanguageVisitor) VisitFunction_param(ctx *Function_paramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleQueryLanguageVisitor) VisitField(ctx *FieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleQueryLanguageVisitor) VisitArray(ctx *ArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleQueryLanguageVisitor) VisitObject(ctx *ObjectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleQueryLanguageVisitor) VisitObject_field(ctx *Object_fieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleQueryLanguageVisitor) VisitName(ctx *NameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleQueryLanguageVisitor) VisitIndex(ctx *IndexContext) interface{} {
	return v.VisitChildren(ctx)
}
