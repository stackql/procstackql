// parser/ast/if.go

package ast

import "token"

type IfExpression struct {
	Token       token.Token // 'if' token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement // Can be nil for 'if' without 'else'
}

type BlockStatement struct {
	Token      token.Token // '{' token
	Statements []Statement
}

func (ie *IfExpression) expressionNode()      {}
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpression) String() string {
	// ... a method to generate a string representation ...
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	// ... a method to generate a string representation ...
}
