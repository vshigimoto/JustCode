package main

import (
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	l := logger.Sugar()
	l = l.With(zap.String("app", "debug-server"))
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Mount("/debug", middleware.Profiler())
	if err := http.ListenAndServe(":8081", r); err != nil {
		l.Fatal("error with run server", err)
	}
}
