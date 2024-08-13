package parse

import (
	"fmt"
	"testing"

	"github.com/auroranou/tree-walk-interpreter/scan"
)

func TestPrettyPrint(t *testing.T) {
	expr := BinaryExpr{
		Left: UnaryExpr{
			Operator: scan.Token{TokenType: scan.MINUS, Lexeme: "-", Literal: "", Line: 1},
			Right:    LiteralExpr{Value: 123},
		},
		Operator: scan.Token{TokenType: scan.STAR, Lexeme: "*", Literal: "", Line: 1},
		Right:    GroupingExpr{Expression: LiteralExpr{Value: 45.67}},
	}

	want := "(* (- 123) (group 45.67))"
	got := AstPrinter{}.Print(expr)

	if fmt.Sprintf("%v", got) != want {
		t.Errorf("Wanted: %v\nGot: %v\n", want, got)
	}
}
