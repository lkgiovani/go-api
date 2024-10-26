package projectError

import (
	"errors"
	"fmt"
	"runtime"
	"time"
)

const (
	ECONFLICT       = "conflict"
	EINTERNAL       = "internal"
	EINVALID        = "invalid"
	ENOTFOUND       = "not_found"
	ENOTIMPLEMENTED = "not_implemented"
	EUNAUTHORIZED   = "unauthorized"

	// Types of error levels
	LogOnly         = 1
	CloseConnection = 2
	ExecuteFunction = 3
)

type Error struct {
	Code      string
	Message   string
	Path      string
	PrevError error
}

func (e *Error) Error() string {
	timeNow := time.Now()

	date := fmt.Sprintf("%02d/%02d/%d %02d:%02d", timeNow.Day(), timeNow.Month(), timeNow.Year(), timeNow.Hour(), timeNow.Minute())

	errorMsg := fmt.Sprintf(
		"\u001B[31m"+
			"\n===============================ERROR=================================\n"+
			"Message: %s\n"+
			"Code: %s\n"+
			"Time: %s\n"+
			"Path: %s"+
			"\n====================================================================="+
			"\u001B[0m", // Formato com códigos de cor
		e.Message, e.Code, date, e.Path,
	)

	if e.PrevError != nil {
		errorMsg += fmt.Sprintf("\nCaused by: %v", e.PrevError)
	}

	return errorMsg
}

func ErrorCode(err error) string {
	var e *Error
	if err == nil {
		return ""
	}

	// Verifica se o erro é do tipo *Error
	if errors.As(err, &e) {
		return e.Code
	}

	return EINTERNAL
}

func ErrorMessage(err error) string {
	var e *Error
	if err == nil {
		return ""
	}

	// Verifica se o erro é do tipo *Error
	if errors.As(err, &e) {
		return e.Message
	}

	return "Internal error"
}

func Errorf(code string, levelError int, format string, args ...interface{}) error {
	// Obtém informações sobre o local do erro
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		file = "unknown"

	}

	// Cria o erro com o horário atual, caminho do arquivo e linha de código
	return &Error{
		Code:      code,
		Message:   fmt.Sprintf(format, args...),
		Path:      file,
		PrevError: nil,
	}
}
