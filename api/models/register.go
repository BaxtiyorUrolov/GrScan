package models

type CreateRegister struct {
	Phone     string `json:"phone"`
	Code      string `json:"code"`
}

type Register struct {
	ID        string `json:"id"`
	Phone     string `json:"phone"`
	Code      string `json:"code"`
	CreatedAT string `json:"created_at"`
}
