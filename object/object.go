package object

import (
	"fmt"
	"strconv"
	"strings"
)
import "../ast"

type Type string

const (
	NUMBER_TYPE       Type = "number"
	STRING_TYPE       Type = "string"
	BOOLEAN_TYPE      Type = "boolean"
	NIL_TYPE          Type = "nil"
	FUNCTION_TYPE     Type = "function"
	LIST_TYPE         Type = "list"
	RECORD_TYPE       Type = "record"
	RETURN_VALUE_TYPE Type = "return_type"
	ERROR_TYPE        Type = "error_type"
)

type Object interface {
	Type() Type
	String() string
	Bool() bool
	Equal(object Object) Boolean
}

// -------------------------------------------
// ----------------- NUMBER ------------------
// -------------------------------------------
type Number float64

func (o Number) Type() Type                     { return NUMBER_TYPE }
func (o Number) Bool() bool                     { return true }
func (o Number) Negate() Number                 { return -o }
func (o Number) Add(other Number) Number        { return o + other }
func (o Number) Sub(other Number) Number        { return o - other }
func (o Number) Mul(other Number) Number        { return o * other }
func (o Number) Div(other Number) Number        { return o / other }
func (o Number) Less(other Number) Boolean      { return o < other }
func (o Number) LessEq(other Number) Boolean    { return o <= other }
func (o Number) Greater(other Number) Boolean   { return o > other }
func (o Number) GreaterEq(other Number) Boolean { return o >= other }
func (o Number) Equal(object Object) Boolean {
	if object.Type() != NUMBER_TYPE {
		return false
	}

	return o == object.(Number)
}
func (o Number) String() string { return strconv.FormatFloat(float64(o), 'f', -1, 64) }

// -------------------------------------------
// ------------------ TEXT -------------------
// -------------------------------------------
type String string

func (o String) Type() Type              { return STRING_TYPE }
func (o String) Bool() bool              { return true }
func (o String) Add(other String) String { return o + other }
func (o String) String() string          { return string(o) }
func (o String) Equal(object Object) Boolean {
	if object.Type() != STRING_TYPE {
		return false
	}

	return o == object.(String)
}

// -------------------------------------------
// ---------------- BOOLEAN ------------------
// -------------------------------------------
type Boolean bool

func (o Boolean) Type() Type     { return BOOLEAN_TYPE }
func (o Boolean) Bool() bool     { return bool(o) }
func (o Boolean) Not() Boolean   { return !o }
func (o Boolean) String() string { return strconv.FormatBool(bool(o)) }
func (o Boolean) Equal(object Object) Boolean {
	if object.Type() != BOOLEAN_TYPE {
		return false
	}

	return o == object.(Boolean)
}

// -------------------------------------------
// ------------------ NIL --------------------
// -------------------------------------------
type Nil struct{}

func (o Nil) Type() Type     { return NIL_TYPE }
func (o Nil) Bool() bool     { return false }
func (o Nil) String() string { return "nil" }
func (o Nil) Equal(object Object) Boolean {
	return object.Type() == NIL_TYPE
}

// -------------------------------------------
// --------------- FUNCTION ------------------
// -------------------------------------------
type Function struct {
	Parameters []string
	Body       *ast.BlockStatement
	Env        *Environment
}

func (o Function) Type() Type { return FUNCTION_TYPE }
func (o Function) Bool() bool { return true }
func (o Function) String() string {
	return fmt.Sprintf("fn (%v) { ... }", strings.Join(o.Parameters, ", "))
}
func (o Function) Equal(object Object) Boolean { return false }

// -------------------------------------------
// ----------- BUILTIN FUNCTION --------------
// -------------------------------------------
type BuiltinFunction func(args ...Object) Object

func (o BuiltinFunction) Type() Type { return FUNCTION_TYPE }
func (o BuiltinFunction) Bool() bool { return true }
func (o BuiltinFunction) String() string {
	return fmt.Sprintf("fn (...args) { builtin }")
}
func (o BuiltinFunction) Equal(object Object) Boolean { return false }

// -------------------------------------------
// ----------------- LIST --------------------
// -------------------------------------------
type List []Object

func (o List) Type() Type { return LIST_TYPE }
func (o List) Bool() bool { return true }
func (o List) String() string {
	if len(o) == 0 {
		return "[]"
	}

	values := []string{}

	for _, val := range o {
		values = append(values, val.String())
	}

	return fmt.Sprintf("[\n%v]", strings.Join(values, ",\n"))
}
func (o List) Equal(object Object) Boolean {
	if object.Type() != LIST_TYPE {
		return false
	}

	list := object.(List)

	if len(list) != len(o) {
		return false
	}

	for i := range o {
		if !o[i].Equal(list[i]) {
			return false
		}
	}

	return true
}

// -------------------------------------------
// ---------------- RECORD -------------------
// -------------------------------------------
type Record struct {
	Stoned bool
	Values map[string]Object
}

func (o Record) Type() Type { return RECORD_TYPE }
func (o Record) Bool() bool { return true }
func (o Record) Equal(object Object) Boolean {
	if object.Type() != RECORD_TYPE {
		return false
	}

	record := object.(Record)

	if len(record.Values) != len(o.Values) {
		return false
	}

	for key := range o.Values {
		if !o.Values[key].Equal(record.Values[key]) {
			return false
		}
	}

	return true
}
func (o *Record) Stone() { o.Stoned = true }
func (o *Record) Add(key string, object Object) {
	if !o.Stoned {
		o.Values[key] = object
	}
}
func (o *Record) Get(key string) Object {
	if obj, ok := o.Values[key]; ok {
		return obj
	}
	return Nil{}
}
func (o Record) String() string {
	if len(o.Values) == 0 {
		return "{}"
	}

	values := []string{}

	for key, val := range o.Values {
		values = append(values, key+" = "+val.String())
	}

	return fmt.Sprintf("{\n%v}", strings.Join(values, "\n"))
}

// -------------------------------------------
// ----------------- ERROR -------------------
// -------------------------------------------
type Error string

func (o Error) Type() Type                  { return ERROR_TYPE }
func (o Error) Bool() bool                  { return false }
func (o Error) String() string              { return "ERROR: " + string(o) }
func (o Error) Equal(object Object) Boolean { return false }

// -------------------------------------------
// ------------- RETURN VALUE ----------------
// -------------------------------------------
type ReturnValue struct {
	Value Object
}

func (o ReturnValue) Type() Type                  { return RETURN_VALUE_TYPE }
func (o ReturnValue) Bool() bool                  { return o.Value.Bool() }
func (o ReturnValue) String() string              { return o.Value.String() }
func (o ReturnValue) Equal(object Object) Boolean { return false }
