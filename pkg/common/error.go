package common

import (
	"errors"

	"github.com/gofiber/fiber"
)

type RespError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// CustomError for handling specific application errors
type CustomError struct {
	StatusCode int
	Message    string
}

// Error satisfies the error interface
func (e *CustomError) Error() string {
	return e.Message
}

func NewError(status int, message string) *CustomError {
	return &CustomError{
		Message:    message,
		StatusCode: status,
	}
}

// Error parser
// parseError determines the HTTP status and message from the error
func ParseError(err error) (int, string) {
	var customErr *CustomError
	if errors.As(err, &customErr) {
		// Handle a specific custom error type
		return customErr.StatusCode, customErr.Message
	}

	// Default to 500 Internal Server Error
	return fiber.StatusInternalServerError, "Internal Server Error"
}
