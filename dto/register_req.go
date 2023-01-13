package dto

type RegisterReq struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Photo    string `json:"photo"`
	Phone    string `json:"phone"`
}
