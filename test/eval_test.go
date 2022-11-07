package test

import (
	"fmt"
	"github.com/stirante/go-simple-eval/eval"
	"github.com/stirante/go-simple-eval/eval/utils"
	"math"
	"reflect"
	"testing"
)

func safeTypeName(v interface{}) string {
	if v == nil {
		return "nil"
	}
	return reflect.TypeOf(v).Name()
}

func compareJsonObject(t *testing.T, expected map[string]interface{}, actual map[string]interface{}, path string, checkExtra bool) {
	t.Helper()
	for key, value1 := range expected {
		if value2, ok := actual[key]; ok {
			newPath := fmt.Sprintf("%s/%s", path, key)
			if v1, ok := value1.(map[string]interface{}); ok {
				if v2, ok := value2.(map[string]interface{}); ok {
					compareJsonObject(t, v1, v2, newPath, true)
				} else {
					t.Errorf("Field %s is not an object (expected %s, got %s)", newPath, utils.ToString(v1), utils.ToString(value2))
				}
			} else if v1, ok := value1.([]interface{}); ok {
				if v2, ok := value2.([]interface{}); ok {
					compareJsonArray(t, v1, v2, newPath)
				} else {
					t.Errorf("Field %s is not an array (expected %s, got %s)", newPath, utils.ToString(v1), utils.ToString(value2))
				}
			} else if v1, ok := value1.(utils.JsonNumber); ok {
				if v2, ok := value2.(utils.JsonNumber); ok {
					if v1.FloatValue() != v2.FloatValue() {
						t.Errorf("Field %s is not equal (expected %s, got %s)", newPath, utils.ToString(v1), utils.ToString(v2))
					}
				} else {
					t.Errorf("Field %s is not a number (expected %s (%s), got %s (%s))", newPath, utils.ToString(v1), safeTypeName(v1), utils.ToString(value2), safeTypeName(value2))
				}
			} else if utils.IsNumber(value1) && utils.IsNumber(value2) {
				if utils.ToNumber(value1).FloatValue() != utils.ToNumber(value2).FloatValue() {
					t.Errorf("Field %s is not equal (expected %s (%s), got %s (%s))", newPath, utils.ToString(value1), safeTypeName(value1), utils.ToString(value2), safeTypeName(value2))
				}
			} else {
				if value1 != value2 {
					t.Errorf("Field %s is not equal (expected %s (%s), got %s (%s))", newPath, utils.ToString(value1), safeTypeName(value1), utils.ToString(value2), safeTypeName(value2))
				}
			}
		} else {
			t.Errorf("Object does not contain key %s", key)
		}
	}
	if checkExtra {
		for key := range actual {
			if _, ok := expected[key]; !ok {
				t.Errorf("Object contains unexpected key %s/%s", path, key)
			}
		}
	}
}

func compareJsonArray(t *testing.T, expected []interface{}, actual []interface{}, path string) {
	t.Helper()
	for i := 0; i < int(math.Min(float64(len(expected)), float64(len(actual)))); i++ {
		newPath := fmt.Sprintf("%s[%d]", path, i)
		value1 := expected[i]
		value2 := actual[i]
		if v1, ok := value1.(map[string]interface{}); ok {
			if v2, ok := value2.(map[string]interface{}); ok {
				compareJsonObject(t, v1, v2, newPath, true)
			} else {
				t.Errorf("Element %s is not an object (expected %s, got %s)", newPath, utils.ToString(v1), utils.ToString(value2))
			}
		} else if v1, ok := value1.([]interface{}); ok {
			if v2, ok := value2.([]interface{}); ok {
				compareJsonArray(t, v1, v2, newPath)
			} else {
				t.Errorf("Element %s is not an array (expected %s, got %s)", newPath, utils.ToString(v1), utils.ToString(value2))
			}
		} else if v1, ok := value1.(utils.JsonNumber); ok {
			if v2, ok := value2.(utils.JsonNumber); ok {
				if v1.FloatValue() != v2.FloatValue() {
					t.Errorf("Element %s is not equal (expected %s, got %s)", newPath, utils.ToString(v1), utils.ToString(v2))
				}
			} else {
				t.Errorf("Element %s is not a number (expected %s (%s), got %s (%s))", newPath, utils.ToString(v1), safeTypeName(v1), utils.ToString(value2), safeTypeName(value2))
			}
		} else if utils.IsNumber(value1) && utils.IsNumber(value2) {
			if utils.ToNumber(value1).FloatValue() != utils.ToNumber(value2).FloatValue() {
				t.Errorf("Element %s is not equal (expected %s (%s), got %s (%s))", newPath, utils.ToString(value1), safeTypeName(value1), utils.ToString(value2), safeTypeName(value2))
			}
		} else {
			if value1 != value2 {
				t.Errorf("Element %s is not equal (expected %s (%s), got %s (%s))", newPath, utils.ToString(value1), safeTypeName(value1), utils.ToString(value2), safeTypeName(value2))
			}
		}
	}
	for i := 0; i < len(actual); i++ {
		if i >= len(expected) {
			t.Errorf("Array contains unexpected element %s[%d]", path, i)
		}
	}
}

func assertArray(t *testing.T, actual interface{}, expected []interface{}) []interface{} {
	t.Helper()
	if actual == nil {
		t.Fatalf("Result is null")
	}
	if !utils.IsArray(actual) {
		t.Fatalf("Result is not an array (%s)", reflect.TypeOf(actual).Name())
	}
	array, ok := actual.([]interface{})
	if !ok {
		t.Fatalf("Result is not a JSON array (%s)", reflect.TypeOf(actual).Name())
	}
	if len(array) != len(expected) {
		t.Fatalf("Array length is not correct (expected %d, got %d)", len(expected), len(array))
	}
	for i := 0; i < len(expected); i++ {
		if utils.IsNumber(expected[i]) && utils.IsNumber(array[i]) {
			if utils.ToNumber(expected[i]).FloatValue() != utils.ToNumber(array[i]).FloatValue() {
				t.Fatalf("Array element %d is not correct (expected %f, got %f)", i, expected[i], array[i])
			}
		} else if utils.IsObject(expected[i]) {
			if array[i] == nil {
				t.Fatalf("Array element %d is null", i)
			}
			if !utils.IsObject(array[i]) {
				t.Fatalf("Array element %d is not an object (%s)", i, reflect.TypeOf(array[i]).Name())
			}
			compareJsonObject(t, expected[i].(map[string]interface{}), array[i].(map[string]interface{}), fmt.Sprintf("#[%d]", i), true)
		} else if utils.IsArray(expected[i]) {
			if array[i] == nil {
				t.Fatalf("Array element %d is null", i)
			}
			if !utils.IsArray(array[i]) {
				t.Fatalf("Array element %d is not an array (%s)", i, reflect.TypeOf(array[i]).Name())
			}
			compareJsonArray(t, expected[i].([]interface{}), array[i].([]interface{}), fmt.Sprintf("#[%d]", i))
		} else if array[i] != expected[i] {
			t.Errorf("Array element %d is not correct (expected %s, got %s)", i, utils.ToString(array[i]), utils.ToString(expected[i]))
		}
	}
	return array
}

func assertObject(t *testing.T, actual interface{}, expected map[string]interface{}) map[string]interface{} {
	t.Helper()
	if actual == nil {
		t.Fatalf("Result is null")
	}
	if !utils.IsObject(actual) {
		t.Fatalf("Result is not an object (%s)", reflect.TypeOf(actual).Name())
	}
	obj, ok := actual.(map[string]interface{})
	if !ok {
		t.Fatalf("Result is not a JSON object (%s)", reflect.TypeOf(actual).Name())
	}
	compareJsonObject(t, expected, obj, "#", true)
	return obj
}

func assertObjectContains(t *testing.T, actual interface{}, expected map[string]interface{}) map[string]interface{} {
	t.Helper()
	if actual == nil {
		t.Fatalf("Result is null")
	}
	if !utils.IsObject(actual) {
		t.Fatalf("Result is not an object (%s)", reflect.TypeOf(actual).Name())
	}
	obj, ok := actual.(map[string]interface{})
	if !ok {
		t.Fatalf("Result is not a JSON object (%s)", reflect.TypeOf(actual).Name())
	}
	compareJsonObject(t, expected, obj, "#", false)
	return obj
}

func assertNumber(t *testing.T, actual interface{}, expected float64) {
	t.Helper()
	if actual == nil {
		t.Fatalf("Result is null")
	}
	if !utils.IsNumber(actual) {
		t.Fatalf("Result is not a number (%s)", reflect.TypeOf(actual).Name())
	}
	number, ok := actual.(float64)
	if !ok {
		iNumber, ok := actual.(int32)
		if !ok {
			t.Fatalf("Result is not a valid number (%s)", reflect.TypeOf(actual).Name())
		} else if float64(iNumber) != expected {
			t.Fatalf("Result is not correct (expected %f, got %d)", expected, iNumber)
		}
	} else if number != expected {
		t.Fatalf("Result is not correct (expected %f, got %f)", expected, number)
	}
}

func assertBool(t *testing.T, actual interface{}, expected bool) {
	t.Helper()
	if actual == nil {
		t.Fatalf("Result is null")
	}
	if _, ok := actual.(bool); !ok {
		t.Fatalf("Result is not a boolean (%s)", reflect.TypeOf(actual).Name())
	}
	if actual.(bool) != expected {
		t.Fatalf("Result is not correct (expected %t, got %t)", expected, actual)
	}
}

func assertString(t *testing.T, actual interface{}, expected string) {
	t.Helper()
	if actual == nil {
		t.Fatalf("Result is null")
	}
	if _, ok := actual.(string); !ok {
		t.Fatalf("Result is not a string (%s)", reflect.TypeOf(actual).Name())
	}
	if actual.(string) != expected {
		t.Fatalf("Result is not correct (expected %s, got %s)", expected, actual)
	}
}

func assertNull(t *testing.T, actual interface{}) {
	t.Helper()
	if actual != nil {
		t.Fatalf("Result is not null (%s)", utils.ToString(actual))
	}
}

func evaluate(t *testing.T, text string) interface{} {
	t.Helper()
	e, err := eval.Eval(text, nil)
	if err != nil {
		t.Fatal(err)
	}
	return e
}

func evaluateWithScope(t *testing.T, text string, scope map[string]interface{}) interface{} {
	t.Helper()
	e, err := eval.Eval(text, scope)
	if err != nil {
		t.Fatal(err)
	}
	return e
}

func TestRangeOperator(t *testing.T) {
	e := evaluate(t, "1..10")
	assertArray(t, e, []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
}

func TestAddition(t *testing.T) {
	e := evaluate(t, "1 + 2")
	assertNumber(t, e, 3)
}

func TestSubtraction(t *testing.T) {
	e := evaluate(t, "1 - 2")
	assertNumber(t, e, -1)
}

func TestMultiplication(t *testing.T) {
	e := evaluate(t, "2 * 3")
	assertNumber(t, e, 6)
}

func TestIntegerDivision(t *testing.T) {
	e := evaluate(t, "6 / 4")
	assertNumber(t, e, 1)
}

func TestFloatDivision(t *testing.T) {
	e := evaluate(t, "6.0 / 4")
	assertNumber(t, e, 1.5)
}

func TestOperationOrder(t *testing.T) {
	e := evaluate(t, "2 + 2 * 2")
	assertNumber(t, e, 6)
}

func TestOperationOrder2(t *testing.T) {
	e := evaluate(t, "2 * 2 + 2")
	assertNumber(t, e, 6)
}

func TestOperationOrder3(t *testing.T) {
	e := evaluate(t, "2 * (2 + 2)")
	assertNumber(t, e, 8)
}

func TestEquality(t *testing.T) {
	e := evaluate(t, "1 == 1")
	assertBool(t, e, true)
}

func TestInequality(t *testing.T) {
	e := evaluate(t, "1 != 1")
	assertBool(t, e, false)
}

func TestLessThan(t *testing.T) {
	e := evaluate(t, "1 < 2")
	assertBool(t, e, true)
}

func TestLessThanOrEqual(t *testing.T) {
	e := evaluate(t, "1 <= 2")
	assertBool(t, e, true)
}

func TestGreaterThan(t *testing.T) {
	e := evaluate(t, "2 > 1")
	assertBool(t, e, true)
}

func TestGreaterThanOrEqual(t *testing.T) {
	e := evaluate(t, "2 >= 1")
	assertBool(t, e, true)
}

func TestAnd(t *testing.T) {
	e := evaluate(t, "true && true")
	assertBool(t, e, true)
}

func TestOr(t *testing.T) {
	e := evaluate(t, "true || false")
	assertBool(t, e, true)
}

func TestNot(t *testing.T) {
	e := evaluate(t, "!true")
	assertBool(t, e, false)
}

func TestNot2(t *testing.T) {
	e := evaluate(t, "!false")
	assertBool(t, e, true)
}

func TestNot3(t *testing.T) {
	e := evaluate(t, "!!true")
	assertBool(t, e, true)
}

func TestNot4(t *testing.T) {
	e := evaluate(t, "!!false")
	assertBool(t, e, false)
}

func TestArrayAccess(t *testing.T) {
	e := evaluate(t, "[1, 2, 3][1]")
	assertNumber(t, e, 2)
}

func TestObjectAccess(t *testing.T) {
	e := evaluate(t, `{"a": 1, "b": 2}["b"]`)
	assertNumber(t, e, 2)
}

func TestObjectAccess2(t *testing.T) {
	e := evaluate(t, `{"a": 1, "b": 2}.b`)
	assertNumber(t, e, 2)
}

func TestScope(t *testing.T) {
	e := evaluateWithScope(t, `b`, map[string]interface{}{
		"a": utils.ToNumber(1),
		"b": utils.ToNumber(2),
	})
	assertNumber(t, e, 2)
}

func TestConcatenation(t *testing.T) {
	e := evaluate(t, `'a' + 'b'`)
	assertString(t, e, "ab")
}

func TestConcatenation2(t *testing.T) {
	e := evaluate(t, `'a' + 1`)
	assertString(t, e, "a1")
}

func TestConcatenation3(t *testing.T) {
	e := evaluate(t, `1 + 'b'`)
	assertString(t, e, "1b")
}

func TestConcatenation4(t *testing.T) {
	e := evaluate(t, `1 + 2 + 'b'`)
	assertString(t, e, "3b")
}

func TestConcatenation5(t *testing.T) {
	e := evaluate(t, `'a' + 1 + 2`)
	assertString(t, e, "a12")
}

func TestConcatenation6(t *testing.T) {
	e := evaluate(t, `'a' + (1 + 2)`)
	assertString(t, e, "a3")
}

func TestNullCoalescing(t *testing.T) {
	e := evaluate(t, `1 ?? 2`)
	assertNumber(t, e, 1)
}

func TestNullCoalescing2(t *testing.T) {
	e := evaluate(t, `null ?? 2`)
	assertNumber(t, e, 2)
}

func TestNullCoalescing3(t *testing.T) {
	e := evaluate(t, `null ?? null`)
	assertNull(t, e)
}

func TestNullCoalescing4(t *testing.T) {
	e := evaluate(t, `null ?? null ?? 3`)
	assertNumber(t, e, 3)
}

func TestTernaryOperator(t *testing.T) {
	e := evaluate(t, `true ? 1 : 2`)
	assertNumber(t, e, 1)
}

func TestTernaryOperator2(t *testing.T) {
	e := evaluate(t, `false ? 1 : 2`)
	assertNumber(t, e, 2)
}

func TestTernaryOperator3(t *testing.T) {
	e := evaluate(t, `true ? 1 : 2 + 3`)
	assertNumber(t, e, 1)
}

func TestTernaryOperator4(t *testing.T) {
	e := evaluate(t, `false ? 1 : 2 + 3`)
	assertNumber(t, e, 5)
}

func TestTernaryOperator5(t *testing.T) {
	e := evaluate(t, `true ? 1 + 2 : 3`)
	assertNumber(t, e, 3)
}

func TestTernaryOperator6(t *testing.T) {
	e := evaluate(t, `false ? 1 + 2 : 3`)
	assertNumber(t, e, 3)
}

func TestTernaryOperator7(t *testing.T) {
	e := evaluate(t, `false ? 1 : true ? 2 : 3`)
	assertNumber(t, e, 2)
}

func TestTernaryOperator8(t *testing.T) {
	e := evaluate(t, `false ? 1 : false ? 2 : 3`)
	assertNumber(t, e, 3)
}

func TestNullForgivingObjectOperator(t *testing.T) {
	e := evaluate(t, `{}?.a`)
	assertNull(t, e)
}

func TestNullForgivingObjectOperator2(t *testing.T) {
	e := evaluate(t, `{"a": 1}?.a`)
	assertNumber(t, e, 1)
}

func TestNullForgivingArrayOperator(t *testing.T) {
	e := evaluate(t, `[]?[0]`)
	assertNull(t, e)
}

func TestNullForgivingArrayOperator2(t *testing.T) {
	e := evaluate(t, `[1]?[0]`)
	assertNumber(t, e, 1)
}
