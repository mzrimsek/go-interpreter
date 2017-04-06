package ast

import (
	"bytes"
	"monkey/token"
)

// WhileExpression : Expression node representing a while loop
type WhileExpression struct {
	Token     token.Token
	Condition Expression
	Block     *BlockStatement
}

func (we *WhileExpression) expressionNode() {}

// TokenLiteral : Returns the literal representation of the token
func (we *WhileExpression) TokenLiteral() string {
	return we.Token.Literal
}

func (we *WhileExpression) String() string {
	var out bytes.Buffer

	out.WriteString("while")
	out.WriteString(we.Condition.String())
	out.WriteString(" ")
	out.WriteString(we.Block.String())

	return out.String()
}
