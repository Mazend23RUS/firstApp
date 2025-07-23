package models

import (
	"fmt"

	"github.com/alexey/firstApp/pkg/common"
	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID
	Name       string
	Email      string
	Password   string
	Status     UserStatus
	Roles      []Role
	IsSelected bool
}

// Value-объекты
type UserStatus string

const (
	StatusActive   UserStatus = "Active"
	StatusNoActive UserStatus = "NotActive"
)

func NewUser(
	email,
	passwors,
	name string,
	role []Role,
	isSelected bool,
) (*User, error) {

	if email == "" {
		return nil, fmt.Errorf("некоректный email")
	}

	if len(passwors) < 8 || passwors == "" {
		return nil, fmt.Errorf("не соответсвует требованием длинны пароля")
	}

	if name == "" {
		return nil, fmt.Errorf("имя роли не может быть пустым")
	}

	uuid := common.GetUUID()

	return &User{
		ID:         uuid,
		Password:   passwors,
		Email:      email,
		Name:       name,
		Roles:      role,
		Status:     StatusActive,
		IsSelected: isSelected,
	}, nil

}

func (us *User) InitRole(roleName string, permission []string) {

	us.Roles = []Role{
		{
			RoleName:   roleName,
			Permitions: permission,
		},
	}
}

func (us *User) ChangeEmail(newEmail string) {
	us.Email = newEmail

}
