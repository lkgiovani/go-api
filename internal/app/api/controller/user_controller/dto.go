package user_controller

type PostSetUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type PostUpDateUserRequest struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
