package grammar

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
	Left     Expr
	Operator Token
	Right    Expr
}

func (expr BinaryExpr) Accept(visitor ExprVisitor) string {
	return visitor.VisitBinaryExpr(expr)
}

type GroupingExpr struct {
	Expression Expr
}

func (expr GroupingExpr) Accept(visitor ExprVisitor) string {
	return visitor.VisitGroupingExpr(expr)
}

type LiteralExpr struct {
	Value interface{}
}

func (expr LiteralExpr) Accept(visitor ExprVisitor) string {
	return visitor.VisitLiteralExpr(expr)
}

type UnaryExpr struct {
	Operator Token
	Right    Expr
}

func (expr UnaryExpr) Accept(visitor ExprVisitor) string {
	return visitor.VisitUnaryExpr(expr)
}
