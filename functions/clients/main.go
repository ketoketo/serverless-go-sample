package main

import (
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/tMatSuZ/serverless-go-sample/pkg/datastore"

	"github.com/tMatSuZ/serverless-go-sample/functions/clients/model"
	helpers "github.com/tMatSuZ/serverless-go-sample/pkg/http"
)

// テストを楽にするため
type ClientRepository interface {
	Get(id string) (*model.Client, error)
	List() (*[]model.Client, error)
	Store(client *model.Client) error
}

type Handler struct {
	repository ClientRepository
}

func (h *Handler) Store(request helpers.Req) (helpers.Res, error) {
	var client *model.Client

	if err := helpers.ParseBody(request, &client); err != nil {
		return helpers.ErrResponse(err, http.StatusBadRequest)
	}

	if err := h.repository.Store(client); err != nil {
		return helpers.ErrResponse(err, http.StatusInternalServerError)
	}

	return helpers.Response(map[string]bool{
		"success": true,
	}, http.StatusCreated)
}

func (h *Handler) Get(id string, request helpers.Req) (helpers.Res, error) {
	clients, err := h.repository.Get(id)
	if err != nil {
		return helpers.ErrResponse(err, http.StatusNotFound)
	}

	return helpers.Response(map[string]interface{}{
		"clients": clients,
	}, http.StatusOK)
}

func (h *Handler) List(request helpers.Req) (helpers.Res, error) {
	clients, err := h.repository.List()
	if err != nil {
		return helpers.ErrResponse(err, http.StatusNotFound)
	}
	return helpers.Response(map[string]interface{}{
		"clients": clients,
	}, http.StatusOK)
}

func main() {
	conn, err := datastore.CreateConnection(os.Getenv("REGION"))
	if err != nil {
		log.Panic(err)
	}

	ddb := datastore.NewDynamoDB(conn, os.Getenv("DB_TABLE"))

	repository := model.NewClientRepository(ddb)

	handler := &Handler{repository}

	router := helpers.Router(handler)

	lambda.Start(router)
}
