package store

import (
	"context"
	"database/sql"

	"github.com/lucaiatropulus/social/internal/dao"
)

type RolesStore struct {
	db *sql.DB
}

func (s *RolesStore) GetRoleByID(ctx context.Context, roleID int64) (*dao.Role, error) {
	query := `
	SELECT id, name, level, description FROM roles
	WHERE id = $1
	`

	role := &dao.Role{}

	err := s.db.QueryRowContext(ctx, query, roleID).Scan(
		&role.ID,
		&role.Name,
		&role.Level,
		&role.Description,
	)

	if err != nil {
		return nil, err
	}

	return role, nil
}

func (s *RolesStore) GetRoleByName(ctx context.Context, roleName string) (*dao.Role, error) {
	query := `
	SELECT id, name, level, description FROM roles
	WHERE name = $1
	`

	role := &dao.Role{}

	err := s.db.QueryRowContext(ctx, query, roleName).Scan(
		&role.ID,
		&role.Name,
		&role.Level,
		&role.Description,
	)

	if err != nil {
		return nil, err
	}

	return role, nil
}
