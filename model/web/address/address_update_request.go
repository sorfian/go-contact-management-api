package address

type AddressUpdateRequest struct {
	Street     string `json:"street" validate:"omitempty,min=1,max=200"`
	City       string `json:"city" validate:"omitempty,min=1,max=100"`
	Province   string `json:"province" validate:"omitempty,min=1,max=100"`
	Country    string `json:"country" validate:"omitempty,min=1,max=100"`
	PostalCode string `json:"postal_code" validate:"omitempty,min=1,max=10"`
}
