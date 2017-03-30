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
)

// Object : Generic object
type Object interface {
	Type() ObjectType
	Inspect() string
}
