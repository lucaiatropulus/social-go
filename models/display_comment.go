package models

type DisplayComment struct {
	ID        int64  `json:"id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	Username  string `json:"username"`
	UserID    int64  `json:"user_id"`
}
