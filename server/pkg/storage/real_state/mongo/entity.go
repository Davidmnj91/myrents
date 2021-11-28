package real_state

import (
	"time"
)

/*RealStateStorage struct representation of a good of a landlord*/
type RealStateStorage struct {
	ID            string    `json:"id" bson:"_id"`
	LandReference string    `json:"land_reference" bson:"land_reference"`
	Street        string    `json:"street" bson:"street"`
	ZipCode       string    `json:"zip_code" bson:"zip_code"`
	Province      string    `json:"province" bson:"province"`
	Country       string    `json:"country" bson:"country"`
	Gateway       string    `json:"gateway" bson:"gateway"`
	Door          string    `json:"door" bson:"door"`
	SqMeters      string    `json:"sq_meters" bson:"sq_meters"`
	Landlord      string    `json:"user_id" bson:"user_id"`
	CreatedAt     time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" bson:"updated_at"`
	DeletedAt     time.Time `json:"deleted_at" bson:"deleted_at"`
}
