package user_controller

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"go-api/pkg/projectError"
	"net/http"
)

func jsonErrorResponse(w http.ResponseWriter, statusCode int, errorMessages []map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	// Coloca as mensagens de erro dentro de um objeto com chave "error"
	response := map[string]interface{}{
		"errors": errorMessages,
	}
	json.NewEncoder(w).Encode(response)
}

func newProjectError(code string, msg string, prevError error, path string) *projectError.Error {
	return &projectError.Error{
		Code:      code,
		Message:   msg,
		PrevError: prevError,
		Path:      path,
	}
}

func formatValidationErrors(err error) []map[string]string {
	var errorList []map[string]string

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, validationErr := range validationErrors {
			field := validationErr.Field()
			errorMessage := validationErr.Error()

			// Adiciona o erro formatado no array
			errorList = append(errorList, map[string]string{field: errorMessage})
		}
	}
	return errorList
}
