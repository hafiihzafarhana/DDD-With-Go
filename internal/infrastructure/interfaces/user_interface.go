package interfaces

import "time"

type NewRegisterResponse struct {
	Id        uint      `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
