package token

type Type string

//Token represents one token passed to our lexer.
//It has a type to distinguish tokens and also a value.
type Token struct {
	Type    Type
	Literal string
}

//NewToken creates a new instance of type Token
func NewToken(t Type, l byte) *Token {
	return &Token{
		Type:    t,
		Literal: string(l),
	}
}


/*
The Monkey language has different types of tokens
that can be expressed as constants in our code.
*/
const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	//identifiers
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT = "INT" // 1343456

	//operators
	ASSIGN = "="
	PLUS = "+"
	MINUS = "-"

	//delimiters
	COMMA = ","
	SEMICOLON = ";"
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//keywords
	FUNCTION = "FUNCTION"
)