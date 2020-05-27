package evaluator

import (
	"path/filepath"

	"../object"
)

var stdPath = object.Record{
	Stoned: true,
	Values: map[string]object.Object{
		"join": object.BuiltinFunction(func(args ...object.Object) object.Object {
			if err := checkAllArgType("env.get", args, object.STRING); err != nil {
				return err
			}

			strs := make([]string, len(args))

			for i, arg := range args {
				strs[i] = arg.String()
			}

			return object.String(filepath.Join(strs...))
		}),
	},
}
