package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/login", login)
	r.HandleFunc("/schedule", profile)
	port := 8090
	log.Println("listening port:", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}

type LoginForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func login(w http.ResponseWriter, r *http.Request) {
	log.Println("login request")
	decoder := json.NewDecoder(r.Body)
	var form LoginForm
	err := decoder.Decode(&form)
	if err != nil {
		panic(err)
	}
	log.Println("email:", form.Email)

	time.Sleep(time.Millisecond * 350)

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		JWT string `json:"jwt_token"`
	}{
		JWT: RandStringRunes(20),
	})
	return
}

func profile(w http.ResponseWriter, r *http.Request) {
	log.Println("profile request")
	log.Println("r.Header", r.Header)
	time.Sleep(time.Millisecond * 100)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		Profile string `json:"profile"`
	}{
		Profile: "profile text",
	})
	return
}
