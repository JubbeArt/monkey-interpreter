package evaluator

import (
	"time"

	"../object"
)

//
//time
//- measure(str)
//- end_measure(str) # should print, or return string

var stdTime = object.Record{
	Stoned: true,
	Values: map[string]object.Object{
		"now": object.BuiltinFunction(func(args ...object.Object) object.Object {
			return object.Number(time.Now().Unix())
		}),
		"sleep": object.BuiltinFunction(func(args ...object.Object) object.Object {
			if err := checkArgLength("time.sleep", args, 1); err != nil {
				return err
			}

			if err := checkFirstArgType("time.sleep", args[0], object.NUMBER); err != nil {
				return err
			}

			time.Sleep(time.Duration(args[0].(object.Number)) * time.Second)
			return object.Nil{}
		}),
	},
}
