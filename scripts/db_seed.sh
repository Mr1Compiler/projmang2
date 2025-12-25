#!/bin/bash

if [ -z "$1" ]; then
  echo "Usage: ./db_seed.sh <DATABASE_CONNECTION_URL>"
  echo "Example: ./db_seed.sh \"postgres://user:pass@host:port/dbname?sslmode=require\""
  echo "You can copy the 'Postgres Connection URL' from Railway -> Database Service -> Connect tab."
  exit 1
fi

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
SEED_FILE="$DIR/../backend/scripts/seed.sql"

if [ ! -f "$SEED_FILE" ]; then
    echo "Error: Seed file not found at $SEED_FILE"
    exit 1
fi

echo "Seeding database from $SEED_FILE..."

# Check if psql is installed
if ! command -v psql &> /dev/null; then
    echo "Error: 'psql' is not installed. Please install postgresql-client."
    echo "Ubuntu/Debian: sudo apt-get install postgresql-client"
    exit 1
fi

psql "$1" -f "$SEED_FILE"

if [ $? -eq 0 ]; then
    echo "✅ Database seeded successfully!"
else
    echo "❌ Seeding failed."
fi
