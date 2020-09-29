package bill

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/Davidmnj91/myrentals/file"
)

/*Bill struct representation of any debt paid or not of the tenant for a realstate rented*/
type Bill struct {
	ID  primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Agreement primitive.ObjectID `bson:"agreementid,omitempty" json:"agreement"`
	Amount float32 `bson:"amount,omitempty" json:"amount"`
	Paid bool `bson:"paid,omitempty" json:"paid"`
	Type string `bson:"type,omitproperty" json:"type"`
	Date time.Time `bson:"date,omitempty" json:"date"`
	Receipt file.File `bson:"receipt,omitempty" json:"receipt"`
	CreatedAt time.Time `bson:"created_at,omitempty" json:"created_at"`
    UpdatedAt time.Time `bson:"updated_at,omitempty" json:"updated_at"`
}