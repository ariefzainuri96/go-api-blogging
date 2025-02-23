## Migration
1. Create migration file:
migrate create -seq -ext sql -dir ././cmd/migrate/migrations create_users

2. Perform migration:
migrate -path ./cmd/migrate/migrations -database="postgres://postgres:110638@localhost:5432/go-crash-course?sslmode=disable" up

## Docker
1. running docker container that we specify in docker-compose.yml:
docker compose up -d

2. stop docker container:
docker compose down

3. remove docker container with its volumes:
docker compose down -v