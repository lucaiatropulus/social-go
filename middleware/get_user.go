package middleware

import (
	"context"
	"github.com/lucaiatropulus/social/cmd/main/application"
	"github.com/lucaiatropulus/social/internal/dao"
)

func GetUser(ctx context.Context, userID int64, app *application.Application) (*dao.User, error) {
	if !app.Config.Redis.Enabled {
		return app.Store.Users.GetByID(ctx, userID)
	}

	user, err := app.Cache.Users.Get(ctx, userID)

	if err != nil {
		return nil, err
	}

	if user == nil {
		user, err = app.Store.Users.GetByID(ctx, userID)

		if err != nil {
			return nil, err
		}

		if err = app.Cache.Users.Set(ctx, user); err != nil {
			return nil, err
		}
	}

	return user, nil
}
