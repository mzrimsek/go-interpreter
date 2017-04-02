// Package object : Object definitions used to evaluate program AST nodes
package object

// ObjectType : String representation of each type of Object
type ObjectType string

// ObjectTypes
const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
	FUNCTION_OBJ     = "FUNCTION"
	STRING_OBJ       = "STRING"
	BUILTIN_OBJ      = "BUILTIN"
	ARRAY_OBJ        = "ARRAY"
	HASH_OBJ         = "HASH"
)

// Object : Generic object
type Object interface {
	Type() ObjectType
	Inspect() string
}

// Hashable : Defines Objects that are usable as keys in a Hash
type Hashable interface {
	HashKey() HashKey
}
