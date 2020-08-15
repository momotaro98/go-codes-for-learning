package main

import (
	"net/http"
	"os"
	"time"

	"github.com/newrelic/go-agent/v3/newrelic"
)

var (
	app    *newrelic.Application
	client = &http.Client{
		Transport: newrelic.NewRoundTripper(nil),
		Timeout:   time.Duration(10) * time.Second,
	}
)

func init() {
	var err error
	app, err = newrelic.NewApplication(
		newrelic.ConfigAppName("ikeda-private-test"),
		newrelic.ConfigLicense(os.Getenv("NEW_RELIC_LICENSE_KEY")),
		newrelic.ConfigDistributedTracerEnabled(true),
		newrelic.ConfigDebugLogger(os.Stdout),
	)
	if err != nil {
		panic(err)
	}
}

// Middleware to create/end NewRelic transaction
func nrt(inner http.Handler) http.Handler {
	mw := func(w http.ResponseWriter, r *http.Request) {
		txn := app.StartTransaction(r.URL.Path)
		defer txn.End()

		r = newrelic.RequestWithTransactionContext(r, txn)

		txn.SetWebRequestHTTP(r)
		w = txn.SetWebResponse(w)
		inner.ServeHTTP(w, r)
	}
	return http.HandlerFunc(mw)
}
