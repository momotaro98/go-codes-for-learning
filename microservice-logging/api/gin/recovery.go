package main

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"

	"github.com/momotaro98/go-codes-for-learning/microservice-logging/logger"
)

var (
	defaultStackSize = 4 << 10 // 4 kb
)

// Recovery is a Gin middleware to handle panic error of application.
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				err, ok := r.(error)
				if !ok {
					err = fmt.Errorf("%v", r)
				}
				var (
					stack  = make([]byte, defaultStackSize)
					length = runtime.Stack(stack, true)
				)
				logger.DefaultLogger().Error(c.Request.Header.Get(logger.XTransactionID),
					ServiceName,
					/* msg */ "panic recovered",
					logger.E(err),
					logger.F("stack", fmt.Sprintf("%s ", stack[:length])),
				)
				c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{
					"errors": []string{"unexpected error"},
				})
			}
		}()
		c.Next()
	}
}
