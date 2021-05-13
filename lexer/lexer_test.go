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
func TestNextTokenWithCode(t *testing.T) {
	is := is2.New(t)
	input := `let five = 5;
		let the_number_ten = 10;
		let add = fn(x, y) {
		  x + y;
		};
		let result = add(five, the_number_ten);
		`
	tests := []struct {
		expectedType token.Type
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "the_number_ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "the_number_ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lx := NewLexer(input)

	for _, tt := range tests {
		tkn := lx.NextToken()

		is.True(tkn.Type == tt.expectedType)
		is.Equal(tkn.Literal, tt.expectedLiteral)
	}

}