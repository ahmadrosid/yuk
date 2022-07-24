package token

type TokenType string

const (
	ILLEGAL  = "ILLEGAL"
	EOF      = "EOF"
	NEW_LINE = "\n"

	// Identifiers + literals
	IDENT  = "IDENT"
	INT    = "INT"
	STRING = "STRING"
	CHAR   = "CHAR"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="

	COMMA      = ","
	DOT        = "."
	SEMICOLON  = ";"
	COLON      = ":"
	UNDERSCORE = "_"
	BACKTICK   = "`"

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "["

	PACKAGE    = "PACKAGE"
	IMPORT     = "IMPORT"
	FUNCTION   = "FUNCTION"
	VAR        = "VAR"
	MAP        = "MAP"
	TRUE       = "TRUE"
	FALSE      = "FALSE"
	INTERFACE  = "INTERFACE"
	IF         = "IF"
	ELSE       = "ELSE"
	RETURN     = "RETURN"
	STRUCT     = "STRUCT"
	SWITCH     = "SWITCH"
	TYPE       = "TYPE"
	STRING_LIT = "STRING_LIT"
)

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"package":   PACKAGE,
	"import":    IMPORT,
	"func":      FUNCTION,
	"var":       VAR,
	"map":       MAP,
	"true":      TRUE,
	"false":     FALSE,
	"interface": INTERFACE,
	"if":        IF,
	"else":      ELSE,
	"return":    RETURN,
	"struct":    STRUCT,
	"switch":    SWITCH,
	"string":    STRING,
	"type":      TYPE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
