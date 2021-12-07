package agreement_list

// swagger:model ListAgreement
type ListAgreement struct {
	// The unique official reference of the real state
	// required: true
	// example: 1234-5678-9012-3456
	RealState string `json:"realState"`
	// The IDNumber code of landlord part that rents the real state
	// required: true
	// example: 01234567-N
	Landlord string `json:"landlord"`
	// The IDNumber code of tenant part that rents the real state
	// required: true
	// example: 89012345-M
	Tenant string `json:"tenant"`
	// The monthly price of the rent
	// required: true
	// example: 800
	MonthlyPrice float32 `json:"monthlyPrice"`
	// The effective start date for the rental
	// required: true
	// example: 2000-12-12
	StartDate string `json:"startDate"`
	// The effective end date for the rental
	// required: true
	// example: 2005-12-12
	EndDate string `json:"endDate"`
}
