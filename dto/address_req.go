package dto

type AddressReq struct {
	ID        int    `json:"id"`
	Recipient string `json:"recipient"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
}
