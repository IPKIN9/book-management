#!/bin/bash
set -e

# Start the PostgreSQL server in the background
docker-entrypoint.sh postgres &

# Wait for PostgreSQL to be ready
until pg_isready -h "localhost" -p "5432"; do
  echo "Waiting for PostgreSQL to be ready..."
  sleep 2
done

# Run the init-db.sh script
/docker-entrypoint-initdb.d/init-db.sh

# Bring PostgreSQL to the foreground
wait
