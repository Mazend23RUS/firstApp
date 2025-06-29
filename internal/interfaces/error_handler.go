package interfaces

import "net/http"

type ErrorHandler interface {
	HandlerError(w http.ResponseWriter, er error)
}
