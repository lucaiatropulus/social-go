package store

import (
	"context"
	"database/sql"
	"errors"
	"github.com/lucaiatropulus/social/internal/dao"
	"github.com/lucaiatropulus/social/models"
)

var (
	ErrNotFound       = errors.New("Resource not found")
	ErrFollowConflict = errors.New("You cannot follow the same user more than once")
)

type Storage struct {
	Posts interface {
		Create(ctx context.Context, post *dao.Post) error
		GetByID(ctx context.Context, postID int64) (*dao.Post, error)
		Update(ctx context.Context, post *dao.Post) error
		Delete(ctx context.Context, postID int64) error
	}

	Comments interface {
		Create(ctx context.Context, comment *dao.Comment) error
		GetByPostID(ctx context.Context, postID int64) ([]models.DisplayComment, error)
		GetCountByID(ctx context.Context, postID int64) (int, error)
	}

	Users interface {
		Create(ctx context.Context, transaction *sql.Tx, user *dao.User) error
		GetByID(ctx context.Context, userID int64) (*dao.User, error)
		GetByEmail(ctx context.Context, email string) (*dao.User, error)
		CreateAndInvite(ctx context.Context, user *dao.User, token string, expiration string) error
		Activate(ctx context.Context, token string) error
		Update(ctx context.Context, user *dao.User) error
		Delete(ctx context.Context, userID int64) error
	}

	Roles interface {
		GetRoleByID(ctx context.Context, roleID int64) (*dao.Role, error)
		GetRoleByName(ctx context.Context, roleName string) (*dao.Role, error)
	}

	Followers interface {
		Follow(ctx context.Context, followedUserID int64, userID int64) error
		Unfollow(ctx context.Context, followedUserID int64, userID int64) error
	}

	Feed interface {
		GetUserFeed(ctx context.Context, userID int64, paginatedQuery PaginatedFeedQuery) ([]models.DisplayPost, error)
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts:     &PostStore{db},
		Comments:  &CommentStore{db},
		Users:     &UserStore{db},
		Roles:     &RolesStore{db},
		Followers: &FollowerStore{db},
		Feed:      &FeedStore{db},
	}
}

func withTransaction(db *sql.DB, ctx context.Context, fn func(*sql.Tx) error) error {
	transaction, err := db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	if err := fn(transaction); err != nil {
		_ = transaction.Rollback()
		return err
	}

	return transaction.Commit()
}
