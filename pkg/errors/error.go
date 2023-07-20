package errors

import "net/http"

type errData struct {
	ErrMessage any    `json:"message"`
	ErrStatus  int    `json:"status"`
	ErrError   string `json:"error"`
}

type MessageErr interface {
	Status() int
	Error() string
}

func (e *errData) Status() int {
	return e.ErrStatus
}
func (e *errData) Error() string {
	return e.ErrError
}

func NewBadRequestError(message string) MessageErr {
	return &errData{
		ErrMessage: message,
		ErrStatus:  http.StatusBadRequest,
		ErrError:   "BAD_REQUEST",
	}
}

func NewInternalServerError(message string) MessageErr {
	return &errData{
		ErrMessage: message,
		ErrStatus:  http.StatusInternalServerError,
		ErrError:   "INTERNAL_SERVER_ERROR",
	}
}

func NewConflictError(message string) MessageErr {
	return &errData{
		ErrMessage: message,
		ErrStatus:  http.StatusConflict,
		ErrError:   "CONFLICT",
	}
}

func NewUnprocessableEntityError(message []string) MessageErr {
	return &errData{
		ErrMessage: message,
		ErrStatus:  http.StatusUnprocessableEntity,
		ErrError:   "UNPROCESSABLE_ENTITY",
	}
}

func NewNotFoundError(message string) MessageErr {
	return &errData{
		ErrMessage: message,
		ErrStatus:  http.StatusNotFound,
		ErrError:   "NOT_FOUND",
	}
}

func NewUnautorizhedError(message string) MessageErr {
	return &errData{
		ErrMessage: message,
		ErrStatus:  http.StatusForbidden,
		ErrError:   "UNAUTHORIZED",
	}
}

func NewUnauthenticatedError(message string) MessageErr {
	return &errData{
		ErrMessage: message,
		ErrStatus:  http.StatusUnauthorized,
		ErrError:   "NOT_AUTHENTICATED",
	}
}
