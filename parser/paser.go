package parser

import (
	"fmt"
	"github.com/ahmadrosid/yuk/ast"
	"github.com/ahmadrosid/yuk/lexer"
	"github.com/ahmadrosid/yuk/token"
)

const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     //*
	PREFIX      // -X or !X
	CALL        // someFuncCall(X)
	INDEX       // array[index]
)

var precedences = map[token.TokenType]int{
	token.EQ:       EQUALS,
	token.NOT_EQ:   EQUALS,
	token.LT:       LESSGREATER,
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.SLASH:    PRODUCT,
	token.ASTERISK: PRODUCT,
	token.LPAREN:   CALL,
	token.LBRACE:   INDEX,
}

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

type Parser struct {
	l              *lexer.Lexer
	errors         []string
	curToken       token.Token
	peekToken      token.Token
	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)
	p.registerPrefix(token.STRING_LIT, p.parseStringLiteral)
	p.registerPrefix(token.STRING, p.parseStringType)
	p.registerPrefix(token.TRUE, p.parseBoolean)
	p.registerPrefix(token.FUNCTION, p.parseFunctionLiteral)
	p.registerPrefix(token.STRUCT, p.parseStructLiteral)
	p.registerPrefix(token.PACKAGE, p.parseExpressionLiteral)
	p.registerPrefix(token.MAP, p.parseMapLiteral)
	p.registerPrefix(token.IF, p.parseIfExpression)

	p.infixParseFns = make(map[token.TokenType]infixParseFn)
	p.registerInfix(token.PLUS, p.parseInfixExpression)
	p.registerInfix(token.MINUS, p.parseInfixExpression)
	p.registerInfix(token.SLASH, p.parseInfixExpression)
	p.registerInfix(token.EQ, p.parseInfixExpression)

	// Set curToken adn peekToken
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) parseIdentifier() ast.Expression {
	if p.peekTokenIs(token.COLON) {
		return p.parseVarExpression()
	}
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseVarExpression() ast.Expression {
	var identifier = p.curToken
	p.nextToken()
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}
	p.nextToken()
	return &ast.VarExpression{
		Token: identifier,
		Ident: identifier,
		Value: p.parseExpression(LOWEST),
	}
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.curToken}
	lit.Value = p.curToken.Literal
	return lit
}

func (p *Parser) parseStringLiteral() ast.Expression {
	return &ast.StringLiteral{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseStringType() ast.Expression {
	return &ast.ExpressionLiteral{
		Token: p.curToken,
		Name: &ast.Identifier{
			Token: p.curToken,
			Value: p.curToken.Literal,
		},
	}
}

func (p *Parser) parseBoolean() ast.Expression {
	return &ast.Boolean{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseFunctionLiteral() ast.Expression {
	lit := &ast.FunctionLiteral{Token: p.curToken}
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	lit.Name = p.curToken.Literal
	if !p.expectPeek(token.LPAREN) {
		return nil
	}
	lit.Params = p.parseFunctionParams()

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	p.nextToken()
	if !p.curTokenIs(token.LBRACE) {
		lit.ReturnType = p.parseExpression(LOWEST)
		p.nextToken()
	}

	if !p.curTokenIs(token.LBRACE) {
		return nil
	}

	lit.Body = p.parseBlockStatement()

	return lit
}

func (p *Parser) parseStructLiteral() ast.Expression {
	lit := &ast.StructStatement{Token: p.curToken}
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	lit.Name = p.curToken
	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	lit.Attributes = p.parseAttributes()

	if !p.curTokenIs(token.RPAREN) {
		return nil
	}

	return lit
}

func (p *Parser) parseExpressionLiteral() ast.Expression {
	lit := &ast.ExpressionLiteral{Token: p.curToken}
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	lit.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	return lit
}

func (p *Parser) parseMapLiteral() ast.Expression {
	lit := &ast.MapLiteral{Token: p.curToken}
	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	p.nextToken()
	lit.Key = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	if !p.expectPeek(token.COMMA) {
		return nil
	}

	p.nextToken()
	lit.Value = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	if p.curTokenIs(token.INTERFACE) {
		lit.Value.Value += "{}"
	}

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	if p.peekTokenIs(token.LBRACE) {
		lit.KeyValue = p.parseHashLiteral()
	}

	return lit
}

func (p *Parser) parseIfExpression() ast.Expression {
	exp := &ast.IfExpression{Token: p.curToken}
	p.nextToken()
	exp.Condition = p.parseExpression(LOWEST)

	if !p.expectPeek(token.LBRACE) {
		return nil
	}
	exp.Consequence = p.parseBlockStatement()
	if p.peekTokenIs(token.ELSE) {
		p.nextToken()
		if !p.expectPeek(token.LBRACE) {
			return nil
		}
		exp.Alternative = p.parseBlockStatement()
	}

	if !p.expectPeek(token.RBRACE) {
		return nil
	}
	return exp
}

func (p *Parser) parseHashLiteral() *ast.HashLiteral {
	lit := &ast.HashLiteral{KeyValue: map[ast.Expression]ast.Expression{}}

	if !p.expectPeek(token.LBRACE) {
		return lit
	}

	for {
		p.nextToken()
		key := p.parseExpression(LOWEST)
		if !p.expectPeek(token.COLON) {
			return nil
		}

		p.nextToken()
		val := p.parseExpression(LOWEST)
		lit.KeyValue[key] = val

		p.nextToken()
		if p.curTokenIs(token.RBRACE) {
			break
		}
	}

	return lit
}

func (p *Parser) parseImportStatement() ast.Statement {
	lit := &ast.ImportStatement{Token: p.curToken}
	if !p.expectPeek(token.STRING_LIT) {
		return nil
	}

	lit.PackageName = p.parseStringLiteral()

	return lit
}

func (p *Parser) parseAttributes() []*ast.StructAttributes {
	attrs := make([]*ast.StructAttributes, 0)
	if p.peekTokenIs(token.RPAREN) {
		return attrs
	}

	p.nextToken()
	for {
		if p.curTokenIs(token.EOF) || p.curTokenIs(token.RPAREN) {
			break
		}
		if p.curTokenIs(token.COMMA) {
			p.nextToken()
			continue
		}

		if p.curTokenIs(token.ILLEGAL) {
			p.errors = append(p.errors, fmt.Sprintf("illegal token %s", p.curToken.Literal))
			return nil
		}

		attr := p.parseStructAttributes()
		if attr == nil {
			return nil
		}

		attrs = append(attrs, attr)
	}

	return attrs
}

func (p *Parser) parseStructAttributes() *ast.StructAttributes {
	attr := &ast.StructAttributes{Name: p.curToken}
	p.nextToken()
	attr.Type = p.curToken
	if !p.peekTokenIs(token.EOF) {
		p.nextToken()
	}

	if p.curTokenIs(token.BACKTICK) {
		meta := &ast.MetaLiteral{Token: p.curToken}
		for {
			p.nextToken()
			if p.curTokenIs(token.BACKTICK) {
				p.nextToken()
				break
			}

			keyVal := &ast.MetaKeyValueLiteral{Key: p.curToken}
			if !p.expectPeek(token.COLON) {
				return nil
			}
			p.nextToken()
			keyVal.Value = p.parseExpression(LOWEST)
			meta.KeyValue = append(meta.KeyValue, keyVal)
		}
		attr.Meta = meta
	}

	return attr
}

func (p *Parser) parseFunctionParams() []*ast.Identifier {
	var identifiers []*ast.Identifier
	if p.peekTokenIs(token.RPAREN) {
		return identifiers
	}

	return identifiers
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{Token: p.curToken}
	p.nextToken()
	for {
		if p.curTokenIs(token.RBRACE) || p.curTokenIs(token.EOF) {
			break
		}
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()
	}

	return block
}

func (p *Parser) parseTypeStatement() *ast.StructAttributes {
	typeToken := p.curToken
	stmt := &ast.StructAttributes{Token: &typeToken}
	p.nextToken()
	stmt.Name = p.curToken
	p.nextToken()
	stmt.Type = p.curToken

	if p.peekTokenIs(token.EOF) {
		return stmt
	}

	return stmt
}

func (p *Parser) parseSwitchStatement() *ast.SwitchStatement {
	stmt := &ast.SwitchStatement{Token: p.curToken}

	p.nextToken()
	stmt.Input = p.curToken

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	stmt.Case = append(stmt.Case, p.parseCaseLiteral())

	if p.curTokenIs(token.COMMA) {
		for {
			stmt.Case = append(stmt.Case, p.parseCaseLiteral())
			if p.curTokenIs(token.RBRACE) {
				break
			}
		}
	}

	return stmt
}

func (p *Parser) parseCaseLiteral() *ast.CaseLiteral {
	p.nextToken()
	lit := &ast.CaseLiteral{Token: p.curToken}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	if !p.expectPeek(token.GT) {
		return nil
	}

	p.nextToken()
	lit.Body = p.parseBlockStatement()

	p.nextToken()
	return lit
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left,
	}
	precedence := p.curPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)
	return expression
}

func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be '%s', got '%s' instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) parseVarStatement() *ast.VarStatement {
	stmt := &ast.VarStatement{Token: p.curToken}
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	p.nextToken()

	stmt.Value = p.parseExpression(LOWEST)

	return stmt
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}
	leftExp := prefix()
	for !p.peekTokenIs(token.NEW_LINE) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}
		p.nextToken()
		leftExp = infix(leftExp)
	}
	return leftExp
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) noPrefixParseFnError(t token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for '%s' found", t)
	p.errors = append(p.errors, msg)
}

func (p *Parser) parseReturnStatement() ast.Statement {
	stmt := &ast.ReturnStatement{Token: p.curToken}
	p.nextToken()
	stmt.ReturnValue = p.parseExpression(LOWEST)
	if p.peekTokenIs(token.NEW_LINE) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseExpressionStatement() ast.Statement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}
	stmt.Expression = p.parseExpression(LOWEST)
	if p.peekTokenIs(token.NEW_LINE) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.VAR:
		return p.parseVarStatement()
	case token.IMPORT:
		return p.parseImportStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	case token.TYPE:
		return p.parseTypeStatement()
	case token.SWITCH:
		return p.parseSwitchStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}
