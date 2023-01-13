package dto

import " hery-ciaputra/final-project-backend/models"

type RegisterRes struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Photo string `json:"photo"`
}

func (rr *RegisterRes) FromRegister(r *models.User) *RegisterRes {
	return &RegisterRes{
		ID:    r.ID,
		Name:  r.Name,
		Email: r.Name,
		Phone: r.Phone,
		Photo: r.Photo,
	}
}
