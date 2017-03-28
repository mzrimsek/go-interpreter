package object

import (
	"fmt"
)

// Integer : Object representing an integer
type Integer struct {
	Value int64
}

// Type : Returns string representation of the object type
func (i *Integer) Type() ObjectType {
	return INTEGER_OBJ
}

// Inspect : Returns string representation of the object value
func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}
