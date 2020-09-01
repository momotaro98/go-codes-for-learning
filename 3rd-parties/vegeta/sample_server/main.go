package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

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

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		JWT string `json:"jwt_token"`
	}{
		JWT: "xxx.yyy.zzz",
	})
	return
}

func profile(w http.ResponseWriter, r *http.Request) {
	log.Println("profile request")
	log.Println("r.Header", r.Header)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		Profile string `json:"profile"`
	}{
		Profile: "profile text",
	})
	return
}
