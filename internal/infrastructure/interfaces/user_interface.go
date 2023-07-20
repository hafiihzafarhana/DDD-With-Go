package interfaces

import "time"

type NewRegisterResponse struct {
	Id        uint      `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type NewLoginResponse struct {
	AccToken     string `json:"acc_token"`
	RefreshToken string `json:"refresh_token"`
}

type NewJWTAuthenticateTransformer struct {
	Id    uint   `json:"id"`
	Email string `json:"email"`
}
