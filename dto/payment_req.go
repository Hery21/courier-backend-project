package dto

type PaymentReq struct {
	ShippingID int `json:"shipping_id"`
	PromoID    int `json:"promo_id"`
}
