package dto

import " hery-ciaputra/final-project-backend/models"

type ProfileRes struct {
	ID      int    `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Photo   string `json:"photo"`
	Balance int    `json:"balance"`
}

func (pr *ProfileRes) FromUser(u *models.User) *ProfileRes {
	return &ProfileRes{
		ID:      u.ID,
		Email:   u.Email,
		Name:    u.Name,
		Phone:   u.Phone,
		Photo:   u.Photo,
		Balance: u.Balance,
	}
}
