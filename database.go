package main

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDB(db_uri string, db_name string) *mongo.Database {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db_uri))

	if err != nil {
		panic(err.Error())
	}

	return client.Database(db_name)
}

// intialize databases
var (
// default_db = InitDB("mongodb://localhost:27017", "learn-clean-arc")
)

// declare collection
var (
// default_coll = &CollectionRepository{default_db.Collection("store")}
)
