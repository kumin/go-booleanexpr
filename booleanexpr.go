package booleanexpr

import "strings"

type TokenType int

const (
	TokenEOF TokenType = iota
	TokenAnd
	TokenOr
	TokenNot
	TokenElement
	TokenLParen
	TokenRParen
)

type Token struct {
	Type  TokenType
	Value string
}

type Lexer struct {
	input string
	pos   int
}

func NewLexer(input string) *Lexer {
	return &Lexer{input: input}
}

func (l *Lexer) NextToken() Token {
	for l.pos < len(l.input) {
		switch l.input[l.pos] {
		case ' ':
			l.pos++
		case '&':
			l.pos++
			return Token{Type: TokenAnd, Value: "&"}
		case '|':
			l.pos++
			return Token{Type: TokenOr, Value: "|"}
		case '!':
			l.pos++
			return Token{Type: TokenNot, Value: "!"}
		case '(':
			l.pos++
			return Token{Type: TokenLParen, Value: "("}
		case ')':
			l.pos++
			return Token{Type: TokenRParen, Value: ")"}
		default:
			start := l.pos
			for l.pos < len(l.input) && !strings.ContainsRune(" &|!()", rune(l.input[l.pos])) {
				l.pos++
			}
			return Token{Type: TokenElement, Value: l.input[start:l.pos]}
		}
	}
	return Token{Type: TokenEOF}
}

func Evaluate(expression string, elements map[string]Element) bool {
	parser := NewParser(expression, elements)
	return parser.Parse()
}
