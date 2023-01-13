package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model `json:"-"`
	ID         int `json:"id"`
	Amount     int `json:"amount"`
}
