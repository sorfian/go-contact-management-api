package web

type TokenResponse struct {
	Token    string `json:"token"`
	TokenExp int64  `json:"token_exp"`
}
