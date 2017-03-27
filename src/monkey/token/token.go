// Package token : Token definitions used to tokenize program input
package token

// TokenType : String representation of each type of Token
type TokenType string

const (
	// ILLEGAL : Illegal TokenType
	ILLEGAL = "ILLEGAL"
	// EOF : End of file TokenType
	EOF = "EOF"

	// IDENT : Identifier literal TokenType
	IDENT = "IDENT"
	// INT : Integer literal TokenType
	INT = "INT"

	// ASSIGN : Assignment operator TokenType
	ASSIGN = "="
	// PLUS : Addition operator TokenType
	PLUS = "+"
	// MINUS : Subtraction and negative number operator TokenType
	MINUS = "-"
	// BANG : Negation operator TokenType
	BANG = "!"
	// ASTERISK : Multiplication operator TokenType
	ASTERISK = "*"
	// SLASH : Division operator TokenType
	SLASH = "/"
	// LT : Less than operator TokenType
	LT = "<"
	// GT : Greater than operator TokenType
	GT = ">"
	// EQ : Equality operator TokenType
	EQ = "=="
	// NOT_EQ : Negated equality operator TokenType
	NOT_EQ = "!="

	// COMMA : Comma delimiter TokenType
	COMMA = ","
	// SEMICOLON : Semicolon delimiter TokenType
	SEMICOLON = ";"
	// LPAREN : Left parenthesis delimiter TokenType
	LPAREN = "("
	// RPAREN : Right parenthesis delimiter TokenType
	RPAREN = ")"
	// LBRACE : Left brace delimiter TokenType
	LBRACE = "{"
	// RBRACE : Right brace delimiter TokenType
	RBRACE = "}"

	// FUNCTION : Function expression keyword TokenType
	FUNCTION = "FUNCTION"
	// LET : Let statement keyword TokenType
	LET = "LET"
	// TRUE : True boolean value keyword TokenType
	TRUE = "TRUE"
	// FALSE : False boolean value keyword TokenType
	FALSE = "FALSE"
	// IF : If expression keyword TokenType
	IF = "IF"
	// ELSE : Else expression keyword TokenType
	ELSE = "ELSE"
	// RETURN : Return statement keyword TokenType
	RETURN = "RETURN"
)

// Token : Defines the type and literal representation for the tokens to be used in program analysis
type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdent : Compares identifier input against list of keywords to return proper TokenType
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
