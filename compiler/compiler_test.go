package compiler

import (
	"github.com/ahmadrosid/yuk/lexer"
	"github.com/ahmadrosid/yuk/parser"
	"testing"
)

func compile(t *testing.T, input string) string {
	lex := lexer.New(input)
	p := parser.New(lex)
	com := New(p)
	res, errs := com.Generate()
	if errs != nil {
		for _, err := range errs {
			t.Errorf("%s\n", input)
			t.Fatalf("error: %q", err.Error())
		}
	}
	return res
}

func TestCompiler_Generate(t *testing.T) {
	tests := []struct {
		expected string
		input    string
	}{
		{"package main", "package main"},
		{"import \"encoding/json\"", "import encoding/json"},
		{"func main() {}", "func main() {}"},
		{"type Token struct {\na Some\nb string\n}", "struct Token(a Some, b string)"},
		{"type TokenType string", "type TokenType string"},
		{"var some = 1", "var some = 1"},
	}

	for _, tt := range tests {
		res := compile(t, tt.input)
		if res != tt.expected {
			t.Errorf("compiler error expected %q. got=%q", tt.expected, res)
		}
	}
}
