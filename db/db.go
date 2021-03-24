package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Connect -
func Connect(mongoDSN, dbName string) (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(mongoDSN)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	db := client.Database(dbName)
	return db, nil
}
