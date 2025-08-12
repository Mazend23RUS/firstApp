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
	msgChan chan<- interface{}
}

func NewController(

	log loggerinterface.Logger,
	auth usecase.UserUsecase,
	requtreader interfaces.RequestReader,
	response interfaces.Responderface,
	msgChanel chan<- interface{},

) UserController {

	return UserController{

		BaseController: BaseController{
			logger: log,
		}, authUseCase: auth,
		requestreader: requtreader,
		response:      response,
		msgChan:       msgChanel,
	}
}

func (uc *UserController) LoginHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// Отправка сообщения в канал с защитой от блокировки
	sendToChannel := func(msg any) {
		select {
		case uc.msgChan <- msg:
		default:
			uc.logger.PrintError(ctx, "Не удалось отправить сообщение в канал: канал заполнен", nil)
		}
	}

	sendToChannel("Начало обработки запроса LoginHandler")

	/* Вычитываем запрос */
	req, err := uc.requestreader.ReadLoginRequest(r)
	if err != nil {
		uc.msgChan <- "ошибка отправленная в канал" + err.Error()
		uc.logger.PrintError(ctx, "Ошибка в чтении запроса", err)
		uc.response.ErrorResponse(w, err)
		return
	}

	sendToChannel("Успешно прочитан запрос от пользователя: " + req.Email)

	user := dto.ModelUserFromDTO(req)

	/* Получаем ответ */
	tok, err := uc.authUseCase.GetUserAuthorities(ctx, user)
	if err != nil {
		uc.logger.PrintError(ctx, "Ошибка в получении ответа", err)
		uc.response.ErrorResponse(w, err)
		return
	}

	sendToChannel("Успешная авторизация пользователя: " + req.Email)
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
