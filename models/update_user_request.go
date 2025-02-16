package models

import "github.com/lucaiatropulus/social/internal/dao"

type UpdateUserRequest struct {
	Username string `json:"username"`
}

func (r *UpdateUserRequest) UpdateUser(user *dao.User) {
	if r.Username != "" {
		user.Username = r.Username
	}
}
