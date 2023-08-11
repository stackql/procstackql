// parser/ast/node.go

package ast

// Node represents a node in the Abstract Syntax Tree.
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement nodes represent actions. Examples: let x = 5, x + y.
type Statement interface {
	Node
	statementNode()
}

// Expression nodes represent values. Examples: literals, variable references.
type Expression interface {
	Node
	expressionNode()
}

// LetStatement represents a statement like "let x = 5;".
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) String() string {
	// ... a method to generate a string representation of the node ...
}

// Identifier represents a variable name.
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

// Other nodes, such as Expressions for numbers, arithmetic operations, 
// function calls, control structures, etc., would also be defined here.
