package real_state

import (
	"context"
	"errors"
	"github.com/Davidmnj91/myrents/pkg/real_state/domain"
	"github.com/Davidmnj91/myrents/pkg/types"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepository struct {
	db *mongo.Collection
}

func NewRepository(db *mongo.Database) domain.Repository {
	return &mongoRepository{db.Collection("real_state")}
}

func (r *mongoRepository) Add(ctx context.Context, realState *domain.RealState) error {
	toInsert := ToRepository(realState)
	toInsert.ID = uuid.New().String()

	_, err := r.db.InsertOne(ctx, toInsert)

	return err
}

func (r *mongoRepository) FindById(ctx context.Context, uuid types.UUID) (*domain.RealState, error) {
	query := bson.M{"_id": uuid.String()}
	found := r.db.FindOne(ctx, query)
	if err := found.Err(); err != nil {
		return nil, err
	}

	var entity RealStateStorage
	err := found.Decode(&entity)
	if err != nil {
		return nil, err
	}

	return ToDomain(entity), nil
}

func (r *mongoRepository) FindByLandReference(ctx context.Context, landReference string) (*domain.RealState, error) {
	query := bson.M{"land_reference": landReference}
	found := r.db.FindOne(ctx, query)
	if err := found.Err(); err != nil {
		return nil, err
	}

	var entity RealStateStorage
	err := found.Decode(&entity)
	if err != nil {
		return nil, err
	}

	return ToDomain(entity), nil
}

func (r *mongoRepository) FindByUserId(ctx context.Context, userUUID types.UUID) ([]domain.RealState, error) {
	query := bson.M{"userid": userUUID}
	found, err := r.db.Find(ctx, query)
	if err != nil {
		return nil, err
	}

	var results []RealStateStorage
	err = found.Decode(&results)
	if err != nil {
		return nil, err
	}

	var realStates []domain.RealState

	for _, result := range results {
		realStates = append(realStates, *ToDomain(result))
	}

	return realStates, nil
}

func (r *mongoRepository) Exists(ctx context.Context, realState *domain.RealState) (bool, error) {
	query := bson.M{
		"$or": []bson.M{
			{"land_reference": realState.LandReference},
			{"$and": []bson.M{
				{"street": realState.Street},
				{"zip_code": realState.ZipCode},
				{"province": realState.Province},
				{"country": realState.Country},
				{"gateway": realState.Gateway},
				{"door": realState.Door}},
			},
		},
	}

	matches, err := r.db.CountDocuments(ctx, query)

	if err != nil {
		return false, err
	}

	return matches != 0, nil
}

func (r *mongoRepository) Update(ctx context.Context, update *domain.RealState) (*domain.RealState, error) {
	toUpdate := ToRepository(update)

	query := bson.M{"$set": bson.M{
		"street":    toUpdate.Street,
		"zip_code":  toUpdate.ZipCode,
		"province":  toUpdate.Province,
		"country":   toUpdate.Country,
		"gateway":   toUpdate.Gateway,
		"door":      toUpdate.Door,
		"sq_meters": toUpdate.SqMeters,
	}}

	updated, err := r.db.UpdateByID(ctx, update.RealStateUUID.String(), query)
	if err != nil {
		return nil, err
	}

	if updated.MatchedCount == 0 {
		return nil, errors.New("could not find real state matching criteria")
	}

	return r.FindById(ctx, update.RealStateUUID)
}
