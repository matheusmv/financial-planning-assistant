package util

import (
	"time"
)

type ErrorResponse struct {
	Error     string    `json:"error"`
	Timestamp time.Time `json:"timestamp"`
}

func NewErrorResponse(message string) ErrorResponse {
	return ErrorResponse{Error: message, Timestamp: time.Now()}
}
