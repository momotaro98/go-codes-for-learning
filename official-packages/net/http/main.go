package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// HTTPサーバ
func hello(w http.ResponseWriter, req *http.Request) {
	b1, _ := ioutil.ReadAll(req.Body) // この時点でbody.sawEOFがtrueになり次回はもう読み取れないようになっている
	b2, _ := ioutil.ReadAll(req.Body) // 読み取れない
	fmt.Println(b1)                   // [123 10 32 32 32 32 34 97 34 10 125]
	fmt.Println(b2)                   // []

	fmt.Fprintf(w, "hello\n")
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}
