package object

import "hash/fnv"

// String : Object representing a string
type String struct {
	Value string
}

// Type : Returns Object's type
func (s *String) Type() ObjectType {
	return STRING_OBJ
}

// Inspect : Returns string representation of Object's value
func (s *String) Inspect() string {
	return s.Value
}

// HashKey : Generates a HashKey
func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))

	return HashKey{Type: s.Type(), Value: h.Sum64()}
}
