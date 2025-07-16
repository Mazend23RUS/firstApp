package dto

type UserDTO struct {
	ID          int    `json:"id,omitempty"`
	Password    string `json:"password" validate:"required,min=8"`
	Name        string `json:"name,omitempty"`
	Role        string `json:"role,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Email       string `json:"email" validate:"required,email"`
	IsSelected  bool   `json:"is_selected"`
}

func NewUserDTO(email, password string, IsSeted bool) *UserDTO {
	return &UserDTO{
		Email:      email,
		Password:   password,
		IsSelected: IsSeted,
	}
}
