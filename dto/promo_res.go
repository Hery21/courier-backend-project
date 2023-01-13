package dto

import (
	" hery-ciaputra/final-project-backend/models"
	"time"
)

type PromoRes struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	MinimumOrder    int       `json:"minimum_order"`
	Discount        float32   `json:"discount"`
	MaximumDiscount int       `json:"maximum_discount"`
	Quota           int       `json:"quota"`
	ExpiryDate      time.Time `json:"expiry_date"`
}

func (pr *PromoRes) FromPromo(p *models.Promo) *PromoRes {
	return &PromoRes{
		ID:              p.ID,
		Name:            p.Name,
		MinimumOrder:    p.MinimumOrder,
		Discount:        p.Discount,
		MaximumDiscount: p.MaximumDiscount,
		Quota:           p.Quota,
		ExpiryDate:      p.ExpiryDate,
	}
}
