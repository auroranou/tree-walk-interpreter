package parse

import (
	"fmt"

	"github.com/auroranou/tree-walk-interpreter/grammar"
)

type ParseError struct {
	token grammar.Token
	msg   string
}

func (e ParseError) Error() string {
	if e.token.TokenType == grammar.EOF {
		return fmt.Sprintf("%d at end, %s", e.token.Line, e.msg)
	} else {
		return fmt.Sprintf("%d at %s, %s", e.token.Line, e.token.Lexeme, e.msg)
	}
}
