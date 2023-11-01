package entity

import "time"

type UserToken struct {
	Id           string    `db:"id"`
	Token        string    `db:"token"`
	RefreshToken string    `db:"refresh_token"`
	UserId       string    `db:"user_id"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

type MyJWTClaims struct {
	Id       string `yaml:"Id"`
	Username string `yaml:"Username"`
}
