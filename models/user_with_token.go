package models

import "github.com/lucaiatropulus/social/internal/dao"

type UserWithToken struct {
	*dao.User
	Token string `json:"token"`
}
