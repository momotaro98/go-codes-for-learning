package main

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/momotaro98/go-codes-for-learning/microservice-logging/logger"
)

var (
	nowFunc     = time.Now
	hostName, _ = os.Hostname()
)

func LoggerWithWriter(l logger.Logger, excludePaths ...string) gin.HandlerFunc {
	var (
		containsExcludePath = func(path string) bool {
			for _, p := range excludePaths {
				if p == path {
					return true
				}
			}
			return false
		}
	)
	return func(c *gin.Context) {
		defer func(ctx *accessLogContext) {
			if containsExcludePath(ctx.Request.URL.Path) {
				return
			}
			fields := make([]logger.Field, len(logFieldEntries))
			for i, entry := range logFieldEntries {
				if val := entry.valueFunc(ctx); val != nil {
					fields[i] = logger.F(entry.key, val)
				}
			}
			l.Info(c.Request.Header.Get(logger.XTransactionID),
				"Sample Service Gin Access",
				fields...,
			)
		}(newAccessLogContext(c))

		c.Next()
	}
}

type accessLogContext struct {
	*gin.Context
	begin time.Time
}

func newAccessLogContext(c *gin.Context) *accessLogContext {
	return &accessLogContext{
		Context: c,
		begin:   nowFunc(),
	}
}

type logFieldEntry struct {
	key       string
	valueFunc func(*accessLogContext) interface{}
}

var logFieldEntries = []logFieldEntry{
	{
		key: "host",
		valueFunc: func(c *accessLogContext) interface{} {
			return hostName
		},
	},
	{
		key: "path",
		valueFunc: func(c *accessLogContext) interface{} {
			return c.Request.URL.EscapedPath()
		},
	},
	{
		key: "response_time",
		valueFunc: func(c *accessLogContext) interface{} {
			return nowFunc().Sub(c.begin).Round(time.Millisecond)
		},
	},
}
