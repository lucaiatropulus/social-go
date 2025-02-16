package testing

import (
	"bytes"
	"encoding/json"
	"github.com/lucaiatropulus/social/cmd/main/lifecycle"
	"github.com/lucaiatropulus/social/internal/dao"
	"github.com/lucaiatropulus/social/models"
	"github.com/lucaiatropulus/social/routing"
	"net/http"
	"testing"
)

func TestGetUser(t *testing.T) {
	app := newTestApplication(t)
	appRouting := routing.NewRouting(app)
	mux := lifecycle.Mount(appRouting)

	t.Run("should not allow unauthenticated requests", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/api/v1/users/1", nil)

		if err != nil {
			t.Fatal(err)
		}

		rr := executeRequest(req, mux, app, nil)

		if rr.Code != http.StatusUnauthorized {
			t.Errorf("expected error code to be %d and we got %d for request path %s", http.StatusUnauthorized, rr.Code, req.URL)
		}
	})

	t.Run("should return the user for authenticated requests if user exists", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/api/v1/users/1", nil)

		if err != nil {
			t.Fatal(err)
		}

		user := &dao.User{ID: 1}

		rr := executeRequest(req, mux, app, user)

		if rr.Code == http.StatusUnauthorized {
			t.Errorf("the user should be considered authenticated, but it isn't for request path %s", req.URL)
			return
		}

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code to be %d and we got %d for request path %s", http.StatusOK, rr.Code, req.URL)
		}
	})

}

func TestUpdateUser(t *testing.T) {
	app := newTestApplication(t)
	appRouting := routing.NewRouting(app)
	mux := lifecycle.Mount(appRouting)

	t.Run("should not allow unauthenticated requests", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/api/v1/users/update", nil)

		if err != nil {
			t.Fatal(err)
		}

		rr := executeRequest(req, mux, app, nil)

		if rr.Code != http.StatusUnauthorized {
			t.Errorf("expected error code to be %d and we got %d for request path %s", http.StatusUnauthorized, rr.Code, req.URL)
		}
	})

	t.Run("should allow authenticated requests", func(t *testing.T) {
		payload := &models.UpdateUserRequest{Username: "new_username"}
		body, err := json.Marshal(payload)

		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPatch, "/api/v1/users/update", bytes.NewBuffer(body))

		if err != nil {
			t.Fatal(err)
		}

		user := &dao.User{ID: 1}

		rr := executeRequest(req, mux, app, user)

		if rr.Code == http.StatusUnauthorized {
			t.Errorf("the user should be considered authenticated, but it isn't for request path %s", req.URL)
			return
		}

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code to be %d and we got %d for request path %s", http.StatusOK, rr.Code, req.URL)
		}
	})

}
