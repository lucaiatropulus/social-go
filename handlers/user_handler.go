package handlers

import (
	"github.com/lucaiatropulus/social/services"
	"net/http"
)

type UserHandler struct {
	service *services.UserService
}

// GetUserHandler GetUser godoc
//
//	@Summary		Fetches a user profile
//	@Description	Fetches a user profile by ID
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"UserID"
//	@Success		200	{object}	dao.User
//	@Failure		400
//	@Failure		404
//	@Failure		500
//	@Security		ApiKeyAuth
//	@Router			/users/{id} [get]
func (handler *UserHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	handler.service.GetUser(w, r)
}

// UpdateUserHandler GetUser godoc
//
//	@Summary		Updates a user profile
//	@Description	Updates a user profile
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			updateUserRequest	body		models.UpdateUserRequest	true	"The update user profile form data"
//	@Success		200					{object}	dao.User
//	@Failure		400
//	@Failure		404
//	@Failure		500
//	@Security		ApiKeyAuth
//	@Router			/users/update [patch]
func (handler *UserHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	handler.service.UpdateUser(w, r)
}

// FollowUserHandler FollowUser godoc
//
//	@Summary		Follows a user
//	@Description	Follows a user by ID
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			userID	path	int	true	"UserID"
//	@Success		204		"User followed"
//	@Failure		400		"User not found"
//	@Security		ApiKeyAuth
//	@Router			/users/{userID}/follow [put]
func (handler *UserHandler) FollowUserHandler(w http.ResponseWriter, r *http.Request) {
	handler.service.FollowUser(w, r)
}

// UnfollowUserHandler UnfollowUser godoc
//
//	@Summary		Unfollows a user
//	@Description	Unfollows a user by ID
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			userID	path	int	true	"UserID"
//	@Success		204		"User unfollowed"
//	@Failure		400		"User not found"
//	@Security		ApiKeyAuth
//	@Router			/users/{userID}/unfollow [put]
func (handler *UserHandler) UnfollowUserHandler(w http.ResponseWriter, r *http.Request) {
	handler.service.UnfollowUser(w, r)
}
