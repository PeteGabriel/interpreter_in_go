package token

import (
	"interpreter_in_go/lexer"
	"testing"

	is2 "github.com/matryer/is"
)

func TestNewToken(t *testing.T) {
	is := is2.New(t)
	intToken := NewToken(lexer.INT, "1")

	is.True(intToken.Type == lexer.INT)
	is.True(intToken.Literal == "1")
}
