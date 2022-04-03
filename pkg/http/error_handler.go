package http

import (
	"encoding/json"
	"net/http"
)

type (
	ErrorHandler func(w http.ResponseWriter, r *http.Request) error

	Error interface {
		Error() string
		ErrorCode() int
		ErrorType() string
	}

	ErrorResponse struct {
		Error     string `json:"error"`
		ErrorCode int    `json:"error_code"`
		ErrorType string `json:"error_type"`
	}
)

func (handler ErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := handler(w, r); err != nil {
		var response ErrorResponse
		value, cast := err.(Error)
		if !cast {
			response = ErrorResponse{
				Error:     err.Error(),
				ErrorCode: -1,
				ErrorType: "undefined",
			}
		} else {
			response = ErrorResponse{
				Error:     value.Error(),
				ErrorCode: value.ErrorCode(),
				ErrorType: value.ErrorType(),
			}
		}

		json.NewEncoder(w).Encode(response)
	}
}
