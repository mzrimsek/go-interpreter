package object

import (
	"fmt"
)

// Integer : Object representing an integer
type Integer struct {
	Value int64
}

func (i *Integer) number() {}

// Type : Returns Object's type
func (i *Integer) Type() ObjectType {
	return INTEGER_OBJ
}

// Inspect : Returns string representation of Object's value
func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

// HashKey : Generates a HashKey
func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}
