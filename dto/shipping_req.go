package dto

type ShippingReq struct {
	SizeID     int `json:"size_id"`
	CategoryID int `json:"category_id"`
	AddressID  int `json:"address_id"`
	PromoID    int `json:"promo_id"`
	AddOnID    int `json:"add_on_id"`
}
