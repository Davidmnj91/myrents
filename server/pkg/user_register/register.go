package user_register

// swagger:model Register
type Register struct {
	// The name of the user in the app
	// required: true
	// example: myUserName
	Username string `json:"username" validate:"required"`
	// The Password of the user
	// required: true
	// example: mySecretPassword
	Password string `json:"password" validate:"required"`
	// The Name of the user
	// required: true
	// example: David
	Name string `json:"name" validate:"required"`
	// The Surname of the user
	// required: true
	// example: Sanchez
	Surname string `json:"surname" validate:"required"`
	// The IDNumber of the user
	// required: true
	// example: 01234567-N
	IDNumber string `json:"id_number" validate:"required"`
	// The Email of the user
	// required: true
	// example: myemail@me.com
	Email string `json:"email" validate:"required"`
	// The Phone of the user
	// required: true
	// example: 666000222
	Phone string `json:"phone" validate:"required"`
	// The BirthDate of the user
	// required: true
	// example: 2000-12-12
	BirthDate string `json:"birth_date" validate:"required,birthDate"`
}
