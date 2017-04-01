package object

// BuiltinFunction : A predefined function included in the interpreter
type BuiltinFunction func(args ...Object) Object

// Builtin : Object representing a built in function
type Builtin struct {
	Fn BuiltinFunction
}

// Type : Returns Object's type
func (b *Builtin) Type() ObjectType {
	return BUILTIN_OBJ
}

// Inspect : Returns string declaring Object is a builtin function
func (b *Builtin) Inspect() string {
	return "builtin function"
}
