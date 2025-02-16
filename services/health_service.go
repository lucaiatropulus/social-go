package services

import (
	"github.com/lucaiatropulus/social/cmd/main/application"
	"github.com/lucaiatropulus/social/responses"
	"net/http"
)

type HealthService struct {
	app *application.Application
}

func NewHealthService(app *application.Application) *HealthService {
	return &HealthService{app}
}

func (s *HealthService) CheckAppHealth(w http.ResponseWriter, r *http.Request) {
	if err := responses.JSONResponse(w, http.StatusOK, "Ok"); err != nil {
		responses.InternalServerError(w, r, err, s.app.Logger)
		return
	}
}
