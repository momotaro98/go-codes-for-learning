package main

import (
	"net/url"
	"os"
	"strings"
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
				"Sample Service Gin Access",
				fields...,
			)
		}(newAccessLogContext(c))

		// process request
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
		key: "client_ip",
		valueFunc: func(c *accessLogContext) interface{} {
			return c.ClientIP()
		},
	},
	{
		key: "method",
		valueFunc: func(c *accessLogContext) interface{} {
			return c.Request.Method
		},
	},
	{
		key: "path",
		valueFunc: func(c *accessLogContext) interface{} {
			return c.Request.URL.EscapedPath()
		},
	},
	{
		key: "query",
		valueFunc: func(c *accessLogContext) interface{} {
			if len(c.Request.URL.RawQuery) == 0 {
				return nil
			}
			escapedQuery, err := url.QueryUnescape(c.Request.URL.RawQuery)
			if err != nil {
				return c.Request.URL.RawQuery
			}
			return escapedQuery
		},
	},
	{
		key: "status",
		valueFunc: func(c *accessLogContext) interface{} {
			return c.Writer.Status()
		},
	},
	{
		key: "response_time",
		valueFunc: func(c *accessLogContext) interface{} {
			return nowFunc().Sub(c.begin).Round(time.Millisecond)
		},
	},
	{
		key: "comment",
		valueFunc: func(c *accessLogContext) interface{} {
			if c.Errors == nil || len(c.Errors) == 0 {
				return nil
			}
			return strings.Join(c.Errors.Errors(), ",")
		},
	},
}
