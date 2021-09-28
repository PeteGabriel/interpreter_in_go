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

	registerParsingFns(p)

	p.curToken = *p.lxr.NextToken()
	p.peekToken = *p.lxr.NextToken()
	return p
}

func registerParsingFns(p *Parser) {
	p.registerPrefix(token.IDENTIFIER, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)

	p.registerPrefix(token.BANG, p.parsePrefixExpression)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)

	p.registerInfix(token.PLUS, p.parseInfixExpression)
	p.registerInfix(token.MINUS, p.parseInfixExpression)
	p.registerInfix(token.SLASH, p.parseInfixExpression)
	p.registerInfix(token.ASTERISK, p.parseInfixExpression)
	p.registerInfix(token.EQUAL, p.parseInfixExpression)
	p.registerInfix(token.NOT_EQUAL, p.parseInfixExpression)
	p.registerInfix(token.LT, p.parseInfixExpression)
	p.registerInfix(token.GT, p.parseInfixExpression)

	p.registerPrefix(token.FALSE, p.parseBooleanLiteral)
	p.registerPrefix(token.TRUE, p.parseBooleanLiteral)

	p.registerPrefix(token.LPAREN, p.parseGroupedExpression)

	p.registerPrefix(token.IF, p.parseIfExpression)

	p.registerPrefix(token.FunctionLiteral, p.parseFunctionExpression)
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
		root.Statements = append(root.Statements, stmt)

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
			Token: p.peekToken,
			Value: p.peekToken.Literal,
		},
		Value: nil,
	}

	//validate that '=' comes after variable name
	if p.curToken.Type == token.IDENTIFIER && p.peekToken.Type != token.ASSIGN {
		return nil
	}

	//get to the right part of the let statement
	for p.curToken.Type != token.ASSIGN {
		p.ReadToken()
	}
	p.ReadToken()

	//parse the right side
	stmt.Value = p.parseExpression(LOWEST)

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
		stmt.ReturnValue = p.parseExpression(LOWEST)
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

func (p *Parser) parseBooleanLiteral() ast.Expression {
	b := &ast.BooleanLiteral{
		Token: p.curToken,
	}

	if val, err := strconv.ParseBool(p.curToken.Literal); err == nil {
		b.Value = val
	} else {
		msg := fmt.Sprintf("could not parse %s as boolean", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	return b
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	prefix := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}

	p.ReadToken() //start what might be an expression

	prefix.Right = p.parseExpression(PREFIX)

	return prefix
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	exp := &ast.InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left,
	}

	precedence := p.currentPrecedence()
	p.ReadToken()
	exp.Right = p.parseExpression(precedence)

	return exp
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

var precedences = map[token.Type]int{
	token.EQUAL:     EQUALS,
	token.NOT_EQUAL: EQUALS,
	token.LT:        LESSGREATER,
	token.GT:        LESSGREATER,
	token.PLUS:      SUM,
	token.MINUS:     SUM,
	token.SLASH:     PRODUCT,
	token.ASTERISK:  PRODUCT,
}

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
	leftExpression := prefix()

	for !(p.peekToken.Type == token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExpression
		}
		p.ReadToken()
		leftExpression = infix(leftExpression)
	}

	return leftExpression
}

func (p *Parser) peekPrecedence() int {
	if pd, ok := precedences[p.peekToken.Type]; ok {
		return pd
	}
	return LOWEST
}

func (p *Parser) currentPrecedence() int {
	if pd, ok := precedences[p.curToken.Type]; ok {
		return pd
	}
	return LOWEST
}

func (p *Parser) parseGroupedExpression() ast.Expression {
	p.ReadToken()

	exp := p.parseExpression(LOWEST)

	if p.peekToken.Type != token.RPAREN {
		return nil
	}

	return exp
}

func (p *Parser) parseIfExpression() ast.Expression {
	ifExp := &ast.IfExpression{
		Token: p.curToken,
	}

	//validate correct syntax
	if p.peekToken.Type != token.LPAREN {
		return nil
	}
	p.ReadToken()
	ifExp.Condition = p.parseExpression(LOWEST)

	if p.peekToken.Type != token.RPAREN {
		return nil
	} else {
		p.ReadToken()
	}
	if p.peekToken.Type != token.LBRACE {
		return nil
	} else {
		p.ReadToken()
	}

	ifExp.Consequence = p.parseBlockStatement()

	//support the 'else' clause
	p.ReadToken()
	if p.peekToken.Type != token.ELSE {
		p.ReadToken()
		if p.peekToken.Type != token.LBRACE {
			return nil
		}

		ifExp.Alternative = p.parseBlockStatement()
	}

	return ifExp
}

// Parse expressions that resolve into a function. For example,
// fn <parameters> <block statement>.
// Functions can also be assigned as `let myFunction = fn(x, y) { return x + y; }`
func (p *Parser) parseFunctionExpression() *ast.Expression {
	funcExp := &ast.FunctionLiteral {
		Token: p.curToken
	}

	if p.peekToken != token.LPAREN {
		return nil
	}

	funcExp.Parameters = p.parseFunctionParameters()

	if p.peekToken != token.LBRACE {
		return nil
	}

	funcExp.Body = p.parseBlockStatement()

	return funcExp
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{
		Token: p.curToken,
	}
	block.Statements = []ast.Statement{}

	p.ReadToken()

	//parse clause until the end
	for p.curToken.Type != token.RBRACE {
		stmt := p.ParseStatement()
		block.Statements = append(block.Statements, stmt)
		p.ReadToken()
	}

	return block
}

func (p *Parser) parseFunctionParameters() []*IdentifierStatement {
	return nil
}