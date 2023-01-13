package dto

import " hery-ciaputra/final-project-backend/models"

type CategoryRes struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

func (cr *CategoryRes) FromCategory(s *models.Category) *CategoryRes {
	return &CategoryRes{
		ID:          s.ID,
		Name:        s.Name,
		Description: s.Description,
		Price:       s.Price,
	}
}
