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
