package evaluator

//
//fs
//- write_file
//- glob
//- trash (???)

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"../object"
)

var stdFs = object.Record{
	Stoned: true,
	Values: map[string]object.Object{
		"mkdir": object.BuiltinFunction(func(args ...object.Object) object.Object {
			if err := checkArgLength("fs.mkdir", args, 1); err != nil {
				return err
			}

			if err := checkFirstArgType("fs.mkdir", args[0], object.STRING); err != nil {
				return err
			}

			dir := args[0].String()
			err := os.MkdirAll(dir, 0777)

			if err != nil {
				return object.Error("fs.mkdir: " + err.Error())
			}

			return object.Nil{}
		}),
		"exists": object.BuiltinFunction(func(args ...object.Object) object.Object {
			if err := checkArgLength("fs.exists", args, 1); err != nil {
				return err
			}

			if err := checkFirstArgType("fs.exists", args[0], object.STRING); err != nil {
				return err
			}

			fullPath := args[0].String()
			_, err := os.Stat(fullPath)

			if err == nil {
				return object.Boolean(true)
			} else if os.IsNotExist(err) {
				return object.Boolean(false)
			} else {
				return object.Error("fs.exists: " + err.Error())
			}
		}),
		"files": object.BuiltinFunction(func(args ...object.Object) object.Object {
			if err := checkArgLength("fs.files", args, 1); err != nil {
				return err
			}

			if err := checkFirstArgType("fs.files", args[0], object.STRING); err != nil {
				return err
			}

			dirPath := args[0].String()
			files := make([]object.Object, 0)

			dir, err := ioutil.ReadDir(dirPath)

			if err != nil {
				return object.Error("fs.files: " + err.Error())
			}

			for _, file := range dir {
				if !file.IsDir() {
					files = append(files, object.String(filepath.Join(dirPath, file.Name())))
				}
			}

			return object.List(files)
		}),
		"folders": object.BuiltinFunction(func(args ...object.Object) object.Object {
			if err := checkArgLength("fs.folders", args, 1); err != nil {
				return err
			}

			if err := checkFirstArgType("fs.folders", args[0], object.STRING); err != nil {
				return err
			}

			dirPath := args[0].String()
			folders := make([]object.Object, 0)

			dir, err := ioutil.ReadDir(dirPath)

			if err != nil {
				return object.Error("fs.folders: " + err.Error())
			}

			for _, file := range dir {
				if file.IsDir() {
					folders = append(folders, object.String(path.Join(dirPath, file.Name())))
				}
			}

			return object.List(folders)
		}),
		"glob": object.BuiltinFunction(func(args ...object.Object) object.Object {
			if err := checkArgLength("fs.glob", args, 1); err != nil {
				return err
			}

			if err := checkFirstArgType("fs.glob", args[0], object.STRING); err != nil {
				return err
			}

			pattern := args[0].String()
			matches, err := filepath.Glob(pattern)

			if err != nil {
				return object.Error("fs.glob: " + err.Error())
			}

			files := make([]object.Object, len(matches))

			for _, match := range matches {
				files = append(files, object.String(match))
			}

			return object.List(files)
		}),
		"home": object.BuiltinFunction(func(args ...object.Object) object.Object {
			if err := checkArgLength("fs.home", args, 0); err != nil {
				return err
			}

			home, err := os.UserHomeDir()

			if err != nil {
				return object.Error("fs.home: " + err.Error())
			}

			return object.String(home)
		}),
		"config": object.BuiltinFunction(func(args ...object.Object) object.Object {
			if err := checkArgLength("fs.config", args, 0); err != nil {
				return err
			}

			home, err := os.UserConfigDir()

			if err != nil {
				return object.Error("fs.config: " + err.Error())
			}

			return object.String(home)
		}),
		"read": object.BuiltinFunction(func(args ...object.Object) object.Object {
			if err := checkArgLength("fs.read", args, 1); err != nil {
				return err
			}

			if err := checkFirstArgType("fs.read", args[0], object.STRING); err != nil {
				return err
			}

			file, err := ioutil.ReadFile(args[1].String())

			if err != nil {
				return object.Error("fs.read: " + err.Error())
			}

			return object.String(file)
		}),
	},
}
