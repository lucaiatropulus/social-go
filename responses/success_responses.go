package responses

import (
	"github.com/lucaiatropulus/social/internal/utils"
	"net/http"
)

func JSONResponse(w http.ResponseWriter, status int, data any) error {
	type envelope struct {
		Data any `json:"data"`
	}

	return utils.WriteJSON(w, status, &envelope{Data: data})
}

func EmptyJSONResponse(w http.ResponseWriter, status int) {
	w.Header().Set("Content-Type", "main/json")
	w.WriteHeader(status)
}
