package models

import (
	"strings"

	"github.com/lucaiatropulus/social/internal/dao"
)

type CreatePostRequest struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

func CreatePostFromPayload(payload CreatePostRequest, userID int64) *dao.Post {
	return &dao.Post{
		Content: payload.Content,
		Title:   payload.Title,
		UserID:  userID,
		Tags:    payload.Tags,
	}
}

func (p *CreatePostRequest) IsValid() (string, bool) {
	isTitleValid := strings.TrimSpace(p.Title) != ""
	isContentValid := strings.TrimSpace(p.Content) != ""

	if !isTitleValid {
		return "Titlul postarii nu poate fi gol", false
	}

	if !isContentValid {
		return "Continutul postarii nu poate fi gol", false
	}

	return "", true
}
