package lexer

import "with_legs/token"

type Lexer struct {
	input        string // entire input
	position     int    // position in input
	readPosition int    // next reading position
	ch           byte   // current char being read, byte only supports Unicode. rune supports Unicode and UTF-8
}

// initializer that return pointer to new Lexer obj
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// helper method to assign position values in Lexer
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 //ASCII nul
	} else {
		l.ch = l.input[l.readPosition] // if ch supported more than Unicode, then this would need to adapt to multiple bytes wide
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	}

	l.readChar()
	return tok
}
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
