package booleanexpr

import (
	"fmt"
)

type Parser struct {
	lexer    *Lexer
	curTok   Token
	elements map[string]Element
}

func NewParser(input string, elements map[string]Element) *Parser {
	lexer := NewLexer(input)
	return &Parser{lexer: lexer, curTok: lexer.NextToken(), elements: elements}
}

func (p *Parser) parsePrimary() bool {
	switch p.curTok.Type {
	case TokenElement:
		element, ok := p.elements[p.curTok.Value]
		if !ok {
			panic(fmt.Sprintf("unknown element: %s", p.curTok.Value))
		}
		p.curTok = p.lexer.NextToken()
		return element.Check()
	case TokenNot:
		p.curTok = p.lexer.NextToken()
		return !p.parsePrimary()
	case TokenLParen:
		p.curTok = p.lexer.NextToken()
		result := p.parseExpression()
		if p.curTok.Type != TokenRParen {
			panic("expected closing parenthesis")
		}
		p.curTok = p.lexer.NextToken()
		return result
	default:
		panic(fmt.Sprintf("unexpected token: %v", p.curTok))
	}
}

func (p *Parser) parseAnd() bool {
	result := p.parsePrimary()
	for p.curTok.Type == TokenAnd {
		p.curTok = p.lexer.NextToken()
		result = result && p.parsePrimary()
	}
	return result
}

func (p *Parser) parseExpression() bool {
	result := p.parseAnd()
	for p.curTok.Type == TokenOr {
		p.curTok = p.lexer.NextToken()
		result = result || p.parseAnd()
	}
	return result
}

func (p *Parser) Parse() bool {
	return p.parseExpression()
}
