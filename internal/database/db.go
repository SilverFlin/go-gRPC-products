package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func Connect() error {
	clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017")
	ctx := context.TODO()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}
	Client = client
	return nil
}

func Close() error {
	if Client != nil {
		err := Client.Disconnect(context.TODO())
		if err != nil {
			return err
		}
	}
	return nil
}
