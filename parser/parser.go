package parser

import (
	"interpreter_in_go/ast"
	"interpreter_in_go/lexer"
	"interpreter_in_go/token"
)

type Parser struct {
	lxr *lexer.Lexer
	errors []string

	curToken token.Token
	peekToken token.Token

	//allo us to check if the appropriate map has a parsing function associated with the token type.
	prefixParseFn map[token.Token]prefixParseFn
	infixParseFn map[token.Token]infixParseFn
}

func NewParser(lxr *lexer.Lexer) *Parser {
	p := &Parser{
		lxr:       lxr,
		curToken:  token.Token{},
		peekToken: token.Token{},
	}

	p.curToken = p.peekToken
	p.peekToken = *p.lxr.NextToken()
	return p
}

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

func (p *Parser) ParseStatement() ast.Statement {

	switch p.curToken.Type {
	case token.LET:
		return p.ParseLetStatement()
	case token.RETURN:
		return p.ParseReturnStatement()
	default:
		return p.ParseExpressionStatement()
	}
}

func (p *Parser) ParseLetStatement() *ast.LetStatement {
	//validate that variable name comes after 'LET'
	if p.curToken.Type == token.LET && p.peekToken.Type != token.IDENTIFIER {
		return nil
	}

	stmt := &ast.LetStatement{
		Token: p.curToken,
		Name:  &ast.IdentifierStatement{
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

func (p *Parser) ParseReturnStatement() *ast.ReturnStatement {
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

func (p *Parser) ParseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}

	stmt.Expression = p.parseExpression(LOWEST)

	//if next token is "the end", read it.
	if p.peekToken.Literal == token.SEMICOLON {
		p.ReadToken()
	}

	return stmt
}

func (p *Parser) ReadToken() {
	p.curToken = p.peekToken
	p.peekToken = *p.lxr.NextToken()
}



type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression)  ast.Expression
)

func (p *Parser) registerPrefix(tokenType token.Token, fn prefixParseFn){
	p.prefixParseFn[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.Token, fn infixParseFn){
	p.infixParseFn[tokenType] = fn
}

func (p *Parser) parseExpression(t string){

}