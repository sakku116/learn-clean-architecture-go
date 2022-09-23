package main

import "go.mongodb.org/mongo-driver/mongo"

type CollectionRepository struct {
	Coll *mongo.Collection
}

func (coll *CollectionRepository) FindAll() *[]Product {
	return &[]Product{}
}

func (coll *CollectionRepository) FindByID() *Product {
	return &Product{}
}

func (coll *CollectionRepository) Save() error {
	return nil
}

func (coll *CollectionRepository) Delete() error {
	return nil
}
