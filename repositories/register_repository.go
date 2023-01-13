package repositories

import (
	" hery-ciaputra/final-project-backend/httperror"
	" hery-ciaputra/final-project-backend/models"
	"gorm.io/gorm"
)

type RegisterRepository interface {
	Register(user *models.User) (*models.User, error)
}

type registerRepository struct {
	db *gorm.DB
}

type RRConfig struct {
	DB *gorm.DB
}

func NewRegisterRepository(c *RRConfig) *registerRepository {
	return &registerRepository{db: c.DB}
}

func (r *registerRepository) Register(user *models.User) (*models.User, error) {
	res := r.db.Select("Name", "Email", "Password", "Phone", "Photo").Create(&user)

	if err := res.Error; err != nil {
		return &models.User{}, httperror.InternalServerError(err.Error())
	}

	registeredUser := &models.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone,
		Photo: user.Photo,
	}

	return registeredUser, res.Error
}
