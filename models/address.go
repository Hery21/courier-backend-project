package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model `json:"-"`
	ID         int    `json:"id"`
	Recipient  string `json:"recipient"`
	Address    string `json:"address"`
	Phone      string `json:"phone"`
	UserID     int    `json:"user_id"`
}
