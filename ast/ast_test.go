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
				Name: &IdentifierStatement{
					Token: token.Token{
						Type:    token.IDENTIFIER,
						Literal: "number_ten",
					},
					Value: "number_ten",
				},
				Value: &IdentifierStatement{
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

func TestString(t *testing.T) {
	is := is2.New(t)
	program := &Root{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &IdentifierStatement{
					Token: token.Token{Type: token.IDENTIFIER, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &IdentifierStatement{
					Token: token.Token{Type: token.IDENTIFIER, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	is.Equal("let myVar = anotherVar;", program.String())
}

func TestPrefixExpressionToString(t *testing.T){
	is := is2.New(t)
	program := &Root{
		Statements: []Statement{
			&ExpressionStatement{
				Token: token.Token{Type: token.BANG, Literal: "!"},
				Expression: &PrefixExpression{
					Token: token.Token{Type: token.BANG, Literal: "!"},
					Operator: "-",
					Right: &IntegerLiteral{
						Token: token.Token{Type: token.INT, Literal: "5"},
						Value: 5,
					},
				},
			},
		},
	}

	is.Equal("(-5)", program.String())
}

func TestInfixExpressionToString(t *testing.T){
	is := is2.New(t)
	program := &Root{
		Statements: []Statement{
			&ExpressionStatement{
				Token: token.Token{Type: token.EXPRESSION, Literal: ""},
				Expression: &InfixExpression{
					Token: token.Token{Type: token.EXPRESSION, Literal: ""},
					Operator: "+",
					Left: &IntegerLiteral{
						Token: token.Token{Type: token.INT, Literal: "1"},
						Value: 5,
					},
					Right: &IntegerLiteral{
						Token: token.Token{Type: token.INT, Literal: "2"},
						Value: 5,
					},
				},
			},
		},
	}

	is.Equal("(1 + 2)", program.String())
}

func TestBooleanLiteralToString(t *testing.T){
	is := is2.New(t)
	program := &Root{
		Statements: []Statement{
			&ExpressionStatement{
				Token: token.Token{Type: token.FALSE, Literal: "false"},
				Expression: &BooleanLiteral{
					Token: token.Token{Type: token.FALSE, Literal: "false"},
					Value: false,
				},
			},
		},
	}

	is.Equal("false", program.String())
}