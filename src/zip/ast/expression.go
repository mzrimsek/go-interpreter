package ast

import (
	"github.com/mzrimsek/zip-lang/src/zip/token"
)

// ExpressionStatement : Statement node representing a high-level expression
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

// TokenLiteral : Returns the literal representation of the token
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
