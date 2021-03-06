// Package token : Token definitions used when tokenizing program input
package token

// TokenType : String representation of each type of Token
type TokenType string

// TokenTypes
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Literals
	IDENT  = "IDENT"
	INT    = "INT"
	FLOAT  = "FLOAT"
	STRING = "STRING"
	CHAR   = "CHAR"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	PERCENT  = "%"
	POWER    = "**"

	LT  = "<"
	GT  = ">"
	LTE = "<="
	GTE = ">="

	EQ     = "=="
	NOT_EQ = "!="

	AND = "&&"
	OR  = "||"

	INCREMENT = "++"
	DECREMENT = "--"

	ADD_ASSIGN  = "+="
	SUB_ASSIGN  = "-="
	MULT_ASSIGN = "*="
	DIV_ASSIGN  = "/="
	MOD_ASSIGN  = "%="
	POW_ASSIGN  = "**="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	WHILE    = "WHILE"
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
	"while":  WHILE,
}

// LookupIdent : Compares identifier input against list of keywords to return proper TokenType
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
