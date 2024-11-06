package user_model

func (user *User) SetId(id string) {
	user.Id = id
}

func (user *User) SetName(name string) {
	user.Name = name
}

func (user *User) SetEmail(email string) {
	user.Email = email
}
