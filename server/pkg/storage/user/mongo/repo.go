package mongo

import (
	"context"
	"github.com/Davidmnj91/myrents/pkg/domain/types"
	"github.com/Davidmnj91/myrents/pkg/domain/user"
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
	return nil
}

func (r *mongoRepository) FindById(ctx context.Context, uuid domain.UUID) (*user.User, error) {
	return &user.User{}, nil
}

func (r *mongoRepository) FindByUsername(ctx context.Context, username string) (*user.User, error) {
	return &user.User{}, nil
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
	return &user.User{}, nil
}

func (r *mongoRepository) Delete(ctx context.Context, uuid domain.UUID) error {
	return nil
}
