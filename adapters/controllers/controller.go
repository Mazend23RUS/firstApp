package controllers

import (
	"net/http"

	usecase "github.com/alexey/boundary/domain/useCase"

	"github.com/alexey/internal/interfaces"
	"github.com/alexey/pkg/logger"
)

type BaseController struct {
	reqvalid interfaces.RequestValidator
	logger   logger.Logger
}

type UserController struct {
	requestreader interfaces.RequestReader
	response      interfaces.Responderface
	errorhand     interfaces.ErrorHandler
	authUseCase   usecase.UserUsecase
	BaseController
}

func NewController(

	log logger.Logger,
	valid interfaces.RequestValidator,
	auth usecase.UserUsecase,
	requtreader interfaces.RequestReader,
	response interfaces.Responderface,
	errorhand interfaces.ErrorHandler,

) UserController {

	return UserController{

		BaseController: BaseController{
			logger:   log,
			reqvalid: valid,
		}, authUseCase: auth,
		requestreader: requtreader,
		response:      response,
		errorhand:     errorhand,
	}
}

func (uc *UserController) LoginHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	/* Вычитываем запрос */
	req, err := uc.requestreader.ReadLoginRequest(r)
	if err != nil {
		uc.response.ErrorResponse(w, err)
		return
	}

	/* Получаем ответ */
	tok, err := uc.authUseCase.GetUserAuthorities(ctx, *req)
	if err != nil {
		uc.logger.PrintError(ctx, "Ошибка в получении ответа", err)
		uc.errorhand.HandlerError(w, err)
		return
	}

	/* отправка ответа */
	uc.response.SuccessResponse(w, tok)
}

func (uc *UserController) ButtonHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req, err := uc.requestreader.ReadLoginRequest(r)
	if err != nil {
		uc.response.ErrorResponse(w, err)
		return
	}

	tok, err := uc.authUseCase.OpenPathGuider(ctx, *req)
	if err != nil {
		uc.logger.PrintError(ctx, "Ошибка в получении ответа", err)
		uc.errorhand.HandlerError(w, err)
		return
	}

	uc.response.SuccessResponse(w, tok)

}
