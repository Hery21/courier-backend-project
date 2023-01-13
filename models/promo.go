package models

import (
	"gorm.io/gorm"
	"time"
)

type Promo struct {
	gorm.Model      `json:"-"`
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	MinimumOrder    int       `json:"minimum_order"`
	Discount        float32   `json:"discount"`
	MaximumDiscount int       `json:"maximum_discount"`
	Quota           int       `json:"quota"`
	ExpiryDate      time.Time `json:"expiry_date"`
}
