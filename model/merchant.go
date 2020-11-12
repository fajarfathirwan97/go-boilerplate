package model

import "time"

type Merchant struct {
	Uuid            string    `json:"uuid" db:"uuid"`
	UserUuid        string    `json:"user_uuid" db:"user_uuid"`
	MerchantName    string    `json:"merchant_name" db:"merchant_name"`
	MerchantAddress string    `json:"merchant_address" db:"merchant_address"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}
