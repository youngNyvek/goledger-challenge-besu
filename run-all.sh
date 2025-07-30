#!/bin/bash

# Exit immediately if any command fails
set -e

echo "🟡 Starting Besu local blockchain network..."

cd besu
./startDev.sh
cd ..

echo "🟡 Starting PostgreSQL database with Docker Compose..."

cd smartcontract-api/docker
docker-compose -f docker-compose.db.yaml up -d
cd ../..

echo "🟢 Running Go application..."
cd smartcontract-api
go run main.go