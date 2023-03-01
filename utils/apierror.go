package apierror

import "net/http"

type Apierror interface {
	Code() int
	Message() string
	Cause() map[string]string
}

type apierr struct {
	code    int               `json:"code"`
	message string            `json:"message"`
	cause   map[string]string `json:"cause"`
}

func (e *apierr) Code() int {
	return e.code
}

func (e *apierr) Message() string {
	return e.message
}

func (e *apierr) Cause() map[string]string {
	return e.cause
}

// atributo no valido status 400
func (e *apierr) NewBadRequest(messageIn string, causeIn map[string]string) *apierr {
	return &apierr{
		code:    http.StatusBadRequest,
		message: messageIn,
		cause:   causeIn,
	}
}

// atributo no valido status 404
func (e *apierr) NewNotFound(massageIn string, causeIn map[string]string) *apierr {
	return &apierr{
		code:    http.StatusNotFound,
		message: massageIn,
		cause:   causeIn,
	}
}
