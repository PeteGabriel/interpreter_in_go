package ast

import (
	"bytes"
	"fmt"
	"interpreter_in_go/token"
	"strings"
)

//Node represents one element in our ast
type Node interface {
	TokenLiteral() string
	String() string
}

//Statement represents a statement in our interpreter.
//Something that does not return a value.
type Statement interface {
	Node
	statementNode()
}

//Expression represents an expression in our interpreter.
//Something that returns a value.
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

//IntegerLiteral can be for example, '-5'.
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

//PrefixExpression represents something like '!isInvalid'.
//It means it has an operator in the left side and an expression in the right side.
type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (p *PrefixExpression) expressionNode() {}
func (p *PrefixExpression) TokenLiteral() string {
	return p.Token.Literal
}
func (p *PrefixExpression) String() string {
	return fmt.Sprintf("(%s%s)", p.Operator, p.Right.String())
}

//InExpression represents something like 'counter++'.
//It means it has an expression in the left side and an operator in the right side.
type InfixExpression struct {
	Token       token.Token
	Left, Right Expression
	Operator    string
}

func (inf *InfixExpression) expressionNode() {}
func (inf *InfixExpression) TokenLiteral() string {
	return inf.Token.Literal
}
func (inf *InfixExpression) String() string {
	return fmt.Sprintf("(%s %s %s)", inf.Left.String(), inf.Operator, inf.Right.String())
}

//BooleanLiteral represents boolean values.
//true; let foo = false;
type BooleanLiteral struct {
	Token token.Token
	Value bool
}

func (b *BooleanLiteral) expressionNode() {}
func (b *BooleanLiteral) TokenLiteral() string {
	return b.Token.Literal
}
func (b *BooleanLiteral) String() string {
	return b.TokenLiteral()
}

type IfExpression struct {
	Token       token.Token //if
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ifEx *IfExpression) expressionNode() {}
func (ifEx *IfExpression) TokenLiteral() string {
	return ifEx.Token.Literal
}
func (ifEx *IfExpression) String() string {
	var out bytes.Buffer
	out.WriteString("if")
	out.WriteString(ifEx.Condition.String())
	out.WriteString(" ")
	out.WriteString(ifEx.Consequence.String())

	if ifEx.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ifEx.Alternative.String())
	}

	return out.String()
}

type BlockStatement struct {
	Token      token.Token //{
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


type FunctionLiteral struct {
	Token token.Token
	Parameters []*IdentifierStatement
	Body *BlockStatement
}

func (f *FunctionLiteral) statementNode() {}
func (f *FunctionLiteral) TokenLiteral() string {
	return f.Token.Literal
}
func (f *FunctionLiteral) String() string {
	var out bytes.Buffer
  params := []string{}
  for _, p := range f.Parameters{
		params = append(params, p.String())
	}

	out.WriteString(f.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(")")
	out.WriteString(f.Body.String())


	return out.String()
}