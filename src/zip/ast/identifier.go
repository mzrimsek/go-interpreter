package ast

import (
	"zip/token"
)

// Identifier : Expression node representing an identifier literal
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral : Returns the literal representation of the token
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Value
}
