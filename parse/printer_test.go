package parse

import (
	"testing"

	"github.com/auroranou/tree-walk-interpreter/grammar"
)

func TestPrettyPrint(t *testing.T) {
	expr := grammar.BinaryExpr{
		Left: grammar.UnaryExpr{
			Operator: grammar.Token{TokenType: grammar.MINUS, Lexeme: "-", Literal: "", Line: 1},
			Right:    grammar.LiteralExpr{Value: 123},
		},
		Operator: grammar.Token{TokenType: grammar.STAR, Lexeme: "*", Literal: "", Line: 1},
		Right:    grammar.GroupingExpr{Expression: grammar.LiteralExpr{Value: 45.67}},
	}

	want := "(* (- 123) (group 45.67))"
	got := AstPrinter{}.Print(expr)

	if got != want {
		t.Errorf("Wanted: %v\nGot: %v\n", want, got)
	}
}
