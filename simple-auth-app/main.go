package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"text/template"

	"github.com/gorilla/pat"
	// "github.com/gorilla/sessions"
	"github.com/markbates/goth"
	// "github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/gplus"
)

func init() {
	goth.UseProviders(
		gplus.New(os.Getenv("GPLUS_KEY"), os.Getenv("GPLUS_SECRET"), "http://localhost:3001/auth/gplus/callback"),
	)
}

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

// var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))

func main() {
	var addr = flag.String("host", ":3001", "Address of the application")
	flag.Parse()

	p := pat.New()

	p.Get("/auth/{provider}/callback", callbackHandler)
	// p.Get("/auth/{provider}", gothic.BeginAuthHandler)
	p.Get("/auth/{provider}", providerHandler)
	// http.HandleFunc("/logout", logoutHandler)
	p.Get("/logout", logoutHandler)
	p.Add("GET", "/login", &templateHandler{filename: "login.html"})
	p.Add("GET", "/", MustAuth(&templateHandler{filename: "index.html"}))

	// WEBサーバーを起動
	log.Println("Webサーバを開始する。ポート:", *addr)
	if err := http.ListenAndServe(*addr, p); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
