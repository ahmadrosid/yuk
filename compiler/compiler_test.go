package compiler

import (
	"testing"

	"github.com/ahmadrosid/yuk/lexer"
	"github.com/ahmadrosid/yuk/parser"
)

func compile(t *testing.T, input string) string {
	lex := lexer.New(input)
	p := parser.New(lex)
	com := New(p)
	res, errs := com.Generate()
	if len(errs) > 0 {
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
		{"type Token struct {\na Some\nb string\n}", "struct Token(a Some, b string)"},
		{"var some = 1", "var some = 1"},
		{"var some = \"Ahmad Rosid\"", "var some = \"Ahmad Rosid\""},
		{"func ReturnFunc() string {\nreturn \"hello\"\n}", "func ReturnFunc() string {return \"hello\"}"},
		{"switch '=' {case '=': {\nvar you = \"me\"\n}\n}", "switch '=' {\n'=' => {\nvar you = \"me\"\n}\n}"},
		{"var data = map[string]interface{}", "var data = map(string, interface)"},
		{"var data = map[string]interface{}{\n\"a\":1,\n}", "var data = map(string, interface){\"a\": 1}"},
		{"var data = map[string]interface{}{\n\"c\":3,\n\"d\":2,\n}", "var data = map(string, interface){\"c\": 3, \"d\": 2}"},
		{"type Post struct {\nTitle string `json:\"input\"`\nCreatedBy User `json:\"created_by\"`\nCreatedAt Date `json:\"created_at\"`\nUpdatedAt Date\n}", "struct Post (\n\tTitle        string      `json:\"input\"`\n\tCreatedBy    User        `json:\"created_by\"`\n\tCreatedAt    Date  `json:\"created_at\"`\n\tUpdatedAt    Date\n)\n"},
	}

	for _, tt := range tests {
		res := compile(t, tt.input)
		expected := tt.expected + "\n"
		if res != expected {
			t.Errorf("compiler error \nexpected='%s'\ngot='%s'", expected, res)
		}
	}
}
