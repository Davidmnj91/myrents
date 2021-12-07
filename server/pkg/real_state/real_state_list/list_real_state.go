package real_state_list

// swagger:model ListRealState
type ListRealState struct {
	// The unique official reference of the real state
	// required: true
	// example: 1234-5678-9012-3456
	LandReference string `json:"landReference"`
	// The street of the real state address
	// required: true
	// example: C/ False
	Street string `json:"street"`
	// The zip code of the real state address
	// required: true
	// example: 057890
	ZipCode string `json:"zipCode"`
	// The province of the real state address
	// required: true
	// example: Madrid
	Province string `json:"province"`
	// The country of the real state address
	// required: true
	// example: Spain
	Country string `json:"country"`
	// The gateway of the real state address
	// required: true
	// example: 19Bis
	Gateway string `json:"gateway"`
	// The Door of the real state address
	// example: 2Iz
	Door string `json:"door"`
	// The Size of the real state
	// example: 120
	SqMeters string `json:"sqMeters"`
}
