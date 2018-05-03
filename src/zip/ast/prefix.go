package ast

import (
	"bytes"

	"github.com/mzrimsek/zip-lang/src/zip/token"
)

// PrefixExpression : Expression node representing prefix expressions like negative numbers, negation, and prefix incrementing/decrementing
type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode() {}

// TokenLiteral : Returns the literal representation of the token
func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}
