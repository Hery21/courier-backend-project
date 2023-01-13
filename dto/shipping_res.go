package dto

import " hery-ciaputra/final-project-backend/models"

type ShippingRes struct {
	ID       int             `json:"id"`
	Status   string          `json:"status"`
	Review   string          `json:"review"`
	Size     models.Size     `json:"size"`
	Category models.Category `json:"category"`
	AddOn    models.AddOn    `json:"add_on"`
	Address  models.Address  `json:"address"`
	//PaymentID int             `json:"payment_id"`
}

func (se *ShippingRes) FromShipping(s *models.Shipping) *ShippingRes {
	return &ShippingRes{
		ID:       s.ID,
		Status:   s.Status,
		Review:   s.Review,
		Size:     s.Size,
		Category: s.Category,
		AddOn:    s.AddOn,
		Address:  s.Address,
		//PaymentID: s.PaymentID,
	}
}
