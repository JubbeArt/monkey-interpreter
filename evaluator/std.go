package evaluator

import (
	"fmt"
	"strings"

	"../object"
)

var (
	builtins = map[string]object.Object{
		"print": object.BuiltinFunction(func(args ...object.Object) object.Object {
			strs := make([]string, len(args))

			for i, arg := range args {
				strs[i] = arg.String()
			}

			fmt.Println(strings.Join(strs, " "))
			return object.Nil{}
		}),
		"conv":   stdConv,
		"env":    stdEnv,
		"fs":     stdFs,
		"json":   stdJson,
		"list":   stdList,
		"log":    stdLog,
		"math":   stdMath,
		"path":   stdPath,
		"record": stdRecord,
		"string": stdString,
		"time":   stdTime,
	}
)

func checkArgLength(name string, args []object.Object, length int) object.Object {
	if len(args) != length {
		return object.Error(fmt.Sprintf("%v expects %v arguments, got %v", name, length, args))
	}
	return nil
}

func checkArgLengthAtLeast(name string, args []object.Object, length int) object.Object {
	if len(args) <= length {
		return object.Error(fmt.Sprintf("%v expects at least %v arguments, got %v", name, length, args))
	}
	return nil
}

func checkAllArgType(name string, args []object.Object, typ object.Type) object.Object {
	for _, arg := range args {
		if arg.Type() != typ {
			return object.Error(fmt.Sprintf("%v expects %v as argument, got %v", name, typ, arg.Type()))
		}
	}
	return nil
}

func checkFirstArgType(name string, arg object.Object, typ object.Type) object.Object {
	if arg.Type() != typ {
		return object.Error(fmt.Sprintf("%v expects a %v as its first argument, got %v", name, typ, arg.Type()))
	}
	return nil
}

func checkSecondArgType(name string, arg object.Object, typ object.Type) object.Object {
	if arg.Type() != typ {
		return object.Error(fmt.Sprintf("%v expects a %v as its second argument, got %v", name, typ, arg.Type()))
	}
	return nil
}

func checkThirdArgType(name string, arg object.Object, typ object.Type) object.Object {
	if arg.Type() != typ {
		return object.Error(fmt.Sprintf("%v expects a %v as its third argument, got %v", name, typ, arg.Type()))
	}
	return nil
}
