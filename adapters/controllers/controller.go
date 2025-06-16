package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/alexey/adapters/controllers/rest/requests"
	usecase "github.com/alexey/boundary/domain/useCase"
	"github.com/alexey/boundary/dto"
	implementationUseCase "github.com/alexey/domain/usecase"
	"github.com/alexey/pkg/logger"
)

type UserController struct {
	authUseCase implementationUseCase.AuthUseCase
	logger      logger.Logger
}

func NewController(auth implementationUseCase.AuthUseCase, log logger.Logger) *UserController {

	return &UserController{
		authUseCase: auth,
		logger:      log,
	}
}

func (uc *UserController) Logger(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req requests.LoginRequest
	var but requests.IsSelected

	// 1. Парсинг запроса
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		uc.logger.PrintError(ctx, "Не получается декодировать запрос", err)
		http.Error(w, "Не верный запрос", http.StatusBadRequest)
		return
	}

	// 2. Валидация
	if err := req.Validate(); err != nil {
		uc.logger.PrintError(ctx, "Валидацию не прошел", err)
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	otputdto := req.MapperOfRequestToDTO()

	// 3. Вызов use case
	output, err := uc.authUseCase.GetUserAuthorities(ctx, dto.UserDTO{
		Email:    otputdto.Email,
		Password: otputdto.Password,
	})
	uc.logger.PtintInfo(ctx, "Вызван usecase GetUserAuthorities")
	if err != nil {
		uc.logger.PrintError(ctx, "Все пошло по пиии", err)
		http.Error(w, "Аунтификации пипец ", http.StatusUnauthorized)
		return
	}

	if err := uc.authUseCase.OpenPathGuider(ctx, but.Selected()); err != nil {
		uc.logger.PrintError(ctx, "Path guider failed", err)
		// Не прерываем выполнение, так как это дополнительная функциональность
	}

	// 4. Формирование ответа
	response := usecase.UserAuthoritiesOutput{
		Email:     output.Email,
		Role:      output.Role,
		Token:     output.Token,
		ExpiresAt: output.ExpiresAt,
	}

	// 5. Сериализация ответа
	jsonSerialaser, err := json.Marshal(response)
	uc.logger.PtintInfo(ctx, "Отарабатывает сериализация")
	if err != nil {
		uc.logger.PrintError(ctx, "Не получилось сериализовать с помощью маршала ", err)
		http.Error(w, "Ошибочка сервера", http.StatusInternalServerError)
		return
	}

	// 7. Отправка ответа
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(jsonSerialaser); err != nil {
		uc.logger.PrintError(ctx, "Failed to write response", err)
	}

}
