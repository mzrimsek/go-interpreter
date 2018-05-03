// Package ast : Statement and Expression node definitions used to build an Abstract Syntax Tree for Monkey programs
package ast

// Node : Generic AST node
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement : Generic AST statement node
type Statement interface {
	Node
	statementNode()
}

// Expression : Generic AST expression node
type Expression interface {
	Node
	expressionNode()
}
