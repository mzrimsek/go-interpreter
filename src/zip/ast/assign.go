package ast

import (
	"bytes"
	"zip/token"
)

// AssignStatement : Statement node representing assigning a value to an existing object
type AssignStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (as *AssignStatement) statementNode() {}

// TokenLiteral : Returns the literal representation of the token
func (as *AssignStatement) TokenLiteral() string {
	return as.Token.Literal
}

func (as *AssignStatement) String() string {
	var out bytes.Buffer

	out.WriteString(as.Name.String())
	out.WriteString(" = ")

	if as.Value != nil {
		out.WriteString(as.Value.String())
	}

	return out.String()
}
