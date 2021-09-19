package user_register

import "time"

// swagger:parameters register
type Register struct {
	Username  string    `json:"user_name" validate:"required"`
	Password  string    `json:"password" validate:"required"`
	Name      string    `json:"name" validate:"required"`
	Surname   string    `json:"surname" validate:"required"`
	IDNumber  string    `json:"id_number" validate:"required"`
	Email     string    `json:"email" validate:"required"`
	Phone     string    `json:"phone" validate:"required"`
	BirthDate time.Time `json:"birth_date" validate:"required,datetime"`
}
