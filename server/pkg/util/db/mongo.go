package util

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoConfiguration struct {
	user   string
	pass   string
	host   string
	port   string
	schema string
}

func NewMongoConfiguration(user string, pass string, host string, port string, table string) *MongoConfiguration {
	return &MongoConfiguration{
		user:   user,
		pass:   pass,
		host:   host,
		port:   port,
		schema: table,
	}
}

// ConnectMongo to a database handle from a mongo configuration.
func ConnectMongo(configuration *MongoConfiguration) (*mongo.Database, error) {
	mongoUri := fmt.Sprintf("mongodb://%s:%s@%s:%s/", configuration.user, configuration.pass, configuration.host, configuration.port)

	clientOptions := options.Client().ApplyURI(mongoUri)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	db := client.Database(configuration.schema)

	return db, nil
}
