package ast

import (
	"bytes"
	"monkey/token"
)

// DivAssignStatement : Statement node representing dividing into the existing object and assigning that new value
type DivAssignStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (das *DivAssignStatement) statementNode() {}

// TokenLiteral : Returns the literal representation of the token
func (das *DivAssignStatement) TokenLiteral() string {
	return das.Token.Literal
}

func (das *DivAssignStatement) String() string {
	var out bytes.Buffer

	out.WriteString(das.Name.String())
	out.WriteString(" /= ")

	if das.Value != nil {
		out.WriteString(das.Value.String())
	}

	return out.String()
}
