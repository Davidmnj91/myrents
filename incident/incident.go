package bill

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/Davidmnj91/myrentals/file"
)

/*Incident struct representation of any debt paid or not of the tenant for a realstate rented*/
type Incident struct {
	ID  primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Agreement primitive.ObjectID `bson:"agreementid,omitempty" json:"agreement"`
	RefurbishmentCompany primitive.ObjectID `bson:"refurbishmentcompanyid,omitempty" json:"refurbishment_company"`
	Title string `bson:"title,omitempty" json:"title"`
	Description string `bson:"description,omitempty" json:"description"`
	StartDate time.Time `bson:"start_date,omitempty" json:"start_date"`
	EndDate time.Time `bson:"end_date,omitempty" json:"end_date"`
	Result string `bson:"result,omitempty" json:"result"`
	Price float32 `bson:"price,omitempty" json:"price"`
	Receipt file.File `bson:"receipt,omitempty" json:"receipt"`
	CreatedAt time.Time `bson:"created_at,omitempty" json:"created_at"`
    UpdatedAt time.Time `bson:"updated_at,omitempty" json:"updated_at"`
}