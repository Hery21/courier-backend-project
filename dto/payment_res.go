package dto

import " hery-ciaputra/final-project-backend/models"

type PaymentRes struct {
	ID         int    `json:"id"`
	Status     string `json:"status"`
	ShippingID int    `json:"shipping_id"`
	Amount     int    `json:"amount"`
	PromoID    int    `json:"promo_id"`
}

func (pr *PaymentRes) FromPayment(p *models.Payment) *PaymentRes {
	return &PaymentRes{
		ID:         p.ID,
		Status:     p.Status,
		ShippingID: p.ShippingID,
		Amount:     p.Amount,
		PromoID:    p.PromoID,
	}
}
