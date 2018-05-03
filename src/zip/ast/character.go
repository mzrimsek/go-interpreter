package ast

import (
	"zip/token"
)

// CharacterLiteral : Expression node representing a character literal
type CharacterLiteral struct {
	Token token.Token
	Value byte
}

func (cl *CharacterLiteral) expressionNode() {}

// TokenLiteral : Returns the literal representation of the token
func (cl *CharacterLiteral) TokenLiteral() string {
	return cl.Token.Literal
}

func (cl *CharacterLiteral) String() string {
	return cl.Token.Literal
}
