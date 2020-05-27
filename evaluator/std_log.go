package evaluator

import (
	"os"
	"path/filepath"

	"../object"
)

var logFolder = "."

var stdLog = object.Record{
	Stoned: true,
	Values: map[string]object.Object{
		"info": object.BuiltinFunction(func(args ...object.Object) object.Object {
			if err := checkArgLength("log.info", args, 1); err != nil {
				return err
			}

			f, err := os.OpenFile(filepath.Join(logFolder, "info.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

			if err != nil {
				return object.Error("log.info: " + err.Error())
			}

			_, _ = f.WriteString(args[0].String())
			_ = f.Close()
			return object.Nil{}
		}),
		"error": object.BuiltinFunction(func(args ...object.Object) object.Object {
			if err := checkArgLength("log.error", args, 1); err != nil {
				return err
			}

			f, err := os.OpenFile(filepath.Join(logFolder, "error.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

			if err != nil {
				return object.Error("log.error: " + err.Error())
			}

			_, _ = f.WriteString(args[0].String())
			_ = f.Close()
			return object.Nil{}
		}),
		"set_folder": object.BuiltinFunction(func(args ...object.Object) object.Object {
			if err := checkArgLength("log.set_folder", args, 1); err != nil {
				return err
			}

			if err := checkFirstArgType("log.set_folder", args[0], object.STRING); err != nil {
				return err
			}

			logFolder = args[0].String()

			return object.Nil{}
		}),
	},
}
