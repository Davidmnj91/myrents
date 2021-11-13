package real_state

import (
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
	"time"
)

/*RealStateStorage struct representation of a good of a landlord*/
type RealStateStorage struct {
	ID        string    `json:"id" bson:"_id"`
	Street    string    `json:"street" bson:"street"`
	ZipCode   string    `json:"zip_code" bson:"zip_code"`
	Province  string    `json:"province" bson:"province"`
	Country   string    `json:"country" bson:"country"`
	Gateway   string    `json:"gateway" bson:"gateway"`
	Door      string    `json:"door" bson:"door"`
	SqMeters  float32   `json:"sq_meters" bson:"sq_meters"`
	Landlord  uuid.UUID `json:"userid" bson:"userid"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`
}
