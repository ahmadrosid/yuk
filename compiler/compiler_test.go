package compiler

import (
	"github.com/ahmadrosid/yuk/lexer"
	"github.com/ahmadrosid/yuk/parser"
	"testing"
)

func TestCompiler_GenerateVar(t *testing.T) {
	expected := "var some = 1"
	input := "var some = 1"
	lex := lexer.New(input)
	p := parser.New(lex)
	com := New(p.ParseProgram())
	res := com.Generate()
	if res != expected {
		t.Errorf("compiler error expected %q. got=%q", expected, res)
	}
}

func TestCompiler_GenerateFunc(t *testing.T) {
	expected := "func main() {}"
	input := "func main() {}"
	lex := lexer.New(input)
	p := parser.New(lex)
	program := p.ParseProgram()
	com := New(program)
	res := com.Generate()
	if res != expected {
		t.Errorf("compiler error expected %q. got=%q", expected, res)
	}
}
