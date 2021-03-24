package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(mongoDSN, dbName string) (*mongo.Database, error) {
	// uri := "mongodb://" + config.GetConfig("MONGODSN") + "/"
	clientOptions := options.Client().ApplyURI(mongoDSN)
	// clientOptions := options.Client().ApplyURI("mongodb://" + config.GetConfig("MONGODSN") + "/")
	// fmt.Println(uri)
	// clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
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
