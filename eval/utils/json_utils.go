package utils

import (
	"encoding/json"
	"math"
	"reflect"
	"strconv"
	"strings"
)

// JsonFunction is a function that can be executed.
type JsonFunction func(args []interface{}) (interface{}, error)

// Number is an interface that represents a number, that can be either integer or decimal.
type Number interface {
	// IntValue returns the number as an integer.
	IntValue() int
	// FloatValue returns the number as a float.
	FloatValue() float64
	// BoolValue returns the number as a boolean.
	BoolValue() bool
	// StringValue returns the number as a string.
	StringValue() string
}

// JsonNumber is a struct that represents a number, that can be either integer or decimal.
type JsonNumber struct {
	Number
	Value   float64
	Decimal bool
}

func (n JsonNumber) IntValue() int32 {
	return int32(n.Value)
}

// toFixed rounds a float to a given precision.
func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return math.Round(num*output) / output
}

func (n JsonNumber) FloatValue() float64 {
	if n.Decimal {
		return toFixed(n.Value, 10)
	} else {
		return math.Floor(n.Value)
	}
}

func (n JsonNumber) BoolValue() bool {
	if toFixed(n.Value, 10) == 0 {
		return false
	}
	return true
}

func (n JsonNumber) StringValue() string {
	if n.Decimal {
		return strconv.FormatFloat(n.Value, 'f', -1, 64)
	}
	return strconv.FormatInt(int64(n.IntValue()), 10)
}

// ToBoolean converts an interface to a boolean.
func ToBoolean(obj interface{}) bool {
	if obj == nil {
		return false
	}
	if b, ok := obj.(bool); ok {
		return b
	}
	if b, ok := obj.(int); ok && b != 0 {
		return true
	}
	if b, ok := obj.(float64); ok && b != 0 {
		return true
	}
	if b, ok := obj.(float32); ok && b != 0 {
		return true
	}
	if b, ok := obj.(string); ok && strings.Trim(b, "\n\r") != "" {
		return true
	}
	if b, ok := obj.(JsonNumber); ok {
		return b.BoolValue()
	}
	return obj != nil
}

func ToSemver(obj interface{}) Semver {
	if obj == nil {
		return Semver{}
	}
	if b, ok := obj.(Semver); ok {
		return b
	}
	return Semver{}
}

// ToNumber converts an interface to a JSON number.
func ToNumber(obj interface{}) JsonNumber {
	if obj == nil {
		return JsonNumber{
			Value:   0,
			Decimal: false,
		}
	}
	if b, ok := obj.(JsonNumber); ok {
		return b
	}
	if b, ok := obj.(float64); ok {
		return JsonNumber{
			Value:   b,
			Decimal: true,
		}
	}
	if b, ok := obj.(float32); ok {
		return JsonNumber{
			Value:   float64(b),
			Decimal: true,
		}
	}
	if b, ok := obj.(int); ok {
		return JsonNumber{
			Value:   float64(b),
			Decimal: false,
		}
	}
	if b, ok := obj.(int32); ok {
		return JsonNumber{
			Value:   float64(b),
			Decimal: false,
		}
	}
	if b, ok := obj.(uint32); ok {
		return JsonNumber{
			Value:   float64(b),
			Decimal: false,
		}
	}
	if b, ok := obj.(int64); ok {
		return JsonNumber{
			Value:   float64(b),
			Decimal: false,
		}
	}
	if b, ok := obj.(bool); ok && b {
		return JsonNumber{
			Value:   1,
			Decimal: false,
		}
	}
	if b, ok := obj.(string); ok {
		result, err := strconv.ParseInt(b, 10, 64)
		if err != nil {
			result1, err := strconv.ParseFloat(b, 64)
			if err != nil {
				return JsonNumber{
					Value:   0,
					Decimal: false,
				}
			}
			return JsonNumber{
				Value:   result1,
				Decimal: true,
			}
		}
		return JsonNumber{
			Value:   float64(result),
			Decimal: false,
		}
	}
	if b, ok := obj.(json.Number); ok {
		result, err := strconv.ParseInt(string(b), 10, 64)
		if err != nil {
			result1, err := strconv.ParseFloat(string(b), 64)
			if err != nil {
				return JsonNumber{
					Value:   0,
					Decimal: false,
				}
			}
			return JsonNumber{
				Value:   result1,
				Decimal: true,
			}
		}
		return JsonNumber{
			Value:   float64(result),
			Decimal: false,
		}
	}
	return JsonNumber{
		Value:   0,
		Decimal: false,
	}
}

// ToString converts an interface to a string.
func ToString(obj interface{}) string {
	if obj == nil {
		return "null"
	}
	if b, ok := obj.(JsonNumber); ok {
		return b.StringValue()
	}
	if b, ok := obj.(Semver); ok {
		return b.String()
	}
	if b, ok := obj.(float64); ok {
		return strconv.FormatFloat(b, 'f', -1, 64)
	}
	if b, ok := obj.(float32); ok {
		return strconv.FormatFloat(float64(b), 'f', -1, 64)
	}
	if b, ok := obj.(int); ok {
		return strconv.FormatInt(int64(b), 10)
	}
	if b, ok := obj.(int32); ok {
		return strconv.FormatInt(int64(b), 10)
	}
	if b, ok := obj.(bool); ok && b {
		return strconv.FormatBool(b)
	}
	if b, ok := obj.(string); ok {
		return b
	}
	str, err := json.Marshal(UnwrapContainers(obj))
	if err != nil {
		return "null"
	}
	return strings.ReplaceAll(string(str), "\n", "")
}

// IsNumber returns true if the given interface is a number.
func IsNumber(obj interface{}) bool {
	if obj == nil {
		return false
	}
	if _, ok := obj.(JsonNumber); ok {
		return true
	}
	if _, ok := obj.(float64); ok {
		return true
	}
	if _, ok := obj.(float32); ok {
		return true
	}
	if _, ok := obj.(int); ok {
		return true
	}
	if _, ok := obj.(int32); ok {
		return true
	}
	if _, ok := obj.(bool); ok {
		return true
	}
	return false
}

// IsSemver returns true if the given interface is a semver object.
func IsSemver(obj interface{}) bool {
	if obj == nil {
		return false
	}
	if _, ok := obj.(Semver); ok {
		return true
	}
	return false
}

// IsArray returns true if the given interface is an array.
func IsArray(obj interface{}) bool {
	if obj == nil {
		return false
	}
	rt := reflect.TypeOf(obj)
	switch rt.Kind() {
	case reflect.Slice, reflect.Array:
		return true
	}
	return false
}

// AsArray returns the given interface as a JSON array.
func AsArray(obj interface{}) []interface{} {
	if obj == nil {
		return nil
	}
	rt := reflect.TypeOf(obj)
	switch rt.Kind() {
	case reflect.Slice, reflect.Array:
		rv := reflect.ValueOf(obj)
		result := make([]interface{}, rv.Len())
		for i := 0; i < rv.Len(); i++ {
			result[i] = rv.Index(i).Interface()
		}
		return result
	}
	return nil
}

// AsObject returns the given interface as a JSON object.
func AsObject(obj interface{}) map[string]interface{} {
	if obj == nil {
		return map[string]interface{}{}
	}
	if b, ok := obj.(map[string]interface{}); ok {
		return b
	}
	rt := reflect.TypeOf(obj)
	switch rt.Kind() {
	case reflect.Map:
		rv := reflect.ValueOf(obj)
		result := map[string]interface{}{}
		for _, key := range rv.MapKeys() {
			result[key.String()] = rv.MapIndex(key).Interface()
		}
		return result
	}
	return map[string]interface{}{}
}

// IsObject returns true if the given interface is an object.
func IsObject(obj interface{}) bool {
	if obj == nil {
		return false
	}
	if _, ok := obj.(map[string]interface{}); ok {
		return true
	}
	rt := reflect.TypeOf(obj)
	switch rt.Kind() {
	case reflect.Map:
		return true
	}
	return false
}

// MergeObject merges two JSON objects into a new JSON object.
// If the same value, that is not an object or an array exists in both objects, the value from the second object will be used.
func MergeObject(template, parent map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	for k, v := range template {
		if IsObject(v) {
			if val, ok := v.(map[string]interface{}); ok {
				result[k] = DeepCopyObject(val)
			}
		} else if IsArray(v) {
			if val, ok := v.([]interface{}); ok {
				result[k] = DeepCopyArray(val)
			} else if val, ok := v.([]interface{}); ok {
				result[k] = DeepCopyArray(val)
			}
		} else {
			result[k] = v
		}
	}
	skipKeys := make([]string, 0)
out:
	for k, v := range parent {
		for _, key := range skipKeys {
			if key == k {
				continue out
			}
		}
		if _, ok := template[k]; !ok {
			if IsObject(v) {
				merge := MergeObject(map[string]interface{}{}, v.(map[string]interface{}))
				result[k] = merge
			} else if IsArray(v) {
				merge := MergeArray(nil, v.([]interface{}))
				result[k] = merge
			} else {
				result[k] = v
			}
		} else {
			if IsObject(v) && IsObject(result[k]) {
				merge := MergeObject(template[k].(map[string]interface{}), v.(map[string]interface{}))
				result[k] = merge
			} else if IsArray(v) && IsArray(template[k]) {
				merge := MergeArray(template[k].([]interface{}), v.([]interface{}))
				result[k] = merge
			} else {
				result[k] = v
			}
		}
	}
	return result
}

// MergeArray merges two JSON arrays into a new JSON array.
func MergeArray(template, parent []interface{}) []interface{} {
	var result []interface{}
	for _, v := range template {
		if IsObject(v) {
			merge := MergeObject(map[string]interface{}{}, v.(map[string]interface{}))
			result = append(result, merge)
		} else if IsArray(v) {
			merge := MergeArray(nil, v.([]interface{}))
			result = append(result, merge)
		} else {
			result = append(result, v)
		}
	}
	for _, v := range parent {
		if IsObject(v) {
			merge := MergeObject(map[string]interface{}{}, v.(map[string]interface{}))
			result = append(result, merge)
		} else if IsArray(v) {
			merge := MergeArray(nil, v.([]interface{}))
			result = append(result, merge)
		} else {
			result = append(result, v)
		}
	}
	return result
}

// IsEqual returns true if the given interfaces are equal.
func IsEqual(a, b interface{}) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if IsSemver(a) && IsSemver(b) {
		return a.(Semver).Equals(b.(Semver))
	}
	if (IsObject(a) != IsObject(b)) || (IsArray(a) != IsArray(b)) || (IsNumber(a) != IsNumber(b)) {
		return false
	}
	if IsNumber(a) && IsNumber(b) {
		return ToNumber(a).FloatValue() == ToNumber(b).FloatValue()
	}
	if IsArray(a) && IsArray(b) {
		return IsEqualArray(a.([]interface{}), b.([]interface{}))
	}
	if IsObject(a) && IsObject(b) {
		return IsEqualObject(a.(map[string]interface{}), b.(map[string]interface{}))
	}
	if a == b {
		return true
	}
	return false
}

// IsEqualObject returns true if the given JSON objects are equal.
func IsEqualObject(a, b map[string]interface{}) bool {
	if len(a) != len(b) {
		return false
	}
	for k, _ := range a {
		if !IsEqual(b[k], a[k]) {
			return false
		}
	}
	return true
}

// IsEqualArray returns true if the given JSON arrays are equal.
func IsEqualArray(a, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if !IsEqual(b[k], v) {
			return false
		}
	}
	return true
}

// CreateRange creates a range of numbers from start to end as a JSON array.
func CreateRange(start, end int32) []interface{} {
	var result []interface{}
	if start > end {
		return result
	}
	for i := start; i <= end; i++ {
		result = append(result, ToNumber(i))
	}
	return result
}

// UnescapeString removes quotes and unescapes a string.
func UnescapeString(str string) string {
	if len(str) < 3 {
		return ""
	}
	str = str[1 : len(str)-1]
	str = strings.ReplaceAll(str, "\\\\\"", "\"")
	str = strings.ReplaceAll(str, "\\\\'", "'")
	str = strings.ReplaceAll(str, "\\\\n", "\n")
	str = strings.ReplaceAll(str, "\\\\\\\\", "\\\\")
	return str
}

// UnwrapContainers removes all containers from the given interface.
// Currently only unpacks JsonNumber into an actual number with correct type.
func UnwrapContainers(obj interface{}) interface{} {
	if obj == nil {
		return nil
	}
	if b, ok := obj.(JsonNumber); ok {
		if b.Decimal {
			return b.FloatValue()
		} else {
			return b.IntValue()
		}
	}
	if IsSemver(obj) {
		return obj.(Semver).String()
	} else if IsObject(obj) {
		object := AsObject(obj)
		c := map[string]interface{}{}
		for k := range object {
			c[k] = UnwrapContainers(object[k])
		}
		return c
	} else if IsArray(obj) {
		array := AsArray(obj)
		c := make([]interface{}, len(array))
		for k, v := range array {
			c[k] = UnwrapContainers(v)
		}
		return c
	}
	return obj
}

// DeepCopyObject creates a deep copy of the given JSON object.
func DeepCopyObject(object map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	for k, v := range object {
		if IsObject(v) {
			result[k] = DeepCopyObject(AsObject(v))
		} else if IsArray(v) {
			result[k] = DeepCopyArray(AsArray(v))
		} else {
			result[k] = v
		}
	}
	return result
}

// DeepCopyArray creates a deep copy of the given JSON array.
func DeepCopyArray(object []interface{}) []interface{} {
	var result []interface{}
	for _, v := range object {
		if IsObject(v) {
			result = append(result, DeepCopyObject(AsObject(v)))
		} else if IsArray(v) {
			result = append(result, DeepCopyArray(AsArray(v)))
		} else {
			result = append(result, v)
		}
	}
	return result
}

// DeleteNulls removes all keys with null values from the given JSON object.
func DeleteNulls(object map[string]interface{}) map[string]interface{} {
	for k, v := range object {
		if IsObject(v) {
			object[k] = DeleteNulls(AsObject(v))
		} else if IsArray(v) {
			object[k] = DeleteNullsFromArray(AsArray(v))
		} else if v == nil {
			delete(object, k)
		}
	}
	return object
}

// DeleteNullsFromArray removes all keys inside elements of JSON object type with null values from the given JSON array.
func DeleteNullsFromArray(array []interface{}) []interface{} {
	for i, v := range array {
		if IsObject(v) {
			array[i] = DeleteNulls(AsObject(v))
		} else if IsArray(v) {
			array[i] = DeleteNullsFromArray(AsArray(v))
		}
	}
	return array
}

func IndexOf[T any](object []T, value T) int {
	for i, v := range object {
		if IsEqual(v, value) {
			return i
		}
	}
	return -1
}

func convertNumbersObject(object map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	for k, v := range object {
		if IsObject(v) {
			result[k] = convertNumbersObject(AsObject(v))
		} else if IsArray(v) {
			result[k] = convertNumbersArray(AsArray(v))
		} else if _, ok := v.(json.Number); ok {
			result[k] = ToNumber(v)
		} else {
			result[k] = v
		}
	}
	return result
}

func convertNumbersArray(object []interface{}) []interface{} {
	var result []interface{}
	for _, v := range object {
		if IsObject(v) {
			result = append(result, convertNumbersObject(AsObject(v)))
		} else if IsArray(v) {
			result = append(result, convertNumbersArray(AsArray(v)))
		} else if _, ok := v.(json.Number); ok {
			result = append(result, ToNumber(v))
		} else {
			result = append(result, v)
		}
	}
	return result
}
