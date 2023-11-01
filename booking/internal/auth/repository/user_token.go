package repository

import (
	"booking/internal/auth/entity"
)

func (r *Repo) CreateUserToken(userToken entity.UserToken) error {
	_, err := r.main.Exec("UPDATE users SET token=$1, refresh_token=$2 WHERE $3", userToken.Token, userToken.RefreshToken, userToken.UserId)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) UpdateUserToken(userToken entity.UserToken) error {
	_, err := r.main.Exec("UPDATE users SET token=$1, refresh_token=$2 WHERE $3", userToken.Token, userToken.RefreshToken, userToken.UserId)
	if err != nil {
		return err
	}
	return nil
}
