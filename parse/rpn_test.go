package parse

import (
	"testing"

	"github.com/auroranou/tree-walk-interpreter/scan"
)

func TestReversePolishNotation(t *testing.T) {
	// (1 + 2) * (4 - 3)
	expr := BinaryExpr{
		Left: BinaryExpr{
			Left:     LiteralExpr{Value: 1},
			Operator: scan.Token{TokenType: scan.PLUS, Lexeme: "+", Literal: "", Line: 1},
			Right:    LiteralExpr{2},
		},
		Operator: scan.Token{TokenType: scan.STAR, Lexeme: "*", Literal: "", Line: 1},
		Right: BinaryExpr{
			Left:     LiteralExpr{Value: 4},
			Operator: scan.Token{TokenType: scan.MINUS, Lexeme: "-", Literal: "", Line: 1},
			Right:    LiteralExpr{Value: 3},
		},
	}

	want := "1 2 + 4 3 - *"
	got := RpnConverter{}.VisitBinaryExpr(expr)

	if got != want {
		t.Errorf("Wanted: %v\nGot: %v\n", want, got)
	}
}
