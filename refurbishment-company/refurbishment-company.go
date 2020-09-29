package refurbishmentCompany

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Contractor struct representation of a worker of a Refurbishment company*/
type Contractor struct {
	Name string `bson:"name,omitempty" json:"name"`
	Surname string `bson:"surname,omitempty" json:"surname"`
	Email string `bson:"email,omitempty" json:"email"`
	Phone string `bson:"phone,omitempty" json:"phone"`
}

/*RefurbishmentCompany struct representation of a commonly used company to repair minor issues in real state properties*/
type RefurbishmentCompany struct {
	ID  primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name string `bson:"name,omitempty" json:"name"`
	IDNumber string `bson:"id_number,omitempty" json:"id_number"`
	Email string `bson:"email,omitempty" json:"email"`
	Phone string `bson:"phone,omitempty" json:"phone"`
	Specialties []string `bson:"specialties,omitempty" json:"specialties"`
	Contractors []Contractor `bson:"contractors,omitempty" json:"contractor"`
	CreatedAt time.Time `bson:"created_at,omitempty" json:"created_at"`
    UpdatedAt time.Time `bson:"updated_at,omitempty" json:"updated_at"`
}