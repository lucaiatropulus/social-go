package models

import "github.com/lucaiatropulus/social/internal/dao"

type DisplayPost struct {
	dao.Post
	Username      string `json:"username"`
	CommentsCount int    `json:"comments_count"`
}
