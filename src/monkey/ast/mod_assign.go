package ast

import (
	"bytes"
	"monkey/token"
)

// ModAssignStatement : Statement node representing executing modulus to the existing object and assigning that new value
type ModAssignStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (mas *ModAssignStatement) statementNode() {}

// TokenLiteral : Returns the literal representation of the token
func (mas *ModAssignStatement) TokenLiteral() string {
	return mas.Token.Literal
}

func (mas *ModAssignStatement) String() string {
	var out bytes.Buffer

	out.WriteString(mas.Name.String())
	out.WriteString(" %= ")

	if mas.Value != nil {
		out.WriteString(mas.Value.String())
	}

	return out.String()
}
