// parser/parser.go

package parser

import "fmt"

// AST nodes represent the parsed structure of the code

type Statement interface {
	Execute() interface{}
}

type Expression interface {
	Evaluate() interface{}
}

type Identifier struct {
	Value string
}

func (i *Identifier) Evaluate() interface{} {
	return i.Value
}

type Program struct {
	Statements []Statement
}

// Parser represents the parser

type Parser struct {
	lexer        *Lexer
	currentToken Token
	peekToken    Token
}

func NewParser(lexer *Lexer) *Parser {
	p := &Parser{lexer: lexer}
	// Initialize both currentToken and peekToken
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) ParseProgram() *Program {
	program := &Program{}
	program.Statements = []Statement{}

	for p.currentToken.Type != TOKEN_EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() Statement {
	switch p.currentToken.Type {
	case TOKEN_IDENTIFIER:
		return p.parseIdentifierStatement()
	default:
		return nil
	}
}

func (p *Parser) parseIdentifierStatement() Statement {
	ident := &Identifier{Value: p.currentToken.Literal}

	if p.peekToken.Type == TOKEN_SEMICOLON {
		p.nextToken()
		return ident
	} else {
		// Handle other cases like assignments or operations
	}

	return nil
}

// ... Potentially more methods for parsing expressions, operations, etc.
