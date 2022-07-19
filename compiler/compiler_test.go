package compiler

import (
	"github.com/ahmadrosid/yuk/lexer"
	"github.com/ahmadrosid/yuk/parser"
	"testing"
)

func compile(input string) string {
	lex := lexer.New(input)
	p := parser.New(lex)
	com := New(p.ParseProgram())
	return com.Generate()
}

func TestCompiler_Generate(t *testing.T) {
	tests := []struct {
		expected string
		input    string
	}{
		{"var some = 1", "var some = 1"},
		{"func main() {}", "func main() {}"},
		{"type Token struct {\na Some\nb string\n}", "struct Token(a Some, b string)"},
		{"package main", "package main"},
		{"import \"encoding/json\"", "import encoding/json"},
	}

	for _, tt := range tests {
		res := compile(tt.input)
		if res != tt.expected {
			t.Errorf("compiler error expected %q. got=%q", tt.expected, res)
		}
	}
}
