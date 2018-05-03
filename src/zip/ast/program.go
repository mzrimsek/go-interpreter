package ast

import (
	"bytes"
)

// Program : Top level construct containing all parsed AST nodes
type Program struct {
	Statements []Statement
}

// TokenLiteral : Returns the literal representations of the tokens for each of the program's statements
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}
