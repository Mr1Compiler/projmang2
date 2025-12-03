# Migrations

```bash
# Apply all migrations
migrate -path ./migrations -database "postgres://user:pass@localhost/dbname?sslmode=disable" up

# Rollback last migration
migrate -path ./migrations -database "postgres://user:pass@localhost/dbname?sslmode=disable" down 1

# Rollback all migrations
migrate -path ./migrations -database "postgres://user:pass@localhost/dbname?sslmode=disable" down
```
