package handlers

import (
	"github.com/lucaiatropulus/social/services"
	"net/http"
)

type FeedHandler struct {
	service *services.FeedService
}

// GetUserFeed godoc
//
//	@Summary		Gets the feed of the logged-in user
//	@Description	Gets the feed of the logged-in user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			limit	path	int		false	"How many items should be returned per page, at least 1, at most 20. Default is 20"
//	@Param			offset	path	int		false	"How many items have already been loaded, min 0. Default is 0"
//	@Param			sort	path	string	false	"Describes how the list should be sorted, ASC or DESC. Default is DESC"
//	@Param			search	path	string	false	"Contains the search term. It searches both by title and content of the post"
//	@Param			tags	path	string	false	"Contains the tags that you want to search by, separated by commas"
//	@Param			since	path	string	false	"Filters the posts by the createdAt date and returns only the posts that were created after the given date"
//	@Param			until	path	string	false	"Filters the posts by the createdAt date and returns only the posts that were created until the given date"
//	@Success		200		{array}	models.DisplayPost
//	@Failure		400
//	@Failure		500
//	@Security		ApiKeyAuth
//	@Router			/users/feed [get]
func (handler *FeedHandler) GetUserFeed(w http.ResponseWriter, r *http.Request) {
	handler.service.GetUserFeed(w, r)
}
