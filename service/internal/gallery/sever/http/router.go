package http

import (
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
	"service/internal/gallery/sever/http/middleware"
)

type router struct {
	logger         *zap.SugaredLogger
	authMiddleware *middleware.JwtV1
}

func NewRouter(logger *zap.SugaredLogger, authMiddleware *middleware.JwtV1) *router {
	return &router{
		logger:         logger,
		authMiddleware: authMiddleware,
	}
}

func (s *router) GetHandler(eh *EndpointHandler) http.Handler {
	router := chi.NewRouter().
		Group(func(r chi.Router) {
			r.Use(s.authMiddleware.AuthV1)
			r.Route("/api/gallery/v1", func(r chi.Router) {
				r.Get("/photo", eh.GetPhotos)
			})
		})

	return router
}
