package dao

type Post struct {
	ID        int64    `json:"id"`
	Content   string   `json:"content"`
	Title     string   `json:"title"`
	UserID    int64    `json:"user_id"`
	Tags      []string `json:"tags"`
	CreatedAt string   `json:"created_at"`
	Version   int      `json:"version"`
	UpdatedAt string   `json:"updated_at"`
}

func CreatePost(content string, title string, userID int64, tags []string) *Post {
	return &Post{
		Content: content,
		Title:   title,
		UserID:  userID,
		Tags:    tags,
	}
}
