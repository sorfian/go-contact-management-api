package web

type ContactCreateRequest struct {
	FirstName string `json:"first_name" validate:"required,min=1,max=100"`
	LastName  string `json:"last_name" validate:"required,min=1,max=100"`
	Email     string `json:"email" validate:"required,email,max=100"`
	Phone     string `json:"phone" validate:"required,min=1,max=20"`
}
