package errors

import "fmt"

type FpError struct {
	ErrorCode FpErrorCode
	Status    int
	Message   string
}

func GetError(code FpErrorCode, status int, msg string) error {
	return &FpError{ErrorCode: code, Status: status, Message: msg}
}

func (e *FpError) Error() string {
	return fmt.Sprintf("message: %s \nerror: %s \nstatus: %d", e.Message, e.ErrorCode.String(), e.Status)
}
