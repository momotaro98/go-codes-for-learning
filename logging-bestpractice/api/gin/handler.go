package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	// "git.rarejob.com/rarejob-platform/student-account/src/services"
)

const (
	transactionCtxKey = "key_account_transaction_id"
	XTransactionID    = "X-Transaction-ID"
)

// HandlerSampleEndpoint ...
var (
	s = NewService()

	HandlerSampleEndpoint = makeSampleEndpoint(s)

	toNormalContext = func(c *gin.Context) context.Context {
		ctx := context.Background()
		ctx = context.WithValue(ctx,
			transactionCtxKey,
			c.Request.Header.Get(XTransactionID))
		return ctx
	}
)

func makeSampleEndpoint(service Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			err error
		)

		output, err := service.SearchSample(toNormalContext(c), "sample-id-test")
		if err != nil {
			//handleError(c, err) // Skipping handle error
			return
		}

		c.JSON(http.StatusOK, output)
	}
}
