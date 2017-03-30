package object

// ReturnValue : Object representing a return value
type ReturnValue struct {
	Value Object
}

// Type : Returns Object's type
func (rv *ReturnValue) Type() ObjectType {
	return RETURN_VALUE_OBJ
}

// Inspect : Returns string representation of the wrapped Object's value
func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}
