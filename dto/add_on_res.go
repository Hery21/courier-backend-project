package dto

import " hery-ciaputra/final-project-backend/models"

type AddOnRes struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

func (ar *AddOnRes) FromAddOn(s *models.AddOn) *AddOnRes {
	return &AddOnRes{
		ID:          s.ID,
		Name:        s.Name,
		Description: s.Description,
		Price:       s.Price,
	}
}
