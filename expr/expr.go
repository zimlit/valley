package expr

import "valley/token"

type Expr interface {
  Accept(visitor Visitor) interface {}
}

type Visitor interface {
   visitBinaryExpr(expr Binary) interface {}
   visitGroupingExpr(expr Grouping) interface {}
   visitLiteralExpr(expr Literal) interface {}
   visitUnaryExpr(expr Unary) interface {}
}

type Binary struct {
    left Expr
    operator token.Token
    right Expr
}

func (binary Binary) Accept(visitor Visitor) interface {} {
    return visitor.visitBinaryExpr(binary)
}

type Grouping struct {
    expression Expr
}

func (grouping Grouping) Accept(visitor Visitor) interface {} {
    return visitor.visitGroupingExpr(grouping)
}

type Literal struct {
    value interface {}
}

func (literal Literal) Accept(visitor Visitor) interface {} {
    return visitor.visitLiteralExpr(literal)
}

type Unary struct {
    operator token.Token
    right Expr
}

func (unary Unary) Accept(visitor Visitor) interface {} {
    return visitor.visitUnaryExpr(unary)
}
