package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// BenchmarkHTTPClientServer benchmarks both the HTTP client and the HTTP server,
// on small requests.
func BenchmarkHTTPClientServer(b *testing.B) {
	msg := []byte("Hello world.\n")
	ts := httptest.NewServer(http.HandlerFunc(Handler))
	defer ts.Close()

	tr := &http.Transport{}
	defer tr.CloseIdleConnections()
	cl := &http.Client{
		Transport: tr,
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, err := cl.Get(ts.URL)
		if err != nil {
			b.Fatal("Get:", err)
		}
		all, err := ioutil.ReadAll(res.Body)
		if err != nil {
			b.Fatal("ReadAll:", err)
		}
		if !bytes.Equal(all, msg) {
			b.Fatalf("Got body %q; want %q", all, msg)
		}
	}
}
