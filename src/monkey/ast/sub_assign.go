package ast

import (
	"bytes"
	"monkey/token"
)

// SubAssignStatement : Statement node representing subtracting from the existing object and assigning that new value
type SubAssignStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (sas *SubAssignStatement) statementNode() {}

// TokenLiteral : Returns the literal representation of the token
func (sas *SubAssignStatement) TokenLiteral() string {
	return sas.Token.Literal
}

func (sas *SubAssignStatement) String() string {
	var out bytes.Buffer

	out.WriteString(sas.Name.String())
	out.WriteString(" -= ")

	if sas.Value != nil {
		out.WriteString(sas.Value.String())
	}

	return out.String()
}
