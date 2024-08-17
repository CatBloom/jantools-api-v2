package types

import "time"

type ReqUser struct {
	Limit int    `query:"limit" validate:"gte=0,lt=20"`
	Order string `query:"order"`
}

type ReqCreateUser struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

type User struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
