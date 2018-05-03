package object

import (
	"fmt"
)

// Boolean : Object representing a boolean
type Boolean struct {
	Value bool
}

// Type : Returns Object's type
func (b *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}

// Inspect : Returns string representation of Object's value
func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

// HashKey : Generates a HashKey
func (b *Boolean) HashKey() HashKey {
	var value uint64

	if b.Value {
		value = 1
	} else {
		value = 0
	}

	return HashKey{Type: b.Type(), Value: value}
}
