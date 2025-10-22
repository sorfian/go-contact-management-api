package web

type UserUpdateRequest struct {
	Name     string `validate:"min=3,max=100" json:"name"`
	Password string `validate:"min=3,max=100" json:"password"`
}
