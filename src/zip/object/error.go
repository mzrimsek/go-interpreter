package object

// Error : Object representing an error during evaluation
type Error struct {
	Message string
}

// Type : Returns Object's type
func (e *Error) Type() ObjectType {
	return ERROR_OBJ
}

// Inspect : Returns string representation of the error message
func (e *Error) Inspect() string {
	return "ERROR: " + e.Message
}
