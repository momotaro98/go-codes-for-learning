package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/comhash", hander)
	log.Fatal(http.ListenAndServe(":5000", r))
}

type resCom struct {
	OK bool `json: "ok"`
}

func hander(w http.ResponseWriter, r *http.Request) {
	stored := r.URL.Query().Get("stored")
	input := r.URL.Query().Get("input")

	err := bcrypt.CompareHashAndPassword([]byte(stored), []byte(input))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		json.NewEncoder(w).Encode(resCom{OK: false})
		return
	}
	if err != nil {
		log.Println("got error in compare hash", err)
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		json.NewEncoder(w).Encode(resCom{OK: false})
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	json.NewEncoder(w).Encode(resCom{OK: true})
	return
}
