package dto

type ProfileReq struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Photo string `json:"photo"`
}
