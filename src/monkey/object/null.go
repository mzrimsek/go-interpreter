package object

// Null : Object representing a null value
type Null struct{}

// Type : Returns string representation of the object type
func (n *Null) Type() ObjectType {
	return NULL_OBJ
}

// Inspect : Returns string representation of the object value
func (n *Null) Inspect() string {
	return "null"
}
