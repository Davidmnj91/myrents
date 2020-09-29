package db

import (
	"time"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbURI string = "mongodb://myrent:myrent123@127.0.0.1:27017/"
const dbName string = "myrent"

/*Db points to the active mongo db */
var Db *mongo.Database
var client *mongo.Client

/*Connect try to stablish a connection to the mongo db, if successfully returns the db connection*/
func Connect() (*mongo.Database, error) {
	clientOptions:= options.Client().ApplyURI(dbURI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	defer Disconnect()

	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	Db = client.Database(dbName)

	return Db, nil
}

/*Disconnect an active mongo client*/
func Disconnect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Disconnect(ctx); err != nil {
		return err
	}

	return nil
}