package eval

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gammazero/deque"
	"github.com/stirante/go-simple-query-language/eval/functions"
	"github.com/stirante/go-simple-query-language/eval/utils"
	"github.com/stirante/go-simple-query-language/parser"
	"reflect"
)

const NaN = "NaN"

var majorAliases = []string{
	"major",
	"a",
	"x",
}
var minorAliases = []string{
	"minor",
	"b",
	"y",
}
var patchAliases = []string{
	"patch",
	"c",
	"z",
}

type ExpressionVisitor struct {
	parser.BaseSimpleQueryLanguageVisitor
	scope deque.Deque[interface{}]
	path  *string
}

func (v *ExpressionVisitor) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *parser.FieldContext:
		return v.VisitField(val)
	case *parser.ArrayContext:
		return v.VisitArray(val)
	case *parser.ObjectContext:
		return v.VisitObject(val)
	case *parser.Object_fieldContext:
		return v.VisitObject_field(val)
	case *parser.Function_paramContext:
		return v.VisitFunction_param(val)
	case *parser.NameContext:
		return v.VisitName(val)
	case *parser.IndexContext:
		return v.VisitIndex(val)
	case *parser.ExpressionContext:
		return v.VisitExpression(val)
	}
	panic("Unknown tree type " + reflect.TypeOf(tree).String())
}

// pushScope pushes a new scope to the stack
func (v *ExpressionVisitor) pushScope(scope map[string]interface{}) {
	v.scope.PushBack(scope)
}

// pushScopePair pushes a new entry to the stack
func (v *ExpressionVisitor) pushScopePair(key string, value interface{}) {
	v.scope.PushBack(map[string]interface{}{key: value})
}

// popScope pops the last scope from the stack
func (v *ExpressionVisitor) popScope() {
	v.scope.PopBack()
}

// negate negates a value
func negate(value interface{}) interface{} {
	if value == nil {
		return nil
	}
	if b, ok := value.(utils.JsonNumber); ok {
		return utils.JsonNumber{
			Value:   -b.Value,
			Decimal: b.Decimal,
		}
	}
	if b, ok := value.(bool); ok {
		return !b
	}
	if b, ok := value.(int); ok {
		return utils.ToNumber(-b)
	}
	if b, ok := value.(float64); ok {
		return utils.ToNumber(-b)
	}
	if b, ok := value.(float32); ok {
		return utils.ToNumber(-b)
	}
	if utils.IsArray(value) {
		result := make([]interface{}, len(value.([]interface{})))
		for i, v := range value.([]interface{}) {
			result[i] = negate(v)
		}
		return result
	}
	return NaN
}

func isError(v interface{}) bool {
	_, err := v.(error)
	return err
}

func getError(v interface{}) error {
	if err, ok := v.(error); ok {
		return err
	}
	return nil
}

// resolveScope resolves a value from the scope by name
func (v *ExpressionVisitor) resolveScope(name string) interface{} {
	for i := v.scope.Len() - 1; i >= 0; i-- {
		m := v.scope.At(i)
		if c, ok := m.(map[string]interface{}); ok {
			if b, ok := c[name]; ok {
				return b
			}
		}
		// Seems like sometimes above cast fails, so we need to check for map[string]interface{} as well
		if c, ok := m.(map[string]interface{}); ok {
			if v, ok := c[name]; ok {
				return v
			}
		}
	}
	return nil
}

func (v *ExpressionVisitor) VisitExpression(context *parser.ExpressionContext) interface{} {
	return v.Visit(context.Field())
}

func (v *ExpressionVisitor) VisitField(context *parser.FieldContext) interface{} {
	if context.Null() != nil {
		return nil
	}
	if context.True() != nil {
		return true
	}
	if context.False() != nil {
		return false
	}
	if context.Not() != nil {
		return !utils.ToBoolean(v.Visit(context.Field(0)))
	}
	// process field composed of two other fields
	if len(context.AllField()) == 2 {
		f1 := v.Visit(context.Field(0))
		if isError(f1) {
			return f1
		}
		// Move AND and OR here to make those operators short-circuiting
		if context.And() != nil {
			if utils.ToBoolean(f1) {
				f2 := v.Visit(context.Field(1))
				if isError(f2) {
					return f2
				}
				return utils.ToBoolean(f2)
			} else {
				return false
			}
		} else if context.Or() != nil {
			b := utils.ToBoolean(f1)
			if !b {
				f2 := v.Visit(context.Field(1))
				if isError(f2) {
					return f2
				}
				return utils.ToBoolean(f2)
			} else {
				return b
			}
		} else if context.Question() != nil {
			if utils.ToBoolean(f1) {
				f2 := v.Visit(context.Field(1))
				if isError(f2) {
					return f2
				}
				return f2
			} else {
				return nil
			}
		}
		f2 := v.Visit(context.Field(1))
		if isError(f2) {
			return f2
		}
		if context.NullCoalescing() != nil {
			if f1 == nil {
				return f2
			} else {
				return f1
			}
		} else if context.Add() != nil {
			if (utils.IsNumber(f1) && utils.IsNumber(f2)) || (utils.IsNumber(f1) && f2 == nil) || (f1 == nil && utils.IsNumber(f2)) {
				n1 := utils.ToNumber(f1)
				n2 := utils.ToNumber(f2)
				return utils.JsonNumber{
					Value:   n1.FloatValue() + n2.FloatValue(),
					Decimal: n1.Decimal || n2.Decimal,
				}
			} else if utils.IsArray(f1) && utils.IsArray(f2) {
				array := utils.MergeArray(f1.([]interface{}), f2.([]interface{}))
				return array
			} else if utils.IsObject(f1) && utils.IsObject(f2) {
				var result []interface{} = nil
				result = append(result, f1.([]interface{})...)
				return append(result, f2.([]interface{})...)
			} else {
				return utils.ToString(f1) + utils.ToString(f2)
			}
		} else if context.Equal() != nil {
			return utils.IsEqual(f1, f2)
		} else if context.NotEqual() != nil {
			return !utils.IsEqual(f1, f2)
		} else if utils.IsNumber(f1) && utils.IsNumber(f2) {
			n1 := utils.ToNumber(f1)
			n2 := utils.ToNumber(f2)
			if context.Range() != nil {
				return utils.CreateRange(n1.IntValue(), n2.IntValue())
			} else if context.Greater() != nil {
				return n1.FloatValue() > n2.FloatValue()
			} else if context.Less() != nil {
				return n1.FloatValue() < n2.FloatValue()
			} else if context.GreaterOrEqual() != nil {
				return n1.FloatValue() >= n2.FloatValue()
			} else if context.LessOrEqual() != nil {
				return n1.FloatValue() <= n2.FloatValue()
			} else if n1.Decimal || n2.Decimal {
				if context.Subtract() != nil {
					return utils.JsonNumber{
						Value:   n1.FloatValue() - n2.FloatValue(),
						Decimal: true,
					}
				} else if context.Divide() != nil {
					return utils.JsonNumber{
						Value:   n1.FloatValue() / n2.FloatValue(),
						Decimal: true,
					}
				} else if context.Multiply() != nil {
					return utils.JsonNumber{
						Value:   n1.FloatValue() * n2.FloatValue(),
						Decimal: true,
					}
				}
			} else {
				if context.Subtract() != nil {
					return utils.JsonNumber{
						Value:   float64(n1.IntValue() - n2.IntValue()),
						Decimal: false,
					}
				} else if context.Divide() != nil {
					return utils.JsonNumber{
						Value:   float64(n1.IntValue() / n2.IntValue()),
						Decimal: false,
					}
				} else if context.Multiply() != nil {
					return utils.JsonNumber{
						Value:   float64(n1.IntValue() * n2.IntValue()),
						Decimal: false,
					}
				}
			}
		} else if utils.IsSemver(f1) && utils.IsSemver(f2) {
			s1 := utils.ToSemver(f1)
			s2 := utils.ToSemver(f2)
			if context.Greater() != nil {
				return s1.CompareTo(s2) > 0
			} else if context.Less() != nil {
				return s1.CompareTo(s2) < 0
			} else if context.GreaterOrEqual() != nil {
				return s1.CompareTo(s2) >= 0
			} else if context.LessOrEqual() != nil {
				return s1.CompareTo(s2) <= 0
			}
		} else if context.Greater() != nil || context.Less() != nil || context.GreaterOrEqual() != nil || context.LessOrEqual() != nil {
			return false
		} else {
			return NaN
		}
	} else if len(context.AllField()) == 3 {
		// handle ternary operator
		if context.Question() != nil {
			f1 := v.Visit(context.Field(0))
			if isError(f1) {
				return f1
			}
			if utils.ToBoolean(f1) {
				return v.Visit(context.Field(1))
			} else {
				return v.Visit(context.Field(2))
			}
		}
	}
	// if the field is another field in parentheses, return the value of that field
	// we need to also check if the first element is not a field, because if it is, it will be a function call
	if context.LeftParen() != nil && len(context.AllField()) == 1 && context.GetChild(0) != context.Field(0) {
		return v.Visit(context.Field(0))
	} else if context.LeftParen() != nil && len(context.AllField()) == 1 && context.GetChild(0) == context.Field(0) {
		lambda := v.Visit(context.Field(0))
		if isError(lambda) {
			return lambda
		}
		params := make([]interface{}, len(context.AllFunction_param()))
		for i, param := range context.AllFunction_param() {
			params[i] = v.Visit(param)
			if isError(params[i]) {
				return params[i]
			}
		}

		if _, ok := lambda.(utils.JsonFunction); ok {
			i, err := lambda.(utils.JsonFunction)(params)
			if err != nil {
				return err
			}
			return i
		} else {
			var methodName *string = nil
			for _, child := range context.Field(0).GetChildren() {
				if b, ok := child.(parser.INameContext); ok {
					text := b.GetText()
					methodName = &text
					break
				}
			}
			if methodName == nil || !functions.HasFunction(*methodName) {
				return utils.WrappedErrorf("Function '%s' not found!", context.Field(0).GetText())
			}
			function, err := functions.CallFunction(*methodName, params)
			if err != nil {
				return utils.WrapErrorf(err, "Error calling function '%s'", *methodName)
			}
			return function
		}
	}
	// handle accessing a property of an object or calling an instance function
	if context.Name() != nil && len(context.AllField()) == 1 {
		text := context.Name().GetText()
		object := v.Visit(context.Field(0))
		if isError(object) {
			return object
		}
		var newScope interface{} = nil
		if object == nil {
			// handle null-forgiving operator
			if context.Question() != nil {
				return nil
			} else {
				return utils.WrappedErrorf("Cannot access %s because %s is %s", context.GetText(), context.Field(0).GetText(), utils.ToString(object))
			}
		}
		if utils.IsSemver(object) {
			u := utils.ToSemver(object)
			if utils.IndexOf(majorAliases, text) != -1 {
				return utils.ToNumber(u.Major)
			}
			if utils.IndexOf(minorAliases, text) != -1 {
				return utils.ToNumber(u.Minor)
			}
			if utils.IndexOf(patchAliases, text) != -1 {
				return utils.ToNumber(u.Patch)
			}
			return utils.WrappedErrorf("Cannot access %s because %s is not a valid semver field", context.GetText(), text)
		} else if utils.IsObject(object) {
			u := object.(map[string]interface{})
			if c, ok := u[text]; ok {
				newScope = c
			}
		} else if functions.HasInstanceFunction(reflect.TypeOf(object), text) {
			return utils.JsonFunction(
				func(o []interface{}) (interface{}, error) {
					function, err := functions.CallInstanceFunction(text, object, o)
					if err != nil {
						return nil, utils.WrapErrorf(err, "Error calling function '%s' on %s", text, utils.ToString(object))
					}
					return function, nil
				},
			)
		}

		if newScope == nil {
			// handle null-forgiving operator
			if context.Question() != nil {
				return nil
			} else {
				return utils.WrappedErrorf("Failed to resolve field '%s' (%s) in %s", text, context.GetText(), utils.ToString(object))
			}
		}
		return newScope
	}
	if context.Name() != nil {
		return v.Visit(context.Name())
	}
	// handle indexed access
	if context.Index() != nil && len(context.AllField()) == 1 {
		i := v.Visit(context.Index())
		if isError(i) {
			return i
		}
		object := v.Visit(context.Field(0))
		if isError(object) {
			return object
		}
		if utils.IsArray(object) {
			// in case of an array, we need an integer index
			if utils.IsNumber(i) {
				value := utils.ToNumber(i).IntValue()
				if value < 0 || value >= int32(len(object.([]interface{}))) {
					// handle null-forgiving operator
					if context.Question() != nil {
						return nil
					} else {
						return utils.WrappedErrorf("Index out of bounds: %d (%s)", value, context.GetText())
					}
				}
				return object.([]interface{})[value]
			} else {
				// handle null-forgiving operator
				if context.Question() != nil {
					return nil
				} else {
					return utils.WrappedErrorf("Index must be a number: %s (%s)", utils.ToString(i), context.GetText())
				}
			}
		} else if utils.IsSemver(object) {
			// in case of an object, we need a string index
			value := utils.ToString(i)
			u := utils.ToSemver(object)
			if utils.IndexOf(majorAliases, value) != -1 {
				return utils.ToNumber(u.Major)
			}
			if utils.IndexOf(minorAliases, value) != -1 {
				return utils.ToNumber(u.Minor)
			}
			if utils.IndexOf(patchAliases, value) != -1 {
				return utils.ToNumber(u.Patch)
			}
			return utils.WrappedErrorf("Cannot access %s because %s is not a valid semver field", context.GetText(), value)
		} else if utils.IsObject(object) {
			// in case of an object, we need a string index
			value := utils.ToString(i)
			u := object.(map[string]interface{})
			if _, ok := u[value]; !ok {
				// handle null-forgiving operator
				if context.Question() != nil {
					return nil
				} else {
					return utils.WrappedErrorf("Property '%s' (%s) not found in %s", value, context.GetText(), utils.ToString(object))
				}
			} else {
				return u[value]
			}
		} else if str, ok := object.(string); ok {
			// in case of a string, we need an integer index
			if utils.IsNumber(i) {
				value := utils.ToNumber(i).IntValue()
				if value < 0 || value >= int32(len(str)) {
					// handle null-forgiving operator
					if context.Question() != nil {
						return nil
					} else {
						return utils.WrappedErrorf("Index out of bounds: %d (%s)", value, context.GetText())
					}
				}
				return string(rune(str[value]))
			} else {
				// handle null-forgiving operator
				if context.Question() != nil {
					return nil
				} else {
					return utils.WrappedErrorf("Index must be a number: %s (%s)", utils.ToString(object), context.GetText())
				}
			}
		} else {
			// handle null-forgiving operator
			if context.Question() != nil {
				return nil
			} else {
				return utils.WrappedErrorf("Cannot access %s because %s is %s", context.GetText(), context.Field(0).GetText(), utils.ToString(object))
			}
		}
	}
	if context.NUMBER() != nil {
		return utils.ToNumber(context.NUMBER().GetText())
	}
	if context.ESCAPED_STRING() != nil {
		return utils.UnescapeString(utils.ToString(context.ESCAPED_STRING().GetText()))
	}
	// literal array notation
	if context.Array() != nil {
		return v.Visit(context.Array())
	}
	// literal object notation
	if context.Object() != nil {
		return v.Visit(context.Object())
	}
	// negation
	if context.Subtract() != nil && len(context.AllField()) == 1 {
		f := v.Visit(context.Field(0))
		if isError(f) {
			return f
		}
		return negate(f)
	}
	return utils.WrappedErrorf("Failed to resolve '%s'", context.GetText())
}

func (v *ExpressionVisitor) VisitName(context *parser.NameContext) interface{} {
	text := context.GetText()
	// "this" is a special case that always refers to the current full scope
	if text == "this" {
		scope := map[string]interface{}{}
		for i := 0; i < v.scope.Len(); i++ {
			s := v.scope.At(i)
			if c, ok := s.(map[string]interface{}); ok {
				for key, _ := range c {
					scope[key] = c[key]
				}
			}
		}
		return scope
	}
	newScope := v.resolveScope(text)

	back := v.scope.Len() - 1
	for newScope == nil && back >= 0 {
		m := v.scope.At(back)
		if c, ok := m.(map[string]interface{}); ok {
			if v, ok := c[text]; ok {
				newScope = v
			}
		}
		back--
	}
	return newScope
}

func (v *ExpressionVisitor) VisitIndex(context *parser.IndexContext) interface{} {
	if context.NUMBER() != nil {
		return utils.ToNumber(context.NUMBER().GetText())
	}
	if context.ESCAPED_STRING() != nil {
		return utils.UnescapeString(utils.ToString(context.ESCAPED_STRING().GetText()))
	}
	if context.Field() != nil {
		return v.Visit(context.Field())
	}
	return utils.WrappedErrorf("Invalid index: %s", context.GetText())
}

func (v *ExpressionVisitor) VisitArray(context *parser.ArrayContext) interface{} {
	result := make([]interface{}, len(context.AllField()))
	for i, f := range context.AllField() {
		result[i] = v.Visit(f)
		if isError(result[i]) {
			return result[i]
		}
	}
	return result
}

func (v *ExpressionVisitor) VisitObject(context *parser.ObjectContext) interface{} {
	result := map[string]interface{}{}
	for _, f := range context.AllObject_field() {
		obj := v.Visit(f)
		if isError(obj) {
			return obj
		}
		u := obj.(map[string]interface{})
		for key, v := range u {
			result[key] = v
		}
	}
	return result
}

func (v *ExpressionVisitor) VisitObject_field(context *parser.Object_fieldContext) interface{} {
	name := ""
	if context.ESCAPED_STRING() != nil {
		name = utils.UnescapeString(utils.ToString(context.ESCAPED_STRING().GetText()))
	} else {
		name = context.Name().GetText()
	}
	field := v.Visit(context.Field())
	if isError(field) {
		return field
	}
	n := map[string]interface{}{}
	n[name] = field
	return n
}

func (v *ExpressionVisitor) VisitFunction_param(ctx *parser.Function_paramContext) interface{} {
	if ctx.Field() != nil {
		return v.Visit(ctx.Field())
	}
	return utils.WrappedErrorf("Invalid function parameter: %s", ctx.GetText())
}
