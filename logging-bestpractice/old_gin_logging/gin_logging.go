package old_gin_logging

import (
	"bytes"
	"fmt"
	"io"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	//authlib "git.rarejob.com/rarejob-platform/golibs/auth"
	//authmiddleware "git.rarejob.com/rarejob-platform/golibs/gin/middlewares/auth"
	"github.com/gin-gonic/gin"
)

var (
	nowFunc     = time.Now
	jst, _      = time.LoadLocation("Asia/Tokyo")
	hostName, _ = os.Hostname()
)

// [Old] The `Logger` is not used in existing Student Account code
// Logger is ...
//func Logger(excludePaths ...string) gin.HandlerFunc {
//	return LoggerWithWriter(gin.DefaultWriter, excludePaths...)
//}

// LoggerWithWriter is ...
func LoggerWithWriter(out io.Writer, excludePaths ...string) gin.HandlerFunc {
	var (
		pool = &sync.Pool{
			New: func() interface{} {
				return new(bytes.Buffer)
			},
		}
		firstFieldKey = func() string {
			if len(logFieldEntries) > 0 {
				return logFieldEntries[0].key
			}
			return ""
		}()
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
			buf := pool.Get().(*bytes.Buffer)
			buf.Reset()
			for _, entry := range logFieldEntries {
				if val := entry.valueFunc(ctx); val != nil {
					if entry.key != firstFieldKey {
						buf.WriteByte('\t')
					}
					buf.WriteString(entry.key)
					buf.WriteByte(':')
					buf.WriteString(fmt.Sprint(val))
				}
			}
			buf.WriteByte('\n')

			// write access log
			buf.WriteTo(out)
			pool.Put(buf)
		}(newAccessLogContext(c))

		// process request
		c.Next()
	}
}

////////////////////////////////////////////
// accessLogContext
////////////////////////////////////////////

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

//func (c *accessLogContext) payload() authlib.Payload {
//	// TODO 本来は不要な実装なのでauthlibリファクタ後に修正
//	if payloadValue, exists := c.Get(authmiddleware.PayloadKeyName); exists {
//		if payload, ok := payloadValue.(authlib.Payload); ok {
//			return payload
//		}
//	}
//	return nil
//}

////////////////////////////////////////////
// logFieldEntry
////////////////////////////////////////////

type logFieldEntry struct {
	key       string
	valueFunc func(*accessLogContext) interface{}
}

var logFieldEntries = []logFieldEntry{
	{
		key: "time",
		valueFunc: func(c *accessLogContext) interface{} {
			now := nowFunc()
			if now.Location() != jst {
				now = now.In(jst)
			}
			return now.Format(time.RFC3339)
		},
	},
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
		key: "transaction",
		valueFunc: func(c *accessLogContext) interface{} {
			if txID := c.Request.Header.Get("X-Transaction-ID"); len(txID) > 0 {
				return txID
			}
			return "unknown"
		},
	},
	//{
	//	key: "product",
	//	valueFunc: func(c *accessLogContext) interface{} {
	//		if payload := c.payload(); payload != nil {
	//			return payload.ProductID()
	//		}
	//		return "unknown"
	//	},
	//},
	//{
	//	key: "user_category",
	//	valueFunc: func(c *accessLogContext) interface{} {
	//		if payload := c.payload(); payload != nil {
	//			return payload.ActorUserCategoryID()
	//		}
	//		return "unknown"
	//	},
	//},
	//{
	//	key: "user_group",
	//	valueFunc: func(c *accessLogContext) interface{} {
	//		if payload := c.payload(); payload != nil {
	//			return payload.ActorUserGroupID()
	//		}
	//		return "unknown"
	//	},
	//},
	//{
	//	key: "user_id",
	//	valueFunc: func(c *accessLogContext) interface{} {
	//		if payload := c.payload(); payload != nil {
	//			return payload.ActorUserID()
	//		}
	//		return "unknown"
	//	},
	//},
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
