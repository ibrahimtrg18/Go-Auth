package dto

type RegisterPartnerUser struct {
	FirstName       string `json:"firstName" validate:"required"`
	LastName        string `json:"lastName" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirmPassword"`
	Name            string `json:"name" validate:"required"`
	Type            string `json:"type" validate:"required"`
}
