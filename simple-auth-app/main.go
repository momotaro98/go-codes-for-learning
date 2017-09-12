package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

type templateHandler struct {
	filename string
	once     sync.Once
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ =
			template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	data := make(map[string]interface{})
	data["UserData"] = "aaa"
	/*
		if authCookie, err := r.Cookie("auth"); err == nil {
			data["UserData"] = objx.MustFromBase64(authCookie.Value)
		}
	*/
	t.templ.Execute(w, data)
}

func main() {
	var addr = flag.String("host", ":3001", "Address of the application")
	flag.Parse()

	http.Handle("/", MustAuth(&templateHandler{filename: "index.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
	// WEBサーバーを起動
	log.Println("Webサーバを開始する。ポート:", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
