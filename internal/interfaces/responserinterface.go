package interfaces

import "net/http"

type Responderface interface {
	SuccessResponse(w http.ResponseWriter, data interface{})
	ErrorResponse(w http.ResponseWriter, err error)
}
