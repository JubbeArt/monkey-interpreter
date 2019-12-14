package object

type Environment struct {
	store map[string]Object
	outer *Environment
}

func NewEnvironment() *Environment {
	return &Environment{store: map[string]Object{}, outer: nil}
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	return &Environment{store: map[string]Object{}, outer: outer}
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]

	if ok || e.outer == nil {
		return obj, ok
	}

	return e.outer.Get(name)
}

func (e *Environment) Set(name string, obj Object) Object {
	e.store[name] = obj
	return obj
}
