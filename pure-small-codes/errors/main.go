package main

import (
	"errors"
	"fmt"
)

func service(text string) error {
	_, err := ParseLibFunc(text)
	if err != nil {
		return fmt.Errorf("service error with '%w'", err) // Unwrapを実装したエラーを返す
	}
	fmt.Println("service finished successfully")
	return err
}

func main() {
	if err := service("2020/02/14"); err != nil {
		fmt.Println(err) // service error with 'lib: parse error'

		// Unwrap関数 でライブラリ関数を取り出す
		wrappedErr := errors.Unwrap(err)
		fmt.Println(wrappedErr)

		// Is
		fmt.Println(errors.Is(wrappedErr, ErrParse)) // true
		fmt.Println(errors.Is(err, ErrParse))        // true

		if errors.Is(err, ErrParse) {
			// handle error
		}

		// As
		var e *LibError
		if errors.As(err, &e) {
			fmt.Println(e)
			fmt.Println(e.Kind())
		}
	}
}
