package token

/*
The Monkey language has different types of tokens
that can be expressed as constants in our code.
*/
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//identifiers
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456

	//operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	EQUAL    = "=="
	NOT_EQUAL    = "!="
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"

	//delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	//keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "IF"
	ELSE     = "ELSE"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	RETURN   = "RETURN"
)

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

var keywords = map[string]Type{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
	"return": RETURN,
}

//GetIdentifier check if given word is an identifier or a keyword.
func GetIdentifier(ident string) Type {
	if t, ok := keywords[ident]; ok == true {
		return t
	}
	return IDENT
}
