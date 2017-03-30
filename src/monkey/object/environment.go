package object

// Environment : Contains information about values that have mapped to a declared identifier so they can be accessed throughout the lifetime of a program
type Environment struct {
	store map[string]Object
	outer *Environment
}

// NewEnvironment : Creates a new Environment at the highest level with no identifier mappings
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

// NewEnclosedEnvironment : Creates a new Environment nested inside a higher level Environment for use in functions
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

// Get : Attempts to retrieve an identifier's value if it exists
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

// Set : Sets an identifier's value
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
