package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

// 失敗しないHTTPサーバ
func hello(w http.ResponseWriter, req *http.Request) {
	buf, _ := ioutil.ReadAll(req.Body)
	bufForLog := bytes.NewBuffer(buf)
	bufForApp := bytes.NewBuffer(buf)

	fmt.Println(bufForLog)

	req.Body = ioutil.NopCloser(bufForApp)
	fmt.Println(req.Body)

	fmt.Fprintf(w, "hello\n")
}

// // 失敗するHTTPサーバ
// func hello(w http.ResponseWriter, req *http.Request) {
// 	bufForLog, _ := ioutil.ReadAll(req.Body) // この時点でbody.sawEOFがtrueになり次回はもう読み取れないようになっている
// 	fmt.Println(bufForLog)                   // [123 10 32 32 32 32 34 97 34 10 125]

// 	bufForApp, _ := ioutil.ReadAll(req.Body) // すでに"消費"済みなので読み取れない
// 	fmt.Println(bufForApp)                   // []

// 	fmt.Fprintf(w, "hello\n")
// }

func main() {
	// 自作ルータをセット
	r := &Router{}
	r.HandleFunc("/hello", hello)
	r.AddMiddleware(LoggingMiddleware)

	http.ListenAndServe(":8090", r)
}
