package parser

import (
	is2 "github.com/matryer/is"
	"interpreter_in_go/ast"
	"interpreter_in_go/lexer"
	"testing"
)

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
		{"foobar"},
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