package implementationUseCase

import (
	"context"
	"fmt"
	"time"

	usecase "github.com/alexey/firstApp/boundary/domain/useCase"

	"github.com/alexey/firstApp/domain/models"
	"github.com/alexey/firstApp/domain/repository"

	loggerinterface "github.com/alexey/firstApp/pkg/logger/interface"
)

type AuthUseCase struct {
	log      loggerinterface.Logger
	userRepo repository.InterfaceUserRepository
}

func NewAuthUseCase(log loggerinterface.Logger) *AuthUseCase {

	return &AuthUseCase{
		log: log,
	}
}

// func (us *AuthUseCase) GetUserAuthorities(c context.Context, input dto.UserDTO) (*usecase.UserAuthoritiesOutput, error) {

func (us *AuthUseCase) GetUserAuthorities(c context.Context, input *models.User) (*usecase.UserAuthoritiesOutput, error) {

	if input.Email == "bboy23@mail.ru" && input.Password == "87654321" {
		us.log.PrintInfo(c, "Введены данные Admin"+" email: "+input.Email)

		input.InitRole("Admin", []string{"read, write"})
		newuser, err := models.NewUser(input.Email, input.Password, "Новый пользователь admin", input.Roles, input.IsSelected)
		if err != nil {
			return nil, fmt.Errorf("Не создался юзер", err)
		}
		fmt.Println(newuser)
		// Сохранение юзера пока не работает, нужно подключение к БД
		// us.userRepo.Save(c, newuser)
		return &usecase.UserAuthoritiesOutput{
			Email:     input.Email,
			Role:      "Admin",
			Token:     "tokke-15654-5631-45$",
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
		}, nil

	}

	us.log.PrintInfo(c, "Введены данные для user")
	input.InitRole("User", []string{"read"})
	newuser, err := models.NewUser(input.Email, input.Password, "Новый пользователь с правами user", input.Roles, input.IsSelected)
	if err != nil {
		return nil, fmt.Errorf("Не создался юзер %w", err)
	}
	fmt.Println(newuser)
	// Сохранение юзера пока не работает, нужно подключение к БД
	// us.userRepo.Save(c, newuser)
	return &usecase.UserAuthoritiesOutput{
		Email:     input.Email,
		Role:      "user",
		Token:     "168456ewq",
		ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
	}, nil
}

func (us *AuthUseCase) OpenPathGuider(c context.Context, input *models.User) (*usecase.UserAuthoritiesOutput, error) {

	// Пока не работает нужно будет подключение к БД
	userFromDB, err := us.userRepo.GetUserByEmail(c, input.Email)
	if err != nil {
		return nil, fmt.Errorf("Не удалось получить пользователя из БД %w", err)
	}

	if input.IsSelected == true && userFromDB.Status == "Active" {
		us.log.PrintInfo(c, "Нажата кнопка открытия проводника")

		return &usecase.UserAuthoritiesOutput{
			Email:      input.Email,
			Token:      "tokke-15654-5631-46$",
			ExpiresAt:  time.Now().Add(12 * time.Hour).Unix(),
			IsSelected: "Поисковик открылся",
		}, nil
	}
	return nil, fmt.Errorf("Ошибка при открытии окна %v", c)
}
