#!/bin/bash

if [ -z "$1" ]; then
  echo "Usage: ./manual_migrate.sh <DATABASE_CONNECTION_URL>"
  exit 1
fi

DB_URL="$1"
MIGRATIONS_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)/../backend/migrations"

echo "Running migrations from $MIGRATIONS_DIR..."

# Sort files to ensure correct order
for file in $(ls "$MIGRATIONS_DIR"/*.up.sql | sort); do
    echo "Applying $(basename "$file")..."
    psql "$DB_URL" -f "$file"
    if [ $? -ne 0 ]; then
        echo "Failed to apply $(basename "$file")"
        exit 1
    fi
done

echo "✅ All migrations applied successfully!"
