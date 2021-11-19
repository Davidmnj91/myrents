package user

import (
	"context"
	"errors"
	"github.com/Davidmnj91/myrents/pkg/types"
	"github.com/Davidmnj91/myrents/pkg/user/domain"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepository struct {
	db *mongo.Collection
}

func NewRepository(db *mongo.Database) domain.Repository {
	return &mongoRepository{db.Collection("users")}
}

func (r *mongoRepository) Add(ctx context.Context, user *domain.User) error {
	toInsert := ToRepository(user)
	toInsert.ID = uuid.New().String()

	_, err := r.db.InsertOne(ctx, toInsert)

	return err
}

func (r *mongoRepository) FindById(ctx context.Context, uuid types.UUID) (*domain.User, error) {
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

func (r *mongoRepository) FindByUsername(ctx context.Context, username string) (*domain.User, error) {
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

func (r *mongoRepository) Exists(ctx context.Context, user *domain.User) (bool, error) {
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

func (r mongoRepository) Update(ctx context.Context, update *domain.User) (*domain.User, error) {
	toUpdate := ToRepository(update)

	query := bson.M{"$set": bson.M{
		"email":      toUpdate.Email,
		"phone":      toUpdate.Phone,
		"birth_date": toUpdate.BirthDate,
		"created_at": toUpdate.CreatedAt,
		"updated_at": toUpdate.UpdatedAt,
		"deleted_at": toUpdate.DeletedAt,
	}}

	updated, err := r.db.UpdateByID(ctx, update.UserUUID.String(), query)
	if err != nil {
		return nil, err
	}

	if updated.MatchedCount == 0 {
		return nil, errors.New("could not find user matching criteria")
	}

	return r.FindById(ctx, update.UserUUID)
}
