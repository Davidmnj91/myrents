package login

// swagger:model Login
type Login struct {
	// The username to login
	// required: true
	// example: myUserName
	Username string `json:"username" validate:"required"`
	// The password to login
	// required: true
	// example: mySecretPassword
	Password string `json:"password" validate:"required"`
}
