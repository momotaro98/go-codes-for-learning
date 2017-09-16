package main

import (
	"fmt"
	"net/http"

	"github.com/markbates/goth/gothic"
	"github.com/stretchr/objx"
)

// MustAuth forces user to be authenticated
func MustAuth(handler http.Handler) http.Handler {
	return &authHandler{next: handler}
}

type authHandler struct {
	next http.Handler
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if cookie, err := r.Cookie("auth"); err == http.ErrNoCookie || cookie.Value == "" {
		// 未承認時はログイン画面へ返す
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
	} else {
		// 承認成功。ラップされたハンドラを呼び出す
		h.next.ServeHTTP(w, r)
	}
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	// アプリ用データをCookieにしこむ
	authCookieValue := objx.New(map[string]interface{}{
		"name":       user.Name,
		"avatar_url": user.AvatarURL,
	}).MustBase64()
	http.SetCookie(w, &http.Cookie{
		Name:  "auth",
		Value: authCookieValue,
		Path:  "/",
	})

	w.Header()["Location"] = []string{"/"}
	w.WriteHeader(http.StatusTemporaryRedirect)
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
