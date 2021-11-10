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

type Router struct {
	PathFuncMap map[string]func(http.ResponseWriter, *http.Request)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// set middle-ware

	// match path
	var handler http.Handler
	for p, f := range r.PathFuncMap {
		fmt.Println(req.URL.Path)
		if p == req.URL.Path {
			handler = http.HandlerFunc(f)
		}
	}
	if handler == nil {
		handler = http.NotFoundHandler()
	}

	// serve http
	handler.ServeHTTP(w, req)
}

func main() {
	r := &Router{
		PathFuncMap: make(map[string]func(http.ResponseWriter, *http.Request)),
	}
	r.PathFuncMap["/hello"] = hello

	http.ListenAndServe(":8090", r)
}
