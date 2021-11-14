package user_profile

// swagger:model Profile
type Profile struct {
	// The name of the user in the app
	// example: myUserName
	Username string `json:"username"`
	// The Name of the user
	// example: David
	Name string `json:"name"`
	// The Surname of the user
	// example: Sanchez
	Surname string `json:"surname"`
	// The IDNumber of the user
	// example: 01234567-N
	IDNumber string `json:"id_number"`
	// The Email of the user
	// example: myemail@me.com
	Email string `json:"email"`
	// The Phone of the user
	// example: 666000222
	Phone string `json:"phone"`
	// The BirthDate of the user
	// example: 2000-12-12
	BirthDate string `json:"birth_date"`
}
