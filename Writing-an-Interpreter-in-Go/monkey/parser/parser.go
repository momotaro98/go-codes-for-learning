package parser

import (
	"github.com/momotaro98/go-codes-for-learning/Writing-an-Interpreter-in-Go/monkey/ast"
	"github.com/momotaro98/go-codes-for-learning/Writing-an-Interpreter-in-Go/monkey/lexer"
	"github.com/momotaro98/go-codes-for-learning/Writing-an-Interpreter-in-Go/monkey/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Read Two tokens. Both curToken and peekToken are set.
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
