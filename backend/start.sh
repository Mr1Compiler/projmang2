#!/bin/sh
set -e

echo "Running migrations..."

# Ensure strictly required variables are present
if [ -z "$DB_HOST" ]; then
    echo "Error: DB_HOST is not set"
    exit 1
fi

# Construct DB_URL string for the migrate tool
# Note: Railway provides PG* variables, but we mapped them to DB_* in the Service settings.
export DB_URL="postgres://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=$DB_SSLMODE"

echo "Connecting to database at $DB_HOST:$DB_PORT..."

# Run migrations
/usr/local/bin/migrate -path /migrations -database "$DB_URL" up

echo "Migrations completed successfully."
echo "Starting application..."

# Run the Go binary
exec /api
