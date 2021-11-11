package main

import (
	"io/ioutil"
	"net/http"
)

// アプリケーションHTTPサーバ
func hello(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)
	if string(b) != `{"a": "b"}` {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"res": "No hello"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"res": "hello"}`))
}

func main() {
	// 自作ルータをセット
	r := &Router{}
	r.HandleFunc("/hello", hello)

	http.ListenAndServe(":8090", r)
}
