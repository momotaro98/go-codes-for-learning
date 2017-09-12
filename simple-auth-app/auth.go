package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/markbates/goth/gothic"
)

func MustAuth(handler http.Handler) http.Handler {
	return &authHandler{next: handler}
}

type authHandler struct {
	next http.Handler
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if cookie, err := r.Cookie("auth"); err == http.ErrNoCookie || cookie.Value == "" {
		// 未承認
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
	} else if err != nil {
		// 何らかのエラーが発生
		panic(err.Error())
	} else {
		// 成功。ラップされたハンドラを呼び出す
		h.next.ServeHTTP(w, r)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	segs := strings.Split(r.URL.Path, "/")
	action := segs[2]
	// provider := segs[3]
	switch action {
	case "login":

		// goth goth goth
		// try to get the user without re-authenticating
		if gothUser, err := gothic.CompleteUserAuth(w, r); err == nil {
			fmt.Println(gothUser)
		} else {
			gothic.BeginAuthHandler(w, r)
		}
		// goth goth goth

	case "callback":

		// goth goth goth
		user, err := gothic.CompleteUserAuth(w, r)
		if err != nil {
			fmt.Fprintln(w, r)
			return
		}
		fmt.Println(user)
		// goth goth goth

		// データをCookieにしこむ
		/*
			authCookieValue := objx.New(map[string]interface{}{
				"userid":     chatUser.UniqueID(),
				"name":       user.Name(),
				"avatar_url": avatarURL,
			}).MustBase64()
			http.SetCookie(w, &http.Cookie{
				Name:  "auth",
				Value: authCookieValue,
				Path:  "/"})
		*/
		w.Header()["Location"] = []string{"/"}
		w.WriteHeader(http.StatusTemporaryRedirect)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "アクション%sには非対応です", action)
	}
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "auth",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
	w.Header()["Location"] = []string{"/"}
	w.WriteHeader(http.StatusTemporaryRedirect)
}
