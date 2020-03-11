package lambda_common

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

// Request is ...
type Request struct {
	*events.APIGatewayProxyRequest
}

// BindJSON is ...
func (req *Request) BindJSON(obj interface{}) error {
	if err := json.Unmarshal([]byte(req.Body), obj); err != nil {
		return err
	}
	//return validator.ValidateStruct(obj) // [Old]

	// [New] Use validation of microservice policy
	return nil
}

// ToHTTPRequest is ...
func (req *Request) ToHTTPRequest() (*http.Request, error) {
	httpReq, err := http.NewRequest(
		strings.ToUpper(req.HTTPMethod),
		req.Path,
		nil,
	)
	if err != nil {
		return nil, err
	}
	// setup request headers
	for h := range req.Headers {
		httpReq.Header.Add(h, req.Headers[h])
	}
	return httpReq, nil
}

// NewRequest is ...
func NewRequest(event *events.APIGatewayProxyRequest) Request {
	return Request{
		APIGatewayProxyRequest: event,
	}
}
