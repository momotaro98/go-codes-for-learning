package hello

// extern はImport

//extern void hello();
import "C"
import "fmt"

// goHelloはhello.cで利用する

//export goHello
func goHello() {
	fmt.Println("hello")
}

// Helloは外のアプリが利用する

func Hello() {
	C.hello()
}
