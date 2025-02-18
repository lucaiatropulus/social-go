export .envrc
MIGRATIONS_PATH=./cmd/migrate/migrations

DOCKER_HOST="ssh://ubuntu@121.032.1.2"
#docker swarm init

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

.PHONY: swarm-init
swarm-init:
	DOCKER_HOST=${DOCKER_HOST} docker swarm init

.PHONY: swarm-deploy-stack
swarm-deploy-stack:
	DOCKER_HOST=${DOCKER_HOST} docker stack deploy -c docker-swarm.yml radcom-project

.PHONY: swarm-ls
swarm-ls:
	DOCKER_HOST=${DOCKER_HOST} docker service ls

.PHONY: swarm-remove-stack
swarm-remove-stack:
	DOCKER_HOST=${DOCKER_HOST} docker stack rm radcom-project

.PHONY: create-secrets
create-secrets:
	printf "secret_password" | DOCKER_HOST=${DOCKER_HOST} docker secret create postgres-passwd -
	printf "postgres://postgres:username@db:5432/postgres" | DOCKER_HOST=${DOCKER_HOST} docker secret create database-url -

.PHONY: remove-secrets
remove-secrets:
	DOCKER_HOST=${DOCKER_HOST} docker secret rm postgres-passwd database-url

