package ast

import (
	"bytes"
	"zip/token"
)

// ShortcutAssignStatement : Statement node representing doing an operator to the existing object and assigning that new value
type ShortcutAssignStatement struct {
	Token    token.Token
	Name     *Identifier
	Value    Expression
	Operator string
}

func (sas *ShortcutAssignStatement) statementNode() {}

// TokenLiteral : Returns the literal representation of the token
func (sas *ShortcutAssignStatement) TokenLiteral() string {
	return sas.Token.Literal
}

func (sas *ShortcutAssignStatement) String() string {
	var out bytes.Buffer

	out.WriteString(sas.Name.String())
	out.WriteString(" ")
	out.WriteString(sas.Operator)
	out.WriteString("=")
	out.WriteString(" ")

	if sas.Value != nil {
		out.WriteString(sas.Value.String())
	}

	return out.String()
}
