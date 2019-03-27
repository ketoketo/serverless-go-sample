package model

import (
	uuid "github.com/satori/go.uuid"
	"github.com/tMatSuZ/serverless-go-sample/pkg/datastore"
)

type ClientRepository struct {
	datastore datastore.Datastore
}

func NewClientRepository(ds datastore.Datastore) *ClientRepository {
	return &ClientRepository{datastore: ds}
}

func (r *ClientRepository) Get(id string) (*Client, error) {
	var client *Client
	if err := r.datastore.Get(id, &client); err != nil {
		return nil, err
	}
	return client, nil
}

func (r *ClientRepository) Store(client *Client) error {
	id, _ := uuid.NewV4()
	client.ID = id.String()
	return r.datastore.Store(client)
}

func (r *ClientRepository) List() (*[]Client, error) {
	var clients *[]Client
	if err := r.datastore.List(&clients); err != nil {
		return nil, err
	}
	return clients, nil
}
