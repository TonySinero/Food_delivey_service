package domain

type Customer struct {
	ID          int    `json:"id" db:"id" validate:"required"`
	FullName    string `json:"fullName" db:"full_name" validate:"required"`
	Address     string `json:"address" db:"address" validate:"required"`
	PhoneNumber int    `json:"phoneNumber" db:"phone_number" validate:"required"`
}

type CustomerInfo struct {
	FullName    string `json:"fullName" db:"full_name"`
	Address     string `json:"address" db:"address"`
	PhoneNumber string `json:"phoneNumber" db:"phone_number"`
	BirthDay    string `json:"birthday" db:"birthday"`
}

type CustomerUpdate struct {
	FullName    *string `json:"fullName" db:"full_name"`
	Address     *string `json:"address" db:"address"`
	PhoneNumber *int    `json:"phoneNumber" db:"phone_number"`
	BirthDay    *string `json:"birthday" db:"birthday"`
}
