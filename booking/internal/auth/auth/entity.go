package auth

type GenerateTokenRequest struct {
	Login    string
	Password string
}

type JwtUserToken struct {
	Token        string
	RefreshToken string
}

type JwtTokenContent struct {
	UserId int
	Name   string
}
