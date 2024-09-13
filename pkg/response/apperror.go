package response

import (
	"net/http"
)

type AppError struct {
	Id       string
	HttpCode int
}

func (e *AppError) Error() string {
	return e.Id
}

func BadRequestError(message string) *AppError {
	return &AppError{message, http.StatusBadRequest}
}

func UnauthorizedError(message string) *AppError {
	return &AppError{message, http.StatusUnauthorized}
}

func ForbiddenError(message string) *AppError {
	return &AppError{message, http.StatusForbidden}
}

func NotFoundError(message string) *AppError {
	return &AppError{message, http.StatusNotFound}
}

func InternalServerError(message string) *AppError {
	return &AppError{message, http.StatusInternalServerError}
}

func ServiceUnavailableError(message string) *AppError {
	return &AppError{message, http.StatusServiceUnavailable}
}
