package dto

import (
	"github.com/alexey/firstApp/domain/models"
	"github.com/google/uuid"
)

type UserDTO struct {
	ID         uuid.UUID     `json:"id,omitempty"`
	Email      string        `json:"email" validate:"required,email"`
	Password   string        `json:"password" validate:"required,min=8"`
	Name       string        `json:"name,omitempty"`
	Role       []models.Role `json:"role,omitempty"`
	IsSelected bool          `json:"is_selected"`
	// PhoneNumber string `json:"phone_number,omitempty"`

}

func NewUserDTO(email, password string, IsSeted bool) *UserDTO {
	return &UserDTO{
		Email:      email,
		Password:   password,
		IsSelected: IsSeted,
	}
}

func ModelUserFromDTO(dto *UserDTO) *models.User {
	return &models.User{
		ID:         dto.ID,
		Email:      dto.Email,
		Password:   dto.Password,
		Roles:      dto.Role,
		IsSelected: dto.IsSelected,
	}
}

func DTOFromModel(user *models.User) *UserDTO {
	return &UserDTO{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
	}
}
