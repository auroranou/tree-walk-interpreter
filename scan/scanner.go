package scan

import (
	"strconv"

	"github.com/auroranou/tree-walk-interpreter/grammar"
)

type Scanner struct {
	current int
	line    int
	source  string
	start   int
	tokens  []grammar.Token
}

func NewScanner(source string) *Scanner {
	scanner := Scanner{current: 0, line: 1, source: source, start: 0, tokens: []grammar.Token{}}
	return &scanner
}

func (s *Scanner) ScanTokens() []grammar.Token {
	for !s.isAtEnd() {
		// We are at the beginning of the next lexeme.
		s.start = s.current
		s.scanToken()
	}

	s.tokens = append(s.tokens, grammar.Token{TokenType: grammar.EOF, Line: s.line})
	return s.tokens
}

func (s *Scanner) scanToken() {
	var nextToken grammar.TokenType

	switch c := s.advance(); c {
	case '(':
		s.addToken(grammar.LEFT_PAREN)
	case ')':
		s.addToken(grammar.RIGHT_PAREN)
	case '{':
		s.addToken(grammar.LEFT_BRACE)
	case '}':
		s.addToken(grammar.RIGHT_BRACE)
	case ',':
		s.addToken(grammar.COMMA)
	case '.':
		s.addToken(grammar.DOT)
	case '-':
		s.addToken(grammar.MINUS)
	case '+':
		s.addToken(grammar.PLUS)
	case ';':
		s.addToken(grammar.SEMICOLON)
	case '*':
		s.addToken(grammar.STAR)
	case '!':
		if s.match('=') {
			nextToken = grammar.BANG_EQUAL
		} else {
			nextToken = grammar.BANG
		}
		s.addToken(nextToken)
	case '=':
		if s.match('=') {
			nextToken = grammar.EQUAL_EQUAL
		} else {
			nextToken = grammar.EQUAL
		}
		s.addToken(nextToken)
	case '<':
		if s.match('=') {
			nextToken = grammar.LESS_EQUAL
		} else {
			nextToken = grammar.LESS
		}
		s.addToken(nextToken)
	case '>':
		if s.match('=') {
			nextToken = grammar.GREATER_EQUAL
		} else {
			nextToken = grammar.GREATER
		}
		s.addToken(nextToken)
	case '/':
		if s.match('/') {
			// A comment goes until the end of the line.
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(grammar.SLASH)
		}
	case ' ':
	case '\r':
	case '\t':
		// Ignore whitespace.
		break
	case '\n':
		s.line++
	case '"':
		s.string()
	default:
		if isDigit(c) {
			s.number()
		} else if isAlpha(c) {
			// Reserved keywords
			s.identifier()
		} else {
			panic("Unexpected character")
		}
	}
}

func (s *Scanner) addToken(tokenType grammar.TokenType) {
	s.addTokenWithLiteral(tokenType, "")
}

func (s *Scanner) addTokenWithLiteral(tokenType grammar.TokenType, literal interface{}) {
	lexeme := s.source[s.start:s.current]
	token := grammar.Token{
		TokenType: tokenType,
		Lexeme:    lexeme,
		Literal:   literal,
		Line:      s.line,
	}
	s.tokens = append(s.tokens, token)
}

func (s *Scanner) advance() rune {
	curr := rune(s.source[s.current])
	s.current++
	return curr
}

func (s *Scanner) match(expected rune) bool {
	if s.isAtEnd() {
		return false
	}

	if rune(s.source[s.current]) != expected {
		return false
	}

	s.current++
	return true
}

func (s *Scanner) peek() rune {
	if s.isAtEnd() {
		return '\000'
	}
	return rune(s.source[s.current])
}

func (s *Scanner) peekNext() rune {
	if s.current+1 > len(s.source) {
		return '\000'
	}
	return rune(s.source[s.current+1])
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) string() {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}

	if s.isAtEnd() {
		panic("Unterminated string")
	}

	// The closing ".
	s.advance()

	// Trim the surrounding quotes.
	strValue := s.source[s.start+1 : s.current-1]
	s.addTokenWithLiteral(grammar.STRING, strValue)
}

func (s *Scanner) number() {
	for isDigit(s.peek()) {
		s.advance()
	}

	// Look for a fractional part.
	if s.peek() == '.' && isDigit(s.peekNext()) {
		// Consume the "."
		s.advance()

		for isDigit(s.peek()) {
			s.advance()
		}
	}

	num, err := strconv.ParseFloat(s.source[s.start:s.current], 64)
	if err != nil {
		panic("Invalid number")
	}

	s.addTokenWithLiteral(grammar.NUMBER, num)
}

func (s *Scanner) identifier() {
	for isAlphaNumeric(s.peek()) {
		s.advance()
	}

	text := s.source[s.start:s.current]
	tokenType := grammar.Keywords[text]

	if tokenType == 0 {
		tokenType = grammar.IDENTIFIER
	}

	s.addToken(tokenType)
}
