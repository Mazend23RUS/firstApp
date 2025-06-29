package errors_domain

import "errors"

var (
	ErrValidationFailed    = errors.New("validation failed")
	ErrInvalidCredentials  = errors.New("invalid credentials")
	ErrUserNotFound        = errors.New("user not found")
	ErrPermissionDenied    = errors.New("permission denied")
	ErrConflict            = errors.New("resource conflict")
	ErrRateLimitExceeded   = errors.New("rate limit exceeded")
	ErrServiceUnavailable  = errors.New("service unavailable")
	ErrInternalServerError = errors.New("internal server error")
)
