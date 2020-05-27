package evaluator

import (
	"strings"

	"../object"
)

var stdString = object.Record{
	Stoned: true,
	Values: map[string]object.Object{
		"has": object.BuiltinFunction(func(args ...object.Object) object.Object {
			if err := checkArgLength("string.has", args, 2); err != nil {
				return err
			}

			if err := checkFirstArgType("string.has", args[0], object.STRING); err != nil {
				return err
			}

			if err := checkSecondArgType("string.has", args[1], object.STRING); err != nil {
				return err
			}

			return object.Boolean(strings.Contains(
				args[0].String(),
				args[1].String()))
		}),
		"has_uncased": object.BuiltinFunction(func(args ...object.Object) object.Object {
			if err := checkArgLength("string.has_uncased", args, 2); err != nil {
				return err
			}

			if err := checkFirstArgType("string.has_uncased", args[0], object.STRING); err != nil {
				return err
			}

			if err := checkSecondArgType("string.has_uncased", args[1], object.STRING); err != nil {
				return err
			}

			return object.Boolean(strings.Contains(
				strings.ToLower(args[0].String()),
				strings.ToLower(args[1].String())))
		}),
	},
}
