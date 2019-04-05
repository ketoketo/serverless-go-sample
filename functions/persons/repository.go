package main

import (
	uuid "github.com/satori/go.uuid"
	"github.com/tMatSuZ/serverless-go-sample/model"
	"github.com/tMatSuZ/serverless-go-sample/pkg/datastore"
)

type DatastoreRepository struct {
	datastore datastore.Datastore
}

func NewDatastoreRepository(ds datastore.Datastore) *DatastoreRepository {
	return &DatastoreRepository{datastore: ds}
}

func (r *DatastoreRepository) Get(id string) (*model.Person, error) {
	var person *model.Person
	if err := r.datastore.Get(id, &person); err != nil {
		return nil, err
	}
	return person, nil
}

func (r *DatastoreRepository) Store(person *model.Person) error {
	id, _ := uuid.NewV4()
	person.ID = id.String()
	return r.datastore.Store(person)
}

func (r *DatastoreRepository) List() (*[]model.Person, error) {
	var persons *[]model.Person
	if err := r.datastore.List(&persons); err != nil {
		return nil, err
	}
	return persons, nil
}
