package ast

import (
	"testing"

	"github.com/momotaro98/go-codes-for-learning/Writing-an-Interpreter-in-Go/monkey/token"
)

/*
> 構文解析器に対して文字列の比較を行うことで、可読性の高いテストのレイヤーを追加する方法を教えてくれる。
*/

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
