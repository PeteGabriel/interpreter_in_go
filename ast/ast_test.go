package ast

import (
	is2 "github.com/matryer/is"
	"interpreter_in_go/token"
	"testing"
)

func TestAstCreation(t *testing.T) {
	is := is2.New(t)

	root := Root{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{
					Type:    token.LET,
					Literal: "let",
				},
				Name: &Identifier{
					Token: token.Token{
						Type:    token.IDENTIFIER,
						Literal: "number_ten",
					},
					Value: "number_ten",
				},
				Value: &Identifier{
					Token: token.Token{
						Type:    token.INT,
						Literal: "10",
					},
					Value: "10",
				},
			},
		},
	}
	is.True(root.Statements[0].TokenLiteral() == "let")
}