package cache

import (
	"context"

	"github.com/lucaiatropulus/social/internal/dao"
)

func NewMockStore() Storage {
	return Storage{
		Users: &MockUserStore{},
	}
}

type MockUserStore struct{}

func (m *MockUserStore) Get(ctx context.Context, userID int64) (*dao.User, error) {
	return &dao.User{}, nil
}

func (m *MockUserStore) Set(ctx context.Context, user *dao.User) error {
	return nil
}

func (m *MockUserStore) Delete(ctx context.Context, userID int64) error {
	return nil
}
