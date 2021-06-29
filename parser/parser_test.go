package parser

import (
	is2 "github.com/matryer/is"
	"interpreter_in_go/ast"
	"interpreter_in_go/lexer"
	"testing"
)


func TestReturnStatements(t *testing.T){
	is := is2.New(t)

	input := `
      return 10;
      return 5;
      return add(13);
    `
	l := lexer.NewLexer(input)
	p := NewParser(l)

	prog := p.ParseProgram()
	is.True(prog != nil)
	is.Equal(len(prog.Statements), 3)

	tests := []struct {
		expectedIdentifier string
	}{
		{"return"},
		{"10"},
		{"return"},
		{"5"},
		{"return"},
	}

	for i, ts := range tests {
		stmt := prog.Statements[i]

		is.True(stmt.TokenLiteral() == "return")

		//convert stmt to specific let stmt
		letStmt, ok := stmt.(*ast.LetStatement)
		is.Equal(ok, true)

		is.True(letStmt.Name.Value == ts.expectedIdentifier)
		is.True(letStmt.Name.TokenLiteral() == ts.expectedIdentifier)
	}
}

func TestParseLetStatements(t *testing.T){
	is := is2.New(t)

	input := `
      let x = 5;
      let y = 10;
      let bar = "foo";
    `

	l := lexer.NewLexer(input)
	p := NewParser(l)

	prog := p.ParseProgram()
	is.True(prog != nil)
	is.True(len(prog.Statements) == 3)

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foo"},
	}

	for i, ts := range tests {
		stmt := prog.Statements[i]
		testLetStatement(is, stmt, ts.expectedIdentifier)
	}

}

func testLetStatement(is *is2.I, stmt ast.Statement, expectedIdent string) {
	is.True(stmt.TokenLiteral() == "let")

	//convert stmt to specific let stmt
	letStmt, ok := stmt.(*ast.LetStatement)
	is.Equal(ok, true)

	is.True(letStmt.Name.Value == expectedIdent)
	is.True(letStmt.Name.TokenLiteral() == expectedIdent)
}

func TestStringOfStatement(t *testing.T){
	is := is2.New(t)
	input := `
      let x = 5;
      return x;
    `
	l := lexer.NewLexer(input)
	p := NewParser(l)
	prog := p.ParseProgram()

	expected := "let x = 5; return x;"
	is.Equal(expected, prog.String())
}

func TestIdentifierExpression(t *testing.T){
	is := is2.New(t)
	input := `foobar;`
	l := lexer.NewLexer(input)
	p := NewParser(l)
	prog := p.ParseProgram()

	//check the len of statements
	is.True(len(prog.Statements) == 1)
	//check if the statement found is actually and expression
	ident, ok := prog.Statements[0].(*ast.ExpressionStatement)
	is.True(ok)
	exp := ident.Expression.(*ast.IdentifierStatement)
	is.Equal("foobar", ident.TokenLiteral())
	is.Equal("foobar", exp.Value)

}