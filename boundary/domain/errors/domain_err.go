package errors_domain

import "errors"

var (
	ErrValidationFailed     = errors.New("validation failed")
	ErrInvalidCredentials   = errors.New("invalid credentials")
	ErrUserNotFound         = errors.New("user not found")
	ErrPermissionDenied     = errors.New("permission denied")
	ErrConflict             = errors.New("resource conflict")
	ErrRateLimitExceeded    = errors.New("rate limit exceeded")
	ErrServiceUnavailable   = errors.New("service unavailable")
	ErrInternalServerError  = errors.New("internal server error")
	ErrValidationRoleName   = errors.New("invalid roleName")
	ErrValidationPermission = errors.New("invalid permission")
	ErrValidationEmail      = errors.New("invalid email")
	ErrLenghPassword        = errors.New("short password")
	ErrPassNotContainNum    = errors.New("password not contain number")
	ErrUpperSimbolInName    = errors.New("lower symbol in name")
)
