package mongo

import (
	"time"
)

type Person struct {
	ID        string    `json:"id" bson:"_id"`
	Username  string    `json:"username" bson:"username"`
	Password  string    `json:"password" bson:"password"`
	Name      string    `json:"name" bson:"name"`
	Surname   string    `json:"surname" bson:"surname"`
	IDNumber  string    `json:"id_number" bson:"id_number"`
	Email     string    `json:"email" bson:"email"`
	Phone     string    `json:"phone" bson:"phone"`
	BirthDate string    `json:"birth_date" bson:"birth_date"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`
}
