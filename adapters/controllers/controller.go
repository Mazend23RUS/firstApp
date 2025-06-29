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
	authUseCase   usecase.UserUsecase
	BaseController
}

func NewController(

	log logger.Logger,
	valid interfaces.RequestValidator,
	auth usecase.UserUsecase,
	requtreader interfaces.RequestReader,
	response interfaces.Responderface,

) UserController {

	return UserController{

		BaseController: BaseController{
			logger:   log,
			reqvalid: valid,
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
		uc.response.ErrorResponse(w, err)
		return
	}

	/* Получаем ответ */
	tok, err := uc.authUseCase.GetUserAuthorities(ctx, *req)

	/* отправка ответа */
	uc.response.SuccessResponse(w, tok)
}
