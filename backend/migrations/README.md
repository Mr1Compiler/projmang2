# Migrations

## Install migrate CLI
```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

## Usage
```bash
# Set your database URL
export DB_URL="postgres://user:pass@localhost:5432/dbname?sslmode=disable"

# Check current migration version
migrate -path ./migrations -database "$DB_URL" version

# Apply all pending migrations
migrate -path ./migrations -database "$DB_URL" up

# Apply N migrations
migrate -path ./migrations -database "$DB_URL" up 2

# Rollback last migration
migrate -path ./migrations -database "$DB_URL" down 1

# Rollback N migrations
migrate -path ./migrations -database "$DB_URL" down 2

# Rollback all migrations
migrate -path ./migrations -database "$DB_URL" down -all

# Force set version (use when migration is dirty)
migrate -path ./migrations -database "$DB_URL" force 1

# Create new migration
migrate create -ext sql -dir ./migrations -seq migration_name
```
