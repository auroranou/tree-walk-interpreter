package parse

import (
	"github.com/auroranou/tree-walk-interpreter/scan"
)

type Expr interface {
	Accept(visitor ExprVisitor) string
}

type ExprVisitor interface {
	VisitBinaryExpr(expr BinaryExpr) string
	VisitGroupingExpr(expr GroupingExpr) string
	VisitLiteralExpr(expr LiteralExpr) string
	VisitUnaryExpr(expr UnaryExpr) string
}

type BinaryExpr struct {
	left     Expr
	operator scan.Token
	right    Expr
}

func (expr *BinaryExpr) accept(visitor ExprVisitor) {
	visitor.VisitBinaryExpr(*expr)
}

type GroupingExpr struct {
	expression Expr
}

func (expr *GroupingExpr) accept(visitor ExprVisitor) {
	visitor.VisitGroupingExpr(*expr)
}

type LiteralExpr struct {
	value interface{}
}

func (expr *LiteralExpr) accept(visitor ExprVisitor) {
	visitor.VisitLiteralExpr(*expr)
}

type UnaryExpr struct {
	operator scan.Token
	right    Expr
}

func (expr *UnaryExpr) accept(visitor ExprVisitor) {
	visitor.VisitUnaryExpr(*expr)
}
