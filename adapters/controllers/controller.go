package controllers

import (
	"net/http"

	usecase "github.com/alexey/firstApp/boundary/domain/useCase"
	"github.com/alexey/firstApp/boundary/dto"
	"github.com/alexey/firstApp/internal/interfaces"
	loggerinterface "github.com/alexey/firstApp/pkg/logger/interface"
)

type BaseController struct {
	logger loggerinterface.Logger
	dto    *dto.UserDTO
}

type UserController struct {
	requestreader interfaces.RequestReader
	response      interfaces.Responderface
	authUseCase   usecase.UserUsecase
	BaseController
}

func NewController(

	log loggerinterface.Logger,
	auth usecase.UserUsecase,
	requtreader interfaces.RequestReader,
	response interfaces.Responderface,

) UserController {

	return UserController{

		BaseController: BaseController{
			logger: log,
		}, authUseCase: auth,
		requestreader: requtreader,
		response:      response,
	}
}

func (uc *UserController) LoginHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	/* Вычитываем запрос */
	req, err := uc.requestreader.ReadLoginRequest(r)
	if err != nil {
		uc.logger.PrintError(ctx, "Ошибка в чтении запроса", err)
		uc.response.ErrorResponse(w, err)
		return
	}

	user := dto.ModelUserFromDTO(req)

	/* Получаем ответ */
	tok, err := uc.authUseCase.GetUserAuthorities(ctx, user)
	if err != nil {
		uc.logger.PrintError(ctx, "Ошибка в получении ответа", err)
		uc.response.ErrorResponse(w, err)
		return
	}

	/* отправка ответа */
	uc.response.SuccessResponse(w, tok)
}

func (uc *UserController) ButtonHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req, err := uc.requestreader.ReadLoginRequest(r)
	if err != nil {
		uc.logger.PrintError(ctx, "Ошибка в чтении запроса", err)
		uc.response.ErrorResponse(w, err)
		return
	}

	user := dto.ModelUserFromDTO(req)

	tok, err := uc.authUseCase.OpenPathGuider(ctx, user)
	if err != nil {
		uc.logger.PrintError(ctx, "Ошибка в получении ответа", err)
		uc.response.ErrorResponse(w, err)
		return
	}

	uc.response.SuccessResponse(w, tok)

}
