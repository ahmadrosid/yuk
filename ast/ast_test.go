package ast

import (
	"github.com/ahmadrosid/yuk/token"
	"testing"
)

func TestVarStatement_String(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&VarStatement{
				Token: token.Token{Type: token.VAR, Literal: "var"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "five"},
					Value: "five",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "other"},
					Value: "5",
				},
			},
		},
	}

	if program.String() != "var five = 5" {
		t.Errorf("program.String() wrong, got=%q", program.String())
	}
}

func TestReturnStatement_String(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&ReturnStatement{
				Token: token.Token{Type: token.RETURN, Literal: "return"},
				ReturnValue: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "five"},
					Value: "five",
				},
			},
		},
	}

	if program.String() != "return five" {
		t.Errorf("program.String() wrong, got=%q", program.String())
	}
}

func TestCaseStatement_String(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&SwitchStatement{
				Token: token.Token{Type: token.SWITCH, Literal: "switch"},
				Input: token.Token{Type: token.CHAR, Literal: "="},
				Case: []*CaseLiteral{
					{
						Token: token.Token{Type: token.CHAR, Literal: "="},
						Body: &BlockStatement{
							Token: token.Token{Type: token.LBRACE, Literal: "{"},
							Statements: []Statement{
								&VarStatement{
									Token: token.Token{Type: token.VAR, Literal: "var"},
									Name: &Identifier{
										Token: token.Token{Type: token.IDENT, Literal: "five"},
										Value: "five",
									},
									Value: &Identifier{
										Token: token.Token{Type: token.IDENT, Literal: "other"},
										Value: "5",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	if program.String() != "switch '=' {case '=': {\nvar five = 5\n}\n}" {
		t.Errorf("program.String() wrong, got=%q", program.String())
	}
}

func TestStructStatement_String(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&StructStatement{
				Token: token.Token{Type: token.STRUCT, Literal: "struct"},
				Name:  token.Token{Type: token.RETURN, Literal: "Token"},
				Attributes: []*TypeStatement{
					{
						Name: token.Token{Type: token.IDENT, Literal: "Type"},
						Type: token.Token{Type: token.IDENT, Literal: "TypeToken"},
					},
					{
						Name: token.Token{Type: token.STRING, Literal: "Literal"},
						Type: token.Token{Type: token.IDENT, Literal: "string"},
					},
				},
			},
		},
	}

	expected := `type Token struct {
Type TypeToken
Literal string
}`
	if program.String() != expected {
		t.Errorf("program.String() wrong \nexpected='%q' \ngot='%q'", expected, program.String())
	}
}
