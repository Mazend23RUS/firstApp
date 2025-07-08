package readerequests

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alexey/adapters/controllers/rest/requests"
	errors_domain "github.com/alexey/boundary/domain/errors"
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
		return nil, fmt.Errorf("%w: %v", errors_domain.ErrValidationFailed, err)
	}

	if err := js.validator.Validate(req); err != nil {
		return nil, fmt.Errorf("%w: %v", errors_domain.ErrValidationFailed, err)
	}

	return req.MapperOfRequestToDTO(), nil

}
