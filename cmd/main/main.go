package main

import (
	"flag"
	"github.com/lucaiatropulus/social/cmd/main/application"
	"github.com/lucaiatropulus/social/cmd/main/lifecycle"
	"github.com/lucaiatropulus/social/config"
	"github.com/lucaiatropulus/social/internal/utils"
	"github.com/lucaiatropulus/social/routing"

	"github.com/lucaiatropulus/social/internal/rate_limiter"
	"github.com/lucaiatropulus/social/internal/store"
	"github.com/lucaiatropulus/social/internal/store/cache"
)

//	@title			Iatropulus Social
//	@description	Aici va dau clasa fraierilor
//	@termsOfService	https://google.com

//	@contact.name	Luca
//	@contact.url	https://google.com
//	@contact.email	iatropulus.luca@gmail.com

//	@license.name	Apache 2.0
//	@license.url	https://google.com

// @securityDefinitions.apiKey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description
func main() {
	environment := flag.String("env", "development", "Defines the environment")
	configuration := config.LoadConfig(environment)

	logger := utils.GetLogger()
	defer logger.Sync()

	database := setupDatabase(configuration, logger)
	defer database.Close()

	logger.Info("Database connection established")

	redisDB := setupRedis(configuration, logger)
	dbStore := store.NewStorage(database)
	cacheStore := cache.NewRedisStorage(redisDB)

	mailClient := setupMailClient(configuration, logger)

	authenticator := setupJWTAuthenticator(configuration)

	rateLimiter := ratelimiter.NewFixedWindowRateLimiter(
		configuration.RateLimiter.RequestsPerTimeFrame,
		configuration.RateLimiter.TimeFrame,
	)

	app := application.NewApplication(
		*configuration,
		dbStore,
		cacheStore,
		logger,
		mailClient,
		authenticator,
		rateLimiter,
	)

	appRouting := routing.NewRouting(app)

	mux := lifecycle.Mount(appRouting)

	logger.Fatal(lifecycle.Run(mux, app))
}
