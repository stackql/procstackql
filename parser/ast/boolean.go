// parser/ast/boolean.go

package ast

import "token"

type Boolean struct {
	Token token.Token // token.TRUE or token.FALSE
	Value bool
}

func (b *Boolean) expressionNode()      {}
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string       { return b.Token.Literal }
