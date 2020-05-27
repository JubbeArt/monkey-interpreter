package evaluator

import (
	"../object"
)

var stdRecord = object.Record{
	Stoned: true,
	Values: map[string]object.Object{
		"keys": object.BuiltinFunction(func(args ...object.Object) object.Object {
			if err := checkArgLength("record.keys", args, 1); err != nil {
				return err
			}

			if err := checkFirstArgType("record.keys", args[0], object.RECORD); err != nil {
				return err
			}

			record := args[0].(object.Record)

			keys := make([]object.Object, len(record.Values))

			i := 0
			for key := range record.Values {
				keys[i] = object.String(key)
				i++
			}

			return object.List(keys)
		}),
		"values": object.BuiltinFunction(func(args ...object.Object) object.Object {
			if err := checkArgLength("record.values", args, 1); err != nil {
				return err
			}

			if err := checkFirstArgType("record.values", args[0], object.RECORD); err != nil {
				return err
			}

			record := args[0].(object.Record)

			values := make([]object.Object, len(record.Values))

			i := 0
			for _, value := range record.Values {
				values[i] = value
				i++
			}

			return object.List(values)
		}),
	},
}
