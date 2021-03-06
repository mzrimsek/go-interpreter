// Package object : Object definitions used to evaluate program AST nodes
package object

// ObjectType : String representation of each type of Object
type ObjectType string

// ObjectTypes
const (
	INTEGER_OBJ      = "INTEGER"
	FLOAT_OBJ        = "FLOAT"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
	FUNCTION_OBJ     = "FUNCTION"
	STRING_OBJ       = "STRING"
	BUILTIN_OBJ      = "BUILTIN"
	ARRAY_OBJ        = "ARRAY"
	HASH_OBJ         = "HASH"
	CHAR_OBJ		 = "CHAR"
)

// Object : Generic object
type Object interface {
	Type() ObjectType
	Inspect() string
}

// Number : Generic number type
type Number interface {
	number()
}

// Hashable : Defines Objects that are usable as keys in a Hash
type Hashable interface {
	HashKey() HashKey
}
