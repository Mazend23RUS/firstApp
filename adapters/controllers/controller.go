package controllers

import (
	"net/http"

	usecase "github.com/alexey/boundary/domain/useCase"
	"github.com/alexey/boundary/dto"

	"github.com/alexey/internal/interfaces"
	"github.com/alexey/pkg/logger"
)

type BaseController struct {
	reqvalid interfaces.RequestValidator
	logger   logger.Logger
}

type UserController struct {
	requestreader interfaces.RequestReader
	// errhand     interfaces.ErrorHandler
	response    interfaces.Responderface
	authUseCase usecase.UserUsecase
	BaseController
}

func NewController(

	log logger.Logger,
	valid interfaces.RequestValidator,
	auth usecase.UserUsecase,
	requtreader interfaces.RequestReader,
	// errhandler interfaces.ErrorHandler,
	response interfaces.Responderface,
	// err  interfaces.ErrorHandler,

) UserController {

	return UserController{

		BaseController: BaseController{
			logger:   log,
			reqvalid: valid,
			// err: err,

		}, authUseCase: auth,
		requestreader: requtreader,
		// errhand: errhandler,
		response: response,
	}
}

func (uc *UserController) LoginHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	/* Парсер запроса */
	req, err := uc.requestreader.ReadLoginRequest(r)
	if err != nil {
		uc.response.ErrorResponse(w, err)
		return
	}
	/* Валидация запроса */
	// if err := uc.reqvalid.Validate(req); err != nil {
	// 	uc.logger.PrintError(ctx, "Ошибка валидации", err)
	// 	http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	// 	return
	// }

	/* Получаем ответ */
	tok, err := uc.authUseCase.GetUserAuthorities(ctx, *req)

	/* отправка ответа */
	uc.response.SuccessResponse(w, tok)
}

// func (uc *UserController) JsonSerialaser(r *http.Request, w http.ResponseWriter, req *requests.LoginRequest) []byte {
// 	ctx := r.Context()
// 	response := uc.GetResponseFromUseCaseAuthorities(w, r, req)
// 	/* 5. Сериализация ответа */
// 	jsonSerialaser, err := json.Marshal(response)
// 	uc.logger.PtintInfo(ctx, "Отарабатывает сериализация")
// 	if err != nil {
// 		uc.logger.PrintError(ctx, "Не получилось сериализовать с помощью маршала ", err)
// 		http.Error(w, "Ошибочка сервера", http.StatusInternalServerError)
// 		return jsonSerialaser
// 	}
// 	return jsonSerialaser
// }

// func (uc *UserController) SendResponse(w http.ResponseWriter, r *http.Request, output usecase.UserAuthoritiesOutput, req *requests.LoginRequest) {
// 	serialas := uc.JsonSerialaser(r, w, req)
// 	ctx := r.Context()
// 	/* 7. Отправка ответа */
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	if _, err := w.Write(serialas); err != nil {
// 		uc.logger.PrintError(ctx, "Failed to write response", err)
// 	}
// }

func (uc *UserController) ActiveteUseCaseAuthorities(outputDto dto.UserDTO, r *http.Request, w http.ResponseWriter) usecase.UserAuthoritiesOutput {
	ctx := r.Context()

	output, err := uc.authUseCase.GetUserAuthorities(ctx, dto.UserDTO{
		Password: outputDto.Password,
		Email:    outputDto.Email,
	})
	uc.logger.PtintInfo(ctx, "Вызван usecase GetUserAuthorities")
	if err != nil {
		uc.logger.PrintError(ctx, "", err)
		http.Error(w, "Аунтификации не прошла ", http.StatusUnauthorized)
		return *output
	}
	return *output
}

// func (uc *UserController) GetResponseFromUseCaseAuthorities(w http.ResponseWriter, r *http.Request, req *requests.LoginRequest) usecase.UserAuthoritiesOutput {

// 	otputdto, err := uc.parser.RequestToDto(req)
// 	if err != nil {
// 		fmt.Println("Ошибка получения дто")
// 	}

// 	/* 3. Вызов use case и получение ответа */
// 	response := uc.ActiveteUseCaseAuthorities(otputdto, r, w)
// 	return response

// }
