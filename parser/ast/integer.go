// parser/ast/integer.go

package ast

import "token"

type IntegerLiteral struct {
	Token token.Token // The token.INT token
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }
