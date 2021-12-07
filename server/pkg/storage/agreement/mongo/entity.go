package agreement

import (
	"time"
)

/*AgreementStorage struct representation of a rental agreement between a landlord and a client for a real-state*/
type AgreementStorage struct {
	ID         string    `json:"id" bson:"_id"`
	RealState  string    `json:"real_state" bson:"real_state"`
	Landlord   string    `json:"landlord" bson:"landlord"`
	Tenant     string    `json:"tenant" bson:"tenant"`
	RentalCost float32   `json:"rental_cost" bson:"rental_cost"`
	StartDate  string    `json:"start_date" bson:"start_date"`
	EndDate    string    `json:"end_date" bson:"end_date"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" bson:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at" bson:"deleted_at"`
}
