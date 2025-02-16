package handlers

import (
	"github.com/lucaiatropulus/social/services"
	"net/http"
)

type AuthHandler struct {
	service *services.AuthService
}

// RegisterUserHandler Register godoc
//
//	@Summary		Registers a user
//	@Description	Registers a user
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		models.RegisterUserRequest	true	"User registration form data"
//	@Success		201		{object}	models.UserWithToken		"User registered"
//	@Failure		400
//	@Failure		500
//	@Security		ApiKeyAuth
//	@Router			/auth/register [post]
func (handler *AuthHandler) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	handler.service.RegisterUser(w, r)
}

// LoginUserHandler Login godoc
//
//	@Summary		Login a user
//	@Description	Login a user
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		models.LoginRequest	true	"User login form data"
//	@Success		200		{string}	string				"Token"
//	@Failure		400
//	@Failure		401
//	@Failure		500
//	@Security		ApiKeyAuth
//	@Router			/auth/login [post]
func (handler *AuthHandler) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	handler.service.LoginUser(w, r)
}

// ActivateAccountHandler ActivateAccount godoc
//
//	@Summary		Activates a user account
//	@Description	Activates a user account
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			token	path	string	true	"The activation token received via email by the user"
//	@Success		204
//	@Failure		400
//	@Failure		500
//	@Security		ApiKeyAuth
//	@Router			/auth/activate/{token} [put]
func (handler *AuthHandler) ActivateAccountHandler(w http.ResponseWriter, r *http.Request) {
	handler.service.ActivateAccount(w, r)
}
