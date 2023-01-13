package handlers

import (
	" hery-ciaputra/final-project-backend/services"
)

type HandlerConfig struct {
	RegisterService services.RegisterService
	AuthService     services.AuthService
}

type Handler struct {
	registerService services.RegisterService
	authService     services.AuthService
}

func New(c *HandlerConfig) *Handler {
	return &Handler{
		registerService: c.RegisterService,
		authService:     c.AuthService,
	}
}
