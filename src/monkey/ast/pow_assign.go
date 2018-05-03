package ast

import (
	"bytes"
	"monkey/token"
)

// PowAssignStatement : Statement node representing putting to the existing object to a power and assigning that new value
type PowAssignStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (pas *PowAssignStatement) statementNode() {}

// TokenLiteral : Returns the literal representation of the token
func (pas *PowAssignStatement) TokenLiteral() string {
	return pas.Token.Literal
}

func (pas *PowAssignStatement) String() string {
	var out bytes.Buffer

	out.WriteString(pas.Name.String())
	out.WriteString(" **= ")

	if pas.Value != nil {
		out.WriteString(pas.Value.String())
	}

	return out.String()
}
