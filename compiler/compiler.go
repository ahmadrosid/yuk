package compiler

import (
	"bytes"
	"github.com/ahmadrosid/yuk/ast"
)

type Compiler struct {
	Program *ast.Program
}

func New(program *ast.Program) *Compiler {
	return &Compiler{Program: program}
}

func (comp *Compiler) Generate() string {
	var out bytes.Buffer
	for _, c := range comp.Program.Statements {
		if c != nil {
			out.WriteString(c.String())
		}
	}
	return out.String()
}
