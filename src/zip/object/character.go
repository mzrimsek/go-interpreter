package object

import (
	"fmt"
)

// Character : Object representing an character
type Character struct {
	Value byte
}

// Type : Returns Object's type
func (c *Character) Type() ObjectType {
	return CHAR_OBJ
}

// Inspect : Returns string representation of Object's value
func (c *Character) Inspect() string {
	return fmt.Sprintf("%c", c.Value)
}

// HashKey : Generates a HashKey
func (c *Character) HashKey() HashKey {
	return HashKey{Type: c.Type(), Value: uint64(c.Value)}
}
