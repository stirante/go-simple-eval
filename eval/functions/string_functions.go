package functions

import (
	"github.com/Bedrock-OSS/go-burrito/burrito"
	"github.com/stirante/go-simple-eval/eval/utils"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"regexp"
	"strings"
)

func RegisterStringFunctions() {
	RegisterFunction(JsonFunction{
		Name:       "replace",
		Body:       replace,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "join",
		Body:       join,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "contains",
		Body:       stringContains,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "split",
		Body:       split,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "indexOf",
		Body:       stringIndexOf,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "lastIndexOf",
		Body:       stringLastIndexOf,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "toUpperCase",
		Body:       toUpperCase,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "toLowerCase",
		Body:       toLowerCase,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "substring",
		Body:       substring,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "substring",
		Body:       substringFrom,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "startsWith",
		Body:       startsWith,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "endsWith",
		Body:       endsWith,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "regexReplace",
		Body:       regexReplace,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "length",
		Body:       length,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "trim",
		Body:       trim,
		IsInstance: true,
	})
}

func replace(str, old, new string) string {
	return strings.Replace(str, old, new, -1)
}

func join(strs []interface{}, sep string) string {
	arr := make([]string, len(strs))
	for i, s := range strs {
		arr[i] = utils.ToString(s)
	}
	return strings.Join(arr, sep)
}

func stringContains(str, substr string) bool {
	return strings.Contains(str, substr)
}

func split(str, sep string) []interface{} {
	strs := strings.Split(str, sep)
	arr := make([]interface{}, len(strs))
	for i, s := range strs {
		arr[i] = s
	}
	return arr
}

func stringIndexOf(str, substr string) utils.JsonNumber {
	return utils.ToNumber(strings.Index(str, substr))
}

func stringLastIndexOf(str, substr string) utils.JsonNumber {
	return utils.ToNumber(strings.LastIndex(str, substr))
}

func toUpperCase(str string) string {
	return cases.Upper(language.Und).String(str)
}

func toLowerCase(str string) string {
	return cases.Lower(language.Und).String(str)
}

func substring(str string, start, end utils.JsonNumber) string {
	return str[start.IntValue():end.IntValue()]
}

func substringFrom(str string, start utils.JsonNumber) string {
	return str[start.IntValue():]
}

func startsWith(str, substr string) bool {
	return strings.HasPrefix(str, substr)
}

func endsWith(str, substr string) bool {
	return strings.HasSuffix(str, substr)
}

func regexReplace(str, pattern, repl string) (string, error) {
	compile, err := regexp.Compile(pattern)
	if err != nil {
		return "", burrito.WrapErrorf(err, "Failed to compile regex pattern")
	}
	return compile.ReplaceAllString(str, repl), nil
}

func length(str string) utils.JsonNumber {
	return utils.ToNumber(len(str))
}

func trim(str string) string {
	return strings.Trim(str, " \t\n\r")
}
