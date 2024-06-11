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

type VerifyCodeRequest struct {
    Phone    string `json:"phone"`
    Code     string `json:"code"`
    Login    string `json:"login"`
    Password string `json:"password"`
}

type CreateUserRequest struct {
    Phone    string `json:"phone"`
    Login    string `json:"login"`
    Password string `json:"password"`
    UserType string `json:"user_type"`
}
