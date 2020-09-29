package realstate

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/Davidmnj91/myrentals/file"
)

/*Furniture struct representation of an item in the room*/
type Furniture struct {
	Name string `bson:"name,omitempty" json:"name"`
	Brand string `bson:"brand,omitempty" json:"brand"`
	Model string `bson:"model,omitempty" json:"model"`
	Value float32 `bson:"value,omitempty" json:"value"`
	PurchaseDate time.Time `bson:"purchase_date,omitempty" json:"purchase_date"`
}

/*Room struct representation of a place in the house*/
type Room struct {
	Name string `bson:"name,omitempty" json:"name"`
	Description string `bson:"description,omitempty" json:"description"`
	Pictures []file.File `bson:"pictures,omitempty" json:"pictures"`
	Furniture []Furniture `bson:"furniture,omitempty" json:"furniture"`
}

/*RealState struct representation of a good of a landlord*/
type RealState struct {
	ID  primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Street string `bson:"street,omitempty" json:"street"`
	ZipCode string `bson:"zip_code,omitempty" json:"zip_code"`
	Province string `bson:"province,omitempty" json:"province"`
	Country string `bson:"country,omitempty" json:"country"`
	Gateway int `bson:"gateway,omitempty" json:"gateway"`
	Door string `bson:"door,omitempty" json:"door"`
	Area float32 `bson:"area,omitempty" json:"area"`
	Landlord primitive.ObjectID `bson:"userid,omitempty" json:"landlord"`
	Rooms []Room `bson:"rooms,omitempty" json:"rooms"`
	CreatedAt time.Time `bson:"created_at,omitempty" json:"created_at"`
    UpdatedAt time.Time `bson:"updated_at,omitempty" json:"updated_at"`
}