package token

type TokenType string

const (
	ILLEGAL = "ILLEGAL" // represents tokens not supported
	EOF     = "EOF"     // represents end of file, lexer stops

	// Identifiers + literals
	IDENT = "IDENT" // functions
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

type Token struct {
	Type    TokenType
	Literal string
}
