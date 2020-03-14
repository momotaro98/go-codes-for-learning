package main

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"

	logger "git.rarejob.com/shintaro.ikeda/platform_logging/new_logger"
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
				// [New]
				logger.DefaultLogger().Error(c.Request.Header.Get(logger.XTransactionID),
					ServiceName,
					/* msg */ "panic recovered",
					logger.E(err),
					logger.F("stack", fmt.Sprintf("%s ", stack[:length])),
				)
				// [Old]
				//logger.Error("panic recovered",
				//	logger.E(err),
				//	logger.F("stack", fmt.Sprintf("%s ", stack[:length])),
				//	logger.F("transaction", c.Request.Header.Get("X-Transaction-ID")),
				//)
				c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{
					"errors": []string{"unexpected error"},
				})
			}
		}()
		// process request
		c.Next()
	}
}
