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
			if err := checkArgLength("json.string", args, 1); err != nil {
				return err
			}
			return object.String(args[0].Json(0))
		}),
	},
}
