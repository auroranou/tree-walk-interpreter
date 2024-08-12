package parse

import "fmt"

type AstPrinter struct {
	ExprVisitor
}

func (printer AstPrinter) print(expr Expr) string {
	return expr.Accept(printer)
}

func (printer AstPrinter) parenthesize(name string, exprs ...Expr) string {
	val := "(" + name
	for _, expr := range exprs {
		val += " "
		val += printer.print(expr)
	}
	val += ")"
	return val
}

func (printer AstPrinter) visitBinaryExpr(expr BinaryExpr) string {
	return printer.parenthesize(expr.operator.Lexeme, expr.left, expr.right)
}

func (printer AstPrinter) visitGroupingExpr(expr GroupingExpr) string {
	return printer.parenthesize("group", expr.expression)
}

func (printer AstPrinter) visitLiteralExpr(expr LiteralExpr) string {
	if expr.value == nil {
		return "nil"
	}
	return fmt.Sprintf("%b", expr.value)
}

func (printer AstPrinter) visitUnaryExpr(expr UnaryExpr) string {
	return printer.parenthesize(expr.operator.Lexeme, expr.right)
}
