package evaluator

import (
	"strconv"

	"../object"
)

var stdConv = object.Record{
	Stoned: true,
	Values: map[string]object.Object{
		"string": object.BuiltinFunction(func(args ...object.Object) object.Object {
			if err := checkArgLength("conv.string", args, 1); err != nil {
				return err
			}
			return object.String(args[0].String())
		}),
		"number": object.BuiltinFunction(func(args ...object.Object) object.Object {
			if err := checkArgLength("conv.number", args, 1); err != nil {
				return err
			}

			value, err := strconv.ParseFloat(args[0].String(), 64)

			if err != nil {
				return object.Error("conv.number: " + err.Error())
			}

			return object.Number(value)
		}),
	},
}
