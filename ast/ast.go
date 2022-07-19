package ast

import (
	"bytes"
	"github.com/ahmadrosid/yuk/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

type VarStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (vs *VarStatement) statementNode()       {}
func (vs *VarStatement) TokenLiteral() string { return vs.Token.Literal }
func (vs *VarStatement) String() string {
	var out bytes.Buffer
	out.WriteString(vs.TokenLiteral() + " ")
	out.WriteString(vs.Name.String())
	out.WriteString(" = ")
	if vs.Value != nil {
		out.WriteString(vs.Value.String())
	}
	return out.String()
}

type ExpressionLiteral struct {
	Token token.Token
	Name  *Identifier
}

func (pl *ExpressionLiteral) expressionNode()      {}
func (pl *ExpressionLiteral) TokenLiteral() string { return pl.Token.Literal }
func (pl *ExpressionLiteral) String() string {
	var out bytes.Buffer
	out.WriteString(pl.TokenLiteral() + " ")
	out.WriteString(pl.Name.String())
	return out.String()
}

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	return out.String()
}

type ImportStatement struct {
	Token       token.Token
	PackageName string
}

func (s *ImportStatement) statementNode()       {}
func (s *ImportStatement) TokenLiteral() string { return s.Token.Literal }
func (s *ImportStatement) String() string {
	var out bytes.Buffer
	out.WriteString(s.TokenLiteral() + " ")
	out.WriteByte('"')
	out.WriteString(s.PackageName)
	out.WriteByte('"')
	return out.String()
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer
	out.WriteString("{")
	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}
	out.WriteString("}")
	return out.String()
}

type StructStatement struct {
	Token      token.Token
	Name       token.Token
	Attributes []*TypeStatement
	Block      *BlockStatement
}

func (ss *StructStatement) expressionNode()      {}
func (ss *StructStatement) statementNode()       {}
func (ss *StructStatement) TokenLiteral() string { return ss.Token.Literal }
func (ss *StructStatement) String() string {
	var out bytes.Buffer
	out.WriteString("type ")
	out.WriteString(ss.Name.Literal + " ")
	out.WriteString(ss.Token.Literal + " ")
	out.WriteString("{")
	out.WriteString("\n")
	for _, s := range ss.Attributes {
		out.WriteString(s.String())
		out.WriteString("\n")
	}
	out.WriteString("}")
	// TODO: handle write block
	return out.String()
}

type TypeStatement struct {
	Token *token.Token
	Name  token.Token
	Type  token.Token
}

func (ts *TypeStatement) statementNode()       {}
func (ts *TypeStatement) TokenLiteral() string { return ts.Token.Literal }
func (ts *TypeStatement) String() string {
	var out bytes.Buffer
	if ts.Token != nil {
		out.WriteString(ts.TokenLiteral() + " ")
	}
	out.WriteString(ts.Name.Literal + " ")
	out.WriteString(ts.Type.Literal)
	return out.String()
}

type IntegerLiteral struct {
	Token token.Token
	Value string
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string {
	var out bytes.Buffer
	out.WriteString(il.Value)
	return out.String()
}

type StringLiteral struct {
	Token token.Token
	Value string
}

func (sl *StringLiteral) expressionNode()      {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StringLiteral) String() string {
	var out bytes.Buffer
	// TODO: format string for StringLiteral
	return out.String()
}

type Boolean struct {
	Token token.Token
	Value string
}

func (b *Boolean) expressionNode()      {}
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string {
	var out bytes.Buffer
	// TODO: format string for Boolean
	return out.String()
}

type InfixExpression struct {
	Token    token.Token
	Operator string
	Left     Expression
	Right    Expression
}

func (ie *InfixExpression) expressionNode()      {}
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *InfixExpression) String() string {
	var out bytes.Buffer
	// TODO: format string for StringLiteral
	return out.String()
}

type FunctionLiteral struct {
	Token  token.Token
	Name   string
	Params []*Identifier
	Body   *BlockStatement
}

func (fl *FunctionLiteral) expressionNode()      {}
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer
	out.WriteString(fl.TokenLiteral() + " ")
	out.WriteString(fl.Name)
	out.WriteString("(")
	for _, s := range fl.Params {
		out.WriteString(s.String())
	}
	out.WriteString(") ")
	out.WriteString(fl.Body.String())
	return out.String()
}
