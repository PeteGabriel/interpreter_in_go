package lexer

import (
	"interpreter_in_go/token"
)

type Lexer struct {
	input   string
	pos     int  // current position in input (points to current char)
	readPos int  // current reading position in input (after current char)
	char    byte // current char under examination
}

//NewLexer creates a new instance of Lexer
func NewLexer(src string) *Lexer {
	lx := &Lexer{input: src}
	lx.next()
	return lx
}

//NextToken retrieves the next token present in the lexer.
//Works like an iterator over a set of elements. Reads a token
//and moves forward.
func (l *Lexer) NextToken() *token.Token {
	var tkn *token.Token

	//skip whitespace
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.next()
	}

	switch l.char {
	case '=':
		tkn = token.NewToken(token.ASSIGN, l.char)
	case ';':
		tkn = token.NewToken(token.SEMICOLON, l.char)
	case '(':
		tkn = token.NewToken(token.LPAREN, l.char)
	case ')':
		tkn = token.NewToken(token.RPAREN, l.char)
	case ',':
		tkn = token.NewToken(token.COMMA, l.char)
	case '+':
		tkn = token.NewToken(token.PLUS, l.char)
	case '{':
		tkn = token.NewToken(token.LBRACE, l.char)
	case '}':
		tkn = token.NewToken(token.RBRACE, l.char)
	case 0:
		tkn = &token.Token{}
		tkn.Literal = ""
		tkn.Type = token.EOF
	default: //keywords and identifiers
		if isLetter(l.char) {
			tkn = &token.Token{}
			tkn.Literal = l.readIdentifier()
			tkn.Type = token.GetIdentifier(tkn.Literal)
			return tkn
		}else if isDigit(l.char){
			tkn = &token.Token{}
			tkn.Literal = l.readNumber(l.char)
			tkn.Type = token.INT
			return tkn
		}else {
			tkn = token.NewToken(token.ILLEGAL, l.char)
		}
	}

	//move forward the read position pointer
	l.next()
	return tkn
}

func (l *Lexer) next() {
	if l.readPos >= len(l.input) {
		l.char = 0
	}else {
		l.char = l.input[l.readPos]
	}
	l.pos = l.readPos
	l.readPos += 1
}

type predicate func(byte) bool

func (l *Lexer) read(pred predicate) string {
	pos := l.pos
	for pred(l.char) {
		l.next()
	}
	return l.input[pos:l.pos]
}

//read entire word
func (l *Lexer) readIdentifier() string {
	return l.read(isLetter)
}

//read entire digit
func (l *Lexer) readNumber(c byte) string {
	return l.read(isDigit)
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}