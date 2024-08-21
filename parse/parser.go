package parse

import (
	"github.com/auroranou/tree-walk-interpreter/grammar"
)

type Parser struct {
	current int
	tokens  []grammar.Token
}

func NewParser(tokens []grammar.Token) *Parser {
	parser := Parser{current: 0, tokens: tokens}
	return &parser
}

func (p *Parser) Parse() (expr grammar.Expr, err error) {
	expr = p.expression()
	defer func() {
		if r := recover(); r != nil {
			_, ok := err.(ParseError)
			if ok {
				// TODO
			} else {
				panic(err)
			}
		}
	}()

	return expr, nil
}

func (p *Parser) expression() grammar.Expr {
	return p.equality()
}

func (p *Parser) equality() grammar.Expr {
	expr := p.comparison()

	for p.match(grammar.BANG_EQUAL, grammar.EQUAL_EQUAL) {
		operator := p.previous()
		right := p.comparison()
		expr = grammar.BinaryExpr{Left: expr, Operator: operator, Right: right}
	}

	return expr
}

func (p *Parser) comparison() grammar.Expr {
	expr := p.term()

	for p.match(grammar.GREATER, grammar.GREATER_EQUAL, grammar.LESS, grammar.LESS_EQUAL) {
		operator := p.previous()
		right := p.term()
		expr = grammar.BinaryExpr{Left: expr, Operator: operator, Right: right}
	}

	return expr
}

func (p *Parser) term() grammar.Expr {
	expr := p.factor()

	for p.match(grammar.MINUS, grammar.PLUS) {
		operator := p.previous()
		right := p.factor()
		expr = grammar.BinaryExpr{Left: expr, Operator: operator, Right: right}
	}

	return expr
}

func (p *Parser) factor() grammar.Expr {
	expr := p.unary()

	for p.match(grammar.SLASH, grammar.STAR) {
		operator := p.previous()
		right := p.unary()
		expr = grammar.BinaryExpr{Left: expr, Operator: operator, Right: right}
	}

	return expr
}

func (p *Parser) unary() grammar.Expr {
	if p.match(grammar.BANG, grammar.MINUS) {
		operator := p.previous()
		right := p.unary()
		return grammar.UnaryExpr{Operator: operator, Right: right}
	}

	return p.primary()
}

func (p *Parser) primary() grammar.Expr {
	if p.match(grammar.FALSE) {
		return grammar.LiteralExpr{Value: false}
	}

	if p.match(grammar.TRUE) {
		return grammar.LiteralExpr{Value: true}
	}

	if p.match(grammar.NIL) {
		return grammar.LiteralExpr{Value: nil}
	}

	if p.match(grammar.NUMBER, grammar.STRING) {
		return grammar.LiteralExpr{Value: p.previous().Literal}
	}

	if p.match(grammar.LEFT_PAREN) {
		expr := p.expression()
		p.consume(grammar.RIGHT_PAREN, "Expect ')' after expression.")
		return grammar.GroupingExpr{Expression: expr}
	}

	p.throwError(p.peek(), "Expect expression.")
	return nil
}

func (p *Parser) match(tokenTypes ...grammar.TokenType) bool {
	for _, tokenType := range tokenTypes {
		if p.check(tokenType) {
			p.advance()
			return true
		}
	}

	return false
}

func (p *Parser) check(tokenType grammar.TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().TokenType == tokenType
}

func (p *Parser) previous() grammar.Token {
	return p.tokens[p.current-1]
}

func (p *Parser) peek() grammar.Token {
	return p.tokens[p.current]
}

func (p *Parser) advance() grammar.Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

func (p *Parser) consume(tokenType grammar.TokenType, msg string) {
	if p.check(tokenType) {
		p.advance()
	} else {
		p.throwError(p.peek(), msg)
	}
}

func (p *Parser) isAtEnd() bool {
	return p.peek().TokenType == grammar.EOF
}

func (p *Parser) synchronize() {
	p.advance()

	for !p.isAtEnd() {
		if p.previous().TokenType == grammar.SEMICOLON {
			return
		}

		switch t := p.peek().TokenType; t {
		case grammar.CLASS:
		case grammar.FUN:
		case grammar.VAR:
		case grammar.FOR:
		case grammar.IF:
		case grammar.WHILE:
		case grammar.PRINT:
		case grammar.RETURN:
			return
		}

		p.advance()
	}
}

func (p *Parser) throwError(token grammar.Token, msg string) {
	panic(ParseError{token: token, msg: msg}.Error())
}
