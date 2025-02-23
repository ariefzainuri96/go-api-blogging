## Migration
1. Create migration file:
migrate create -seq -ext sql -dir ././cmd/migrate/migrations create_users

2. Perform migration:
migrate -path ./cmd/migrate/migrations -database="postgres://postgres:110638@localhost:5432/go-crash-course?sslmode=disable" up