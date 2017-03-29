package object

// ReturnValue : Object representing a return value
type ReturnValue struct {
	Value Object
}

// Type : Returns string representation of the object type
func (rv *ReturnValue) Type() ObjectType {
	return RETURN_VALUE_OBJ
}

// Inspect : Returns string representation of the object value
func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}
