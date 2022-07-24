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

type MapLiteral struct {
	Token    token.Token
	Key      *Identifier
	Value    *Identifier
	KeyValue *HashLiteral
}

func (m *MapLiteral) expressionNode()      {}
func (m *MapLiteral) TokenLiteral() string { return m.Token.Literal }
func (m *MapLiteral) String() string {
	var out bytes.Buffer
	out.WriteString(m.TokenLiteral())
	out.WriteString("[")
	out.WriteString(m.Key.String())
	out.WriteString("]")
	out.WriteString(m.Value.String())
	if m.KeyValue != nil {
		out.WriteString(m.KeyValue.String())
	}
	return out.String()
}

type HashLiteral struct {
	KeyValue map[Expression]Expression
}

func (h *HashLiteral) expressionNode()      {}
func (h *HashLiteral) TokenLiteral() string { return "" }
func (h *HashLiteral) String() string {
	var out bytes.Buffer
	out.WriteString("{\n")
	for k, v := range h.KeyValue {
		out.WriteString(k.String())
		out.WriteString(":")
		out.WriteString(v.String())
		out.WriteString(",\n")
	}
	out.WriteString("}")
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
	PackageName Expression
}

func (s *ImportStatement) statementNode()       {}
func (s *ImportStatement) TokenLiteral() string { return s.Token.Literal }
func (s *ImportStatement) String() string {
	var out bytes.Buffer
	out.WriteString(s.TokenLiteral() + " ")
	out.WriteString(s.PackageName.String())
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

type MetaLiteral struct {
	Token    token.Token
	KeyValue []*MetaKeyValueLiteral
}

func (m *MetaLiteral) statementNode()       {}
func (m *MetaLiteral) TokenLiteral() string { return m.Token.Literal }
func (m *MetaLiteral) String() string {
	var out bytes.Buffer
	out.WriteString(m.TokenLiteral())
	for _, kv := range m.KeyValue {
		out.WriteString(kv.String())
	}
	out.WriteString(m.TokenLiteral())
	return out.String()
}

type MetaKeyValueLiteral struct {
	Key   token.Token
	Value Expression
}

func (m *MetaKeyValueLiteral) statementNode()       {}
func (m *MetaKeyValueLiteral) TokenLiteral() string { return "" }
func (m *MetaKeyValueLiteral) String() string {
	var out bytes.Buffer
	out.WriteString(m.Key.Literal)
	out.WriteString(":")
	out.WriteString(m.Value.String())
	return out.String()
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
		out.WriteString("\n")
		out.WriteString(s.String())
		out.WriteString("\n")
	}
	out.WriteString("}")
	return out.String()
}

type StructStatement struct {
	Token      token.Token
	Name       token.Token
	Attributes []*StructAttributes
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
	return out.String()
}

type StructAttributes struct {
	Token *token.Token
	Name  token.Token
	Type  token.Token
	Meta  *MetaLiteral
}

func (ts *StructAttributes) statementNode()       {}
func (ts *StructAttributes) TokenLiteral() string { return ts.Token.Literal }
func (ts *StructAttributes) String() string {
	var out bytes.Buffer
	if ts.Token != nil {
		out.WriteString(ts.TokenLiteral() + " ")
	}
	out.WriteString(ts.Name.Literal + " ")
	out.WriteString(ts.Type.Literal)
	if ts.Meta != nil {
		out.WriteString(" ")
		out.WriteString(ts.Meta.String())
	}
	return out.String()
}

type SwitchStatement struct {
	Token token.Token
	Input token.Token
	Case  []*CaseLiteral
}

func (ss *SwitchStatement) statementNode()       {}
func (ss *SwitchStatement) TokenLiteral() string { return ss.Token.Literal }
func (ss *SwitchStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ss.TokenLiteral() + " ")
	if ss.Input.Type == token.CHAR {
		out.WriteString("'")
		out.WriteString(ss.Input.Literal)
		out.WriteString("'")
	} else if ss.Input.Type == token.STRING {
		out.WriteString("\"")
		out.WriteString(ss.Input.Literal)
		out.WriteString("\"")
	} else {
		out.WriteString(ss.Input.Literal)
	}
	out.WriteByte(' ')

	out.WriteString("{")
	for _, c := range ss.Case {
		out.WriteString(c.String())
		out.WriteString("\n")
	}
	out.WriteString("}")

	return out.String()
}

type CaseLiteral struct {
	Token token.Token
	Body  *BlockStatement
}

func (cl *CaseLiteral) expressionNode()      {}
func (cl *CaseLiteral) TokenLiteral() string { return cl.Token.Literal }
func (cl *CaseLiteral) String() string {
	var out bytes.Buffer

	if cl.Token.Type == token.UNDERSCORE {
		out.WriteString("default")
	} else {
		out.WriteString("case ")
		if cl.Token.Type == token.CHAR {
			out.WriteString("'")
			out.WriteString(cl.Token.Literal)
			out.WriteString("'")
		} else if cl.Token.Type == token.STRING {
			out.WriteString("\"")
			out.WriteString(cl.Token.Literal)
			out.WriteString("\"")
		} else {
			out.WriteString(cl.Token.Literal)
		}
	}

	out.WriteString(": ")

	out.WriteString(cl.Body.String())

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
	out.WriteString("\"")
	out.WriteString(sl.Value)
	out.WriteString("\"")
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
	Token      token.Token
	Name       string
	Params     []*Identifier
	ReturnType Expression
	Body       *BlockStatement
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
	if fl.ReturnType != nil {
		out.WriteString(fl.ReturnType.TokenLiteral())
		out.WriteString(" ")
	}
	out.WriteString(fl.Body.String())
	return out.String()
}
