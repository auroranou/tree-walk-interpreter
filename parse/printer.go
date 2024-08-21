package parse

import (
	"fmt"

	"github.com/auroranou/tree-walk-interpreter/grammar"
)

type AstPrinter struct {
	grammar.ExprVisitor
}

func (printer AstPrinter) Print(expr grammar.Expr) string {
	return expr.Accept(printer)
}

func (printer AstPrinter) parenthesize(name string, exprs ...grammar.Expr) string {
	val := "(" + name
	for _, expr := range exprs {
		val += " "
		val += printer.Print(expr)
	}
	val += ")"
	return val
}

func (printer AstPrinter) VisitBinaryExpr(expr grammar.BinaryExpr) string {
	return printer.parenthesize(expr.Operator.Lexeme, expr.Left, expr.Right)
}

func (printer AstPrinter) VisitGroupingExpr(expr grammar.GroupingExpr) string {
	return printer.parenthesize("group", expr.Expression)
}

func (printer AstPrinter) VisitLiteralExpr(expr grammar.LiteralExpr) string {
	if expr.Value == nil {
		return "nil"
	}
	return fmt.Sprintf("%v", expr.Value)
}

func (printer AstPrinter) VisitUnaryExpr(expr grammar.UnaryExpr) string {
	return printer.parenthesize(expr.Operator.Lexeme, expr.Right)
}
