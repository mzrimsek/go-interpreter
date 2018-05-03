package object

import (
	"bytes"
	"strings"
)

// Array : Object representing an array of elements
type Array struct {
	Elements []Object
}

// Type : Returns Object's type
func (ao *Array) Type() ObjectType {
	return ARRAY_OBJ
}

// Inspect : Returns string representation of Object's value
func (ao *Array) Inspect() string {
	var out bytes.Buffer

	elements := []string{}
	for _, el := range ao.Elements {
		elements = append(elements, el.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}
