package parse

import "fmt"

type AstPrinter struct {
	ExprVisitor
}

func (printer AstPrinter) Print(expr Expr) string {
	return expr.Accept(printer)
}

func (printer AstPrinter) parenthesize(name string, exprs ...Expr) string {
	val := "(" + name
	for _, expr := range exprs {
		val += " "
		val += printer.Print(expr)
	}
	val += ")"
	return val
}

func (printer AstPrinter) VisitBinaryExpr(expr BinaryExpr) string {
	return printer.parenthesize(expr.Operator.Lexeme, expr.Left, expr.Right)
}

func (printer AstPrinter) VisitGroupingExpr(expr GroupingExpr) string {
	return printer.parenthesize("group", expr.Expression)
}

func (printer AstPrinter) VisitLiteralExpr(expr LiteralExpr) string {
	if expr.Value == nil {
		return "nil"
	}
	return fmt.Sprintf("%v", expr.Value)
}

func (printer AstPrinter) VisitUnaryExpr(expr UnaryExpr) string {
	return printer.parenthesize(expr.Operator.Lexeme, expr.Right)
}
