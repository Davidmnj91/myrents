package login

// swagger:model Login
type Login struct {
	// The username to login
	// required: true
	Username string `json:"username" validate:"required"`
	// The password to login
	// required: true
	Password string `json:"password" validate:"required"`
}
