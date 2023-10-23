package config

type Config struct {
	Database Database `yaml:"Database"`
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
