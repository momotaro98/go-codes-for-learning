package lambda_common

import (
	"os"
	"time"

	logger "git.rarejob.com/shintaro.ikeda/platform_logging/new_logger"
)

var (
	jst, _      = time.LoadLocation("Asia/Tokyo")
	hostName, _ = os.Hostname()
)

func writeAccessLog(l logger.Logger, begin time.Time, req Request, resp Response) {
	var (
		aCtx = newAccessLogContext(begin, req, resp)
		//firstFieldKey = logFieldEntries[0].key
		//buf           = new(bytes.Buffer)
		fields = make([]logger.Field, len(logFieldEntries))
	)
	for i, entry := range logFieldEntries {
		if val := entry.valueFunc(aCtx); val != nil {
			fields[i] = logger.F(entry.key, val)
			// [Old]
			//buf.WriteString(entry.key)
			//buf.WriteByte(':')
			//buf.WriteString(fmt.Sprint(val))
		}
	}
	//buf.WriteByte('\n')
	// write access log
	//buf.WriteTo(os.Stderr)

	// [New]
	l.Info(req.Headers[logger.XTransactionID],
		req.Resource,
		"Request to lambda function",
		fields...,
	)
}

type accessLogContext struct {
	begin time.Time
	req   Request
	resp  Response
}

func newAccessLogContext(begin time.Time, req Request, resp Response) *accessLogContext {
	return &accessLogContext{
		begin: begin,
		req:   req,
		resp:  resp,
	}
}

////////////////////////////////////////////
// logFieldEntry
////////////////////////////////////////////

type logFieldEntry struct {
	key       string
	valueFunc func(*accessLogContext) interface{}
}

var logFieldEntries = []logFieldEntry{
	//{
	//	key: "time",
	//	valueFunc: func(c *accessLogContext) interface{} {
	//		now := nowFunc()
	//		if now.Location() != jst {
	//			now = now.In(jst)
	//		}
	//		return now.Format(time.RFC3339)
	//	},
	//},
	{
		key: "host",
		valueFunc: func(c *accessLogContext) interface{} {
			return hostName
		},
	},
	{
		key: "client_ip",
		valueFunc: func(c *accessLogContext) interface{} {
			return c.req.RequestContext.Identity.SourceIP
		},
	},
	{
		key: "method",
		valueFunc: func(c *accessLogContext) interface{} {
			return c.req.HTTPMethod
		},
	},
	{
		key: "path",
		valueFunc: func(c *accessLogContext) interface{} {
			return c.req.Path
		},
	},
	{
		key: "query",
		valueFunc: func(c *accessLogContext) interface{} {
			if len(c.req.QueryStringParameters) == 0 {
				return nil
			}
			return c.req.QueryStringParameters
		},
	},
	{
		key: "status",
		valueFunc: func(c *accessLogContext) interface{} {
			return c.resp.status
		},
	},
	{
		key: "response_time",
		valueFunc: func(c *accessLogContext) interface{} {
			return nowFunc().Sub(c.begin).Round(time.Millisecond)
		},
	},
	//{
	//	key: "transaction",
	//	valueFunc: func(c *accessLogContext) interface{} {
	//		if txID, exists := c.req.Headers["X-Transaction-ID"]; exists {
	//			return txID
	//		}
	//		return "unknown"
	//	},
	//},
	{
		key: "stage",
		valueFunc: func(c *accessLogContext) interface{} {
			return c.req.RequestContext.Stage
		},
	},
	{
		key: "comment",
		valueFunc: func(c *accessLogContext) interface{} {
			if c.resp.err == nil {
				return nil
			}
			return c.resp.err.Error()
		},
	},
}
