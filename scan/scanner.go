package scan

import "strconv"

type Scanner struct {
	current int
	line    int
	source  string
	start   int
	tokens  []Token
}

func NewScanner(source string) *Scanner {
	scanner := Scanner{current: 0, line: 1, source: source, start: 0, tokens: []Token{}}
	return &scanner
}

func (s *Scanner) ScanTokens() []Token {
	for !s.isAtEnd() {
		// We are at the beginning of the next lexeme.
		s.start = s.current
		s.scanToken()
	}

	s.tokens = append(s.tokens, Token{tokenType: EOF, line: s.line})
	return s.tokens
}

func (s *Scanner) scanToken() {
	var nextToken TokenType

	switch c := s.advance(); c {
	case '(':
		s.addToken(LEFT_PAREN)
	case ')':
		s.addToken(RIGHT_PAREN)
	case '{':
		s.addToken(LEFT_BRACE)
	case '}':
		s.addToken(RIGHT_BRACE)
	case ',':
		s.addToken(COMMA)
	case '.':
		s.addToken(DOT)
	case '-':
		s.addToken(MINUS)
	case '+':
		s.addToken(PLUS)
	case ';':
		s.addToken(SEMICOLON)
	case '*':
		s.addToken(STAR)
	case '!':
		if s.match('=') {
			nextToken = BANG_EQUAL
		} else {
			nextToken = BANG
		}
		s.addToken(nextToken)
	case '=':
		if s.match('=') {
			nextToken = EQUAL_EQUAL
		} else {
			nextToken = EQUAL
		}
		s.addToken(nextToken)
	case '<':
		if s.match('=') {
			nextToken = LESS_EQUAL
		} else {
			nextToken = LESS
		}
		s.addToken(nextToken)
	case '>':
		if s.match('=') {
			nextToken = GREATER_EQUAL
		} else {
			nextToken = GREATER
		}
		s.addToken(nextToken)
	case '/':
		if s.match('/') {
			// A comment goes until the end of the line.
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(SLASH)
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

func (s *Scanner) addToken(tokenType TokenType) {
	s.addTokenWithLiteral(tokenType, "")
}

func (s *Scanner) addTokenWithLiteral(tokenType TokenType, literal interface{}) {
	lexeme := s.source[s.start:s.current]
	token := Token{
		tokenType: tokenType,
		lexeme:    lexeme,
		literal:   literal,
		line:      s.line,
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
	s.addTokenWithLiteral(STRING, strValue)
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

	s.addTokenWithLiteral(NUMBER, num)
}

func (s *Scanner) identifier() {
	for isAlphaNumeric(s.peek()) {
		s.advance()
	}

	text := s.source[s.start:s.current]
	tokenType := keywords[text]

	if tokenType == 0 {
		tokenType = IDENTIFIER
	}

	s.addToken(tokenType)
}
