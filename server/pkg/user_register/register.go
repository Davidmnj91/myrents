package user_register

// swagger:model Register
type Register struct {
	// The name of the user in the app
	// required: true
	Username string `json:"username" validate:"required"`
	// The Password of the user
	// required: true
	Password string `json:"password" validate:"required"`
	// The Name of the user
	// required: true
	Name string `json:"name" validate:"required"`
	// The Surname of the user
	// required: true
	Surname string `json:"surname" validate:"required"`
	// The IDNumber of the user
	// required: true
	IDNumber string `json:"id_number" validate:"required"`
	// The Email of the user
	// required: true
	Email string `json:"email" validate:"required"`
	// The Phone of the user
	// required: true
	Phone string `json:"phone" validate:"required"`
	// The BirthDate of the user
	// required: true
	BirthDate string `json:"birth_date" validate:"required,birthDate"`
}
