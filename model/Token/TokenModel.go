package Token

type Token struct {
	Token               string `json:"token"`
	RefreshToken        string `json:"refresh_token"`
	TokenExpired        int64  `json:"token_expired"`
	RefreshTokenExpired int64  `json:"refresh_token_expired"`
}
