package parse

import "fmt"

// Chapter 5, challenge 3

type RpnConverter struct{ ExprVisitor }

func (r RpnConverter) VisitBinaryExpr(expr BinaryExpr) string {
	return fmt.Sprintf("%v %v %v", expr.Left.Accept(r), expr.Right.Accept(r), expr.Operator.Lexeme)
}

func (r RpnConverter) VisitGroupingExpr(expr GroupingExpr) string {
	return expr.Expression.Accept(r)
}

func (r RpnConverter) VisitLiteralExpr(expr LiteralExpr) string {
	return fmt.Sprintf("%v", expr.Value)
}

func (r RpnConverter) VisitUnaryExpr(expr UnaryExpr) string {
	return fmt.Sprintf("%v %v", expr.Right.Accept(r), expr.Operator)
}
