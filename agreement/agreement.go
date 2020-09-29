package agreement

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/Davidmnj91/myrentals/file"
)

/*Agreement struct representation of a rental agreement between a landlord and a client for a realstate*/
type Agreement struct {
	ID  primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	RealState primitive.ObjectID `bson:"realstateid,omitempty" json:"realstate"`
	Landlord primitive.ObjectID `bson:"userid,omitempty" json:"landlord"`
	Tenant primitive.ObjectID `bson:"userid,omitempty" json:"tenant"`
	RentalCost float32 `bson:"rental_cost,omitempty" json:"rental_cost"`
	RentalDuration int `bson:"rental_duration,omitempty" json:"rental_duration"`
	StartDate time.Time `bson:"start_date,omitempty" json:"start_date"`
	Agreement file.File `bson:"agreement,omitempty" json:"agreement"`
	CreatedAt time.Time `bson:"created_at,omitempty" json:"created_at"`
    UpdatedAt time.Time `bson:"updated_at,omitempty" json:"updated_at"`
}