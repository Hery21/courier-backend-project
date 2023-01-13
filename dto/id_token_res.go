package dto

type TokenResponse struct {
	IDToken string `json:"idToken"`
}

func (_ *TokenResponse) FromUser(token string) *TokenResponse {
	return &TokenResponse{
		IDToken: token,
	}
}
