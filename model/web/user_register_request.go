package web

type UserRegisterRequest struct {
	Username string `validate:"required,min=3,max=100" json:"username"`
	Name     string `validate:"required,min=3,max=100" json:"name"`
	Password string `validate:"required,min=3,max=100" json:"password"`
}
