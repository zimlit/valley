package expr

import "fmt"

type AstPrinter struct {}

func (astPrinter AstPrinter) Print(expr Expr) interface {} {
	return expr.Accept(astPrinter)
}

func (astPrinter AstPrinter) visitBinaryExpr(expr Binary) interface {} {
	return astPrinter.parenthesize(expr.Operator.Literal, expr.Left, expr.Right)	
}

func (astPrinter AstPrinter) visitGroupingExpr(expr Grouping) interface {} {
	return astPrinter.parenthesize("group", expr.Expression)	
}

func (astPrinter AstPrinter) visitLiteralExpr(expr Literal) interface {} {
	return fmt.Sprint(expr.Value)
}

func (astPrinter AstPrinter) visitUnaryExpr(expr Unary) interface {} {
	return astPrinter.parenthesize(expr.Operator.Literal, expr.Right)	
}

func (astPrinter AstPrinter) parenthesize(name string, exprs ...Expr) string {
	str := ""
	str += "("
	str += name
	for _, exp  := range exprs {
		str += " "
		str += fmt.Sprint(exp.Accept(astPrinter))
	}
	str += ")"

	return str
}
