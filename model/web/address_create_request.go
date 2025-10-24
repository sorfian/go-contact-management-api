package web

type AddressCreateRequest struct {
	Street     string `json:"street" validate:"required,min=1,max=200"`
	City       string `json:"city" validate:"required,min=1,max=100"`
	Province   string `json:"province" validate:"required,min=1,max=100"`
	Country    string `json:"country" validate:"required,min=1,max=100"`
	PostalCode string `json:"postal_code" validate:"required,min=1,max=10"`
}
