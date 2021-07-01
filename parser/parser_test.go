package parser

import (
	"interpreter_in_go/ast"
	"interpreter_in_go/lexer"
	"testing"

	is2 "github.com/matryer/is"
)

func TestReturnStatements(t *testing.T) {
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

func TestParseLetStatements(t *testing.T) {
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

func TestStringOfStatement(t *testing.T) {
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

func TestParseExpression(t *testing.T) {
	is := is2.New(t)
	input := `
	  foobar;	
	`
	l := lexer.NewLexer(input)
	p := NewParser(l)
	prog := p.ParseProgram()

	//check the len of statements
	is.True(len(prog.Statements) == 1)
	//check that first statement found is actually an expression
	ident, ok := prog.Statements[0].(*ast.ExpressionStatement)
	is.True(ok)
	exp := ident.Expression.(*ast.IdentifierStatement)
	is.Equal("foobar", ident.TokenLiteral())
	is.Equal("foobar", exp.Value)

}

func TestParseIntegerLiteral(t *testing.T) {
	is := is2.New(t)
	input := `
	  5;	
	`
	l := lexer.NewLexer(input)
	p := NewParser(l)
	prog := p.ParseProgram()

	//check the len of statements
	is.True(len(prog.Statements) == 1)

	//check that first statement found is actually an integer literal
	intLiteralExp, ok := prog.Statements[0].(*ast.ExpressionStatement)
	is.True(ok)
	intLiteral := intLiteralExp.Expression.(*ast.IntegerLiteral)
	is.Equal("5", intLiteral.TokenLiteral())
	is.Equal(int64(5), intLiteral.Value)
}

func TestParsingPrefixExpressions(t *testing.T) {
	is := is2.New(t)
	prefixTests := []struct {
		input        string
		operator     string
		integerValue int64
	}{
		{"!5;", "!", 5},
		{"-15;", "-", 15},
	}

	for _, tt := range prefixTests {
		l := lexer.NewLexer(tt.input)
		p := NewParser(l)
		prog := p.ParseProgram()

		is.True(len(prog.Statements) == 1)

		exp, ok := prog.Statements[0].(*ast.ExpressionStatement)
		is.True(ok)
		prefix := exp.Expression.(*ast.PrefixExpression)
		intVal, ok := prefix.Right.(*ast.IntegerLiteral)
		is.True(ok)
		is.Equal(tt.integerValue, intVal.Value)
		is.Equal(tt.operator, prefix.Operator)
	}
}

func TestParsingInfixExpressions(t *testing.T) {
	is := is2.New(t)
	prefixTests := []struct {
		input string
    leftValue int64
    operator string
    rightValue int64
	}{
		{"5 + 5;", 5, "+", 5},
		{"5 - 5;", 5, "-", 5},
		{"5 * 5;", 5, "*", 5},
		{"5 / 5;", 5, "/", 5},
		{"5 > 5;", 5, ">", 5},
		{"5 < 5;", 5, "<", 5},
		{"5 == 5;", 5, "==", 5},
		{"5 != 5;", 5, "!=", 5},
	}

	for _, tt := range prefixTests {
		l := lexer.NewLexer(tt.input)
		p := NewParser(l)
		prog := p.ParseProgram()

		is.True(len(prog.Statements) == 1)

		exp, ok := prog.Statements[0].(*ast.ExpressionStatement)
		is.True(ok)
		prefix := exp.Expression.(*ast.InfixExpression)
		leftVal, ok := prefix.Left.(*ast.IntegerLiteral)
		is.True(ok)
		rightVal, ok := prefix.Right.(*ast.IntegerLiteral)
		is.True(ok)
		is.Equal(tt.leftValue, leftVal.Value)
		is.Equal(tt.rightValue, rightVal.Value)
		is.Equal(tt.operator, prefix.Operator)
	}
}

func TestBooleanLiteralParsing(t *testing.T){
	is := is2.New(t)
	input := `let isMonday = false`
	l := lexer.NewLexer(input)
	p := NewParser(l)
	prog := p.ParseProgram()

	is.True(len(prog.Statements) == 1)
	
	letStmt, ok := prog.Statements[0].(*ast.LetStatement)
	is.True(ok)
	is.Equal(letStmt.Name.Value, "isMonday")
	is.Equal(letStmt.Value.String(), "false")
}