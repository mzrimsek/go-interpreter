package ast

import (
	"bytes"
	"monkey/token"
)

// AddAssignStatement : Statement node representing adding to the existing object and assigning that new value
type AddAssignStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (aas *AddAssignStatement) statementNode() {}

// TokenLiteral : Returns the literal representation of the token
func (aas *AddAssignStatement) TokenLiteral() string {
	return aas.Token.Literal
}

func (aas *AddAssignStatement) String() string {
	var out bytes.Buffer

	out.WriteString(aas.Name.String())
	out.WriteString(" += ")

	if aas.Value != nil {
		out.WriteString(aas.Value.String())
	}

	return out.String()
}
