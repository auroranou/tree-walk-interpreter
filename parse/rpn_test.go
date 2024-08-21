package parse

import (
	"testing"

	"github.com/auroranou/tree-walk-interpreter/grammar"
)

func TestReversePolishNotation(t *testing.T) {
	// (1 + 2) * (4 - 3)
	expr := grammar.BinaryExpr{
		Left: grammar.BinaryExpr{
			Left:     grammar.LiteralExpr{Value: 1},
			Operator: grammar.Token{TokenType: grammar.PLUS, Lexeme: "+", Literal: "", Line: 1},
			Right:    grammar.LiteralExpr{Value: 2},
		},
		Operator: grammar.Token{TokenType: grammar.STAR, Lexeme: "*", Literal: "", Line: 1},
		Right: grammar.BinaryExpr{
			Left:     grammar.LiteralExpr{Value: 4},
			Operator: grammar.Token{TokenType: grammar.MINUS, Lexeme: "-", Literal: "", Line: 1},
			Right:    grammar.LiteralExpr{Value: 3},
		},
	}

	want := "1 2 + 4 3 - *"
	got := RpnConverter{}.VisitBinaryExpr(expr)

	if got != want {
		t.Errorf("Wanted: %v\nGot: %v\n", want, got)
	}
}
