package ast

import (
	"github.com/mzrimsek/zip-lang/src/zip/token"
)

// IntegerLiteral : Expression node representing an integer literal
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}

// TokenLiteral : Returns the literal representation of the token
func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}
