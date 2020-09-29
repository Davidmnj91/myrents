package useraccount

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*UserAccount struct representation of an user in the system*/
type UserAccount struct {
	ID  primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name string `bson:"name,omitempty" json:"name"`
	Surname string `bson:"surname,omitempty" json:"surname"`
	IDNumber string `bson:"id_number,omitempty" json:"id_number"`
	Email string `bson:"email,omitempty" json:"email"`
	Phone string `bson:"phone,omitempty" json:"phone"`
	BirthDate time.Time `bson:"birth_date,omitempty" json:"birth_date"`
	CreatedAt time.Time `bson:"created_at,omitempty" json:"created_at"`
    UpdatedAt time.Time `bson:"updated_at,omitempty" json:"updated_at"`
}