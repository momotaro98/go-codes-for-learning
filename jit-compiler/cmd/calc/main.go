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

		/*
			evaluated := evaluator.Eval(program, env)
			if evaluated != nil {
				io.WriteString(out, evaluated.Inspect())
				io.WriteString(out, "\n")
			}
		*/
	}
}

func main() {
	Start(os.Stdin, os.Stdout)
}
