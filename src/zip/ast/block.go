package ast

import (
	"bytes"

	"github.com/mzrimsek/zip-lang/src/zip/token"
)

// BlockStatement : Statement node representing a block of statements
type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (bs *BlockStatement) statementNode() {}

// TokenLiteral : Returns the literal representation of the token
func (bs *BlockStatement) TokenLiteral() string {
	return bs.Token.Literal
}

func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}
