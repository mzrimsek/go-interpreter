package object

import (
	"fmt"
)

// Boolean : Object representing a boolean
type Boolean struct {
	Value bool
}

// Type : Returns string representation of the object type
func (b *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}

// Inspect : Returns string representation of the object value
func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}
