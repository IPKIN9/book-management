#!/bin/bash
set -e

# Function to check if a database exists
function database_exists() {
  psql -U "$POSTGRES_USER" -tc "SELECT 1 FROM pg_database WHERE datname = '$1'" | grep -q 1
}

# Create multiple databases if they don't exist
for db in book_database author_database category_database user_database; do
  if database_exists $db; then
    echo "Database $db already exists"
  else
    psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
      CREATE DATABASE $db;
EOSQL
    echo "Database $db created"
  fi
done
