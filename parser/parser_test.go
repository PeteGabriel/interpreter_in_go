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
			return 13;
			return 1;
      return add(13);
    `
	l := lexer.NewLexer(input)
	p := NewParser(l)

	prog := p.ParseProgram()
	is.True(prog != nil)
	is.Equal(len(prog.Statements), 3)

	tests := []struct {
		expectedToken string
		expectedIdentifier string
	}{
		{"return", "13"},
		{"return", "1"},
		{"return", "add(13)"},
	}

	for i, ts := range tests {
		stmt := prog.Statements[i]

		is.True(stmt.TokenLiteral() == "return")

		//convert stmt to specific let stmt
		letStmt, ok := stmt.(*ast.ReturnStatement)
		is.Equal(ok, true)

		is.True(letStmt.Token.Literal == ts.expectedToken)
		is.True(letStmt.ReturnValue.String() == ts.expectedIdentifier)
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
	infixTests := []struct {
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

	for _, tt := range infixTests {
		l := lexer.NewLexer(tt.input)
		p := NewParser(l)
		prog := p.ParseProgram()

		is.True(len(prog.Statements) == 1)

		exp, ok := prog.Statements[0].(*ast.ExpressionStatement)
		is.True(ok)
		infix := exp.Expression.(*ast.InfixExpression)
		leftVal, ok := infix.Left.(*ast.IntegerLiteral)
		is.True(ok)
		rightVal, ok := infix.Right.(*ast.IntegerLiteral)
		is.True(ok)
		is.Equal(tt.leftValue, leftVal.Value)
		is.Equal(tt.rightValue, rightVal.Value)
		is.Equal(tt.operator, infix.Operator)
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

func TestInfixBooleanExpressions(t *testing.T){
	is := is2.New(t)
	boolExps := []struct{
		input string
		left bool
		operator string
		right bool
	}{
		{"false == true", false, "==", true},
		{"false == false", false, "==", false},
		{"false != true", false, "!=", true},
		{"false != false", false, "!=", false},
	}

	for _, t := range boolExps {
		l := lexer.NewLexer(t.input)
		p := NewParser(l)
		prog := p.ParseProgram()

		is.True(len(prog.Statements) == 1)
		exp, ok := prog.Statements[0].(*ast.ExpressionStatement)
		is.True(ok)
		infix := exp.Expression.(*ast.InfixExpression)

    leftVal, ok := infix.Left.(*ast.BooleanLiteral)
		is.True(ok)
		rightVal, ok := infix.Right.(*ast.BooleanLiteral)
		is.True(ok)
		is.Equal(t.left, leftVal.Value)
		is.Equal(t.right, rightVal.Value)
		is.Equal(t.operator, infix.Operator)
	}
}


func TestPrefixBooleanExpressions(t * testing.T){
	is := is2.New(t)
	boolExps := []struct{
		input string
		operator string
		value interface{}
	}{
		{"!true;", "!", true},
    {"!false;", "!", false},
	}

	for _, t := range boolExps {
		l := lexer.NewLexer(t.input)
		p := NewParser(l)
		prog := p.ParseProgram()

		is.True(len(prog.Statements) == 1)
		exp, ok := prog.Statements[0].(*ast.ExpressionStatement)
		is.True(ok)
		
		prefix := exp.Expression.(*ast.PrefixExpression)
		prefixValue := prefix.Right.(*ast.BooleanLiteral)
    
		
		is.Equal(t.value, prefixValue.Value)
		is.Equal(t.operator, prefix.Operator)
	}
}


func  TestOperatorPrecedenceParsing(t * testing.T){
	is := is2.New(t)
	exps := []struct{
		input string
		expected string
	}{
		{
			"3 + 4; -5 * 5",
			"(3 + 4)((-5) * 5)",
		},
		{
			"2 / (5 + 5)",
      "(2 / (5 + 5))",
		},
		{
			"-(5 + 5)",
			"(-(5 + 5))",
		},
		{
			"!(true == true)",
			"(!(true == true))",
		},
	}

	for _, t := range exps {
		l := lexer.NewLexer(t.input)
		p := NewParser(l)
		prog := p.ParseProgram()

		parsedExp := prog.String()				
		is.Equal(t.expected, parsedExp)
	}
}


func TestIfExpressionParsing(t *testing.T){
	input := `if (x < y) { x }`
	is := is2.New(t)

	l := lexer.NewLexer(input)
	p := NewParser(l)
	prog := p.ParseProgram()

	exp := prog.Statements[0].(*ast.ExpressionStatement)
	ifExp := exp.Expression.(*ast.IfExpression)
	is.Equal(ifExp.Token.Literal, "if")
	is.Equal(ifExp.Condition.String(), "(x < y)")
	is.Equal(ifExp.Consequence.String(), "x")
}

func TestFunctionLiteralParsing(t *testing.T){
	is := is2.New(t)
	input := `fn(x, y) { x + y; }`
	l := lexer.NewLexer(input)
	p := NewParser(l)
	program := p.ParseProgram()

	
	is.Equal(len(program.Statements), 1)
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	is.True(ok)

	function, ok := stmt.Expression.(*ast.FunctionLiteral)
	is.True(ok)

	is.Equal(len(function.Parameters), 2)
	is.Equal(len(function.Body.Statements), 1)
}