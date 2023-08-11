// parser/lexer.go

package parser

import (
	"bufio"
	"io"
	"strings"
	"unicode"
)

type TokenType int

const (
	// Define your token types here
	TOKEN_EOF TokenType = iota
	TOKEN_ILLEGAL
	TOKEN_IDENTIFIER
	TOKEN_SEMICOLON
	// ... other tokens like TOKEN_ASSIGN, TOKEN_PLUS, etc.
)

type Token struct {
	Type    TokenType
	Literal string
}

type Lexer struct {
	reader *bufio.Reader
	ch     rune // current character
}

func NewLexer(r io.Reader) *Lexer {
	lexer := &Lexer{reader: bufio.NewReader(r)}
	lexer.readChar() // Initialize ch
	return lexer
}

func (l *Lexer) readChar() {
	ch, _, err := l.reader.ReadRune()
	if err != nil {
		l.ch = 0 // EOF
	} else {
		l.ch = ch
	}
}

func (l *Lexer) NextToken() Token {
	l.skipWhitespace()

	switch l.ch {
	case ';':
		return Token{TOKEN_SEMICOLON, string(l.ch)}
	case 0:
		return Token{TOKEN_EOF, ""}
	default:
		if isLetter(l.ch) {
			ident := l.readIdentifier()
			return Token{TOKEN_IDENTIFIER, ident}
		}
		return Token{TOKEN_ILLEGAL, string(l.ch)}
	}
}

func (l *Lexer) readIdentifier() string {
	start := l.ch
	for isLetter(l.ch) {
		l.readChar()
	}
	return string(start)
}

func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(l.ch) {
		l.readChar()
	}
}

func isLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// ... Potentially more helper methods and token types ...
