package success

import "net/http"

type successData struct {
	SuccessMessage any         `json:"message"`
	SuccessStatus  int         `json:"status"`
	Data           interface{} `json:"data"`
}

type MessageSuccess interface {
	Status() int
}

func (e *successData) Status() int {
	return e.SuccessStatus
}

func NewCreatedSuccess(message string, data interface{}) MessageSuccess {
	return &successData{
		SuccessMessage: message,
		SuccessStatus:  http.StatusCreated,
		Data:           data,
	}
}
