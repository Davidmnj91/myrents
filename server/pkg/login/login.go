package login

// swagger:parameters login
type Login struct {
	// The username to login
	//
	// in: body
	// required: true
	Username string `json:"username" validate:"required"`
	// The password to login
	//
	// in: body
	// required: true
	Password string `json:"password" validate:"required"`
}
