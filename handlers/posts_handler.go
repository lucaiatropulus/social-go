package handlers

import (
	"github.com/lucaiatropulus/social/services"
	"net/http"
)

type PostHandler struct {
	service *services.PostService
}

// CreatePostHandler CreatePost godoc
//
//	@Summary		Creates a post
//	@Description	Creates a post for the signed in user
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			createPost	body		models.CreatePostRequest	true	"CreatePostRequest"
//	@Success		201			{object}	dao.Post					"Post created"
//	@Failure		400
//	@Failure		500
//	@Security		ApiKeyAuth
//	@Router			/posts [post]
func (handler *PostHandler) CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	handler.service.CreatePost(w, r)
}

// GetPostHandler GetPost godoc
//
//	@Summary		Gets a post
//	@Description	Gets a post by ID
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			postID	path		int	true	"PostID"
//	@Success		200		{object}	models.DisplayPost
//	@Failure		400
//	@Failure		500
//	@Security		ApiKeyAuth
//	@Router			/posts/{postID} [get]
func (handler *PostHandler) GetPostHandler(w http.ResponseWriter, r *http.Request) {
	handler.service.GetPost(w, r)
}

// UpdatePostHandler UpdatePost godoc
//
//	@Summary		Updates a post
//	@Description	Updates a post by ID
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			postID		path		int					true	"PostID"
//	@Param			updatePost	body		models.UpdatePost	true	"UpdatePost"
//	@Success		200			{object}	dao.Post
//	@Failure		400
//	@Failure		500
//	@Security		ApiKeyAuth
//	@Router			/posts/{postID} [patch]
func (handler *PostHandler) UpdatePostHandler(w http.ResponseWriter, r *http.Request) {
	handler.service.UpdatePost(w, r)
}

// DeletePostHandler DeletePost godoc
//
//	@Summary		Deletes a post
//	@Description	Deletes a post by ID
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			postID	path	int	true	"PostID"
//	@Success		204
//	@Failure		400
//	@Failure		500
//	@Security		ApiKeyAuth
//	@Router			/posts/{postID} [delete]
func (handler *PostHandler) DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	handler.service.DeletePost(w, r)
}
