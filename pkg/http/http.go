package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type ResponseError map[string]error

type Req events.APIGatewayProxyRequest

type Res events.APIGatewayProxyResponse

func Response(data interface{}, code int) (Res, error) {
	body, _ := json.Marshal(data)
	return Res{
		Body:       string(body),
		StatusCode: code,
	}, nil
}

func ErrResponse(err error, code int) (Res, error) {
	data := map[string]string{
		"err": err.Error(),
	}
	body, _ := json.Marshal(data)
	return Res{
		Body:       string(body),
		StatusCode: code,
	}, err
}

type RestHandler interface {
	Get(id string, request Req) (Res, error)
	Store(request Req) (Res, error)
	List(request Req) (Res, error)
}

func ParseBody(request Req, castTo interface{}) error {
	return json.Unmarshal([]byte(request.Body), &castTo)
}

type RequestHandleFunc func(request Req) (Res, error)

func Router(h RestHandler) RequestHandleFunc {
	return func(request Req) (Res, error) {
		switch request.HTTPMethod {
		case "GET":
			id := request.PathParameters["id"]
			if id != "" {
				return h.Get(id, request)
			}
			return h.List(request)
		case "POST":
			return h.Store(request)
		default:
			return ErrResponse(errors.New("method not allowed"), http.StatusMethodNotAllowed)
		}
	}
}
