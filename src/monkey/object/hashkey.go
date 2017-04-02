package object

import (
	"hash/fnv"
)

// HashKey : Definitions for how to determine the keys for a hash for each type of object
type HashKey struct {
	Type  ObjectType
	Value uint64
}

// HashKey : Generates a HashKey for Boolean objects
func (b *Boolean) HashKey() HashKey {
	var value uint64

	if b.Value {
		value = 1
	} else {
		value = 0
	}

	return HashKey{Type: b.Type(), Value: value}
}

// HashKey : Generates a HashKey for Integer objects
func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

// HashKey : Generates a HashKey for String objects
func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))

	return HashKey{Type: s.Type(), Value: h.Sum64()}
}
