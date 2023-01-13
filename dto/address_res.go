package dto

import " hery-ciaputra/final-project-backend/models"

type AddressRes struct {
	ID        int    `json:"id"`
	Address   string `json:"address"`
	Recipient string `json:"recipient"`
	Phone     string `json:"phone"`
}

func (ar *AddressRes) FromAddress(a *models.Address) *AddressRes {
	return &AddressRes{
		ID:        a.ID,
		Address:   a.Address,
		Recipient: a.Recipient,
		Phone:     a.Phone,
	}
}
