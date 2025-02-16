package store

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"errors"
	"github.com/lucaiatropulus/social/internal/utils"
	"time"

	"github.com/lucaiatropulus/social/internal/dao"
)

type UserStore struct {
	db *sql.DB
}

func (s *UserStore) Create(ctx context.Context, transaction *sql.Tx, user *dao.User) error {
	query := `
	INSERT INTO users (username, password, email, role_id)
	VALUES($1, $2, $3, $4)
	RETURNING id, created_at
	`
	err := transaction.QueryRowContext(
		ctx,
		query,
		user.Username,
		user.Password.Hash,
		user.Email,
		user.RoleID,
	).Scan(
		&user.ID,
		&user.CreatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *UserStore) CreateAndInvite(ctx context.Context, user *dao.User, token string, expiration string) error {
	return withTransaction(s.db, ctx, func(tx *sql.Tx) error {
		if err := s.Create(ctx, tx, user); err != nil {
			return err
		}

		expirationDuration, ok := utils.ParseStringToDuration(expiration)

		if !ok {
			return errors.New("Invalid invitation expiration")
		}

		if err := s.createUserInvitation(ctx, tx, user.ID, token, expirationDuration); err != nil {
			return err
		}

		return nil
	})
}

func (s *UserStore) Activate(ctx context.Context, token string) error {
	return withTransaction(s.db, ctx, func(tx *sql.Tx) error {
		user, err := s.getUserFromInvitation(ctx, tx, token)

		if err != nil {
			return err
		}

		user.IsActive = true

		if err := s.updateUser(ctx, tx, user); err != nil {
			return err
		}

		if err := s.deleteUserInvitations(ctx, tx, user.ID); err != nil {
			return err
		}

		return nil
	})
}

func (s *UserStore) Update(ctx context.Context, user *dao.User) error {
	query := `
	UPDATE users
	SET username = $1
	WHERE id = $2
	`
	_, err := s.db.ExecContext(ctx, query, user.Username, user.ID)

	if err != nil {
		return err
	}

	return nil
}

func (s *UserStore) Delete(ctx context.Context, userID int64) error {
	return withTransaction(s.db, ctx, func(tx *sql.Tx) error {
		if err := s.deleteUser(ctx, tx, userID); err != nil {
			return err
		}

		if err := s.deleteUserInvitations(ctx, tx, userID); err != nil {
			return err
		}

		return nil
	})
}

func (s *UserStore) GetByID(ctx context.Context, userID int64) (*dao.User, error) {
	query := `
	SELECT id, username, email, password, created_at, role_id
	FROM users
	WHERE id = $1
	`

	var user dao.User

	err := s.db.QueryRowContext(ctx, query, userID).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password.Hash,
		&user.CreatedAt,
		&user.RoleID,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}

func (s *UserStore) GetByEmail(ctx context.Context, email string) (*dao.User, error) {
	query := `
	SELECT id, username, email, password, created_at, role_id
	FROM users
	WHERE email = $1 AND is_active = true
	`

	var user dao.User

	err := s.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password.Hash,
		&user.CreatedAt,
		&user.RoleID,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}

func (s *UserStore) createUserInvitation(ctx context.Context, transaction *sql.Tx, userID int64, token string, expiration time.Duration) error {
	query := `
		INSERT INTO user_invitations (token, user_id, expiration_date)
		VALUES ($1, $2, $3)
	`

	expirationDate := time.Now().Add(expiration)

	_, err := transaction.ExecContext(ctx, query, token, userID, expirationDate)

	if err != nil {
		return err
	}

	return nil
}

func (s *UserStore) getUserFromInvitation(ctx context.Context, transaction *sql.Tx, token string) (*dao.User, error) {
	query := `
	SELECT u.id, u.username, u.email, u.created_at, u.is_active, u.role_id
	FROM users u
	JOIN user_invitations ui ON u.id = ui.user_id
	WHERE ui.token = $1 AND ui.expiration_date > $2
	`

	hash := sha256.Sum256([]byte(token))
	hashedToken := hex.EncodeToString(hash[:])

	user := &dao.User{}

	err := transaction.QueryRowContext(ctx, query, hashedToken, time.Now()).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.CreatedAt,
		&user.IsActive,
		&user.RoleID,
	)

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return user, nil
}

func (s *UserStore) updateUser(ctx context.Context, transaction *sql.Tx, user *dao.User) error {
	query := `
	UPDATE users
	SET username = $1, email = $2, is_active = $3
	WHERE id = $4
	`

	_, err := transaction.ExecContext(ctx, query, user.Username, user.Email, user.IsActive, user.ID)

	if err != nil {
		return err
	}

	return nil
}

func (s *UserStore) deleteUser(ctx context.Context, transaction *sql.Tx, userID int64) error {
	query := `DELETE FROM users WHERE id = $1`

	_, err := transaction.ExecContext(ctx, query, userID)

	if err != nil {
		return err
	}

	return nil
}

func (s *UserStore) deleteUserInvitations(ctx context.Context, transaction *sql.Tx, userID int64) error {
	query := `DELETE FROM user_invitations WHERE user_id = $1`

	_, err := transaction.ExecContext(ctx, query, userID)

	if err != nil {
		return err
	}

	return nil
}
