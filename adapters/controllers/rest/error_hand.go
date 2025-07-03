package readerequests

import (
	"errors"
	"net/http"

	errors_domain "github.com/alexey/boundary/domain/errors"
)

type ErrorStatus struct{}

func NewErrorStatus() *ErrorStatus {
	return &ErrorStatus{}
}

func mapErrorToStatus(err error) int {
	switch {
	case errors.Is(err, errors_domain.ErrValidationFailed):
		return http.StatusBadRequest
	case errors.Is(err, errors_domain.ErrInvalidCredentials):
		return http.StatusUnauthorized
	case errors.Is(err, errors_domain.ErrUserNotFound):
		return http.StatusNotFound
	case errors.Is(err, errors_domain.ErrPermissionDenied):
		return http.StatusForbidden
	case errors.Is(err, errors_domain.ErrConflict):
		return http.StatusConflict
	case errors.Is(err, errors_domain.ErrRateLimitExceeded):
		return http.StatusTooManyRequests
	default:
		// Для неизвестных ошибок возвращаем 500
		return http.StatusInternalServerError
	}

}

func (er *ErrorStatus) HandlerError(w http.ResponseWriter, err error) {
	w.WriteHeader(mapErrorToStatus(err))

}
