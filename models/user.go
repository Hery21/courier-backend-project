package models

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"-"`
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Role       string `json:"role"`
	Phone      string `json:"phone"`
	Photo      string `json:"photo"`
	Balance    int    `json:"balance"`
	Password   string `json:"-"`
}
