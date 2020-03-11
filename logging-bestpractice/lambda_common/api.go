package lambda_common

import (
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	//"git.rarejob.com/shintaro.ikeda/platform_logging/logger/fio"
	logger "git.rarejob.com/shintaro.ikeda/platform_logging/new_logger"
)

var (
	nowFunc = time.Now
	// Version is replaced by go build ldflags and used to show application version.
	Version = "unknown"
	// Revision is replaced by go build ldflags and used to show application revision.
	Revision = "unknown"
	// log is logger
	log = logger.DefaultLogger()
	//log := logger.NewLogger(logger.NewConfig(
	//	logger.WithMinLevel(logger.Levels.Info),
	//	logger.WithOut(os.Stdout)))
)

// Func is ...
type Func func(Request) Response

// Start is ...
func Start(apiFunc Func, options ...ResponseOption) {
	// [Old]
	// setup logger
	//logger.SetupRootLogger(logger.NewConfig("lambda-log",
	//	logger.WithLevel(logger.Levels.Info),
	//	logger.WithOut(fio.NewBufferedWriter()),
	//))
	lambdaFunc := func(event events.APIGatewayProxyRequest) (response events.APIGatewayProxyResponse, err error) {
		var (
			begin = nowFunc()
			req   = NewRequest(&event)
			resp  Response
		)
		defer func() {
			if r := recover(); r != nil {
				err, ok := r.(error)
				if !ok {
					err = fmt.Errorf("%v", r)
				}
				// [Old]
				//logger.Error("unexpected error",
				//	logger.E(err),
				//)
				// [New]
				log.Error(req.Headers[logger.XTransactionID], req.Resource, "unexpected error", logger.E(err))
			}
			writeAccessLog(log, begin, req, resp)
		}()

		// process api function
		resp = apiFunc(req)

		response = events.APIGatewayProxyResponse{
			StatusCode: resp.StatusCode(),
			Headers: map[string]string{
				"Content-Type": "application/json",
				//"X-Version":    versioning.Version,
				//"X-Revision":   versioning.Revision,
				"X-Version":  Version,
				"X-Revision": Revision,
			},
		}
		response.Body, err = resp.BodyJSONString()
		for _, option := range options {
			option(&response)
		}
		return response, err
	}
	// process lambda function
	lambda.Start(lambdaFunc)
}

// ResponseOption is ...
type ResponseOption func(*events.APIGatewayProxyResponse)

// WithHeader is ...
func WithHeader(key, value string) ResponseOption {
	return func(response *events.APIGatewayProxyResponse) {
		response.Headers[key] = value
	}
}
