package evaluator

import (
	"../object"
)

var stdList = object.Record{
	Stoned: true,
	Values: map[string]object.Object{
		"each": object.BuiltinFunction(func(args ...object.Object) object.Object {
			if err := checkArgLength("list.map", args, 2); err != nil {
				return err
			}

			if err := checkFirstArgType("list.map", args[0], object.LIST); err != nil {
				return err
			}

			if err := checkSecondArgType("list.map", args[1], object.FUNCTION); err != nil {
				return err
			}

			list := args[0].(object.List)

			switch fun := args[1].(type) {
			case object.BuiltinFunction:
				for _, item := range list {
					fun(item)
				}
			case object.Function:

			}

			return object.Nil{}
		}),
		"map": object.BuiltinFunction(func(args ...object.Object) object.Object {
			if err := checkArgLength("list.map", args, 2); err != nil {
				return err
			}

			if err := checkFirstArgType("list.map", args[0], object.LIST); err != nil {
				return err
			}

			if err := checkSecondArgType("list.map", args[1], object.FUNCTION); err != nil {
				return err
			}

			list := args[0].(object.List)
			mappedList := make([]object.Object, len(list))

			switch fun := args[1].(type) {
			case object.BuiltinFunction:
				for i, item := range list {
					mappedList[i] = fun(item)
				}
			case object.Function:

			}

			return object.List(mappedList)
		}),
		"filter": object.BuiltinFunction(func(args ...object.Object) object.Object {
			if err := checkArgLength("list.filter", args, 2); err != nil {
				return err
			}

			if err := checkFirstArgType("list.filter", args[0], object.LIST); err != nil {
				return err
			}

			if err := checkSecondArgType("list.filter", args[1], object.FUNCTION); err != nil {
				return err
			}

			list := args[0].(object.List)
			filteredList := make([]object.Object, 0)

			switch fun := args[1].(type) {
			case object.BuiltinFunction:
				for _, item := range list {
					if fun(item).Bool() {
						filteredList = append(filteredList, item)

					}
				}
			case object.Function:

			}

			return object.List(filteredList)
		}),
		"reduce": object.BuiltinFunction(func(args ...object.Object) object.Object {
			if err := checkArgLength("list.reduce", args, 3); err != nil {
				return err
			}

			if err := checkFirstArgType("list.reduce", args[0], object.LIST); err != nil {
				return err
			}

			if err := checkSecondArgType("list.reduce", args[1], object.FUNCTION); err != nil {
				return err
			}

			list := args[0].(object.List)
			accumulator := args[1]

			switch fun := args[1].(type) {
			case object.BuiltinFunction:
				for _, item := range list {
					accumulator = fun(accumulator, item)
				}
			case object.Function:

			}

			return accumulator
		}),
		"flat": object.BuiltinFunction(func(args ...object.Object) object.Object {
			if err := checkArgLength("list.flat", args, 1); err != nil {
				return err
			}

			if err := checkFirstArgType("list.flat", args[0], object.LIST); err != nil {
				return err
			}

			list := args[0].(object.List)
			flattenedList := make([]object.Object, 0, len(list))

			for _, item := range list {
				if item.Type() == object.LIST {
					flattenedList = append(flattenedList, item.(object.List)...)
				} else {
					flattenedList = append(flattenedList, item)
				}
			}

			return object.List(flattenedList)
		}),
		"has": object.BuiltinFunction(func(args ...object.Object) object.Object {
			if err := checkArgLength("list.has", args, 2); err != nil {
				return err
			}

			if err := checkFirstArgType("list.has", args[0], object.LIST); err != nil {
				return err
			}

			list := args[0].(object.List)
			check := args[1]

			for _, item := range list {
				if item.Equal(check) {
					return object.Boolean(true)
				}
			}

			return object.Boolean(false)
		}),
	},
}
