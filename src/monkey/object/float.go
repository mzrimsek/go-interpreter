package object

import (
	"fmt"
)

// Float : Object representing a float
type Float struct {
	Value float64
}

// Type : Returns Object's type
func (f *Float) Type() ObjectType {
	return FLOAT_OBJ
}

// Inspect : Returns string representation of Object's value
func (f *Float) Inspect() string {
	return fmt.Sprintf("%f", f.Value)
}
