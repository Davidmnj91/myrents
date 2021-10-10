package mongo

import (
	"context"
	"errors"
	"github.com/Davidmnj91/myrents/pkg/domain/types"
	"github.com/Davidmnj91/myrents/pkg/domain/user"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepository struct {
	db *mongo.Collection
}

func NewRepository(db *mongo.Collection) user.Repository {
	return &mongoRepository{db}
}

func (r *mongoRepository) Add(ctx context.Context, user *user.User) error {
	toInsert := ToRepository(user)
	toInsert.ID = uuid.New().String()

	_, err := r.db.InsertOne(ctx, toInsert)
	if err != nil {
		return err
	}

	return nil
}

func (r *mongoRepository) FindById(ctx context.Context, uuid domain.UUID) (*user.User, error) {
	query := bson.M{"_id": uuid.String()}
	found := r.db.FindOne(ctx, query)
	if err := found.Err(); err != nil {
		return nil, err
	}

	var person Person
	err := found.Decode(&person)
	if err != nil {
		return nil, err
	}

	return ToDomain(person), nil
}

func (r *mongoRepository) FindByUsername(ctx context.Context, username string) (*user.User, error) {
	query := bson.M{"username": username}
	found := r.db.FindOne(ctx, query)
	if err := found.Err(); err != nil {
		return nil, err
	}

	var person Person
	err := found.Decode(&person)
	if err != nil {
		return nil, err
	}

	return ToDomain(person), nil
}

func (r *mongoRepository) Exists(ctx context.Context, user *user.User) (bool, error) {
	query := bson.M{
		"$or": []bson.M{
			{"id_number": user.IDNumber},
			{"username": user.Username},
			{"email": user.Email},
			{"phone": user.Phone},
		},
	}

	matches, err := r.db.CountDocuments(ctx, query)

	if err != nil {
		return false, err
	}

	return matches != 0, nil
}

func (r mongoRepository) Update(ctx context.Context, update *user.User) (*user.User, error) {
	updated, err := r.db.UpdateByID(ctx, update.UserUUID.String(), ToRepository(update))
	if err != nil {
		return nil, err
	}

	if updated.ModifiedCount == 0 {
		return nil, errors.New("could not find user matching criteria")
	}

	return r.FindById(ctx, update.UserUUID)
}
