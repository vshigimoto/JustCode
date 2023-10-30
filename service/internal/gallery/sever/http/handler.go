package http

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"service/internal/gallery/sever/http/middleware"
)

type EndpointHandler struct {
	logger *zap.SugaredLogger
}

func NewEndpointHandler(
	logger *zap.SugaredLogger,
) *EndpointHandler {
	return &EndpointHandler{
		logger: logger,
	}
}

func (h *EndpointHandler) GetPhotos(w http.ResponseWriter, r *http.Request) {
	user, err := middleware.GetContextUser(r.Context())
	if err != nil {
		h.logger.Errorf("cannot find user in context")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	response := struct {
		UserId int64 `json:"user_id"`
	}{
		UserId: user.ID,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
