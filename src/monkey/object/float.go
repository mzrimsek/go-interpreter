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
	unformatted := fmt.Sprintf("%f", f.Value)
	trimmed := strings.TrimRight(unformatted, "0")
	if trimmed[len(trimmed)-1] == '.' {
		trimmed += "0"
	}
	return trimmed
}
