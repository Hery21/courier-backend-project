package services

import (
	" hery-ciaputra/final-project-backend/config"
	" hery-ciaputra/final-project-backend/dto"
	" hery-ciaputra/final-project-backend/models"
	" hery-ciaputra/final-project-backend/repositories"
	"golang.org/x/crypto/bcrypt"
)

type RegisterService interface {
	Register(req *dto.RegisterReq) (*dto.RegisterRes, error)
}

type registerService struct {
	registerRepository repositories.RegisterRepository
	appConfig          config.AppConfig
}

type RSConfig struct {
	RegisterRepository repositories.RegisterRepository
}

func NewRegisterService(r *RSConfig) RegisterService {
	return &registerService{
		registerRepository: r.RegisterRepository,
	}
}

func (r *registerService) Register(req *dto.RegisterReq) (*dto.RegisterRes, error) {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	req.Password = string(bytes)
	registeringUser := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Photo:    req.Photo,
		Password: req.Password,
	}

	registeredUser, err := r.registerRepository.Register(registeringUser)

	if err != nil {
		return new(dto.RegisterRes), err
	}

	return new(dto.RegisterRes).FromRegister(registeredUser), err
}
