package object

import (
	"fmt"
	"strings"
)

// Float : Object representing a float
type Float struct {
	Value float64
}

func (f *Float) number() {}

// Type : Returns Object's type
func (f *Float) Type() ObjectType {
	return FLOAT_OBJ
}

// Inspect : Returns string representation of Object's value
func (f *Float) Inspect() string {
	return strings.TrimRight(fmt.Sprintf("%f", f.Value), "0")
}
