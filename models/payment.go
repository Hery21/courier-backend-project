package models

import "gorm.io/gorm"

type Payment struct {
	gorm.Model `json:"-"`
	ID         int    `json:"id"`
	Status     string `json:"status"`
	ShippingID int    `json:"shipping_id"`
	Amount     int    `json:"amount"`
	PromoID    int    `json:"promo_id"`
}
