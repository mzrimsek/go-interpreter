package ast

import (
	"bytes"
	"zip/token"
)

// PostfixExpression : Expression node representing postfix expressions like incrementing and decrementing
type PostfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
}

func (pe *PostfixExpression) expressionNode() {}

// TokenLiteral : Returns the literal representation of the token
func (pe *PostfixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

func (pe *PostfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Left.String())
	out.WriteString(pe.Operator)
	out.WriteString(")")

	return out.String()
}
