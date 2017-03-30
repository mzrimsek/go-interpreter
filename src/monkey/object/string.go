package object

// String : Object representing a string
type String struct {
	Value string
}

// Type : Returns Object's type
func (s *String) Type() ObjectType {
	return STRING_OBJ
}

// Inspect : Returns string representation of Object's value
func (s *String) Inspect() string {
	return s.Value
}
