package parser

import (
	"github.com/ahmadrosid/yuk/ast"
	"github.com/ahmadrosid/yuk/lexer"
	"testing"
)

func TestVarStatement(t *testing.T) {
	tests := []struct {
		input              string
		expectedIdentifier string
		expectedValue      interface{}
	}{
		{"var x = 5", "x", 5},
		{"var y = true", "y", true},
		{"var foo = x", "foo", "x"},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParseErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain 1 statement. god=%q", len(program.Statements))
		}

		stmt := program.Statements[0]
		if !testVarStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testVarStatement(t *testing.T, stmt ast.Statement, identifier string) bool {
	if stmt.TokenLiteral() != "var" {
		t.Errorf("stmt.TokenLiteral no 'var'. got-%q", stmt.TokenLiteral())
		return false
	}
	varStmt, ok := stmt.(*ast.VarStatement)
	if !ok {
		t.Errorf("stmt not *ast.VarStatemtnt. got=%q", stmt)
		return false
	}
	if varStmt.Name.Value != identifier {
		t.Errorf("varStmt.Value not '%s'. got '%s'", identifier, varStmt.Name.Value)
		return false
	}

	if varStmt.Name.TokenLiteral() != identifier {
		t.Errorf("varStmt.TokenLiteral not '%s'. got '%s'", identifier, varStmt.Name.TokenLiteral())
		return false
	}

	return true
}

func checkParseErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parse error: %q", msg)
	}
	t.FailNow()
}
