package object

import (
	"bytes"
	"fmt"
	"strings"
)

// HashKey : Definitions for how to determine the keys for a hash for each type of object
type HashKey struct {
	Type  ObjectType
	Value uint64
}

// HashPair : Represents a key-value pair used in the Hash object
type HashPair struct {
	Key   Object
	Value Object
}

// Hash : Object representing a map of key-value pairs
type Hash struct {
	Pairs map[HashKey]HashPair
}

// Type : Returns Object's type
func (h *Hash) Type() ObjectType {
	return HASH_OBJ
}

// Inspect : Returns string representation of Object's value
func (h *Hash) Inspect() string {
	var out bytes.Buffer

	pairs := []string{}
	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s", pair.Key.Inspect(), pair.Value.Inspect()))
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}
