package calculator

import (
	"bytes"
	"fmt"
	"strconv"
)

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	INT = "INT"

	// Operators
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
)

type Token struct {
	Type    TokenType
	Literal string
}

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()

	switch l.ch {
	case '+':
		tok = newToken(PLUS, l.ch)
	case '-':
		tok = newToken(MINUS, l.ch)
	case '/':
		tok = newToken(SLASH, l.ch)
	case '*':
		tok = newToken(ASTERISK, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = EOF
	default:
		if isDigit(l.ch) {
			tok.Type = INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func newToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}

// AST

type Node interface {
	TokenLiteral() string
	String() string
}

type IntegerLiteral struct {
	Token Token
	Value int64
}

func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }

type InfixExpression struct {
	Token    Token // The operator token, e.g. +
	Left     Node
	Operator string
	Right    Node
}

func (oe *InfixExpression) TokenLiteral() string { return oe.Token.Literal }
func (oe *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(oe.Left.String())
	out.WriteString(" " + oe.Operator + " ")
	out.WriteString(oe.Right.String())
	out.WriteString(")")

	return out.String()
}

// Parser

type Parser struct {
	l *Lexer

	errors []string

	curToken  Token
	peekToken Token
}

func NewParser(l *Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}
	// Read Two tokens. Both curToken and peekToken are set.
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) Parse() (Node, error) {
	ast, err := p.AddOrSub()
	if err != nil {
		return nil, err
	}
	return ast, nil
}

func (p *Parser) AddOrSub() (Node, error) {
	left, err := p.MulOrDiv()
	if err != nil {
		return nil, err
	}
loop:
	for p.curToken.Type != EOF {
		switch p.curToken.Type {
		case PLUS, MINUS:
			token := p.curToken
			p.nextToken()
			right, err := p.MulOrDiv()
			if err != nil {
				return nil, err
			}
			node := &InfixExpression{
				Token:    token,
				Left:     left,
				Operator: token.Literal,
				Right:    right,
			}
			left = node
		default:
			break loop
		}
	}
	return left, nil
}

func (p *Parser) MulOrDiv() (Node, error) {
	left, err := p.Num()
	if err != nil {
		return nil, err
	}
loop:
	for p.curToken.Type != EOF {
		switch p.curToken.Type {
		case ASTERISK, SLASH:
			token := p.curToken
			p.nextToken()
			right, err := p.Num()
			if err != nil {
				return nil, err
			}
			node := &InfixExpression{
				Token:    token,
				Left:     left,
				Operator: token.Literal,
				Right:    right,
			}
			left = node
		default:
			break loop
		}
	}
	return left, nil
}

func (p *Parser) Num() (Node, error) {
	if p.curToken.Type == INT {
		token := p.curToken
		p.nextToken()

		value, err := strconv.ParseInt(token.Literal, 0, 64)
		if err != nil {
			return nil, err
		}

		node := &IntegerLiteral{
			Token: token,
			Value: value,
		}

		return node, nil
	}
	return nil, fmt.Errorf("error in Num()")
}

// Eval

func Eval(ast Node) int64 {
	switch node := ast.(type) {
	case *IntegerLiteral:
		return node.Value
	case *InfixExpression:
		left := Eval(node.Left)
		right := Eval(node.Right)
		switch node.Operator {
		case "+":
			return left + right
		case "-":
			return left - right
		case "*":
			return left * right
		case "/":
			return left / right
		}
	}
	return 0
}
