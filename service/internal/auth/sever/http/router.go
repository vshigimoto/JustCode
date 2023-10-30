package http

import (
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
)

type router struct {
	logger *zap.SugaredLogger
}

func NewRouter(logger *zap.SugaredLogger) *router {
	return &router{
		logger: logger,
	}
}

func (s *router) GetHandler(eh *EndpointHandler) http.Handler {
	router := chi.NewRouter().
		Group(func(r chi.Router) {
			r.Route("/api/auth/v1", func(r chi.Router) {
				r.Post("/register", eh.Register)
				r.Post("/user-confirm", eh.ConfirmUser)
				r.Post("/login", eh.Login)
				r.Post("/renew-token", eh.RenewToken)
			})
		})

	return router
}
