package ast

import (
	"bytes"
	"fmt"
	"interpreter_in_go/token"
)

//Node represents one element in our ast
type Node interface {
	TokenLiteral() string
	String() string
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
	} else {
		return ""
	}
}

//String representation of a set of statements that compose a program.
func (p *Root) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

//LetStatement represents a statement in the form of `let <identifier> = <expression>`
type LetStatement struct {
	Token token.Token
	Name  *IdentifierStatement
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

//String representation of a let statement. "let x = 1;"
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")
	if ls.Value != nil { //todo Remove this
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

//IdentifierStatement of statement
type IdentifierStatement struct {
	Token token.Token
	Value string
}

func (i *IdentifierStatement) expressionNode()      {}
func (i *IdentifierStatement) TokenLiteral() string { return i.Token.Literal }

//String representation of an identifier.
func (i *IdentifierStatement) String() string {
	return i.Value
}

//ReturnStatement represents a statement in the form of `return <expression>;`
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

//String representation of a return statement. "return x;"
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil { //todo Remove this
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

//ExpressionStatement represents an expression`
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}
func (es *ExpressionStatement) String() string {
	if es.Expression != nil { //todo Remove this
		return es.Expression.String()
	}
	return ""
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (i *IntegerLiteral) expressionNode() {}
func (i *IntegerLiteral) TokenLiteral() string {
	return i.Token.Literal
}
func (i *IntegerLiteral) String() string {
	return i.Token.Literal
}



type PrefixExpression struct {
	Token token.Token
	Operator string
	Right Expression
}
func (p *PrefixExpression) expressionNode() {}
func (p *PrefixExpression) TokenLiteral() string {
	return p.Token.Literal
}
func (p *PrefixExpression) String() string {
	return fmt.Sprintf("(%s%s)", p.Operator, p.Right.String())
}