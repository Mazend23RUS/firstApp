package dto

type UserDTO struct {
	ID          int    `json:"id,omitempty"`
	Password    string `json:"password" validate:"required,password"`
	Name        string `json:"name,omitempty"`
	Role        string `json:"role,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Email       string `json:"email" validate:"required,email"`
}

func NewUserDTO(email, password string) *UserDTO {
	return &UserDTO{
		Email:    email,
		Password: password,
	}
}
