#!/bin/bash

# Exit immediately if any command fails
set -e

echo "🟡 Starting Besu local blockchain network..."

cd besu
./startDev.sh
cd ..

echo "🟡 Starting PostgreSQL database with Docker Compose..."

cd app/
docker-compose up -d

echo "🟢 Running Go application..."