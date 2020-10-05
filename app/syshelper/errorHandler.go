package syshelper

import (
	"errors"
	"fmt"
)

// ErrorMessage struct
type ErrorMessage struct {
	StringError []error
}

// SystemErrorHandler function
func (em *ErrorMessage) SystemErrorHandler() {
	message := recover()
	if message == nil {
		em.StringError = nil
	}
	em.StringError = append(em.StringError, errors.New("Something went wrong! "+fmt.Sprint(message)))
}

// CustomErrorHandler function
func (em *ErrorMessage) CustomErrorHandler(sError error) {
	em.StringError = append(em.StringError, errors.New("Something went wrong! "+fmt.Sprint(sError)))
}
