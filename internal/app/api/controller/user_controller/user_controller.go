package user_controller

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gofrs/uuid"
	"go-api/internal/app/api/model/user_model"
	"go-api/internal/app/repository/user_repo"
	"go-api/pkg/projectError"
	"io"
	"net/http"
)

type userController struct {
	db *sql.DB
}

func NewUserController(db *sql.DB) UserControllerInterface {
	return &userController{
		db: db,
	}
}

func (uc *userController) SetUser(w http.ResponseWriter, r *http.Request) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)

	}

	var user PostSetUserRequest
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)

	}

	// Gerando o UUID para o usuário
	id, err := uuid.NewV7()
	if err != nil {
		http.Error(w, "Failed to generate UUID", http.StatusInternalServerError)

	}

	userDB := user_repo.NewUserRepository(uc.db)

	err = userDB.SetUser(context.Background(), user_model.User{Id: id.String(), Name: user.Name, Email: user.Email})
	if err != nil {
		return &projectError.Error{
			Code:      projectError.EINTERNAL,
			Message:   "failed to set user in database",
			PrevError: err,
		}
	}

	response := map[string]string{"message": "User created successfully!"}
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

	return nil
}

func (uc *userController) GetAllUser(w http.ResponseWriter, r *http.Request) error {
	userDB := user_repo.NewUserRepository(uc.db)

	users, err := userDB.GetAllUser(context.Background())
	if err != nil {
		return &projectError.Error{
			Code:      projectError.EINTERNAL,
			Message:   "failed to set user in database",
			PrevError: err,
		}
	}

	fmt.Println("salve")

	response := map[string][]user_model.User{"users": users}
	jsonResponse, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

	return nil
}

func (uc *userController) GetUserById(w http.ResponseWriter, r *http.Request) error {

	id := r.URL.Query().Get("id")

	if id == "" {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, `{"error": "Missing id"}`, http.StatusBadRequest)
		return &projectError.Error{
			Code:    projectError.EINTERNAL,
			Message: "Missing id",
		}

	}

	userDB := user_repo.NewUserRepository(uc.db)

	users, err := userDB.GetUserById(id)
	if err != nil {
		return &projectError.Error{
			Code:      projectError.EINTERNAL,
			Message:   "failed to set user in database",
			PrevError: err,
		}
	}

	fmt.Printf("Users: %+v\n", users)

	response := user_model.User{Id: id, Name: users.Name, Email: users.Email}
	jsonResponse, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

	return nil

}

func (uc *userController) DeleteUserById(w http.ResponseWriter, r *http.Request) error {

	id := r.URL.Query().Get("id")

	if id == "" {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, `{"error": "Missing id"}`, http.StatusBadRequest)
		return &projectError.Error{
			Code:    projectError.EINTERNAL,
			Message: "Missing id",
		}

	}

	userDB := user_repo.NewUserRepository(uc.db)

	err := userDB.DeleteUserById(context.Background(), id)
	if err != nil {
		return &projectError.Error{
			Code:      projectError.EINTERNAL,
			Message:   "failed to set user in database",
			PrevError: err,
		}
	}

	response := map[string]string{"message": "User deleted successfully!"}
	jsonResponse, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

	return nil
}

func (uc *userController) UpdateUserById(w http.ResponseWriter, r *http.Request) error {
	//body, err := io.ReadAll(r.Body)
	//if err != nil {
	//	http.Error(w, "Failed to read request body", http.StatusInternalServerError)
	//
	//}
	//
	//var user PostUpDateUserRequest
	//err = json.Unmarshal(body, &user)
	//if err != nil {
	//	http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
	//
	//}
	//
	//userDB := user_repo.NewUserRepository(uc.db)

	//err = userDB.SetUser(context.Background())
	//if err != nil {
	//	return &projectError.Error{
	//		Code:      projectError.EINTERNAL,
	//		Message:   "failed to set user in database",
	//		PrevError: err,
	//	}
	//}

	response := map[string]string{"message": "User created successfully!"}
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

	return nil
}
