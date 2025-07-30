#!/bin/bash

set -e

echo "🟡 Starting Besu local blockchain network..."
cd besu
./startDev.sh
cd ..

echo "🟡 Starting PostgreSQL database with Docker Compose..."
cd app/docker
docker-compose -f docker-compose.db.yaml up -d
cd ../..

echo "⏳ Waiting for Postgres (docker exec)..."
until docker exec goledger-db pg_isready -U admin -d goledger >/dev/null 2>&1; do
  echo "Postgres is still unavailable - sleeping"
  sleep 1
done
echo "🟢 Postgres is up!"

echo "🟢 Running Go application..."
cd app
go run main.go