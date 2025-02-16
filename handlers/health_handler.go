package handlers

import (
	"github.com/lucaiatropulus/social/services"
	"net/http"
)

type HealthHandler struct {
	service *services.HealthService
}

func (h *HealthHandler) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	h.service.CheckAppHealth(w, r)
}
