package lexer

import (
	is2 "github.com/matryer/is"
	"interpreter_in_go/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	is := is2.New(t)

	input := `=+(){},;`

	tests := []token.Token {
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.PLUS, Literal: "+"},
		{Type: token.LPAREN, Literal: "("},
		{Type: token.RPAREN, Literal: ")"},
		{Type: token.LBRACE, Literal: "{"},
		{Type: token.RBRACE, Literal: "}"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.EOF, Literal: ""},
	}

	lx := NewLexer(input)

	for _, tt := range tests {
		tkn := lx.NextToken()

		is.True(tkn.Type == tt.Type)
		is.Equal(tkn.Literal, tt.Literal)
	}
}
