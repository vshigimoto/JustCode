package usecase

import "booking/internal/user/entity"

func UserBuilder(id int, name string, email string, login string, password string) *entity.User {
	var user entity.User
	if id != 0 {
		user.Id = id
	}
	if name != "" {
		user.Name = name
	}
	if email != "" {
		user.Email = email
	}
	if login != "" {
		user.Login = login
	}
	if password != "" {
		user.Password = password
	}
	return &user
}
