package model

import (
	"go-docker/helper"
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

func (u User) CreateUser(m Merchant) error {
	db, err := helper.GetDBClient()
	if err != nil {
		return err
	}
	defer db.Close()
	tx := db.MustBegin()
	if _, err = tx.NamedExec(`
		INSERT INTO users (
			uuid,
			email,
			password,
			created_at,
			updated_at
		) VALUES (
			:uuid,
			:email,
			:password,
			:created_at,
			:updated_at
		)`, u); err != nil {
		return err
	}
	if _, err = tx.NamedExec(`
		INSERT INTO merchants (
			uuid,
			user_uuid,
			merchant_name,
			merchant_address,
			created_at,
			updated_at
		) VALUES (
			:uuid,
			:user_uuid,
			:merchant_name,
			:merchant_address,
			:created_at,
			:updated_at
		)`,
		m); err != nil {

		return err
	}
	tx.Commit()
	return nil
}

func (u User) GetUserByEmail() (User, error) {
	if db, err := helper.GetDBClient(); err != nil {
		return User{}, err
	} else {
		defer db.Close()
		err := db.Get(&u, "SELECT * FROM users WHERE email=$1 LIMIT 1", u.Email)
		if err != nil {
			return User{}, nil
		}
		return u, nil
	}
}
