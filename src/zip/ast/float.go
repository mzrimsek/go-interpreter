package ast

import "zip/token"

// FloatLiteral : Expression node representing an float literal
type FloatLiteral struct {
	Token token.Token
	Value float64
}

func (fl *FloatLiteral) expressionNode() {}

// TokenLiteral : Returns the literal representation of the token
func (fl *FloatLiteral) TokenLiteral() string {
	return fl.Token.Literal
}

func (fl *FloatLiteral) String() string {
	return fl.Token.Literal
}
