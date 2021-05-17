package token

import (
	"testing"

	is2 "github.com/matryer/is"
)

func TestNewToken(t *testing.T) {
	is := is2.New(t)
	intToken := NewToken(INT, '1')

	is.True(intToken.Type == INT)
	is.True(intToken.Literal == "1")
}
