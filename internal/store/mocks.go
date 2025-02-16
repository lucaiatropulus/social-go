package store

import (
	"context"
	"database/sql"
	"github.com/lucaiatropulus/social/internal/dao"
)

func NewMockStore() Storage {
	return Storage{
		Users: &MockUserStore{},
	}
}

type MockUserStore struct{}

func (m *MockUserStore) Create(ctx context.Context, tx *sql.Tx, user *dao.User) error {
	return nil
}

func (m *MockUserStore) GetByID(ctx context.Context, userID int64) (*dao.User, error) {
	return &dao.User{}, nil
}

func (m *MockUserStore) GetByEmail(ctx context.Context, email string) (*dao.User, error) {
	return &dao.User{}, nil
}

func (m *MockUserStore) CreateAndInvite(ctx context.Context, user *dao.User, token string, expiration string) error {
	return nil
}

func (m *MockUserStore) Activate(ctx context.Context, token string) error {
	return nil
}

func (m *MockUserStore) Update(ctx context.Context, user *dao.User) error {
	return nil
}

func (m *MockUserStore) Delete(ctx context.Context, userID int64) error {
	return nil
}
