package main

import (
	"net/http"
)

// シンプルHTTPルータ
type Router struct {
	pathFuncMap map[string]func(http.ResponseWriter, *http.Request)
	middlewares []func(next http.Handler) http.Handler
}

func (r *Router) HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) {
	if r.pathFuncMap == nil {
		r.pathFuncMap = make(map[string]func(http.ResponseWriter, *http.Request))
	}
	r.pathFuncMap[path] = f
}

func (r *Router) AddMiddleware(mf func(handler http.Handler) http.Handler) {
	r.middlewares = append(r.middlewares, mf)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// match path
	var handler http.Handler
	for p, f := range r.pathFuncMap {
		if p == req.URL.Path {
			handler = http.HandlerFunc(f)
		}
	}
	if handler == nil {
		handler = http.NotFoundHandler()
	} else {
		// set middleware
		for i := len(r.middlewares) - 1; i >= 0; i-- {
			handler = r.middlewares[i](handler)
		}
	}

	// serve http
	handler.ServeHTTP(w, req)
}
