package usecase

import (
    "context"

)


    // Интерфейсы для работы с сценариями пользователя
type (
     

	// Получаем авторизацию пользователя 
     GetUserAuthoritiesUseCase interface {
		Execute(cont context.Context, input GetUserAuthoritiesInput) (GetUserAuthoritiesOutput, error)
	}
     // отправляем команду на открытие проводника
    ComandToOpenPathguiderUseCase interface {
        Execute(cont context.Context, isOpenGuider bool) (ResultOfComand error) 
	}
     
)


type GetUserAuthoritiesInput struct {
        Email string
		Password string
}

type GetUserAuthoritiesOutput struct {
        Email string
		Password string
		Role string 
}