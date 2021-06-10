package parser

import (
	"interpreter_in_go/ast"
	"interpreter_in_go/lexer"
	"interpreter_in_go/token"
)

type Parser struct {
	lxr *lexer.Lexer
	curToken token.Token
	peekToken token.Token
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
		return nil
	}
}

func (p *Parser) ParseLetStatement() ast.Statement {
	//validate that variable name comes after 'LET'
	if p.curToken.Type == token.LET && p.peekToken.Type != token.IDENTIFIER {
		return nil
	}

	var stmt ast.Statement
	stmt = &ast.LetStatement{
		Token: p.curToken,
		Name:  &ast.Identifier{
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

func (p *Parser) ParseReturnStatement() ast.Statement {
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


func (p *Parser) ReadToken() {
	p.curToken = p.peekToken
	p.peekToken = *p.lxr.NextToken()
}

