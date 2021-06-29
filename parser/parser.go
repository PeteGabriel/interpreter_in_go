package parser

import (
	"fmt"
	"interpreter_in_go/ast"
	"interpreter_in_go/lexer"
	"interpreter_in_go/token"
	"strconv"
)

//Parser struct represents our parser from te pov of our program.
//It contains a lexer to read tokens, an error list where parsing errors
//are logged into, a pointer to the current token and another to the next token.
//Also, depending on the expression found we apply a parsing function for
//infix expressions or prefix expressions.
type Parser struct {
	lxr    *lexer.Lexer
	errors []string

	curToken  token.Token
	peekToken token.Token

	//allo us to check if the appropriate map has a parsing function associated with the token type.
	prefixParseFns map[token.Type]prefixParseFn
	infixParseFns  map[token.Type]infixParseFn
}

//NewParser returns a new instance of the type Parser.
func NewParser(lxr *lexer.Lexer) *Parser {
	p := &Parser{
		lxr:            lxr,
		curToken:       token.Token{},
		peekToken:      token.Token{},
		errors:         make([]string, 10),
		infixParseFns:  make(map[token.Type]infixParseFn),
		prefixParseFns: make(map[token.Type]prefixParseFn),
	}

	//registration of parsing functions
	p.registerPrefix(token.IDENTIFIER, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)

	p.curToken = *p.lxr.NextToken()
	p.peekToken = *p.lxr.NextToken()
	return p
}

//ParseProgram returns the root of our program
//in the form of an AST.
func (p *Parser) ParseProgram() *ast.Root {
	root := &ast.Root{
		Statements: []ast.Statement{},
	}

	//while we do not reach the end of file
	for p.curToken.Type != token.EOF {
		stmt := p.ParseStatement()
		if stmt != nil {
			root.Statements = append(root.Statements, stmt)
		}
		//get another token from lexer
		p.ReadToken()
	}

	return root
}

//ParseStatement decides which kind of parsing method to
// apply based on the type of current token.
func (p *Parser) ParseStatement() ast.Statement {

	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

//parseLetStatement parses a statement of the type let.
//let x = 9;
//let myFn = sum;
func (p *Parser) parseLetStatement() *ast.LetStatement {
	//validate that variable name comes after 'LET'
	if p.curToken.Type == token.LET && p.peekToken.Type != token.IDENTIFIER {
		return nil
	}

	stmt := &ast.LetStatement{
		Token: p.curToken,
		Name: &ast.IdentifierStatement{
			Token: p.curToken,
			Value: p.curToken.Literal,
		},
		Value: nil,
	}

	//validate that '=' comes after variable name
	if p.curToken.Type == token.IDENTIFIER && p.peekToken.Type != token.ASSIGN {
		return nil
	}

	for p.curToken.Type != token.SEMICOLON {
		p.ReadToken()
	}

	return stmt
}

//parseReturnStatement parses a statement of the type return.
//return <something>;
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	//validate we have found a return statement
	if p.curToken.Type != token.RETURN {
		return nil
	}

	stmt := &ast.ReturnStatement{
		Token:       p.curToken,
		ReturnValue: nil,
	}

	p.ReadToken()
	for p.curToken.Literal != token.SEMICOLON {
		p.ReadToken()
	}

	return stmt
}

//parseExpressionStatement parses a statement of the type expression.
//5 + 5
//if (something)
func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}

	stmt.Expression = p.parseExpression(LOWEST)

	//if next token is "the end", read it.
	//having ';' is optional for something like 5 + 5 to work well.
	if p.peekToken.Literal == token.SEMICOLON {
		p.ReadToken()
	}

	return stmt
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.IdentifierStatement{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	il := &ast.IntegerLiteral{
		Token: p.curToken,
	}
	literal, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	il.Value = literal
	return il
}

func (p *Parser) ReadToken() {
	p.curToken = p.peekToken
	p.peekToken = *p.lxr.NextToken()
}

//This constants help the parser understand the rule of
//high precedence between operators.
const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(X)
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

func (p *Parser) registerPrefix(tokenType token.Type, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.Type, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		return nil
	}
	return prefix()
}
