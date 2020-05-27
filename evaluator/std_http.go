package evaluator

//
//http
//- get
//- download

import (
	"io/ioutil"
	"net/http"

	"../object"
)

var stdHttp = object.Record{
	Stoned: true,
	Values: map[string]object.Object{
		"get": object.BuiltinFunction(func(args ...object.Object) object.Object {
			if err := checkArgLength("env.get", args, 1); err != nil {
				return err
			}
			if err := checkFirstArgType("env.get", args[0], object.STRING); err != nil {
				return err
			}

			url := args[0].String()
			resp, err := http.Get(url)

			if err != nil {
				return object.Error("http.get: " + err.Error())
			}

			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)

			if err != nil {
				return object.Error("http.get: " + err.Error())
			}

			return object.String(body)
		}),
	},
}
