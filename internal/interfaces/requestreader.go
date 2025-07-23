package interfaces

import (
	"net/http"

	"github.com/alexey/firstApp/boundary/dto"
)

type RequestReader interface {
	ReadLoginRequest(r *http.Request) (*dto.UserDTO, error)
}
