package parser

import (
	"with_legs/ast"
	"with_legs/lexer"
	"with_legs/token"
)

type Parser struct {
	l         *lexer.Lexer // instance of lexer that has nextToken
	curToken  token.Token  // pointers for our position in the input
	peekToken token.Token  // peekToken allows us to bring many tokens into one expression
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// Create new AST node as program
// advance tokens
// loop until input is EOF
// Parse tokens according to their tokens as statement
// push statement to program graph
func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
