package models

import "gorm.io/gorm"

type Shipping struct {
	gorm.Model `json:"-"`
	ID         int      `json:"id"`
	Status     string   `json:"status"`
	Review     string   `json:"review"`
	SizeID     int      `json:"size_id"`
	CategoryID int      `json:"category_id"`
	AddOnID    int      `json:"add_on_id"`
	AddressID  int      `json:"address_id"`
	Size       Size     `gorm:"foreignkey:SizeID"`
	Category   Category `gorm:"foreignkey:CategoryID"`
	AddOn      AddOn    `gorm:"foreignkey:AddOnID"`
	Address    Address  `gorm:"foreignkey:AddressID"`
	PaymentID  int      `json:"payment_id"`
}
