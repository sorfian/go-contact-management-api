package user

type UserLoginRequest struct {
	Username string `validate:"required,min=3,max=100" json:"username"`
	Password string `validate:"required,min=3,max=100" json:"password"`
}
