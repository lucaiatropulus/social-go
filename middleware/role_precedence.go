package middleware

import (
	"context"
	"github.com/lucaiatropulus/social/cmd/main/application"
	"github.com/lucaiatropulus/social/internal/dao"
)

func checkRolePrecedence(ctx context.Context, user *dao.User, requiredRoleName string, app *application.Application) (bool, error) {
	requiredRole, err := app.Store.Roles.GetRoleByName(ctx, requiredRoleName)

	if err != nil {
		return false, err
	}

	userRole, err := app.Store.Roles.GetRoleByID(ctx, user.RoleID)

	if err != nil {
		return false, err
	}

	return userRole.Level >= requiredRole.Level, nil
}
