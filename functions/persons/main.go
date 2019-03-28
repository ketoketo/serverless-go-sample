package main

import (
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/tMatSuZ/serverless-go-sample/pkg/datastore"

	"github.com/tMatSuZ/serverless-go-sample/functions/persons/model"
	helpers "github.com/tMatSuZ/serverless-go-sample/pkg/http"
)

type PersonRepository interface {
	Get(id string) (*model.Person, error)
	List() (*[]model.Person, error)
	Store(person *model.Person) error
}

type Handler struct {
	repository PersonRepository
}

func (h *Handler) Store(request helpers.Req) (helpers.Res, error) {
	var person *model.Person

	if err := helpers.ParseBody(request, &person); err != nil {
		return helpers.ErrResponse(err, http.StatusBadRequest)
	}

	if err := h.repository.Store(person); err != nil {
		return helpers.ErrResponse(err, http.StatusInternalServerError)
	}

	return helpers.Response(map[string]bool{
		"success": true,
	}, http.StatusCreated)
}

func (h *Handler) Get(id string, request helpers.Req) (helpers.Res, error) {
	persons, err := h.repository.Get(id)
	if err != nil {
		return helpers.ErrResponse(err, http.StatusNotFound)
	}

	return helpers.Response(map[string]interface{}{
		"persons": persons,
	}, http.StatusOK)
}

func (h *Handler) List(request helpers.Req) (helpers.Res, error) {
	persons, err := h.repository.List()
	if err != nil {
		return helpers.ErrResponse(err, http.StatusNotFound)
	}
	return helpers.Response(map[string]interface{}{
		"persons": persons,
	}, http.StatusOK)
}

func main() {
	conn, err := datastore.CreateConnection(os.Getenv("REGION"))
	if err != nil {
		log.Panic(err)
	}

	ddb := datastore.NewDynamoDB(conn, os.Getenv("DB_TABLE"))

	repository := model.NewPersonRepository(ddb)

	handler := &Handler{repository}

	router := helpers.Router(handler)

	lambda.Start(router)
}
