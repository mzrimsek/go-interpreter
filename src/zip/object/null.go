package object

// Null : Object representing a null value
type Null struct{}

// Type : Returns Object's type
func (n *Null) Type() ObjectType {
	return NULL_OBJ
}

// Inspect : Returns string representation of a null value
func (n *Null) Inspect() string {
	return "null"
}
