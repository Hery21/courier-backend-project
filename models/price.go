package models

import "gorm.io/gorm"

type Price struct {
	gorm.Model      `json:"-"`
	SizesPrice      int `gorm:"column:sizes.price"`
	CategoriesPrice int `gorm:"column:categories.price"`
	AddOnsPrice     int `gorm:"column:add_ons.price"`
	Discount        int `gorm:"column:discount.price"`
}
