package compiler

import (
	"bytes"
	"fmt"
	"github.com/ahmadrosid/yuk/ast"
	"github.com/ahmadrosid/yuk/parser"
)

type Compiler struct {
	Parser  *parser.Parser
	Program *ast.Program
}

func New(p *parser.Parser) *Compiler {
	return &Compiler{Parser: p}
}

func (c *Compiler) Generate() (string, []error) {
	c.Program = c.Parser.ParseProgram()
	if len(c.Parser.Errors()) > 0 {
		var errors []error
		for _, s := range c.Parser.Errors() {
			errors = append(errors, fmt.Errorf("%s", s))
		}
		return "", errors
	}

	var out bytes.Buffer
	for _, stmt := range c.Program.Statements {
		out.WriteString(stmt.String())
	}
	return out.String(), nil
}
