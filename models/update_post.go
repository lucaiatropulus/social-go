package models

import "github.com/lucaiatropulus/social/internal/dao"

type UpdatePost struct {
	Title   *string `json:"title"`
	Content *string `json:"content"`
}

func (p *UpdatePost) IsValid() bool {
	isTitleValid := p.Title != nil
	isContentValid := p.Content != nil

	return isTitleValid || isContentValid
}

func (p *UpdatePost) SetUpdatedContent(post *dao.Post) {
	if p.Title != nil {
		post.Title = *p.Title
	}

	if p.Content != nil {
		post.Content = *p.Content
	}
}
