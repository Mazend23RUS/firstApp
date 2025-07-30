package models

import (
	"fmt"
	"regexp"
	"unicode"

	"github.com/alexey/firstApp/pkg/common"
	"github.com/google/uuid"
)

type User struct {
	id         uuid.UUID
	name       string
	email      string
	password   string
	status     UserStatus
	roles      []Role
	isSelected bool
}

// Доменные примитивы
type Email string
type Password string
type Name string

func NewEmail(value string) (Email, error) {
	if value == "" {
		return "", fmt.Errorf("field cant be empty")
	}

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	if !regexp.MustCompile(emailRegex).MatchString(value) {
		return "", fmt.Errorf("not valid email")
	}

	return Email(value), nil
}

func NewPassword(value string) (Password, error) {
	if len(value) < 8 {
		return "", fmt.Errorf("lengh cant be < 8 sibols")
	}

	UpperSimbol := false
	Number := false

	for _, c := range value {

		switch {
		case unicode.IsUpper(c):
			UpperSimbol = true
		case unicode.IsNumber(c):
			Number = true
		}
	}

	if !UpperSimbol || !Number {
		return "", fmt.Errorf("password must have one Upper simbol and Number")
	}

	return Password(value), nil
}

func NewName(value string) (Name, error) {

	if len(value) < 3 {
		return "", fmt.Errorf("name must be not short then 3 simbols")
	}

	runes := []rune(value)

	if !unicode.IsUpper(runes[0]) {
		return "", fmt.Errorf("firs simbol bust be Upper ")
	}

	return Name(value), nil
}

// Value-объекты
type UserStatus string

const (
	StatusActive   UserStatus = "Active"
	StatusNoActive UserStatus = "NotActive"
)

func NewUserModelFromDTO(id uuid.UUID, name string, email string, password string, ro []Role, isSel bool) *User {
	return &User{
		id:         id,
		name:       name,
		email:      email,
		password:   password,
		roles:      ro,
		isSelected: isSel,
	}
}

func NewModelForTest(email string, pass string) *User {
	return &User{
		email:    email,
		password: pass,
	}
}

func NewUser(
	email,
	password,
	name string,
	role []Role,
	isSelected bool,
) (*User, error) {

	emailValue, err := NewEmail(email)
	if err != nil {
		return nil, fmt.Errorf("email validation failed: %w", err)
	}

	passwordValue, err := NewPassword(password)
	if err != nil {
		return nil, fmt.Errorf("password validation failed: %w", err)
	}

	nameValue, err := NewName(name)
	if err != nil {
		return nil, fmt.Errorf("name validation failed: %w", err)
	}

	// if email == "" {
	// 	return nil, fmt.Errorf("некоректный email")
	// }

	// if len(passwors) < 8 || passwors == "" {
	// 	return nil, fmt.Errorf("не соответсвует требованием длинны пароля")
	// }

	// if name == "" {
	// 	return nil, fmt.Errorf("имя роли не может быть пустым")
	// }

	uuid := common.GetUUID()

	return &User{
		id:         uuid,
		password:   string(passwordValue),
		email:      string(emailValue),
		name:       string(nameValue),
		roles:      role,
		status:     StatusActive,
		isSelected: isSelected,
	}, nil

}

func (us *User) InitRole(roleName string, permission []string) {

	us.roles = []Role{
		{
			roleName:    roleName,
			permissions: permission,
		},
	}
}

func (us *User) ChangeEmail(newEmail string) {
	us.email = newEmail
}

func (us *User) Id() uuid.UUID {
	return us.id
}

func (us *User) Name() string {
	return us.name
}

func (us *User) Email() string {
	return us.email
}

func (us *User) Password() string {
	return us.password
}

func (us *User) Status() UserStatus {
	return us.status
}

func (us *User) Roles() []Role {
	return us.roles
}

func (us *User) IsSelected() bool {
	return us.isSelected
}
