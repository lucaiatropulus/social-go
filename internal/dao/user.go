package dao

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID        int64    `json:"id"`
	Username  string   `json:"username"`
	Email     string   `json:"email"`
	Password  password `json:"-"`
	CreatedAt string   `json:"created_at"`
	IsActive  bool     `json:"is_active"`
	RoleID    int64    `json:"role_id"`
}

type password struct {
	text *string
	Hash []byte
}

func (p *password) Set(text string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	p.text = &text
	p.Hash = hash

	return nil
}

func (p *password) CheckPassword(text string) bool {
	if err := bcrypt.CompareHashAndPassword(p.Hash, []byte(text)); err != nil {
		return false
	}

	return true
}
