package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	goji "goji.io"
	"goji.io/pat"

	"golang.org/x/crypto/bcrypt"

	"time"

	"github.com/newrelic/go-agent/v3/newrelic"
)

func compServer() {
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/comhash"), compHandler)
	log.Fatal(http.ListenAndServe(":5000", mux))
}

type resCom struct {
	OK bool `json:"ok"`
}

func compHandler(w http.ResponseWriter, r *http.Request) {
	stored := r.URL.Query().Get("stored")
	input := r.URL.Query().Get("input")

	err := bcrypt.CompareHashAndPassword([]byte(stored), []byte(input))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		// log.Println("got error in compare hash", err, "stored:", stored, "input:", input)
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		json.NewEncoder(w).Encode(resCom{OK: false})
		return
	}
	if err != nil {
		// log.Println("got error in compare hash", err, "stored:", stored, "input:", input)
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		json.NewEncoder(w).Encode(resCom{OK: false})
		return
	}

	// log.Println("comp OK!", "stored:", stored, "input:", input)

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	json.NewEncoder(w).Encode(resCom{OK: true})
	return
}

func main() {
	go compServer()

	mux := goji.NewMux()
	mux.Use(nrt)
	mux.HandleFunc(pat.Get("/sample"), sampleHandler)
	log.Fatal(http.ListenAndServe(":8000", mux))
}

func sampleHandler(w http.ResponseWriter, r *http.Request) {
	checklogic(r.Context())

	ok, err := conCompareHashAndPassword(r, "abc", "def")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func checklogic(ctx context.Context) {
	s := newrelic.FromContext(ctx).StartSegment("checklogic")
	defer s.End()
	time.Sleep(time.Millisecond * 5)
}

func conCompareHashAndPassword(r *http.Request, stored, input string) (bool, error) {
	stored = url.QueryEscape(stored)
	input = url.QueryEscape(input)
	url := fmt.Sprintf("http://localhost:5000/comhash?stored=%s&input=%s", stored, input)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req = req.WithContext(r.Context())
	if err != nil {
		return false, err
	}
	// res, err := http.DefaultClient.Do(req)
	res, err := client.Do(req) // client is wrapped by new relic
	if err != nil {
		log.Printf("err in http request in conCompareHashAndPassword")
		return false, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Printf("status code: %d; body: %s\n", res.StatusCode, res.Body)
		return false, fmt.Errorf("status code: %d; body: %s", res.StatusCode, res.Body)
	}

	type Response struct {
		OK bool `json:"ok"`
	}
	ress := &Response{}
	err = json.NewDecoder(res.Body).Decode(ress)
	if err != nil {
		log.Printf("decode failed in conCompareHashAndPassword")
		return false, err
	}
	return ress.OK, nil
}
