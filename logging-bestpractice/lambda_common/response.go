package lambda_common

import (
	"encoding/json"
	"net/http"
)

// Response is ...
type Response struct {
	status int
	body   interface{}
	err    error
}

// StatusCode is ...
func (resp *Response) StatusCode() int {
	return resp.status
}

// Body is ...
func (resp *Response) Body() interface{} {
	return resp.body
}

// BodyJSONString is ...
func (resp *Response) BodyJSONString() (string, error) {
	if resp.body == nil {
		return "", nil
	}
	b, err := json.Marshal(resp.body)
	if err != nil {
		return "", err
	}
	return string(b), nil
	/*
		var buf = new(bytes.Buffer)
		if err := json.NewEncoder(buf).Encode(resp.body); err != nil {
			return "", err
		}
		return buf.String(), nil
	*/
}

// Error is ...
func (resp *Response) Error() error {
	return resp.err
}

// OK is ...
func OK(obj interface{}) Response {
	return Response{
		status: http.StatusOK,
		body:   obj,
	}
}

// Created is ...
func Created(obj interface{}) Response {
	return Response{
		status: http.StatusCreated,
		body:   obj,
	}
}

// BadRequest is ...
func BadRequest(err error) Response {
	return buildErrorResponse(http.StatusBadRequest, err)
}

// Unauthorized is ...
func Unauthorized(err error) Response {
	return buildErrorResponse(http.StatusUnauthorized, err)
}

// Forbidden is ...
func Forbidden(err error) Response {
	return buildErrorResponse(http.StatusForbidden, err)
}

// NotFound is ...
func NotFound(err error) Response {
	return buildErrorResponse(http.StatusNotFound, err)
}

// Conflict is ...
func Conflict(err error) Response {
	return buildErrorResponse(http.StatusConflict, err)
}

// InternalServerError is ...
func InternalServerError(err error) Response {
	return buildErrorResponse(http.StatusInternalServerError, err)
}

func buildErrorResponse(status int, err error) Response {
	var messages interface{}
	switch e := err.(type) {
	case multiErrors:
		messages = e.Errors()
	default:
		messages = []string{
			err.Error(),
		}
	}
	return Response{
		status: status,
		body: map[string]interface{}{
			"messages": messages,
		},
		err: err,
	}
}

type multiErrors interface {
	error
	Errors() []string
}
