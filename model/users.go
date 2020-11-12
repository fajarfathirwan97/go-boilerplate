package model

import (
	"time"
)

type User struct {
	Uuid            string    `json:"uuid" db:"uuid" validate:"required"`
	Email           string    `json:"email" db:"email" validate:"email"`
	Password        string    `json:"password" db:"password" validate:"required,eqfield=ConfirmPassword"`
	ConfirmPassword string    `json:"confirm_password" validate:"required,eqfield=Password"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}
