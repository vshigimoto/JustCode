package config

import "time"

type Config struct {
	Database   Database   `yaml:"Database"`
	HttpServer HttpServer `yaml:"HttpServer"`
	Transport  Transport  `yaml:"Transport"`
	Auth       Auth       `yaml:"Auth"`
}

type Database struct {
	Main    DbNone `yaml:"Main"`
	Replica DbNone `yaml:"Replica"`
}

type DbNone struct {
	Host     string `yaml:"Host"`
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
	DbName   string `yaml:"DbName"`
	SslMode  string `yaml:"SslMode"`
}

type HttpServer struct {
	Port int `yaml:"Port"`
}

type Transport struct {
	UserTransport UserTransport `yaml:"User"`
}

type UserTransport struct {
	Host    string        `yaml:"Host"`
	Timeout time.Duration `yaml:"Timeout"`
}

type Auth struct {
	PasswordSecretKey string `yaml:"passwordSecretKey"`
	JwtSecretKey      string `yaml:"jwtSecretKey"`
}
