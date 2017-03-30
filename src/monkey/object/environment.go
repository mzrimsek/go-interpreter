package object

// Environment : Contains information about values that have mapped to a declared identifier so they can be accessed throughout the lifetime of a program
type Environment struct {
	store map[string]Object
}

// NewEnvironment : Creates a new Environment with no identifier mappings
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
}

// Get : Attempts to retrieve an identifier's value if it exists
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	return obj, ok
}

// Set : Sets an identifier's value
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
