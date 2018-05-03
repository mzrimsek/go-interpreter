package ast

import (
	"bytes"
	"monkey/token"
)

// MultAssignStatement : Statement node representing multiplying to the existing object and assigning that new value
type MultAssignStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (mas *MultAssignStatement) statementNode() {}

// TokenLiteral : Returns the literal representation of the token
func (mas *MultAssignStatement) TokenLiteral() string {
	return mas.Token.Literal
}

func (mas *MultAssignStatement) String() string {
	var out bytes.Buffer

	out.WriteString(mas.Name.String())
	out.WriteString(" *= ")

	if mas.Value != nil {
		out.WriteString(mas.Value.String())
	}

	return out.String()
}
