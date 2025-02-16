package main

import (
	"database/sql"
	"github.com/go-redis/redis/v8"
	"github.com/lucaiatropulus/social/config"
	"github.com/lucaiatropulus/social/internal/auth"
	"github.com/lucaiatropulus/social/internal/db"
	"github.com/lucaiatropulus/social/internal/mailer"
	"github.com/lucaiatropulus/social/internal/store/cache"
	"go.uber.org/zap"
)

func setupDatabase(configuration *config.Config, logger *zap.SugaredLogger) *sql.DB {
	database, err := db.New(
		configuration.DB.Address,
		configuration.DB.MaxOpenConnections,
		configuration.DB.MaxIdleConnections,
		configuration.DB.MaxIdleTime,
	)

	if err != nil {
		logger.Fatal(err)
	}

	return database
}

func setupRedis(configuration *config.Config, logger *zap.SugaredLogger) *redis.Client {
	if configuration.Redis.Enabled {
		logger.Info("Redis connection established")
		return cache.NewRedisClient(configuration.Redis.Address, configuration.Redis.Password, configuration.Redis.Database)
	}

	return nil
}

func setupMailClient(configuration *config.Config, logger *zap.SugaredLogger) mailer.MailtrapClient {
	mailerClient, err := mailer.NewMailTrapClient(configuration.Mail.ApiKey, configuration.Mail.Email)

	if err != nil {
		logger.Fatal(err)
	}

	return mailerClient
}

func setupJWTAuthenticator(configuration *config.Config) *auth.JWTAuthenticator {
	return auth.NewJWTAuthenticator(configuration.Auth.Secret, configuration.Auth.Audience, configuration.Auth.Issuer)
}
