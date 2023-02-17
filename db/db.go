package db

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoOnce sync.Once
var mongoClient *mongo.Client
var mongoClientError error

const (
	uri                    = "mongodb://mongo:27017"
	Database               = "bikes-db"
	BikesCollection string = "bikes"
)

func GetMongoClient() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		mongoClient, mongoClientError = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	})
	return mongoClient, mongoClientError
}
