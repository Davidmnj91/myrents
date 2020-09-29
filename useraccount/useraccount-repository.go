package useraccount

import (
	"context"	
	"go.mongodb.org/mongo-driver/mongo"
)

const userAccountCollection = "useraccount"

func newUserAccountRepo(db *mongo.Database) Repo {
	return Repo{db}
}

/*Repo struct to bypass repositories*/
type Repo struct {
	db *mongo.Database
}

/*Create a user in the database*/
func (ur *Repo) Create(ctx context.Context, u *UserAccount) (*mongo.InsertOneResult, error) {
	userAccountCollection := ur.db.Collection(userAccountCollection)
	res, err := userAccountCollection.InsertOne(ctx, u)

	return res, err
}