package web

type AddressResponse struct {
	ID         int64  `json:"id"`
	Street     string `json:"street"`
	City       string `json:"city"`
	Province   string `json:"province"`
	Country    string `json:"country"`
	PostalCode string `json:"postal_code"`
}
