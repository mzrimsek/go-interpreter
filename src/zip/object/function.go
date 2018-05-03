package object

import (
	"bytes"
	"strings"
	"zip/ast"
)

// Function : Object representing a function including its parameters and body statements
type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

// Type : Returns Object's type
func (f *Function) Type() ObjectType {
	return FUNCTION_OBJ
}

// Inspect : Returns string representation of function's parameters and body
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}
