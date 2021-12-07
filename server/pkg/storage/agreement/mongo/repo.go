package agreement

import (
	"context"
	"github.com/Davidmnj91/myrents/pkg/agreement/domain"
	"github.com/Davidmnj91/myrents/pkg/types"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepository struct {
	db *mongo.Collection
}

func NewRepository(db *mongo.Database) domain.Repository {
	return &mongoRepository{db.Collection("agreement")}
}

func (r *mongoRepository) Add(ctx context.Context, agreement *domain.Agreement) error {
	toInsert := ToRepository(agreement)
	toInsert.ID = uuid.New().String()

	_, err := r.db.InsertOne(ctx, toInsert)

	return err
}

func (r *mongoRepository) FindById(ctx context.Context, uuid types.UUID) (*domain.Agreement, error) {
	query := bson.M{"_id": uuid.String()}
	found := r.db.FindOne(ctx, query)
	if err := found.Err(); err != nil {
		return nil, err
	}

	var entity AgreementStorage
	err := found.Decode(&entity)
	if err != nil {
		return nil, err
	}

	return ToDomain(entity), nil
}

func (r *mongoRepository) FindByLandReference(ctx context.Context, landReference string) (*domain.Agreement, error) {
	query := bson.M{"real_state": landReference}
	found := r.db.FindOne(ctx, query)
	if err := found.Err(); err != nil {
		return nil, err
	}

	var entity AgreementStorage
	err := found.Decode(&entity)
	if err != nil {
		return nil, err
	}

	return ToDomain(entity), nil
}

func (r *mongoRepository) FindByLandlordOrTenant(ctx context.Context, userUUID types.UUID) ([]domain.Agreement, error) {
	query := bson.M{
		"$or": []bson.M{
			{"landlord": userUUID.String()},
			{"tenant": userUUID.String()},
		},
	}
	found, err := r.db.Find(ctx, query)
	if err != nil {
		return nil, err
	}

	var agreements []domain.Agreement

	for found.Next(context.TODO()) {
		var elem AgreementStorage
		if err := found.Decode(&elem); err != nil {
			return nil, err
		}

		agreements = append(agreements, *ToDomain(elem))
	}

	if err := found.Err(); err != nil {
		return nil, err
	}

	if err := found.Close(context.TODO()); err != nil {
		return nil, err
	}

	return agreements, nil
}
