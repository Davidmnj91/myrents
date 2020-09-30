package useraccount

import (
	"context"	
	"go.mongodb.org/mongo-driver/mongo"
)

const userAccountCollection = "useraccount"

func newUserAccountRepo(db *mongo.Database) Repo {
	return &MongoRepo{db}
}

type Repo interface {
	Create(context.Context, *UserAccount)(interface{}, error)
}

/*MongoRepo struct to bypass repositories*/
type MongoRepo struct {
	db *mongo.Database
}

/*Create a user in the database*/
func (ur *MongoRepo) Create(ctx context.Context, u *UserAccount) (interface{}, error) {
	userAccountCollection := ur.db.Collection(userAccountCollection)
	res, err := userAccountCollection.InsertOne(ctx, u)

	return res.InsertedID, err
}