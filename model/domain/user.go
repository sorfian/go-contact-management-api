package domain

type User struct {
	Username string
	Password string
	Name     string
	Token    string
	TokenExp int64
}
