package parse

import (
	"fmt"

	"github.com/auroranou/tree-walk-interpreter/grammar"
)

// Chapter 5, challenge 3

type RpnConverter struct{ grammar.ExprVisitor }

func (r RpnConverter) VisitBinaryExpr(expr grammar.BinaryExpr) string {
	return fmt.Sprintf("%v %v %v", expr.Left.Accept(r), expr.Right.Accept(r), expr.Operator.Lexeme)
}

func (r RpnConverter) VisitGroupingExpr(expr grammar.GroupingExpr) string {
	return expr.Expression.Accept(r)
}

func (r RpnConverter) VisitLiteralExpr(expr grammar.LiteralExpr) string {
	return fmt.Sprintf("%v", expr.Value)
}

func (r RpnConverter) VisitUnaryExpr(expr grammar.UnaryExpr) string {
	return fmt.Sprintf("%v %v", expr.Right.Accept(r), expr.Operator)
}
