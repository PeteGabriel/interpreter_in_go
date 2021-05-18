package ast

import "interpreter_in_go/token"

//Node represents one element in our ast
type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

//Root of every ast out parser produces.
type Root struct {
	Statements []Statement
}

func (p *Root) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}else {
		return ""
	}
}

//LetStatement represents a statement in the form of `let <identifier> = <expression>`
type LetStatement struct {
	Token token.Token
	Name *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode(){}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

//Identifier of statement
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }