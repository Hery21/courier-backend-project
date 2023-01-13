package dto

import "time"

type PromoReq struct {
	Name            string    `json:"name"`
	MinimumOrder    int       `json:"minimum_order"`
	Discount        float32   `json:"discount"`
	MaximumDiscount int       `json:"maximum_discount"`
	Quota           int       `json:"quota"`
	ExpiryDate      time.Time `json:"expiry_date"`
}
