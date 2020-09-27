package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/momotaro98/go-codes-for-learning/jit-compiler/calculator"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		fmt.Println(line)
		l := calculator.NewLexer(line)
		p := calculator.NewParser(l)

		ast, err := p.Parse()
		if err != nil {
			panic(err)
		}
		fmt.Println("ast:", ast)

		evaluated := calculator.Eval(ast)
		fmt.Println(evaluated)
	}
}

func main() {
	Start(os.Stdin, os.Stdout)
}
