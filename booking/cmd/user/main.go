package main

import (
	"booking/internal/user/config"
	"booking/internal/user/database"
	"booking/internal/user/repository"
	"booking/internal/user/server/http"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" //
	"github.com/spf13/viper"
)

func main() {
	r := gin.Default()
	cfg, err := loadConfig("config/user")
	if err != nil {
		return
	}
	mainDB, err := database.New(cfg.Database.Main)
	replicaDB, err := database.New(cfg.Database.Replica)
	rep := repository.NewRepository(mainDB, replicaDB)
	http.InitRouter(rep, r)
	err = r.Run(":8080")
	if err != nil {
		return
	}
}

func loadConfig(path string) (config config.Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return config, fmt.Errorf("failed to ReadInConfig err: %w", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return config, fmt.Errorf("failed to Unmarshal config err: %w", err)
	}

	return config, nil
}
