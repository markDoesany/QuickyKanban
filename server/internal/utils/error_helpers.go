package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type APIError struct {
	StatusCode int `json:"status_code"`
	Message    any `json:"message"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("api error: %s", e.Message)
}

func NewAPIError(statusCode int, errorMessage string) APIError {
	return APIError{
		StatusCode: statusCode,
		Message:    errorMessage,
	}
}

func WriteAPIError(w http.ResponseWriter, apiErr APIError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(apiErr.StatusCode)
	json.NewEncoder(w).Encode(apiErr)
}
