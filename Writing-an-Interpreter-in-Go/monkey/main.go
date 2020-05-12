package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/momotaro98/go-codes-for-learning/Writing-an-Interpreter-in-Go/monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n",
		user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
