package lexer

import (
	"testing"

	"github.com/ahmadrosid/yuk/token"
)

func TestNextToken(t *testing.T) {
	input := `package main
type TokenType string
struct Token(Type TypeToken, Literal string)
func main() {
	var five = 5
}`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.PACKAGE, "package"},
		{token.IDENT, "main"},
		{token.TYPE, "type"},
		{token.IDENT, "TokenType"},
		{token.STRING, "string"},
		{token.STRUCT, "struct"},
		{token.IDENT, "Token"},
		{token.LPAREN, "("},
		{token.IDENT, "Type"},
		{token.IDENT, "TypeToken"},
		{token.COMMA, ","},
		{token.IDENT, "Literal"},
		{token.STRING, "string"},
		{token.RPAREN, ")"},
		{token.FUNCTION, "func"},
		{token.IDENT, "main"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.VAR, "var"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.RBRACE, "}"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - token type wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("test[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
