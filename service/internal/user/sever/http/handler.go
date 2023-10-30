package http

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
	"service/internal/user/repository"
	"service/internal/user/user"
)

type EndpointHandler struct {
	logger      *zap.SugaredLogger
	userService user.UseCase
}

func NewEndpointHandler(
	logger *zap.SugaredLogger,
	userService user.UseCase,
) *EndpointHandler {
	return &EndpointHandler{
		logger:      logger,
		userService: userService,
	}
}

func (h *EndpointHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	login := chi.URLParam(r, "login")

	userEntity, err := h.userService.GetUserByLogin(r.Context(), login)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			h.logger.Warnf("user with login: %s not found", login)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := struct {
		Id          int    `json:"id"`
		Login       string `json:"login"`
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		IsConfirmed bool   `json:"is_confirmed"`
		Password    string `json:"password"`
	}{
		Id:          userEntity.Id,
		Login:       userEntity.Login,
		FirstName:   userEntity.FirstName,
		LastName:    userEntity.LastName,
		IsConfirmed: userEntity.IsConfirmed,
		Password:    userEntity.Password,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
