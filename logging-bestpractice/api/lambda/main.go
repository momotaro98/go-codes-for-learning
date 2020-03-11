package main

import (
	"net/http"

	//"git.rarejob.com/rarejob-platform/auth/errors"
	//"git.rarejob.com/rarejob-platform/auth/services"
	//"git.rarejob.com/rarejob-platform/golibs/aws/lambda/api"
	api "git.rarejob.com/shintaro.ikeda/platform_logging/lambda_common"
	//"git.rarejob.com/rarejob-platform/golibs/aws/lambda/validator"
	//"git.rarejob.com/rarejob-platform/golibs/jwt"
)

var (
//service = services.NewService()
)

/*
curl -XGET \
  localhost:3000/token/verify \
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer $TOKEN' \
  | jq .
*/
func handleLambdaEvent(req api.Request) api.Response {
	var (
	//input   services.VerifyToken
	//payload jwt.Payload
	)

	// convert to http request
	//httpReq, err := req.ToHTTPRequest()
	//if err != nil {
	//	// api.Request -> http.Request変換時のエラーは通常起こり得ない
	//	return api.InternalServerError(err)
	//}

	// verify token
	//input, payload, err = verifyToken(httpReq)
	//if err != nil {
	//	return errors.Handlers(err)
	//}

	// validate
	//if err = validator.ValidateStruct(&input); err != nil {
	//	return api.BadRequest(err)
	//}

	// input
	//if err = service.VerifyToken(input); err != nil {
	//	return errors.Handlers(err)
	//}

	// common claimsを出力させないため詰め替えて返却
	//return api.OK(&jwt.Payload{
	//	ProductID:  payload.ProductID,
	//	CategoryID: payload.CategoryID,
	//	GroupID:    payload.GroupID,
	//	ID:         payload.ID,
	//})

	return api.OK(struct{ ProductID string }{ProductID: "product-ID"})
}

func main() {
	api.Start(handleLambdaEvent)
}

func verifyToken(req *http.Request) (err error) {
	return nil
}
