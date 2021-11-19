package real_state_register

// swagger:model RegisterRealState
type Register struct {
	// The owner of the real state uuid
	// required: true
	// example: 1234-5678-9012-3456
	LandReference string `json:"land_reference" validate:"required"`
	// The street of the real state address
	// required: true
	// example: C/ False
	Street string `json:"street" validate:"required"`
	// The zip code of the real state address
	// required: true
	// example: 057890
	ZipCode string `json:"zip_code" validate:"required"`
	// The province of the real state address
	// required: true
	// example: Madrid
	Province string `json:"province" validate:"required"`
	// The country of the real state address
	// required: true
	// example: Spain
	Country string `json:"country" validate:"required"`
	// The gateway of the real state address
	// required: true
	// example: 19Bis
	Gateway string `json:"gateway" validate:"required"`
	// The Door of the real state address
	// example: 2Iz
	Door string `json:"door"`
	// The Size of the real state
	// example: 120
	SqMeters string `json:"sq_meters"`
	// The owner's uuid of the real state
	// required: true
	// example: 1234-5678-9012-3456
	Landlord string `json:"landlord" validate:"required"`
}
