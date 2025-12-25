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

# Debug: Print DB Host (partially masked)
echo "----------------------------------------"
echo "STARTING BACKEND SERVICE"
echo "DB_HOST: $DB_HOST"
echo "DB_PORT: $DB_PORT"
echo "DB_USER: $DB_USER"
echo "Listening on PORT: $PORT"
echo "----------------------------------------"

# Run migrations
echo "Running migrations..."
/usr/local/bin/migrate -path /migrations -database "$DB_URL" up || {
    echo "❌ MIGRATION FAILED!"
    echo "Check your DB credentials and connection."
    # We exit here because if migrations fail, app will likely panic
    exit 1
}

echo "Migrations completed. Starting application..."

# Run the Go binary and capture output
/api 2>&1
EXIT_CODE=$?

echo "❌ Application exited with code $EXIT_CODE"
exit $EXIT_CODE
