package user_controller

import (
	"go-api/pkg/projectError"
	"net/http"
)

func jsonError(w http.ResponseWriter, statusCode int, errMsg string) {
	w.Header().Set("Content-Type", "application/json")
	http.Error(w, `{"error": "`+errMsg+`"}`, statusCode)
}

func newProjectError(code string, msg string, prevError error, path string) *projectError.Error {
	return &projectError.Error{
		Code:      code,
		Message:   msg,
		PrevError: prevError,
		Path:      path,
	}
}
