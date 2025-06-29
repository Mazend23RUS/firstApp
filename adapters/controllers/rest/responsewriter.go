package renderequests

import (
	"encoding/json"
	"net/http"
)

type ResponseWriter struct{}

func NewResponseWriter() ResponseWriter {
	return ResponseWriter{}
}

func (rw ResponseWriter) SuccessResponse(w http.ResponseWriter, data interface{}) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (rw ResponseWriter) ErrorResponse(w http.ResponseWriter, err error) {
	status := mapErrorToStatus(err)

	w.WriteHeader(status)
	json.NewEncoder(w).Encode(struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	})

}
