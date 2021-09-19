package mongo

import (
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
	"time"
)

/*Agreement struct representation of a rental agreement between a landlord and a client for a real-state*/
type Agreement struct {
	ID         uuid.UUID `json:"id" bson:"_id"`
	RealState  uuid.UUID `json:"realstate" bson:"realstateid"`
	Landlord   uuid.UUID `json:"landlord" bson:"userid"`
	Tenant     uuid.UUID `json:"tenant" bson:"userid"`
	RentalCost float32   `json:"rental_cost" bson:"rental_cost"`
	StartDate  time.Time `json:"start_date" bson:"start_date"`
	EndDate    time.Time `json:"end_date" bson:"end_date"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" bson:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at" bson:"deleted_at"`
}
