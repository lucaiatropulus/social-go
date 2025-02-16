package store

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
	"github.com/lucaiatropulus/social/models"
)

type FeedStore struct {
	db *sql.DB
}

func (s *FeedStore) GetUserFeed(ctx context.Context, userID int64, paginatedQuery PaginatedFeedQuery) ([]models.DisplayPost, error) {
	println(len(paginatedQuery.Tags))
	query := `
	SELECT p.id, p.user_id, u.username, p.title, p.content, p.created_at, p.version, p.tags, COUNT(c.id) AS comments_count
	FROM posts p
	LEFT JOIN comments c ON c.post_id = p.id
	LEFT JOIN users u ON u.id = p.user_id
	JOIN followers f ON f.follower_id = p.user_id OR p.user_id = $1
	WHERE f.user_id = $1 AND
		(p.title ILIKE '%' || $4 || '%' OR p.content ILIKE '%' || $4 || '%') AND
		(p.tags @> $5 OR array_upper($5, 1) is null)
	GROUP BY p.id, u.username
	ORDER BY p.created_at ` + paginatedQuery.Sort + `
	LIMIT $2 OFFSET $3
	`

	rows, err := s.db.QueryContext(
		ctx,
		query,
		userID,
		paginatedQuery.Limit,
		paginatedQuery.Offset,
		paginatedQuery.Search,
		pq.Array(paginatedQuery.Tags),
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	feed := []models.DisplayPost{}

	for rows.Next() {
		var post models.DisplayPost

		err := rows.Scan(
			&post.Post.ID,
			&post.Post.UserID,
			&post.Username,
			&post.Post.Title,
			&post.Post.Content,
			&post.Post.CreatedAt,
			&post.Post.Version,
			pq.Array(&post.Post.Tags),
			&post.CommentsCount,
		)

		if err != nil {
			return nil, err
		}

		feed = append(feed, post)
	}

	return feed, nil
}
