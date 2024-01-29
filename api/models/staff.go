package models

type Staff struct {
	ID         string `json:"id"`
	BranchID   string `json:"branch_id"`
	TariffID   string `json:"tariff_id"`
	StaffType  string `json:"staff_type"`
	Name       string `json:"name"`
	Balance    uint   `json:"balance"`
	Age        uint   `json:"age"`
	BirthDate  string `json:"birth_date"`
	Login      string `json:"login"`
	Password   string `json:"password"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	DeletedAt  string `json:"deleted_at"`
}

type CreateStaff struct {
	BranchID   string `json:"branch_id"`
	TariffID   string `json:"tariff_id"`
	StaffType  string `json:"staff_type"`
	Name       string `json:"name"`
	Balance    uint   `json:"balance"`
	BirthDate  string `json:"birth_date"`
	Login      string `json:"login"`
	Password   string `json:"password"`
}

type UpdateStaff struct {
	ID         string `json:"id"`
	BranchID   string `json:"branch_id"`
	TariffID   string `json:"tariff_id"`
	StaffType  string `json:"staff_type"`
	Name       string `json:"name"`
	Balance    uint   `json:"balance"`
	Login      string `json:"login"`
}

type StaffsResponse struct {
	Staffs    []Staff `json:"staffs"`
	Count     int     `json:"count"`
}

type UpdateStaffPassword struct {
	ID          string `json:"id"`
	NewPassword string `json:"new_password"`
	OldPassword string `json:"old_password"`
}