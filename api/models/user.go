package models

type User struct {
	ID 		    string `json:"id"`
	UserID      uint   `json:"user_id"`
	Phone       string `json:"phone"`
	Login       string `json:"login"`
	Password    string `json:"password"`
	Balance     uint   `json:"balance"`
	Count       uint   `json:"count"`
	Key         string `json:"key"`
	UserType    string `json:"user_type"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
}

type CreateUser struct {
	Phone       string `json:"phone"`
	Login       string `json:"login"`
	Password    string `json:"password"`
	UserType    string `json:"user_type"`
}

type UpdateUser struct {
	ID 		    string `json:"id"`
	Phone       string `json:"phone"`
	Login       string `json:"login"`
	Password    string `json:"password"`
}

type UsersResponse struct {
	Users       []User `json:"logins"`
	Count       int    `json:"count"`
}

type UpdateBalance struct {
	ID 		    string `json:"id"`
	Balance     uint   `json:"balance"`
}

type UpdateUserPassword struct {
	ID          string `json:"id"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}