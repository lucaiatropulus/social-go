package store

import (
	"context"
	"database/sql"

	"github.com/lucaiatropulus/social/internal/dao"
	models "github.com/lucaiatropulus/social/models"
)

type CommentStore struct {
	db *sql.DB
}

func (s *CommentStore) Create(ctx context.Context, comment *dao.Comment) error {
	query := `
	INSERT INTO comments (post_id, user_id, content)
	VALUES ($1, $2, $3)
	RETURNING id, created_at
	`

	err := s.db.QueryRowContext(
		ctx,
		query,
		comment.PostID,
		comment.UserID,
		comment.Content,
	).Scan(
		&comment.ID,
		&comment.CreatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *CommentStore) GetByPostID(ctx context.Context, postID int64) ([]models.DisplayComment, error) {
	query := `
	SELECT c.id, c.user_id, c.content, c.created_at, u.username FROM comments c
	JOIN users u on u.id = c.user_id
	WHERE c.post_id = $1
	ORDER BY c.created_at DESC
	`

	rows, err := s.db.QueryContext(ctx, query, postID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	comments := []models.DisplayComment{}

	for rows.Next() {
		var c models.DisplayComment
		err := rows.Scan(
			&c.ID,
			&c.UserID,
			&c.Content,
			&c.CreatedAt,
			&c.Username,
		)

		if err != nil {
			return nil, err
		}

		comments = append(comments, c)
	}

	return comments, nil
}

func (s *CommentStore) GetCountByID(ctx context.Context, postID int64) (int, error) {
	query := `
	SELECT count(*) FROM comments
	WHERE post_id = $1
	`

	var count int

	err := s.db.QueryRowContext(ctx, query, postID).Scan(&count)

	if err != nil {
		return 0, err
	}

	return count, nil
}
