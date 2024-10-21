package user_model

type User struct {
	Id    string
	Name  string
	Email string
}

func NewUser(id, name, email string) *User {
	return &User{
		Id:    id,
		Name:  name,
		Email: email,
	}
}
