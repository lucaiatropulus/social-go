package store

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

type FollowerStore struct {
	db *sql.DB
}

func (s *FollowerStore) Follow(ctx context.Context, followedUserID int64, userID int64) error {
	query := `
	INSERT INTO followers (user_id, follower_id)
	VALUES ($1, $2)
	`

	_, err := s.db.ExecContext(ctx, query, userID, followedUserID)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return ErrFollowConflict
		}

		return err
	}

	return nil
}

func (s *FollowerStore) Unfollow(ctx context.Context, followedUserID int64, userID int64) error {
	query := `
	DELETE FROM followers
	WHERE user_id = $1 AND follower_id = $2
	`

	_, err := s.db.ExecContext(ctx, query, userID, followedUserID)

	return err
}
