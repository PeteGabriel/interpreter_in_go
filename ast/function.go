package ast

import (
	"bytes"
	"interpreter_in_go/token"
	"strings"
)

type FunctionLiteral struct {
	Token      token.Token
	Parameters []*IdentifierStatement
	Body       *BlockStatement
}

func (f *FunctionLiteral) expressionNode() {}

func (f *FunctionLiteral) TokenLiteral() string {
	return f.Token.Literal
}

func (f *FunctionLiteral) String() string {
	var out bytes.Buffer
	params := []string{}

	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(f.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(")")
	out.WriteString(f.Body.String())

	return out.String()
}

type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (b *BlockStatement) statementNode() {}
func (b *BlockStatement) TokenLiteral() string {
	return b.Token.Literal
}
func (b *BlockStatement) String() string {
	var out bytes.Buffer

	for _, stm := range b.Statements {
		out.WriteString(stm.String())
	}

	return out.String()
}
