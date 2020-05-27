package evaluator

import (
	"os"

	"../object"
)

var stdEnv = object.Record{
	Stoned: true,
	Values: map[string]object.Object{
		"get": object.BuiltinFunction(func(args ...object.Object) object.Object {
			if err := checkArgLength("env.get", args, 1); err != nil {
				return err
			}
			if err := checkFirstArgType("env.get", args[0], object.STRING); err != nil {
				return err
			}

			key := args[0].String()
			if value, ok := os.LookupEnv(key); ok {
				return object.String(value)
			}
			return object.Nil{}
		}),
		"set": object.BuiltinFunction(func(args ...object.Object) object.Object {
			if err := checkArgLength("env.set", args, 2); err != nil {
				return err
			}

			err := os.Setenv(args[0].String(), args[1].String())

			if err != nil {
				return object.Error("env.set: " + err.Error())
			}
			return object.Nil{}
		}),
	},
}
