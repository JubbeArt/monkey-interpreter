package evaluator

import (
	"../object"
)

var stdMath = object.Record{
	Stoned: true,
	Values: map[string]object.Object{
		"max": object.BuiltinFunction(func(args ...object.Object) object.Object {
			if err := checkArgLengthAtLeast("math.max", args, 1); err != nil {
				return err
			}
			if err := checkAllArgType("math.max", args, object.NUMBER); err != nil {
				return err
			}

			res := args[0].(object.Number)

			for _, arg := range args[1:] {
				a := arg.(object.Number)

				if a > res {
					res = a
				}
			}

			return res
		}),
		"min": object.BuiltinFunction(func(args ...object.Object) object.Object {
			if err := checkArgLengthAtLeast("math.min", args, 1); err != nil {
				return err
			}
			if err := checkAllArgType("math.min", args, object.NUMBER); err != nil {
				return err
			}

			res := args[0].(object.Number)

			for _, arg := range args[1:] {
				a := arg.(object.Number)

				if a < res {
					res = a
				}
			}

			return res
		}),
	},
}
