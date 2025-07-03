package readerequests

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alexey/adapters/controllers/rest/requests"
	"github.com/alexey/boundary/dto"
	"github.com/alexey/infrastructure/http/validator"
)

type JSONRequestReader struct{ validator *validator.Validator }

func NewJSONRequestReader() *JSONRequestReader {
	v := validator.NewValidator()

	return &JSONRequestReader{
		validator: v,
	}
}

func (js *JSONRequestReader) ReadLoginRequest(r *http.Request) (*dto.UserDTO, error) {
	var req requests.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, fmt.Errorf("Error decoding RenderJSONRequest()")
	}

	if err := js.validator.Validate(req); err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}

	return req.MapperOfRequestToDTO(), nil

}
