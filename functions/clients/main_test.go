package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tMatSuZ/serverless-go-sample/functions/clients/model"
	httpdelivery "github.com/tMatSuZ/serverless-go-sample/pkg/http"
)

type MockClientRepository struct{}

func (r *MockClientRepository) Get(id string) (*model.Client, error) {
	return &model.Client{
		ID:          "123",
		Name:        "tatsu",
		Rate:        10,
		Description: "des",
	}, nil
}

func (r *MockClientRepository) List() (*[]model.Client, error) {
	return &[]model.Client{
		model.Client{
			ID:          "123",
			Name:        "tatsu",
			Rate:        10,
			Description: "des",
		},
	}, nil
}

func (r *MockClientRepository) Store(client *model.Client) error {
	return nil
}

func TestCanFetchClient(t *testing.T) {
	request := httpdelivery.Req{
		HTTPMethod:     "GET",
		PathParameters: map[string]string{"id": "123"},
	}
	h := &Handler{&MockClientRepository{}}
	router := httpdelivery.Router(h)
	response, err := router(request)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func TestCanCreateClient(t *testing.T) {
	request := httpdelivery.Req{
		HTTPMethod: "POST",
		Body:       `{"name": "test", "description": "dest", "rate": 40}`,
	}
	h := &Handler{&MockClientRepository{}}
	router := httpdelivery.Router(h)
	responce, err := router(request)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, responce.StatusCode)
}

func TestCanListClients(t *testing.T) {
	request := httpdelivery.Req{
		HTTPMethod: "GET",
	}
	h := &Handler{&MockClientRepository{}}
	router := httpdelivery.Router(h)
	responce, err := router(request)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, responce.StatusCode)
}

func TestHandleInvalidJSON(t *testing.T) {
	request := httpdelivery.Req{
		HTTPMethod:"POST",
		Body: "",
	}
	h := &Handler{&MockClientRepository{}}
	router := httpdelivery.Router(h)
	responce, err := router(request)
	assert.Error(t, err)
	assert.Equal(t, http.StatusBadRequest, responce.StatusCode)
}