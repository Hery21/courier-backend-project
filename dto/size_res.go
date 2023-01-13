package dto

import " hery-ciaputra/final-project-backend/models"

type SizeRes struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

func (sr *SizeRes) FromSize(s *models.Size) *SizeRes {
	return &SizeRes{
		ID:          s.ID,
		Name:        s.Name,
		Description: s.Description,
		Price:       s.Price,
	}
}
