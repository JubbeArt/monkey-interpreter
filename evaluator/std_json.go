package evaluator

import (
	"../object"
)

var stdJson = object.Record{
	Stoned: true,
	Values: map[string]object.Object{
		"parse": object.BuiltinFunction(func(args ...object.Object) object.Object {
			return object.Nil{}
		}),
		"string": object.BuiltinFunction(func(args ...object.Object) object.Object {
			return object.Nil{}
		}),
	},
}
