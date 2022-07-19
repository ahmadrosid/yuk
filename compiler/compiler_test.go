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
		{"import \"encoding/json\"", "import \"encoding/json\""},
		{"func main() {}", "func main() {}"},
		{"type TokenType string", "type TokenType string"},
		{"type Token struct {\n\ta Some\n\tb string\n}", "struct Token(a Some, b string)"},
		{"var some = 1", "var some = 1"},
		{"var some = \"Ahmad Rosid\"", "var some = \"Ahmad Rosid\""},
		{"func ReturnFunc() string {\n\treturn \"hello\"\n}", "func ReturnFunc() string {\treturn \"hello\"}"},
	}

	for _, tt := range tests {
		res := compile(t, tt.input)
		expected := tt.expected + "\n"
		if res != expected {
			t.Errorf("compiler error \nexpected=%q\ngot=%q", expected, res)
		}
	}
}
