export .envrc
MIGRATIONS_PATH=./cmd/migrate/migrations

.PHONY: migrate-create
migration:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@, $(MAKECMDGOALS))

.PHONY: migrate-up
migrate-up:
	echo $(DB_MIGRATOR_ADDR)
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_MIGRATOR_ADDR) up

.PHONY: migrate-down
migrate-down:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_MIGRATOR_ADDR) down $(filter-out $@, $(MAKECMDGOALS))

.PHONY: seed
seed:
	@source .envrc && go run cmd/migrate/seed/main.go

.PHONY: gendocs
.gendocs:
	@swag init -g ./api/main.go -d ./cmd/api && swag fmt

.PHONY: air-development
.air-development:
	@air && -env "development"

.PHONY: air-production
.air-production:
	@air && -env "production"