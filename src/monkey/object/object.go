// Package object : Object definitions used to evaluate program AST nodes
package object

// ObjectType : String representation of each type of Object
type ObjectType string

// ObjectTypes
const (
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
	NULL_OBJ    = "NULL"
)

// Object : Generic object
type Object interface {
	Type() ObjectType
	Inspect() string
}
